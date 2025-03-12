package convert

import (
	"github.com/naceto/tempstation/internal/generated/db"
	"github.com/naceto/tempstation/internal/storage/models"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:enum:unknown @ignore
type Convert interface {
	// goverter:useZeroValueOnPointerInconsistency
	CreateUserModelToDB(input *models.CreateUser) db.CreateUserParams
	UserModelFromDB(input db.User) *models.User
	UserModelsFromDB(input []db.User) []*models.User

	// goverter:useZeroValueOnPointerInconsistency
	CreateSensorModelToDB(input *models.CreateSensor) db.CreateSensorParams
	SensorModelFromDB(input db.Sensor) *models.Sensor
	SensorModelsFromDB(input []db.Sensor) []*models.Sensor
}
