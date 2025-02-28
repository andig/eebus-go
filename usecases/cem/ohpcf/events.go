package ohpcf

import (
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

// handle SPINE events
func (o *OHPCF) HandleEvent(payload spineapi.EventPayload) {
	// only about events from an EV entity or device changes for this remote device

	if !o.IsCompatibleEntityType(payload.Entity) {
		return
	}

	if payload.Data == nil {
		return
	}

	switch payload.Data.(type) {
	case *model.SmartEnergyManagementPsDataType:
		o.loadSmartEnergyManagementPsDataType(payload)
		break
	}
}

func (o *OHPCF) loadSmartEnergyManagementPsDataType(payload spineapi.EventPayload) {
	data := payload.Data.(*model.SmartEnergyManagementPsDataType)

	if len(data.Alternatives) == 1 {
		if len(data.Alternatives[0].PowerSequence) == 1 &&
			o.EventCB != nil {

			if len(data.Alternatives[0].PowerSequence[0].PowerTimeSlot) == 1 &&
				data.Alternatives[0].PowerSequence[0].PowerTimeSlot[0].ValueList != nil &&
				len(data.Alternatives[0].PowerSequence[0].PowerTimeSlot[0].ValueList.Value) == 1 &&
				data.Alternatives[0].PowerSequence[0].PowerTimeSlot[0].ValueList.Value[0].Value != nil {
				o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdatePower)
			}

			if data.Alternatives[0].PowerSequence[0].OperatingConstraintsInterrupt != nil &&
				data.Alternatives[0].PowerSequence[0].OperatingConstraintsInterrupt.IsStoppable != nil {
				// [OHPCF-011/5]
				// [OHPCF-012/3]
				o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateConsumptionIsStoppable)
			}

			if data.Alternatives[0].PowerSequence[0].OperatingConstraintsInterrupt != nil &&
				data.Alternatives[0].PowerSequence[0].OperatingConstraintsInterrupt.IsPausable != nil {
				// [OHPCF-011/6]
				// [OHPCF-012/3]
				o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateConsumptionIsPausable)
			}

			if data.Alternatives[0].PowerSequence[0].Schedule != nil &&
				data.Alternatives[0].PowerSequence[0].Schedule.StartTime != nil {
				// [OHPCF-012/1]
				o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateConsumptionStartTime)
			}

			if data.Alternatives[0].PowerSequence[0].State != nil &&
				data.Alternatives[0].PowerSequence[0].State.State != nil {
				// [OHPCF-012/2]
				o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateConsumptionState)

				if *data.Alternatives[0].PowerSequence[0].State.State == model.PowerSequenceStateTypeInactive {
					// [OHPCF-011/1]
					o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOptionalPowerConsumptionAvailable)
				}

				if *data.Alternatives[0].PowerSequence[0].State.State == model.PowerSequenceStateTypeScheduled {
					// [OHPCF-012/2/1]
					o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOptionalPowerConsumptionScheduled)
				}

				if *data.Alternatives[0].PowerSequence[0].State.State == model.PowerSequenceStateTypeRunning {
					//[OHPCF-012/2/2], [OHPCF-012/4], [OHPCF-022/3]
					o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOptionalPowerConsumptionRunning)
				}

				if *data.Alternatives[0].PowerSequence[0].State.State == model.PowerSequenceStateTypePaused {
					// [OHPCF-012/2/3], [OHPCF-022/2]
					o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOptionalPowerConsumptionPaused)
				}

				if *data.Alternatives[0].PowerSequence[0].State.State == model.PowerSequenceStateTypeCompleted {
					// [OHPCF-006/3], [OHPCF-012/2/5]
					o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOptionalPowerConsumptionCompleted)
				}

				if *data.Alternatives[0].PowerSequence[0].State.State == model.PowerSequenceStateTypeInvalid {
					// [OHPCF-006/1], [OHPCF-012/2/4], [OHPCF-022/1]
					o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOptionalPowerConsumptionStopped)
				}
			}

			if data.Alternatives[0].PowerSequence[0].OperatingConstraintsDuration != nil {
				if data.Alternatives[0].PowerSequence[0].OperatingConstraintsDuration.ActiveDurationMin != nil {
					// [OHPCF-008]
					o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateMinimalRunDuration)
				}
				if data.Alternatives[0].PowerSequence[0].OperatingConstraintsDuration.PauseDurationMin != nil {
					// [OHPCF-009]
					o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateMinimalPauseDuration)
				}
			}
		} else {
			if o.EventCB != nil {
				// [OHPCF-003], [OHPCF-006/2]
				o.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateOptionalPowerConsumptionStopped)
			}
		}
	}
}
