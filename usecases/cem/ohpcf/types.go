package ohpcf

import "github.com/enbility/eebus-go/api"

const (
	// Update of the list of remote entities supporting the Use Case
	//
	// Use `RemoteEntities` to get the current data
	UseCaseSupportUpdate api.EventType = "cem-ohpcf-UseCaseSupportUpdate"

	// SmartEnergyManagementPs data of a compressor was updated
	//
	// Use `PowerConsumptionAlternatives` to get the current data
	//
	// Use Case OHPCF, Scenario 1
	DataUpdatePowerConsumptionAlternatives api.EventType = "cem-ohpcf-DataUpdatePowerConsumptionAlternatives"
)
