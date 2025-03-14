package convert

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/naceto/tempstation/internal/generated/db"
	"github.com/naceto/tempstation/internal/storage/models"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:enum:unknown @ignore
// goverter:extend ConvertTimeToTimestamptz
// goverter:extend ConvertTimestamptzToTime
type Convert interface {
	// goverter:useZeroValueOnPointerInconsistency
	CreateUserModelToDB(input *models.CreateUser) db.CreateUserParams
	UserModelFromDB(input db.User) *models.User
	UserModelsFromDB(input []db.User) []*models.User

	// goverter:useZeroValueOnPointerInconsistency
	CreateSensorModelToDB(input *models.CreateSensor) db.CreateSensorParams
	SensorModelFromDB(input db.Sensor) *models.Sensor
	SensorModelsFromDB(input []db.Sensor) []*models.Sensor

	// goverter:useZeroValueOnPointerInconsistency
	CreateSensorDataModelToDB(input *models.CreateSensorData) db.CreateSensorDataParams
	CreateSensorDataModelFromDB(input db.SensorDatum) *models.SensorData
}

// ConvertTimeToTimestamptz converts time.Time to pgtype.Timestamptz.
func ConvertTimeToTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t, Valid: true}
}

// ConvertTimestamptzToTime converts pgtype.Timestamptz to time.Time.
func ConvertTimestamptzToTime(tz pgtype.Timestamptz) time.Time {
	if tz.Valid {
		return tz.Time
	}

	return time.Time{}
}
