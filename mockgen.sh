#! /bin/bash -e

go get github.com/aporeto-inc/mock/mockgen
go get -u github.com/golang/mock/gomock
go get -u golang.org/x/tools/cmd/goimports

goimport_sanitize () {
  goimports $1 > $1.bk
  cat $1.bk | sed  $'s/^func /\/\/ nolint\\\nfunc /g' | sed  $'s/^type /\/\/ nolint\\\ntype /g' > $1
  rm -f $1.bk
}

echo "Cgnetcls Mocks"
mkdir -p cgnetcls/mock
mockgen -source cgnetcls/interfaces.go -destination cgnetcls/mock/mockcgnetcls.go -package mockcgnetcls -source_package github.com/aporeto-inc/trireme-lib/cgnetcls
goimport_sanitize cgnetcls/mock/mockcgnetcls.go

echo "Enforcer/PolicyEnforcer Mocks"
mkdir -p enforcer/policyenforcer/mock
mockgen -source enforcer/policyenforcer/interfaces.go -destination enforcer/policyenforcer/mock/mockpolicyenforcer.go -package mockpolicyenforcer -source_package github.com/aporeto-inc/trireme-lib/enforcer
goimport_sanitize enforcer/policyenforcer/mock/mockpolicyenforcer.go

echo "Supervisor Mocks"
mkdir -p supervisor/mock
mockgen -source supervisor/interfaces.go -destination supervisor/mock/mocksupervisor.go -package mocksupervisor
goimport_sanitize supervisor/mock/mocksupervisor.go

echo "Internal/ContextStore Mocks"
mkdir -p internal/contextstore/mock
mockgen -source internal/contextstore/interfaces.go -destination internal/contextstore/mock/mockcontextstore.go -package mockcontextstore
goimport_sanitize internal/contextstore/mock/mockcontextstore.go

echo "Internal/Processmon Mocks"
mkdir -p internal/processmon/mock
mockgen -source internal/processmon/interfaces.go -destination internal/processmon/mock/mockprocessmon.go -package mockprocessmon
goimport_sanitize internal/processmon/mock/mockprocessmon.go

echo "Internal/RemoteEnforcer Mocks"
mkdir -p internal/remoteenforcer/mock
mockgen -source internal/remoteenforcer/interfaces.go -destination internal/remoteenforcer/mock/mockremoteenforcer.go -package mockremoteenforcer
goimport_sanitize internal/remoteenforcer/mock/mockremoteenforcer.go

echo "Internal/RemoteEnforcer/StatsClient Mocks"
mkdir -p internal/remoteenforcer/internal/statsclient/mock
mockgen -source internal/remoteenforcer/internal/statsclient/interfaces.go -destination internal/remoteenforcer/internal/statsclient/mock/mockstatsclient.go -package mockstatsclient
goimport_sanitize internal/remoteenforcer/internal/statsclient/mock/mockstatsclient.go

echo "Internal/RemoteEnforcer/StatsCollector Mocks"
mkdir -p internal/remoteenforcer/internal/statscollector/mock
mockgen -source internal/remoteenforcer/internal/statscollector/interfaces.go -aux_files collector=collector/interfaces.go -destination internal/remoteenforcer/internal/statscollector/mock/mockstatscollector.go -package mockstatscollector
goimport_sanitize internal/remoteenforcer/internal/statscollector/mock/mockstatscollector.go

echo "Collector Mocks"
mkdir -p collector/mock
mockgen -source collector/interfaces.go -destination collector/mock/mockcollector.go -package mockcollector -source_package github.com/aporeto-inc/trireme-lib/collector
goimport_sanitize collector/mock/mockcollector.go

echo "Monitor Mocks"
mkdir -p monitor/mock
mockgen -source monitor/interfaces.go -destination monitor/mock/mockmonitor.go -package mockmonitor -source_package github.com/aporeto-inc/trireme-lib/monitor
goimport_sanitize monitor/mock/mockmonitor.go

echo "Monitor/Instance Mocks"
mkdir -p monitor/instance/mock
mockgen -source monitor/instance/interfaces.go -destination monitor/instance/mock/mockinstance.go -aux_files collector=collector/interfaces.go -package mockinstance -source_package github.com/aporeto-inc/trireme-lib/monitor/instance
goimport_sanitize monitor/instance/mock/mockinstance.go

echo "Monitor/RPC/Processor Mocks"
mkdir -p monitor/rpc/processor/mock
mockgen -source monitor/rpc/processor/interfaces.go -destination monitor/rpc/processor/mock/mockprocessor.go -aux_files collector=collector/interfaces.go -package mockprocessor -source_package github.com/aporeto-inc/trireme-lib/monitor/rpc/processor
goimport_sanitize monitor/rpc/processor/mock/mockprocessor.go

echo "Trireme Mocks"
mkdir -p mock
mockgen -source interfaces.go -destination mock/mocktrireme.go -package mocktrireme  -aux_files constants=constants/constants.go events=monitor/rpc/events/interfaces.go policy=policy/interfaces.go processor=monitor/rpc/processor/interfaces.go supervisor=supervisor/interfaces.go -source_package github.com/aporeto-inc/trireme-lib
goimport_sanitize mock/mocktrireme.go

echo >&2 "OK"