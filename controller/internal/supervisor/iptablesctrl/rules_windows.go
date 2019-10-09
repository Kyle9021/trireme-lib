// +build windows

package iptablesctrl

var triremChains = `
-t OUTPUT  -N GlobalRules-OUTPUT
-t INPUT   -N GlobalRules-INPUT
-t OUTPUT  -N HostSvcRules-OUTPUT
-t INPUT   -N HostSvcRules-INPUT
-t OUTPUT  -N HostPU-OUTPUT
-t INPUT   -N HostPU-INPUT
`
var globalRules = `
-A  GlobalRules-INPUT -m set  --match-set {{.ExclusionsSet}} srcIP -j ACCEPT
-A  GlobalRules-OUTPUT -m set  --match-set {{.ExclusionsSet}} dstIP -j ACCEPT
`

var containerChains = ``

// TODO(windows): add drop/log rule for host pu at end of chain (ipv6 support needed first)
// -A  HostPU-OUTPUT -p tcp -j DROP -j NFLOG
// -A  HostPU-INPUT -p tcp -j DROP -j NFLOG

// cgroupCaptureTemplate are the list of iptables commands that will hook traffic and send it to a PU specific
// chain. The hook method depends on the type of PU.
var cgroupCaptureTemplate = `
{{if isHostPU}}
-A HostPU-OUTPUT -p tcp -m set --match-set {{.TargetTCPNetSet}} dstIP -m set --match-set {{.DestIPSet}} dstIP,dstPort -j REDIRECT  --to-ports {{.ProxyPort}}
-A HostPU-OUTPUT -p tcp -m set --match-set {{.TargetTCPNetSet}} dstIP -j NFQUEUE -j MARK {{.Mark}}
-A HostPU-OUTPUT -p udp -m set --match-set {{.TargetUDPNetSet}} dstIP -j NFQUEUE -j MARK {{.Mark}}
-A HostPU-INPUT -p tcp -m set --match-set {{.SrvIPSet}} dstPort -j REDIRECT --to-ports {{.ProxyPort}}
-A HostPU-INPUT -p tcp -m set --match-set {{.TargetTCPNetSet}} srcIP -j NFQUEUE -j MARK {{.Mark}}
-A HostPU-INPUT -p udp -m set --match-set {{.TargetUDPNetSet}} srcIP -m string --string {{.UDPSignature}} --offset 2 -j NFQUEUE -j MARK {{.Mark}}
{{else}}
-A HostSvcRules-INPUT -p tcp -m set --match-set {{.SrvIPSet}} dstPort -j REDIRECT --to-ports {{.ProxyPort}}
-A HostSvcRules-INPUT -p tcp --dport {{.TCPPorts}} -j NFQUEUE -j MARK {{.Mark}}
-A HostSvcRules-INPUT -p udp --dport {{.UDPPorts}} -m string --string {{.UDPSignature}} --offset 2 -j NFQUEUE -j MARK {{.Mark}}
-A HostSvcRules-OUTPUT -p tcp --sport {{.TCPPorts}} -j NFQUEUE -j MARK {{.Mark}}
-A HostSvcRules-OUTPUT -p udp --sport {{.UDPPorts}} -j NFQUEUE -j MARK {{.Mark}}
{{end}}
`

// containerChainTemplate will hook traffic towards the container specific chains.
var containerChainTemplate = ``

var uidChainTemplate = ``

var acls = ``

// packetCaptureTemplate are the rules that trap traffic towards the user space.
var packetCaptureTemplate = ``

var proxyChainTemplate = ``

var deleteChains = ``

var globalHooks = ``

var legacyProxyRules = ``