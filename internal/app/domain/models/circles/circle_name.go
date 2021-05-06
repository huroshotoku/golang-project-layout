package circles

import (
	"fmt"
	"reflect"
)

type CircleName string

func NewName(value string) (*CircleName, error) {
	if value == "" {
		return nil, fmt.Errorf("ValueError: CircleNameがnullです。")
	}
	if len(value) < 3 {
		return nil, fmt.Errorf("ValueError: サークル名は3文字以上です。")
	}
	if len(value) > 20 {
		return nil, fmt.Errorf("ValueError: サークル名は20文字以下です。")
	}
	name := CircleName(value)
	return &name, nil
}

func (n *CircleName) Equals(other interface{}) (bool, error) {
	if (other == nil) || reflect.ValueOf(other).IsNil() {
		return false, nil
	}
	return *n == other, nil
}
