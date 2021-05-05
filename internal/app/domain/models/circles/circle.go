package circles

import (
	"fmt"

	"github.com/huroshotoku/golang-project-layout/internal/app/domain/models/users"
)

type Circle struct {
	id      CircleId // idの公開はアリ
	name    CircleName
	owner   users.UserId
	members []users.UserId
}

func NewCircle(id CircleId, name CircleName, owner users.UserId, members []users.UserId) (*Circle, error) {
	if id == "" {
		return nil, fmt.Errorf("ValueError: CircleIdがnullです。")
	}
	if name == "" {
		return nil, fmt.Errorf("ValueError: CircleNameがnullです。")
	}
	if owner == "" {
		return nil, fmt.Errorf("ValueError: ownerがnullです。")
	}
	circle := Circle{
		id:      id,
		name:    name,
		owner:   owner,
		members: members,
	}
	return &circle, nil
}

func (c *Circle) Join(memberId users.UserId) error {
	if memberId == "" {
		return fmt.Errorf("ArgumentNullError: memberId")
	}
	if c.IsFull() {
		return fmt.Errorf("CircleFullError: %s", c.id)
	}
	c.members = append(c.members, memberId)
	return nil
}

func (c Circle) IsFull() bool {
	return c.CountMembers() >= 30
}

func (c Circle) CountMembers() int {
	return len(c.members) + 1 // ownerユーザ分を足している
}