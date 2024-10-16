package phalanx

import (
	"strconv"
)

func Must[T int | string](v T, _ error) T {
	return v
}

func ValidatedPort[T uint32 | int | string](v T) (int, error) {
	switch value := any(v).(type) {
	case uint32:
		return int(value), nil
	case int:
		if value < 0 || value >= 65535 {
			return 0, ErrorInvalidPort(value)
		}
		return value, nil
	case string:
		v, err := strconv.Atoi(value)
		if err != nil {
			return 0, ErrorInvalidPort(value)
		}
		return v, nil
	}
	return 0, ErrorIllegalValue()
}

func NonNullValue[T int | string](v string, defaultValue T) T {
	if v == "" {
		return defaultValue
	}
	switch any(defaultValue).(type) {
	case int:
		v, err := strconv.Atoi(v)
		if err != nil {
			return any(0).(T)
		}
		return any(v).(T)
	}
	return any(v).(T)
}
