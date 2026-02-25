package internal

import (
	"github.com/enbility/eebus-go/features/server"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/util"

	"github.com/stretchr/testify/assert"
)

func (s *InternalSuite) Test_ConfigurationWriteRequiresApproval() {
	msg := spineapi.Message{}
	configsToApprove := map[model.DeviceConfigurationKeyNameType]struct{}{
		model.DeviceConfigurationKeyNameTypeFailsafeConsumptionActivePowerLimit: {},
		model.DeviceConfigurationKeyNameTypeFailsafeDurationMinimum:             {},
	}
	// Header missing
	_, err := ConfigurationWriteRequiresApproval(&msg, s.localEntity, configsToApprove)
	assert.NotNil(s.T(), err)

	// MsgCounter missing
	header := model.HeaderType{}
	msg = spineapi.Message{RequestHeader: &header}
	_, err = ConfigurationWriteRequiresApproval(&msg, s.localEntity, configsToApprove)
	assert.NotNil(s.T(), err)

	// DeviceConfigurationKeyValueListData missing
	header = model.HeaderType{MsgCounter: util.Ptr(model.MsgCounterType(1))}
	msg = spineapi.Message{RequestHeader: &header}
	_, err = ConfigurationWriteRequiresApproval(&msg, s.localEntity, configsToApprove)
	assert.NotNil(s.T(), err)

	// DeviceConfigurationKeyValueListData.DeviceConfigurationKeyValueData is nil/length of 0
	cmd := model.CmdType{DeviceConfigurationKeyValueListData: util.Ptr(model.DeviceConfigurationKeyValueListDataType{})}
	msg = spineapi.Message{RequestHeader: &header, Cmd: cmd}
	_, err = ConfigurationWriteRequiresApproval(&msg, s.localEntity, configsToApprove)
	assert.NotNil(s.T(), err)

	// Not all elements in slice of DeviceConfigurationKeyValueDataType have KeyId set
	deviceConfigList := []model.DeviceConfigurationKeyValueDataType{{KeyId: nil}}
	cmd = model.CmdType{DeviceConfigurationKeyValueListData: util.Ptr(model.DeviceConfigurationKeyValueListDataType{DeviceConfigurationKeyValueData: deviceConfigList})}
	msg = spineapi.Message{RequestHeader: &header, Cmd: cmd}
	_, err = ConfigurationWriteRequiresApproval(&msg, s.localEntity, configsToApprove)
	assert.NotNil(s.T(), err)

	// Valid message but not a KeyId we care about => no approval required
	deviceConfigList = []model.DeviceConfigurationKeyValueDataType{{KeyId: util.Ptr(model.DeviceConfigurationKeyIdType(0))}}
	cmd = model.CmdType{DeviceConfigurationKeyValueListData: util.Ptr(model.DeviceConfigurationKeyValueListDataType{DeviceConfigurationKeyValueData: deviceConfigList})}
	msg = spineapi.Message{RequestHeader: &header, Cmd: cmd}
	approvalRequired, err := ConfigurationWriteRequiresApproval(&msg, s.localEntity, configsToApprove)
	assert.Nil(s.T(), err)
	assert.False(s.T(), approvalRequired)

	// Valid message with KeyId we care about => approval required
	if dcs, err := server.NewDeviceConfiguration(s.localEntity); err == nil {
		dcs.AddKeyValueDescription(
			model.DeviceConfigurationKeyValueDescriptionDataType{
				KeyName:   util.Ptr(model.DeviceConfigurationKeyNameTypeFailsafeConsumptionActivePowerLimit),
				ValueType: util.Ptr(model.DeviceConfigurationKeyValueTypeTypeScaledNumber),
				Unit:      util.Ptr(model.UnitOfMeasurementTypeW),
			},
		)

		value := &model.DeviceConfigurationKeyValueValueType{
			ScaledNumber: model.NewScaledNumberType(0),
		}
		_ = dcs.UpdateKeyValueDataForFilter(
			model.DeviceConfigurationKeyValueDataType{
				Value:             value,
				IsValueChangeable: util.Ptr(true),
			},
			nil,
			model.DeviceConfigurationKeyValueDescriptionDataType{
				KeyName: util.Ptr(model.DeviceConfigurationKeyNameTypeFailsafeConsumptionActivePowerLimit),
			},
		)
	}
	approvalRequired, err = ConfigurationWriteRequiresApproval(&msg, s.localEntity, configsToApprove)
	assert.Nil(s.T(), err)
	assert.True(s.T(), approvalRequired)
}

