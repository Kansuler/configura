package configura

import (
	"errors"
	"maps"
	"sync"
)

var ErrMissingVariable = errors.New("missing configuration variables")

type constraint interface {
	string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | []byte | []rune | float32 | float64 | bool
}

type Variable[T constraint] string

// Write is a generic function that writes configuration values to the provided configuration struct.
// It uses type assertions to determine the type of the values and writes them to the appropriate map in the
// configuration struct.
func Write[T constraint](cfg *Config, values map[Variable[T]]T) error {
	if cfg == nil {
		return errors.New("Config cannot be nil")
	}

	cfg.rwLock.Lock()
	defer cfg.rwLock.Unlock()
	switch v := any(values).(type) {
	case map[Variable[string]]string:
		maps.Copy(cfg.regString, v)
	case map[Variable[int]]int:
		maps.Copy(cfg.regInt, v)
	case map[Variable[int8]]int8:
		maps.Copy(cfg.regInt8, v)
	case map[Variable[int16]]int16:
		maps.Copy(cfg.regInt16, v)
	case map[Variable[int32]]int32:
		maps.Copy(cfg.regInt32, v)
	case map[Variable[int64]]int64:
		maps.Copy(cfg.regInt64, v)
	case map[Variable[uint]]uint:
		maps.Copy(cfg.regUint, v)
	case map[Variable[uint8]]uint8:
		maps.Copy(cfg.regUint8, v)
	case map[Variable[uint16]]uint16:
		maps.Copy(cfg.regUint16, v)
	case map[Variable[uint32]]uint32:
		maps.Copy(cfg.regUint32, v)
	case map[Variable[uint64]]uint64:
		maps.Copy(cfg.regUint64, v)
	case map[Variable[uintptr]]uintptr:
		maps.Copy(cfg.regUintptr, v)
	case map[Variable[[]byte]][]byte:
		maps.Copy(cfg.regBytes, v)
	case map[Variable[[]rune]][]rune:
		maps.Copy(cfg.regRunes, v)
	case map[Variable[float32]]float32:
		maps.Copy(cfg.regFloat32, v)
	case map[Variable[float64]]float64:
		maps.Copy(cfg.regFloat64, v)
	case map[Variable[bool]]bool:
		maps.Copy(cfg.regBool, v)
	default:
		return errors.New("unsupported values type")
	}

	return nil
}

// Load is a generic function that loads an environment variable into the provided configuration,
// using the specified key and fallback value. It uses type assertions to determine the type of the key
// and fallback value, and registers the variable in the appropriate map of the configuration struct.
func Load[T constraint](cfg *Config, key Variable[T], fallback T) {
	cfg.rwLock.Lock()
	defer cfg.rwLock.Unlock()
	switch any(key).(type) {
	case Variable[string]:
		cfg.regString[any(key).(Variable[string])] = String(any(key).(Variable[string]), any(fallback).(string))
	case Variable[int]:
		cfg.regInt[any(key).(Variable[int])] = Int(any(key).(Variable[int]), any(fallback).(int))
	case Variable[int8]:
		cfg.regInt8[any(key).(Variable[int8])] = Int8(any(key).(Variable[int8]), any(fallback).(int8))
	case Variable[int16]:
		cfg.regInt16[any(key).(Variable[int16])] = Int16(any(key).(Variable[int16]), any(fallback).(int16))
	case Variable[int32]:
		cfg.regInt32[any(key).(Variable[int32])] = Int32(any(key).(Variable[int32]), any(fallback).(int32))
	case Variable[int64]:
		cfg.regInt64[any(key).(Variable[int64])] = Int64(any(key).(Variable[int64]), any(fallback).(int64))
	case Variable[uint]:
		cfg.regUint[any(key).(Variable[uint])] = Uint(any(key).(Variable[uint]), any(fallback).(uint))
	case Variable[uint8]:
		cfg.regUint8[any(key).(Variable[uint8])] = Uint8(any(key).(Variable[uint8]), any(fallback).(uint8))
	case Variable[uint16]:
		cfg.regUint16[any(key).(Variable[uint16])] = Uint16(any(key).(Variable[uint16]), any(fallback).(uint16))
	case Variable[uint32]:
		cfg.regUint32[any(key).(Variable[uint32])] = Uint32(any(key).(Variable[uint32]), any(fallback).(uint32))
	case Variable[uint64]:
		cfg.regUint64[any(key).(Variable[uint64])] = Uint64(any(key).(Variable[uint64]), any(fallback).(uint64))
	case Variable[uintptr]:
		cfg.regUintptr[any(key).(Variable[uintptr])] = Uintptr(any(key).(Variable[uintptr]), any(fallback).(uintptr))
	case Variable[[]byte]:
		cfg.regBytes[any(key).(Variable[[]byte])] = Bytes(any(key).(Variable[[]byte]), any(fallback).([]byte))
	case Variable[[]rune]:
		cfg.regRunes[any(key).(Variable[[]rune])] = Runes(any(key).(Variable[[]rune]), any(fallback).([]rune))
	case Variable[float32]:
		cfg.regFloat32[any(key).(Variable[float32])] = Float32(any(key).(Variable[float32]), any(fallback).(float32))
	case Variable[float64]:
		cfg.regFloat64[any(key).(Variable[float64])] = Float64(any(key).(Variable[float64]), any(fallback).(float64))
	case Variable[bool]:
		cfg.regBool[any(key).(Variable[bool])] = Bool(any(key).(Variable[bool]), any(fallback).(bool))
	}
}

