// +build windows

package provider

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"github.com/DavidGamba/go-getoptions"
	"go.aporeto.io/trireme-lib/controller/internal/windows/frontman"
	"go.uber.org/zap"
)

// IptablesProvider is an abstraction of all the methods an implementation of userspace
// iptables need to provide.
type IptablesProvider interface {
	BaseIPTables
	// Commit will commit changes if it is a batch provider.
	Commit() error
	// RetrieveTable allows a caller to retrieve the final table.
	RetrieveTable() map[string]map[string][]string
}

// BaseIPTables is the base interface of iptables functions.
type BaseIPTables interface {
	// Append apends a rule to chain of table
	Append(table, chain string, rulespec ...string) error
	// Insert inserts a rule to a chain of table at the required pos
	Insert(table, chain string, pos int, rulespec ...string) error
	// Delete deletes a rule of a chain in the given table
	Delete(table, chain string, rulespec ...string) error
	// ListChains lists all the chains associated with a table
	ListChains(table string) ([]string, error)
	// ClearChain clears a chain in a table
	ClearChain(table, chain string) error
	// DeleteChain deletes a chain in the table. There should be no references to this chain
	DeleteChain(table, chain string) error
	// NewChain creates a new chain
	NewChain(table, chain string) error
}

// BatchProvider uses iptables-restore to program ACLs
type BatchProvider struct {
}

// NewGoIPTablesProviderV4 returns an IptablesProvider interface based on the go-iptables
// external package.
func NewGoIPTablesProviderV4(batchTables []string) (*BatchProvider, error) {
	return &BatchProvider{}, nil
}

// NewGoIPTablesProviderV6 returns an IptablesProvider interface based on the go-iptables
// external package.
func NewGoIPTablesProviderV6(batchTables []string) (*BatchProvider, error) {
	return &BatchProvider{}, nil
}

// NewCustomBatchProvider is a custom batch provider wher the downstream
// iptables utility is provided by the caller. Very useful for testing
// the ACL functions with a mock.
func NewCustomBatchProvider(ipt BaseIPTables, commit func(buf *bytes.Buffer) error, batchTables []string) *BatchProvider {
	return &BatchProvider{}
}

type windowsRuleSpec struct {
	protocol        int
	matchPortBegin  int
	matchPortEnd    int
	matchSet        string // ipset name
	matchSetNegate  bool
	matchSetDstIp   bool
	matchSetDstPort bool
	matchSetSrcIp   bool
	matchSetSrcPort bool
	action          int // FilterAction (allow, drop, nfq, proxy)
	proxyPort       int
	mark            int
}

func (b *BatchProvider) parseRuleSpec(rulespec ...string) (*windowsRuleSpec, error) {

	opt := getoptions.New()

	protocolOpt := opt.String("p", "")
	sdPortOpt := opt.String("sport", "", opt.Alias("dport"))
	actionOpt := opt.StringSlice("j", 1, 10, opt.Required())
	modeOpt := opt.StringSlice("m", 1, 2)
	matchSetOpt := opt.StringSlice("match-set", 2, 2)
	redirectPortOpt := opt.Int("to-ports", 0)

	_, err := opt.Parse(rulespec)
	if err != nil {
		return nil, err
	}

	result := &windowsRuleSpec{}

	// protocol: either "tcp" or "udp"
	if strings.HasPrefix(*protocolOpt, "T") || strings.HasPrefix(*protocolOpt, "t") {
		result.protocol = 6
	} else if strings.HasPrefix(*protocolOpt, "U") || strings.HasPrefix(*protocolOpt, "u") {
		result.protocol = 17
	}

	// source or dest port: either port or port range
	if *sdPortOpt != "" {
		result.matchPortBegin, err = strconv.Atoi(*sdPortOpt)
		if err != nil {
			sdPortRange := strings.SplitN(*sdPortOpt, ":", 2)
			if len(sdPortRange) != 2 {
				return nil, errors.New("rulespec not valid: invalid match port")
			}
			result.matchPortBegin, err = strconv.Atoi(sdPortRange[0])
			if err != nil {
				return nil, errors.New("rulespec not valid: invalid match port")
			}
			result.matchPortEnd, err = strconv.Atoi(sdPortRange[1])
			if err != nil {
				return nil, errors.New("rulespec not valid: invalid match port")
			}
		}
		if result.matchPortEnd == 0 {
			result.matchPortEnd = result.matchPortBegin
		}
	}

	// match ipset
	if len(*modeOpt) > 0 && (strings.HasPrefix((*modeOpt)[0], "S") || strings.HasPrefix((*modeOpt)[0], "s")) {
		// see if negate of --match-set occurred
		if len(*modeOpt) > 1 && (*modeOpt)[1] == "!" {
			result.matchSetNegate = true
		}
		// now check --match-set opt
		if len(*matchSetOpt) != 2 {
			return nil, errors.New("rulespec not valid: no ipset name")
		}
		// first part is the ipset name
		result.matchSet = (*matchSetOpt)[0]
		// second part is the dst/src match specifier
		ipPortSpecLower := (*matchSetOpt)[1]
		if strings.HasPrefix(ipPortSpecLower, "dstip") {
			result.matchSetDstIp = true
		} else if strings.HasPrefix(ipPortSpecLower, "srcip") {
			result.matchSetSrcIp = true
		}
		if strings.HasSuffix(ipPortSpecLower, "dstport") {
			result.matchSetDstPort = true
			if result.protocol == 0 {
				return nil, errors.New("rulespec not valid: ipset match on port requires protocol be set")
			}
		} else if strings.HasSuffix(ipPortSpecLower, "srcport") {
			result.matchSetSrcPort = true
			if result.protocol == 0 {
				return nil, errors.New("rulespec not valid: ipset match on port requires protocol be set")
			}
		}
	}

	// action: required, either NFQUEUE, REDIRECT, MARK, ACCEPT, DROP
	for i := 0; i < len(*actionOpt); i++ {
		a := (*actionOpt)[i]
		if strings.HasPrefix(a, "N") || strings.HasPrefix(a, "n") {
			result.action = frontman.FilterActionNfq
		} else if strings.HasPrefix(a, "R") || strings.HasPrefix(a, "r") {
			result.action = frontman.FilterActionProxy
		} else if strings.HasPrefix(a, "M") || strings.HasPrefix(a, "m") {
			i++
			if i >= len(*actionOpt) {
				return nil, errors.New("rulespec not valid: no mark given")
			}
			result.mark, err = strconv.Atoi((*actionOpt)[i])
			if err != nil {
				return nil, errors.New("rulespec not valid: mark should be int32")
			}
		} else if strings.HasPrefix(a, "A") || strings.HasPrefix(a, "a") {
			result.action = frontman.FilterActionAllow
		} else if strings.HasPrefix(a, "D") || strings.HasPrefix(a, "d") {
			result.action = frontman.FilterActionBlock
		} else {
			return nil, errors.New("rulespec not valid: invalid action")
		}
	}

	// redirect port
	result.proxyPort = *redirectPortOpt
	if result.action == frontman.FilterActionProxy && result.proxyPort == 0 {
		return nil, errors.New("rulespec not valid: no redirect port given")
	}

	return result, nil
}

