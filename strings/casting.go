package strings

import (
	"strconv"
)

func StringToBool(s string, defaultValue bool) bool {
	result, err := strconv.ParseBool(s)
	if err != nil {
		return defaultValue
	}
	return bool(result)
}

func StringToUint64(s string, defaultValue uint64) uint64 {
	result, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return defaultValue
	}
	return uint64(result)
}

func StringToUint32(s string, defaultValue uint32) uint32 {
	result, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return defaultValue
	}
	return uint32(result)
}

func StringToUint(s string, defaultValue uint) uint {
	result, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return defaultValue
	}
	return uint(result)
}

func StringToInt64(s string, defaultValue int64) int64 {
	result, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return defaultValue
	}
	return int64(result)
}

func StringToInt32(s string, defaultValue int32) int32 {
	result, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return defaultValue
	}
	return int32(result)
}

func StringToInt(s string, defaultValue int) int {
	result, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return defaultValue
	}
	return int(result)
}

func StringToFloat32(s string, defaultValue float32) float32 {
	result, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return defaultValue
	}
	return float32(result)
}

func StringToFloat64(s string, defaultValue float64) float64 {
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}
	return float64(result)
}