// config is a concrete implementation of the Config interface, holding maps for each type of configuration
// variable. It provides methods to retrieve values for each type and checks if all required keys are registered.
type Config struct {
	rwLock     sync.RWMutex
	regString  map[Variable[string]]string
	regInt     map[Variable[int]]int
	regInt8    map[Variable[int8]]int8
	regInt16   map[Variable[int16]]int16
	regInt32   map[Variable[int32]]int32
	regInt64   map[Variable[int64]]int64
	regUint    map[Variable[uint]]uint
	regUint8   map[Variable[uint8]]uint8
	regUint16  map[Variable[uint16]]uint16
	regUint32  map[Variable[uint32]]uint32
	regUint64  map[Variable[uint64]]uint64
	regUintptr map[Variable[uintptr]]uintptr
	regBytes   map[Variable[[]byte]][]byte
	regRunes   map[Variable[[]rune]][]rune
	regFloat32 map[Variable[float32]]float32
	regFloat64 map[Variable[float64]]float64
	regBool    map[Variable[bool]]bool
}

func New() *Config {
	return &Config{
		regString:  make(map[Variable[string]]string),
		regInt:     make(map[Variable[int]]int),
		regInt8:    make(map[Variable[int8]]int8),
		regInt16:   make(map[Variable[int16]]int16),
		regInt32:   make(map[Variable[int32]]int32),
		regInt64:   make(map[Variable[int64]]int64),
		regUint:    make(map[Variable[uint]]uint),
		regUint8:   make(map[Variable[uint8]]uint8),
		regUint16:  make(map[Variable[uint16]]uint16),
		regUint32:  make(map[Variable[uint32]]uint32),
		regUint64:  make(map[Variable[uint64]]uint64),
		regUintptr: make(map[Variable[uintptr]]uintptr),
		regBytes:   make(map[Variable[[]byte]][]byte),
		regRunes:   make(map[Variable[[]rune]][]rune),
		regFloat32: make(map[Variable[float32]]float32),
		regFloat64: make(map[Variable[float64]]float64),
		regBool:    make(map[Variable[bool]]bool),
	}
}

func (c *Config) String(key Variable[string]) string {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regString[key]; exists {
		return value
	}
	return ""
}

