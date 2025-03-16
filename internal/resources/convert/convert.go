package convert

import (
	"time"

	api "github.com/naceto/tempstation/internal/generated/api/sensors"
	apiS "github.com/naceto/tempstation/internal/generated/api/sensors"
	apiU "github.com/naceto/tempstation/internal/generated/api/users"
	"github.com/naceto/tempstation/internal/storage/models"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:extend ListUsersAPIFromStorageExtend
// goverter:extend ListSensorsAPIFromStorage
// goverter:enum:unknown @ignore
// goverter:extend ConvertTimeToString
// goverter:extend ConvertTime
// goverter:extend ConvertTimePointer
// goverter:extend GetSensorDataAPIFromStorage
type Convert interface {
	// goverter:useZeroValueOnPointerInconsistency
	// goverter:map Body.Name Name
	// goverter:map Body.Email Email
	PostUserAPIToStorage(input *apiU.PostV1UsersRequestObject) *models.CreateUser

	// goverter:map . UserResponseJSONResponse
	PostUserAPIFromStorage(input *models.User) *apiU.PostV1Users200JSONResponse

	// goverter:map . UserResponseJSONResponse
	GetUserAPIFromStorage(input *models.User) *apiU.GetV1UsersId200JSONResponse

	// goverter:map ID Id
	UserAPIFromStorageEmbedded(input models.User) apiU.UserResponseJSONResponse

	// goverter:useZeroValueOnPointerInconsistency
	// goverter:map Params.Offset Offset
	// goverter:map Params.Limit Limit
	ListUsersAPIToStorage(input apiU.GetV1UsersRequestObject) *models.ListUsersParams

	ListUsersAPIFromStorage(input []*models.User) apiU.GetV1Users200JSONResponse

	// goverter:useZeroValueOnPointerInconsistency
	// goverter:map Params.Offset Offset
	// goverter:map Params.Limit Limit
	// goverter:map Params.Mac Mac
	ListSensorsAPIToStorage(input apiS.GetV1SensorsRequestObject) *models.ListSensorsParams
	ListSensorsAPIFromStorage(input []*models.Sensor) apiS.GetV1Sensors200JSONResponse

	// goverter:useZeroValueOnPointerInconsistency
	// goverter:map Body.Mac MacAddress
	// goverter:map Body.UserId UserID
	// goverter:map Body.Name Name
	// goverter:map Body.Type Type
	PostSensorAPIToStorage(input api.PostV1SensorsRequestObject) *models.CreateSensor

	// goverter:map . SensorResponseJSONResponse
	PostSensorAPIFromStorage(input *models.Sensor) *apiS.PostV1Sensors200JSONResponse

	// goverter:map ID Id
	// goverter:map MacAddress Mac
	// goverter:map UserID UserId
	SensorAPIFromStorageEmbedded(input models.Sensor) apiS.SensorResponseJSONResponse

	// goverter:map . SensorResponseJSONResponse
	GetSensorAPIFromStorage(input *models.Sensor) *apiS.GetV1SensorsId200JSONResponse

	// goverter:useZeroValueOnPointerInconsistency
	// goverter:map Body.SensorId SensorID
	// goverter:map Body.Temperature Temperature
	// goverter:map Body.Humidity Humidity
	PostSensorDataAPIToStorage(input apiS.PostV1SensorsIdDataRequestObject) *models.CreateSensorData

	// goverter:map . SensorDataResponseJSONResponse
	SensorDataAPIFromStorage(input *models.SensorData) *apiS.PostV1SensorsIdData200JSONResponse

	// goverter:map ID Id
	// goverter:map SensorID SensorId
	// goverter:map ReadingTime | ConvertTimeToString
	SensorDataAPIFromStorageEmbedded(input models.SensorData) apiS.SensorDataResponseJSONResponse

	// goverter:map Id SensorID
	// goverter:map Params.Start Start | ConvertTime
	// goverter:map Params.End End | ConvertTimePointer
	GetSensorDataAPIToStorage(input apiS.GetV1SensorsIdDataRequestObject) *models.GetSensorDataParams

	GetSensorDataAPIFromStorage(input []*models.SensorData) apiS.GetV1SensorsIdData200JSONResponse
}

func ListUsersAPIFromStorageExtend(input []*models.User) apiU.GetV1Users200JSONResponse {
	users := make([]apiU.User, 0, len(input))
	for _, u := range input {
		users = append(users, apiU.User{
			Id:    &u.ID,
			Name:  &u.Name,
			Email: &u.Email,
		})
	}

	return apiU.GetV1Users200JSONResponse{
		UsersResponseJSONResponse: apiU.UsersResponseJSONResponse{
			Users: users,
		},
	}
}

func ListSensorsAPIFromStorage(input []*models.Sensor) apiS.GetV1Sensors200JSONResponse {
	sensors := make([]apiS.Sensor, 0, len(input))
	for _, s := range input {
		sensors = append(sensors, apiS.Sensor{
			Id:     s.ID,
			Mac:    s.MacAddress,
			Name:   s.Name,
			Type:   apiS.SensorType(s.Type),
			UserId: s.UserID,
		})
	}

	return apiS.GetV1Sensors200JSONResponse{
		SensorsResponseJSONResponse: apiS.SensorsResponseJSONResponse{
			Sensors: sensors,
		},
	}
}

func ConvertTimeToString(input time.Time) string {
	return input.Format(time.RFC3339)
}

func ConvertTime(t time.Time) time.Time {
	return t
}

func ConvertTimePointer(t *time.Time) *time.Time {
	return t
}

func GetSensorDataAPIFromStorage(input []*models.SensorData) apiS.GetV1SensorsIdData200JSONResponse {
	data := make([]apiS.SensorData, 0, len(input))
	for _, sd := range input {
		data = append(data, apiS.SensorData{
			Id:          sd.ID,
			SensorId:    sd.SensorID,
			Temperature: sd.Temperature,
			Humidity:    sd.Humidity,
			ReadingTime: sd.ReadingTime.Format(time.RFC3339),
		})
	}

	return apiS.GetV1SensorsIdData200JSONResponse{
		SensorsDataResponseJSONResponse: apiS.SensorsDataResponseJSONResponse{
			Data: data,
		},
	}
}
