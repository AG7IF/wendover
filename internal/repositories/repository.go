package repositories

import (
	"github.com/google/uuid"

	"github.com/ag7if/wendover/pkg/auth"
	"github.com/ag7if/wendover/pkg/org"
)

type Repository interface {
	InsertActivity(activity org.Activity) (org.Activity, error)
	SelectActivities() ([]org.Activity, error)
	SelectActivity(key string) (org.Activity, error)
	UpdateActivity(key string, activity org.Activity) (org.Activity, error)
	DeleteActivity(key string) error

	InsertUser(user auth.User) (auth.User, error)
	SelectUsers() ([]auth.User, error)
	SelectUser(username string) (auth.User, error)
	UpdateUser(username string, user auth.User) (auth.User, error)
	DeleteUser(username string) error
	AddUserRole(username, activityKey string, userRole auth.Role) (auth.User, error)
	RemoveUserRole(username, activityKey string) (auth.User, error)

	InsertActivityUnits(activityId, superiorUnitId uuid.UUID, activityUnit []org.ActivityUnit) ([]org.ActivityUnit, error)
	SelectActivityUnitsByActivity(activity org.Activity) ([]org.ActivityUnit, error)
	SelectActivityUnit(id uuid.UUID) (org.ActivityUnit, error)
	UpdateActivityUnit(id uuid.UUID, activity org.ActivityUnit) (org.ActivityUnit, error)
	DeleteActivityUnit(id uuid.UUID) error
}
