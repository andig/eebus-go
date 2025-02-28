package ohpcf

import "github.com/enbility/eebus-go/api"

const (
	UseCaseSupportUpdate api.EventType = "cem-ohpcf-UseCaseSupportUpdate"

	// Scenario 1

	DataUpdatePower api.EventType = "cem-ohpcf-DataUpdatePowerChanged"

	DataUpdateConsumptionIsStoppable api.EventType = "cem-ohpcf-DataUpdateConsumptionIsStoppable"

	DataUpdateConsumptionIsPausable api.EventType = "cem-ohpcf-DataUpdateConsumptionIsPausable"

	DataUpdateConsumptionStartTime api.EventType = "cem-ohpcf-DataUpdateConsumptionStartTime"

	DataUpdateConsumptionState api.EventType = "cem-ohpcf-DataUpdateConsumptionState"

	DataUpdateOptionalPowerConsumptionAvailable api.EventType = "cem-ohpcf-DataUpdateOptionalPowerConsumptionAvailable"

	DataUpdateOptionalPowerConsumptionScheduled api.EventType = "cem-ohpcf-DataUpdateOptionalPowerConsumptionScheduled"

	DataUpdateOptionalPowerConsumptionRunning api.EventType = "cem-ohpcf-DataUpdateOptionalPowerConsumptionRunning"

	DataUpdateOptionalPowerConsumptionPaused api.EventType = "cem-ohpcf-DataUpdateOptionalPowerConsumptionPaused"

	DataUpdateOptionalPowerConsumptionCompleted api.EventType = "cem-ohpcf-DataUpdateOptionalPowerConsumptionCompleted"

	DataUpdateOptionalPowerConsumptionStopped api.EventType = "cem-ohpcf-DataUpdateOptionalPowerConsumptionStopped"

	DataUpdateMinimalRunDuration api.EventType = "cem-ohpcf-DataUpdateMinimalRunDuration"

	DataUpdateMinimalPauseDuration api.EventType = "cem-ohpcf-DataUpdateMinimalPauseDuration"
)
