package configura

import (
	"errors"
	"maps"
)

var ErrMissingVariable = errors.New("missing configuration variables")

type constraint interface {
	string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | []byte | []rune | float32 | float64 | bool
}

type Variable[T constraint] string

// Config is an interface that defines methods for accessing configuration variables of various types.
type Config interface {
	String(key Variable[string]) string
	Int(key Variable[int]) int
	Int8(key Variable[int8]) int8
	Int16(key Variable[int16]) int16
	Int32(key Variable[int32]) int32
	Int64(key Variable[int64]) int64
	Uint(key Variable[uint]) uint
	Uint8(key Variable[uint8]) uint8
	Uint16(key Variable[uint16]) uint16
	Uint32(key Variable[uint32]) uint32
	Uint64(key Variable[uint64]) uint64
	Uintptr(key Variable[uintptr]) uintptr
	Bytes(key Variable[[]byte]) []byte
	Runes(key Variable[[]rune]) []rune
	Float32(key Variable[float32]) float32
	Float64(key Variable[float64]) float64
	Bool(key Variable[bool]) bool
	ConfigurationKeysRegistered(keys ...any) error
}

// LoadEnvironment is a generic function that loads an environment variable into the provided configuration,
// using the specified key and fallback value. It uses type assertions to determine the type of the key
// and fallback value, and registers the variable in the appropriate map of the configuration struct.
func LoadEnvironment[T constraint](config *ConfigImpl, key Variable[T], fallback T) {
	switch any(key).(type) {
	case Variable[string]:
		config.RegString[any(key).(Variable[string])] = String(any(key).(Variable[string]), any(fallback).(string))
	case Variable[int]:
		config.RegInt[any(key).(Variable[int])] = Int(any(key).(Variable[int]), any(fallback).(int))
	case Variable[int8]:
		config.RegInt8[any(key).(Variable[int8])] = Int8(any(key).(Variable[int8]), any(fallback).(int8))
	case Variable[int16]:
		config.RegInt16[any(key).(Variable[int16])] = Int16(any(key).(Variable[int16]), any(fallback).(int16))
	case Variable[int32]:
		config.RegInt32[any(key).(Variable[int32])] = Int32(any(key).(Variable[int32]), any(fallback).(int32))
	case Variable[int64]:
		config.RegInt64[any(key).(Variable[int64])] = Int64(any(key).(Variable[int64]), any(fallback).(int64))
	case Variable[uint]:
		config.RegUint[any(key).(Variable[uint])] = Uint(any(key).(Variable[uint]), any(fallback).(uint))
	case Variable[uint8]:
		config.RegUint8[any(key).(Variable[uint8])] = Uint8(any(key).(Variable[uint8]), any(fallback).(uint8))
	case Variable[uint16]:
		config.RegUint16[any(key).(Variable[uint16])] = Uint16(any(key).(Variable[uint16]), any(fallback).(uint16))
	case Variable[uint32]:
		config.RegUint32[any(key).(Variable[uint32])] = Uint32(any(key).(Variable[uint32]), any(fallback).(uint32))
	case Variable[uint64]:
		config.RegUint64[any(key).(Variable[uint64])] = Uint64(any(key).(Variable[uint64]), any(fallback).(uint64))
	case Variable[uintptr]:
		config.RegUintptr[any(key).(Variable[uintptr])] = Uintptr(any(key).(Variable[uintptr]), any(fallback).(uintptr))
	case Variable[[]byte]:
		config.RegBytes[any(key).(Variable[[]byte])] = Bytes(any(key).(Variable[[]byte]), any(fallback).([]byte))
	case Variable[[]rune]:
		config.RegRunes[any(key).(Variable[[]rune])] = Runes(any(key).(Variable[[]rune]), any(fallback).([]rune))
	case Variable[float32]:
		config.RegFloat32[any(key).(Variable[float32])] = Float32(any(key).(Variable[float32]), any(fallback).(float32))
	case Variable[float64]:
		config.RegFloat64[any(key).(Variable[float64])] = Float64(any(key).(Variable[float64]), any(fallback).(float64))
	case Variable[bool]:
		config.RegBool[any(key).(Variable[bool])] = Bool(any(key).(Variable[bool]), any(fallback).(bool))
	}
}

