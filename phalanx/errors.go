package phalanx

import (
	"errors"
	"fmt"
)

func ErrorInvalidPort[T string | int](v T) error {
	if fmt.Sprint(v) == "" {
		return errors.New("port is empty")
	}
	return fmt.Errorf("invalid port: %v", v)
}

func ErrorIllegalValue() error {
	return errors.New("illegal value error")
}

func ErrorRequired(text string) error {
	return fmt.Errorf("%s is required", text)
}

func ErrorUnsupport(text string) error {
	if text == "" {
		return errors.New("unsupport")
	}
	return fmt.Errorf("unsupport %s", text)
}
