package configura

import (
	"os"
	"strconv"
)

// Bool takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Bool(key Variable[bool], fallback bool) bool {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vBool, err := strconv.ParseBool(vStr); err == nil {
			return vBool
		}
	}
	return fallback
}

// String takes an environment key, and a fallback value. Returns environment value if it isn't unset.
func String(key Variable[string], fallback string) string {
	if v, ok := os.LookupEnv(string(key)); ok {
		return v
	}
	return fallback
}

// Int takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int(key Variable[int], fallback int) int {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vInt, err := strconv.Atoi(vStr); err == nil {
			return vInt
		}
	}
	return fallback
}

// Int8 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int8(key Variable[int8], fallback int8) int8 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vInt, err := strconv.ParseInt(vStr, 10, 8); err == nil {
			return int8(vInt)
		}
	}
	return fallback
}

// Int16 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int16(key Variable[int16], fallback int16) int16 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vInt, err := strconv.ParseInt(vStr, 10, 16); err == nil {
			return int16(vInt)
		}
	}
	return fallback
}

// Int32 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int32(key Variable[int32], fallback int32) int32 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vInt, err := strconv.ParseInt(vStr, 10, 32); err == nil {
			return int32(vInt)
		}
	}
	return fallback
}

// Int64 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int64(key Variable[int64], fallback int64) int64 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vInt, err := strconv.ParseInt(vStr, 10, 64); err == nil {
			return vInt
		}
	}
	return fallback
}

// Uint takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint(key Variable[uint], fallback uint) uint {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vUint, err := strconv.ParseUint(vStr, 10, 0); err == nil { // 0 for uint
			return uint(vUint)
		}
	}
	return fallback
}

// Uint8 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint8(key Variable[uint8], fallback uint8) uint8 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vUint, err := strconv.ParseUint(vStr, 10, 8); err == nil {
			return uint8(vUint)
		}
	}
	return fallback
}

// Uint16 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint16(key Variable[uint16], fallback uint16) uint16 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vUint, err := strconv.ParseUint(vStr, 10, 16); err == nil {
			return uint16(vUint)
		}
	}
	return fallback
}

// Uint32 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint32(key Variable[uint32], fallback uint32) uint32 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vUint, err := strconv.ParseUint(vStr, 10, 32); err == nil {
			return uint32(vUint)
		}
	}
	return fallback
}

// Uint64 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint64(key Variable[uint64], fallback uint64) uint64 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vUint, err := strconv.ParseUint(vStr, 10, 64); err == nil {
			return vUint
		}
	}
	return fallback
}

// Uintptr takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uintptr(key Variable[uintptr], fallback uintptr) uintptr {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		// strconv.ParseUint is appropriate, bitSize 0 means infer from system (32 or 64)
		if vUint, err := strconv.ParseUint(vStr, 10, strconv.IntSize); err == nil {
			return uintptr(vUint)
		}
	}
	return fallback
}

// Bytes takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Bytes(key Variable[[]byte], fallback []byte) []byte {
	if v, ok := os.LookupEnv(string(key)); ok {
		return []byte(v)
	}
	return fallback
}

// Runes takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Runes(key Variable[[]rune], fallback []rune) []rune {
	if v, ok := os.LookupEnv(string(key)); ok {
		return []rune(v)
	}
	return fallback
}

// Float32 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Float32(key Variable[float32], fallback float32) float32 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vFloat, err := strconv.ParseFloat(vStr, 32); err == nil {
			return float32(vFloat)
		}
	}
	return fallback
}

// Float64 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Float64(key Variable[float64], fallback float64) float64 {
	if vStr, ok := os.LookupEnv(string(key)); ok {
		if vFloat, err := strconv.ParseFloat(vStr, 64); err == nil {
			return vFloat
		}
	}
	return fallback
}