func (s *InternalSuite) Test_GroupPendingDeviceConfigurations() {
	failsafeLimitDesc := model.DeviceConfigurationKeyValueDescriptionDataType{
		KeyName:   util.Ptr(model.DeviceConfigurationKeyNameTypeFailsafeConsumptionActivePowerLimit),
		ValueType: util.Ptr(model.DeviceConfigurationKeyValueTypeTypeScaledNumber),
		Unit:      util.Ptr(model.UnitOfMeasurementTypeW),
		KeyId:     util.Ptr(model.DeviceConfigurationKeyIdType(0)),
	}
	failsafeDurationMinDesc := model.DeviceConfigurationKeyValueDescriptionDataType{
		KeyName:   util.Ptr(model.DeviceConfigurationKeyNameTypeFailsafeDurationMinimum),
		ValueType: util.Ptr(model.DeviceConfigurationKeyValueTypeTypeDuration),
		KeyId:     util.Ptr(model.DeviceConfigurationKeyIdType(1)),
	}
	if dcs, err := server.NewDeviceConfiguration(s.localEntity); err == nil {
		dcs.AddKeyValueDescription(failsafeLimitDesc)

		// only add if it doesn't exist yet
		filter := model.DeviceConfigurationKeyValueDescriptionDataType{
			KeyName: util.Ptr(model.DeviceConfigurationKeyNameTypeFailsafeDurationMinimum),
		}
		if data, err := dcs.GetKeyValueDescriptionsForFilter(filter); err == nil && len(data) == 0 {
			dcs.AddKeyValueDescription(failsafeDurationMinDesc)
		}

		value := &model.DeviceConfigurationKeyValueValueType{
			ScaledNumber: model.NewScaledNumberType(0),
		}
		_ = dcs.UpdateKeyValueDataForFilter(
			model.DeviceConfigurationKeyValueDataType{
				Value:             value,
				IsValueChangeable: util.Ptr(true),
			},
			nil,
			model.DeviceConfigurationKeyValueDescriptionDataType{
				KeyName: util.Ptr(model.DeviceConfigurationKeyNameTypeFailsafeConsumptionActivePowerLimit),
			},
		)

		value = &model.DeviceConfigurationKeyValueValueType{
			Duration: model.NewDurationType(0),
		}
		_ = dcs.UpdateKeyValueDataForFilter(
			model.DeviceConfigurationKeyValueDataType{
				Value:             value,
				IsValueChangeable: util.Ptr(true),
			},
			nil,
			model.DeviceConfigurationKeyValueDescriptionDataType{
				KeyName: util.Ptr(model.DeviceConfigurationKeyNameTypeFailsafeDurationMinimum),
			},
		)
	}

	deviceConfigList := []model.DeviceConfigurationKeyValueDataType{{KeyId: util.Ptr(model.DeviceConfigurationKeyIdType(0))}, {KeyId: util.Ptr(model.DeviceConfigurationKeyIdType(1))}, {KeyId: util.Ptr(model.DeviceConfigurationKeyIdType(2))}}
	cmd := model.CmdType{DeviceConfigurationKeyValueListData: util.Ptr(model.DeviceConfigurationKeyValueListDataType{DeviceConfigurationKeyValueData: deviceConfigList})}
	msg := spineapi.Message{Cmd: cmd}
	pendingDeviceConfigs := map[model.MsgCounterType]*spineapi.Message{model.MsgCounterType(1): &msg}
	groupedConfigurations := GroupPendingDeviceConfigurations(pendingDeviceConfigs, s.localEntity)
	// For one of the KeyIds no corresponding device configuration exists, that element should thus be skipped
	expected := map[model.MsgCounterType][]ucapi.PendingDeviceConfiguration{
		model.MsgCounterType(1): {
			{Description: &failsafeLimitDesc},
			{Description: &failsafeDurationMinDesc},
		},
	}
	assert.Equal(s.T(), groupedConfigurations, expected)

}
