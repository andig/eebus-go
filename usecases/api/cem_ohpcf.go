package api

import (
	"github.com/enbility/eebus-go/api"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"time"
)

type CemOHPCFInterface interface {
	api.UseCaseInterface

	// Scenario 1

	// The availability of an optional consumption of power [OHPCF-011/1].
	//
	// return true if the optional consumption of power is available
	OptionalPowerConsumptionAvailable(entity spineapi.EntityRemoteInterface) (bool, error)

	// The power value [OHPCF-011/2/1 or 2].
	//
	// return the power value
	Power(entity spineapi.EntityRemoteInterface) (float64, error)

	// Indication whether the consumption may be stopped by the CEM [OHPCF-011/5].
	//
	// return true if the consumption may be stopped
	ConsumptionIsStoppable(entity spineapi.EntityRemoteInterface) (bool, error)

	// Indication whether the consumption may be paused and resumed by the CEM [OHPCF-011/6]
	//
	// return true if the consumption may be paused
	ConsumptionIsPausable(entity spineapi.EntityRemoteInterface) (bool, error)

	// The start time of the process [OHPCF-012/1].
	//
	// return the start time of the process
	PowerConsumptionProcessStartTime(entity spineapi.EntityRemoteInterface) (time.Time, error)

	// The current state of this power consumption process [OHPCF-012/2].
	//
	// return the current state of this power consumption process
	PowerConsumptionProcessState(entity spineapi.EntityRemoteInterface) (CompressorPowerConsumptionStateType, error)

	// The minimal time a consumption process must last [OHPCF-008].
	//
	// return the minimal time a consumption process must last
	PowerConsumptionMinimalRunDuration(entity spineapi.EntityRemoteInterface) (time.Duration, error)

	// The minimal time a pause of a consumption process must last [OHPCF-009].
	//
	// return the minimal time a pause of a consumption process must last
	PowerConsumptionMinimalPauseDuration(entity spineapi.EntityRemoteInterface) (time.Duration, error)

	// Scenario 2

	// Schedule an optional power consumption process [OHPCF-004].
	//
	// note:
	// A re-schedule of an already scheduled power consumption process is possible as long as the
	// scheduled process did not start.
	//
	// parameters:
	//   - start: The start time of the power consumption
	SchedulePowerConsumptionProcess(entity spineapi.EntityRemoteInterface, start time.Time, callback func(msg spineapi.ResponseMessage)) (*model.MsgCounterType, error)

	// stop (abort) the process [OHPCF-022/1].
	StopAbortPowerConsumptionProcess(entity spineapi.EntityRemoteInterface, callback func(msg spineapi.ResponseMessage)) (*model.MsgCounterType, error)

	// pause the process [OHPCF-022/2].
	PausePowerConsumptionProcess(entity spineapi.EntityRemoteInterface, callback func(msg spineapi.ResponseMessage)) (*model.MsgCounterType, error)

	// resume the process [OHPCF-022/3].
	ResumePowerConsumptionProcess(entity spineapi.EntityRemoteInterface, callback func(msg spineapi.ResponseMessage)) (*model.MsgCounterType, error)
}
