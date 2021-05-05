package circles

import (
	"github.com/huroshotoku/golang-project-layout/internal/app/domain/models/users"
)

type ICircleFactory interface {
	Create(name CircleName, owner users.UserId) (*Circle, error)
}