func (c *Config) Int(key Variable[int]) int {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regInt[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Int8(key Variable[int8]) int8 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regInt8[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Int16(key Variable[int16]) int16 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regInt16[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Int32(key Variable[int32]) int32 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regInt32[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Int64(key Variable[int64]) int64 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regInt64[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Uint(key Variable[uint]) uint {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regUint[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Uint8(key Variable[uint8]) uint8 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regUint8[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Uint16(key Variable[uint16]) uint16 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regUint16[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Uint32(key Variable[uint32]) uint32 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regUint32[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Uint64(key Variable[uint64]) uint64 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regUint64[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Uintptr(key Variable[uintptr]) uintptr {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regUintptr[key]; exists {
		return value
	}
	return 0
}

func (c *Config) Bytes(key Variable[[]byte]) []byte {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regBytes[key]; exists {
		return value
	}
	return nil
}

func (c *Config) Runes(key Variable[[]rune]) []rune {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regRunes[key]; exists {
		return value
	}
	return nil
}

func (c *Config) Float32(key Variable[float32]) float32 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regFloat32[key]; exists {
		return value
	}
	return 0.0
}

func (c *Config) Float64(key Variable[float64]) float64 {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regFloat64[key]; exists {
		return value
	}
	return 0.0
}

func (c *Config) Bool(key Variable[bool]) bool {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, exists := c.regBool[key]; exists {
		return value
	}
	return false
}

// MissingVariableError is an error type that holds a list of missing configuration variable keys.
type MissingVariableError struct {
	Keys []string
}

// Error implements the error interface for missingVariableError.
func (e MissingVariableError) Error() string {
	return "missing configuration variables: " + formatKeys(e.Keys)
}

// Unwrap implements the Unwrap method for the error interface, allowing the error to be unwrapped to ErrMissingVariable.
func (e MissingVariableError) Unwrap() error {
	return ErrMissingVariable
}

// formatKeys formats the keys into a string for error messages. If no keys are provided, it returns "none".
func formatKeys(keys []string) string {
	if len(keys) == 0 {
		return "none"
	}
	result := ""
	for i, key := range keys {
		if i > 0 {
			result += ", "
		}
		result += string(key)
	}
	return result
}

var _ error = (*MissingVariableError)(nil)

// checkKey checks if the provided key exists in the configuration. It uses type assertion to determine the type of the
// key and checks the corresponding map in the configuration struct.
func (c *Config) checkKey(key any) (string, bool) {
	var exists bool
	var keyName string
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	switch k := key.(type) {
	case Variable[string]:
		_, exists = c.regString[k]
		keyName = string(k)
	case Variable[int]:
		_, exists = c.regInt[k]
		keyName = string(k)
	case Variable[int8]:
		_, exists = c.regInt8[k]
		keyName = string(k)
	case Variable[int16]:
		_, exists = c.regInt16[k]
		keyName = string(k)
	case Variable[int32]:
		_, exists = c.regInt32[k]
		keyName = string(k)
	case Variable[int64]:
		_, exists = c.regInt64[k]
		keyName = string(k)
	case Variable[uint]:
		_, exists = c.regUint[k]
		keyName = string(k)
	case Variable[uint8]:
		_, exists = c.regUint8[k]
		keyName = string(k)
	case Variable[uint16]:
		_, exists = c.regUint16[k]
		keyName = string(k)
	case Variable[uint32]:
		_, exists = c.regUint32[k]
		keyName = string(k)
	case Variable[uint64]:
		_, exists = c.regUint64[k]
		keyName = string(k)
	case Variable[uintptr]:
		_, exists = c.regUintptr[k]
		keyName = string(k)
	case Variable[[]byte]:
		_, exists = c.regBytes[k]
		keyName = string(k)
	case Variable[[]rune]:
		_, exists = c.regRunes[k]
		keyName = string(k)
	case Variable[float32]:
		_, exists = c.regFloat32[k]
		keyName = string(k)
	case Variable[float64]:
		_, exists = c.regFloat64[k]
		keyName = string(k)
	case Variable[bool]:
		_, exists = c.regBool[k]
		keyName = string(k)
	}

	return keyName, exists
}

// Exists checks if all provided keys are registered in the configuration. To ensure that the
// client of the package have taken all required keys into consideration when building the configuration object.
func (c *Config) Exists(keys ...any) error {
	var missingKeys []string
	for _, key := range keys {
		if keyName, ok := c.checkKey(key); !ok {
			missingKeys = append(missingKeys, keyName)
		}
	}

	if len(missingKeys) > 0 {
		return MissingVariableError{Keys: missingKeys}
	}

	return nil
}

// Fallback is a helper function that returns the fallback value if the provided value is empty.
// Only works on comparable types, which includes basic types like int, string, bool, etc.
func Fallback[T comparable](value T, fallback T) T {
	var emptyValue T
	if value == emptyValue {
		return fallback
	}
	return value
}

// Merge combines multiple Config instances into a single Config instance.
// To ensure a consistent view of the source configurations, it locks all
// configuration types for reading during the merge operation.
func Merge(cfgs ...*Config) *Config {
	merged := New()
	merged.rwLock.Lock()
	defer merged.rwLock.Unlock()

	for _, cfg := range cfgs {
		cfg.rwLock.RLock()
		maps.Copy(merged.regString, cfg.regString)
		maps.Copy(merged.regInt, cfg.regInt)
		maps.Copy(merged.regInt8, cfg.regInt8)
		maps.Copy(merged.regInt16, cfg.regInt16)
		maps.Copy(merged.regInt32, cfg.regInt32)
		maps.Copy(merged.regInt64, cfg.regInt64)
		maps.Copy(merged.regUint, cfg.regUint)
		maps.Copy(merged.regUint8, cfg.regUint8)
		maps.Copy(merged.regUint16, cfg.regUint16)
		maps.Copy(merged.regUint32, cfg.regUint32)
		maps.Copy(merged.regUint64, cfg.regUint64)
		maps.Copy(merged.regUintptr, cfg.regUintptr)
		maps.Copy(merged.regBytes, cfg.regBytes)
		maps.Copy(merged.regRunes, cfg.regRunes)
		maps.Copy(merged.regFloat32, cfg.regFloat32)
		maps.Copy(merged.regFloat64, cfg.regFloat64)
		maps.Copy(merged.regBool, cfg.regBool)
		cfg.rwLock.RUnlock()
	}
	return merged
}