// Append will append the provided rule to the local cache or call
// directly the iptables command depending on the table.
func (b *BatchProvider) Append(table, chain string, rulespec ...string) error {
	zap.L().Error("Append",
		zap.String("Table", table),
		zap.String("Chain", chain),
		zap.Strings("Rules", rulespec))

	_, err := b.parseRuleSpec(rulespec...)
	if err != nil {
		return err
	}

	//driverHandle, err := frontman.GetDriverHandle()
	//if err != nil {
	//	return err
	//}

	// TODO
	//dllRet, _, err := frontman.SetFilterCriteriaProc.Call(driverHandle,
	//	uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(chain))),
	//	0x40, //criteria.mark,
	//	uintptr(frontman.FilterActionNfq),
	//	0,      // log=false
	//	0,      //criteria.proxyPort, // proxyPort,
	//	0,      // matchIpSet
	//	6,      // tcp=6
	//	80, 80) //criteria.matchPortBegin, criteria.matchPortEnd)

	//if dllRet == 0 {
	//	return fmt.Errorf("%s failed (%v)", frontman.SetFilterCriteriaProc.Name, err)
	//}

	return nil
}

// Insert will insert the rule in the corresponding position in the local
// cache or call the corresponding iptables command, depending on the table.
func (b *BatchProvider) Insert(table, chain string, pos int, rulespec ...string) error {
	return nil
}

// Delete will delete the rule from the local cache or the system.
func (b *BatchProvider) Delete(table, chain string, rulespec ...string) error {
	zap.L().Error("Delete",
		zap.String("Table", table),
		zap.String("Chain", chain),
		zap.Strings("Rules", rulespec))

	return nil
}

// ListChains will provide a list of the current chains.
func (b *BatchProvider) ListChains(table string) ([]string, error) {
	return []string{}, nil
}

// ClearChain will clear the chains.
func (b *BatchProvider) ClearChain(table, chain string) error {
	driverHandle, err := frontman.GetDriverHandle()
	if err != nil {
		return err
	}

	dllRet, _, err := frontman.EmptyFilterProc.Call(driverHandle, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(chain))))
	if dllRet == 0 {
		return fmt.Errorf("%s failed (%v)", frontman.EmptyFilterProc.Name, err)
	}

	return nil
}

// DeleteChain will delete the chains.
func (b *BatchProvider) DeleteChain(table, chain string) error {
	driverHandle, err := frontman.GetDriverHandle()
	if err != nil {
		return err
	}

	dllRet, _, err := frontman.DestroyFilterProc.Call(driverHandle, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(chain))))
	if dllRet == 0 {
		return fmt.Errorf("%s failed (%v)", frontman.DestroyFilterProc.Name, err)
	}

	return nil
}

// NewChain creates a new chain.
func (b *BatchProvider) NewChain(table, chain string) error {
	driverHandle, err := frontman.GetDriverHandle()
	if err != nil {
		return err
	}

	var outbound uintptr
	if strings.HasPrefix(table, "O") || strings.HasPrefix(table, "o") {
		outbound = 1
	} else if strings.HasPrefix(table, "I") || strings.HasPrefix(table, "i") {
		outbound = 0
	} else {
		return fmt.Errorf("'%s' is not a valid table for NewChain", table)
	}

	dllRet, _, err := frontman.AppendFilterProc.Call(driverHandle, outbound, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(chain))))
	if dllRet == 0 {
		return fmt.Errorf("%s failed (%v)", frontman.AppendFilterProc.Name, err)
	}

	return nil
}

// Commit commits the rules to the system
func (b *BatchProvider) Commit() error {
	return nil
}

// RetrieveTable allows a caller to retrieve the final table. Mostly
// needed for debuging and unit tests.
func (b *BatchProvider) RetrieveTable() map[string]map[string][]string {
	return map[string]map[string][]string{}
}
