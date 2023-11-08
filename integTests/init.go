package integTests

import (
	"github.com/xdblab/xdb-golang-sdk/integTests/basic"
	"github.com/xdblab/xdb-golang-sdk/integTests/command_request"
	"github.com/xdblab/xdb-golang-sdk/integTests/failure_recovery"
	"github.com/xdblab/xdb-golang-sdk/integTests/global_attribute"
	"github.com/xdblab/xdb-golang-sdk/integTests/multi_states"
	"github.com/xdblab/xdb-golang-sdk/integTests/process_timeout"
	"github.com/xdblab/xdb-golang-sdk/integTests/state_decision"
	"github.com/xdblab/xdb-golang-sdk/integTests/stateretry"
	"github.com/xdblab/xdb-golang-sdk/xdb"
)

var registry = xdb.NewRegistry()
var client = xdb.NewClient(registry, getTestClientOptions())
var workerService = xdb.NewWorkerService(registry, nil)

func init() {
	err := registry.AddProcesses(
		&basic.IOProcess{},
		&failure_recovery.StateFailureRecoveryTestExecuteProcess{},
		&failure_recovery.StateFailureRecoveryTestWaitUntilProcess{},
		&failure_recovery.StateFailureRecoveryTestExecuteNoWaitUntilProcess{},
		&failure_recovery.StateFailureRecoveryTestExecuteFailedAtStartProcess{},
		&multi_states.MultiStatesProcess{},
		&state_decision.GracefulCompleteProcess{},
		&state_decision.ForceCompleteProcess{},
		&state_decision.ForceFailProcess{},
		&state_decision.DeadEndProcess{},
		&stateretry.BackoffProcess{},
		&global_attribute.SingleTableProcess{},
		&global_attribute.MultiTablesProcess{},
		&process_timeout.TimeoutProcess{},
		&command_request.AnyOfTimerLocalQProcess{},
	)
	if err != nil {
		panic(err)
	}
}

func getTestClientOptions() *xdb.ClientOptions {
	options := xdb.GetLocalDefaultClientOptions()
	options.DefaultProcessTimeoutSecondsOverride = 10
	return options
}