// ConfigImpl is a concrete implementation of the Config interface, holding maps for each type of configuration
// variable. It provides methods to retrieve values for each type and checks if all required keys are registered.
type ConfigImpl struct {
	RegString  map[Variable[string]]string
	RegInt     map[Variable[int]]int
	RegInt8    map[Variable[int8]]int8
	RegInt16   map[Variable[int16]]int16
	RegInt32   map[Variable[int32]]int32
	RegInt64   map[Variable[int64]]int64
	RegUint    map[Variable[uint]]uint
	RegUint8   map[Variable[uint8]]uint8
	RegUint16  map[Variable[uint16]]uint16
	RegUint32  map[Variable[uint32]]uint32
	RegUint64  map[Variable[uint64]]uint64
	RegUintptr map[Variable[uintptr]]uintptr
	RegBytes   map[Variable[[]byte]][]byte
	RegRunes   map[Variable[[]rune]][]rune
	RegFloat32 map[Variable[float32]]float32
	RegFloat64 map[Variable[float64]]float64
	RegBool    map[Variable[bool]]bool
}

func NewConfigImpl() *ConfigImpl {
	return &ConfigImpl{
		RegString:  make(map[Variable[string]]string),
		RegInt:     make(map[Variable[int]]int),
		RegInt8:    make(map[Variable[int8]]int8),
		RegInt16:   make(map[Variable[int16]]int16),
		RegInt32:   make(map[Variable[int32]]int32),
		RegInt64:   make(map[Variable[int64]]int64),
		RegUint:    make(map[Variable[uint]]uint),
		RegUint8:   make(map[Variable[uint8]]uint8),
		RegUint16:  make(map[Variable[uint16]]uint16),
		RegUint32:  make(map[Variable[uint32]]uint32),
		RegUint64:  make(map[Variable[uint64]]uint64),
		RegUintptr: make(map[Variable[uintptr]]uintptr),
		RegBytes:   make(map[Variable[[]byte]][]byte),
		RegRunes:   make(map[Variable[[]rune]][]rune),
		RegFloat32: make(map[Variable[float32]]float32),
		RegFloat64: make(map[Variable[float64]]float64),
		RegBool:    make(map[Variable[bool]]bool),
	}
}

var _ Config = (*ConfigImpl)(nil)

func (c ConfigImpl) String(key Variable[string]) string {
	if value, exists := c.RegString[key]; exists {
		return value
	}
	return ""
}

