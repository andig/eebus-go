package server

import (
	"errors"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/features/internal"
	"github.com/enbility/eebus-go/util"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

type LoadControl struct {
	*Feature

	*internal.LoadControlCommon
}

func NewLoadControl(localEntity spineapi.EntityLocalInterface) (*LoadControl, error) {
	feature, err := NewFeature(model.FeatureTypeTypeLoadControl, localEntity)
	if err != nil {
		return nil, err
	}

	lc := &LoadControl{
		Feature:           feature,
		LoadControlCommon: internal.NewLocalLoadControl(feature.featureLocal),
	}

	return lc, nil
}

// Add a new description data set and return the limitId
//
// NOTE: the limitId may not be provided
//
// will return nil if the data set could not be added
func (l *LoadControl) AddLimitDescription(
	description model.LoadControlLimitDescriptionDataType,
) *model.LoadControlLimitIdType {
	if description.LimitId != nil {
		return nil
	}

	data, err := l.GetLimitDescriptionsForFilter(model.LoadControlLimitDescriptionDataType{})
	if err != nil {
		data = []model.LoadControlLimitDescriptionDataType{}
	}

	maxId := model.LoadControlLimitIdType(0)

	for _, item := range data {
		if item.LimitId != nil && *item.LimitId >= maxId {
			maxId = *item.LimitId + 1
		}
	}

	limitId := util.Ptr(maxId)
	description.LimitId = limitId

	partial := model.NewFilterTypePartial()
	datalist := &model.LoadControlLimitDescriptionListDataType{
		LoadControlLimitDescriptionData: []model.LoadControlLimitDescriptionDataType{description},
	}

	if err := l.featureLocal.UpdateData(model.FunctionTypeLoadControlLimitDescriptionListData, datalist, partial, nil); err != nil {
		return nil
	}

	return limitId
}

// Set or update data set for a limitId
// Elements provided in deleteElements will be removed from the data set before the update
//
// Will return an error if the data set could not be updated
func (l *LoadControl) UpdateLimitDataForId(
	data model.LoadControlLimitDataType,
	deleteElements *model.LoadControlLimitDataElementsType,
	limitId model.LoadControlLimitIdType,
) (resultErr error) {
	return l.UpdateLimitDataForFilter(data, deleteElements, model.LoadControlLimitDescriptionDataType{LimitId: &limitId})
}

// Set or update data set for a filter
// Elements provided in deleteElements will be removed from the data set before the update
//
// Will return an error if the data set could not be updated
func (l *LoadControl) UpdateLimitDataForFilter(
	data model.LoadControlLimitDataType,
	deleteElements *model.LoadControlLimitDataElementsType,
	filter model.LoadControlLimitDescriptionDataType,
) (resultErr error) {
	resultErr = api.ErrDataNotAvailable

	descriptions, err := l.GetLimitDescriptionsForFilter(filter)
	if err != nil || descriptions == nil || len(descriptions) != 1 {
		return
	}

	description := descriptions[0]
	data.LimitId = description.LimitId

	datalist := &model.LoadControlLimitListDataType{
		LoadControlLimitData: []model.LoadControlLimitDataType{data},
	}

	partial := model.NewFilterTypePartial()
	var delete *model.FilterType
	if deleteElements != nil {
		delete = &model.FilterType{
			LoadControlLimitListDataSelectors: &model.LoadControlLimitListDataSelectorsType{
				LimitId: description.LimitId,
			},
			LoadControlLimitDataElements: deleteElements,
		}
	}

	if err := l.featureLocal.UpdateData(model.FunctionTypeLoadControlLimitListData, datalist, partial, delete); err != nil {
		return errors.New(err.String())
	}

	return nil
}