func (c ConfigImpl) Int(key Variable[int]) int {
	if value, exists := c.RegInt[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Int8(key Variable[int8]) int8 {
	if value, exists := c.RegInt8[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Int16(key Variable[int16]) int16 {
	if value, exists := c.RegInt16[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Int32(key Variable[int32]) int32 {
	if value, exists := c.RegInt32[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Int64(key Variable[int64]) int64 {
	if value, exists := c.RegInt64[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Uint(key Variable[uint]) uint {
	if value, exists := c.RegUint[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Uint8(key Variable[uint8]) uint8 {
	if value, exists := c.RegUint8[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Uint16(key Variable[uint16]) uint16 {
	if value, exists := c.RegUint16[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Uint32(key Variable[uint32]) uint32 {
	if value, exists := c.RegUint32[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Uint64(key Variable[uint64]) uint64 {
	if value, exists := c.RegUint64[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Uintptr(key Variable[uintptr]) uintptr {
	if value, exists := c.RegUintptr[key]; exists {
		return value
	}
	return 0
}

func (c ConfigImpl) Bytes(key Variable[[]byte]) []byte {
	if value, exists := c.RegBytes[key]; exists {
		return value
	}
	return nil
}

func (c ConfigImpl) Runes(key Variable[[]rune]) []rune {
	if value, exists := c.RegRunes[key]; exists {
		return value
	}
	return nil
}

func (c ConfigImpl) Float32(key Variable[float32]) float32 {
	if value, exists := c.RegFloat32[key]; exists {
		return value
	}
	return 0.0
}

func (c ConfigImpl) Float64(key Variable[float64]) float64 {
	if value, exists := c.RegFloat64[key]; exists {
		return value
	}
	return 0.0
}

func (c ConfigImpl) Bool(key Variable[bool]) bool {
	if value, exists := c.RegBool[key]; exists {
		return value
	}
	return false
}

// missingVariableError is an error type that holds a list of missing configuration variable keys.
type missingVariableError struct {
	Keys []string
}

// Error implements the error interface for missingVariableError.
func (e missingVariableError) Error() string {
	return "missing configuration variables: " + formatKeys(e.Keys)
}

// Unwrap implements the Unwrap method for the error interface, allowing the error to be unwrapped to ErrMissingVariable.
func (e missingVariableError) Unwrap() error {
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

var _ error = (*missingVariableError)(nil)

// checkKey checks if the provided key exists in the configuration. It uses type assertion to determine the type of the
// key and checks the corresponding map in the configuration struct.
func (c ConfigImpl) checkKey(key any) (string, bool) {
	var exists bool
	var keyName string
	switch any(key).(type) {
	case Variable[string]:
		_, exists = c.RegString[key.(Variable[string])]
		keyName = string(key.(Variable[string]))
	case Variable[int]:
		_, exists = c.RegInt[key.(Variable[int])]
		keyName = string(key.(Variable[int]))
	case Variable[int8]:
		_, exists = c.RegInt8[key.(Variable[int8])]
		keyName = string(key.(Variable[int8]))
	case Variable[int16]:
		_, exists = c.RegInt16[key.(Variable[int16])]
		keyName = string(key.(Variable[int16]))
	case Variable[int32]:
		_, exists = c.RegInt32[key.(Variable[int32])]
		keyName = string(key.(Variable[int32]))
	case Variable[int64]:
		_, exists = c.RegInt64[key.(Variable[int64])]
		keyName = string(key.(Variable[int64]))
	case Variable[uint]:
		_, exists = c.RegUint[key.(Variable[uint])]
		keyName = string(key.(Variable[uint]))
	case Variable[uint8]:
		_, exists = c.RegUint8[key.(Variable[uint8])]
		keyName = string(key.(Variable[uint8]))
	case Variable[uint16]:
		_, exists = c.RegUint16[key.(Variable[uint16])]
		keyName = string(key.(Variable[uint16]))
	case Variable[uint32]:
		_, exists = c.RegUint32[key.(Variable[uint32])]
		keyName = string(key.(Variable[uint32]))
	case Variable[uint64]:
		_, exists = c.RegUint64[key.(Variable[uint64])]
		keyName = string(key.(Variable[uint64]))
	case Variable[uintptr]:
		_, exists = c.RegUintptr[key.(Variable[uintptr])]
		keyName = string(key.(Variable[uintptr]))
	case Variable[[]byte]:
		_, exists = c.RegBytes[key.(Variable[[]byte])]
		keyName = string(key.(Variable[[]byte]))
	case Variable[[]rune]:
		_, exists = c.RegRunes[key.(Variable[[]rune])]
		keyName = string(key.(Variable[[]rune]))
	case Variable[float32]:
		_, exists = c.RegFloat32[key.(Variable[float32])]
		keyName = string(key.(Variable[float32]))
	case Variable[float64]:
		_, exists = c.RegFloat64[key.(Variable[float64])]
		keyName = string(key.(Variable[float64]))
	case Variable[bool]:
		_, exists = c.RegBool[key.(Variable[bool])]
		keyName = string(key.(Variable[bool]))
	}

	return keyName, exists
}

// ConfigurationKeysRegistered checks if all provided keys are registered in the configuration. To ensure that the
// client of the package have taken all required keys into consideration when building the configuration object.
func (c ConfigImpl) ConfigurationKeysRegistered(keys ...any) error {
	var missingKeys []string
	for _, key := range keys {
		if keyName, ok := c.checkKey(key); !ok {
			missingKeys = append(missingKeys, keyName)
		}
	}

	if len(missingKeys) > 0 {
		return missingVariableError{Keys: missingKeys}
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
func Merge(cfgs ...Config) Config {
	merged := NewConfigImpl()

	for _, cfg := range cfgs {
		if strMap, ok := cfg.(*ConfigImpl); ok {
			maps.Copy(merged.RegString, strMap.RegString)
			maps.Copy(merged.RegInt, strMap.RegInt)
			maps.Copy(merged.RegInt8, strMap.RegInt8)
			maps.Copy(merged.RegInt16, strMap.RegInt16)
			maps.Copy(merged.RegInt32, strMap.RegInt32)
			maps.Copy(merged.RegInt64, strMap.RegInt64)
			maps.Copy(merged.RegUint, strMap.RegUint)
			maps.Copy(merged.RegUint8, strMap.RegUint8)
			maps.Copy(merged.RegUint16, strMap.RegUint16)
			maps.Copy(merged.RegUint32, strMap.RegUint32)
			maps.Copy(merged.RegUint64, strMap.RegUint64)
			maps.Copy(merged.RegUintptr, strMap.RegUintptr)
			maps.Copy(merged.RegBytes, strMap.RegBytes)
			maps.Copy(merged.RegRunes, strMap.RegRunes)
			maps.Copy(merged.RegFloat32, strMap.RegFloat32)
			maps.Copy(merged.RegFloat64, strMap.RegFloat64)
			maps.Copy(merged.RegBool, strMap.RegBool)
		} else {
			panic("unsupported config type")
		}
	}
	return merged
}
