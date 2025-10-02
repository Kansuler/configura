package configura

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// --- Test Suite Definitions ---

type ConfigSuite struct {
	suite.Suite
	config *Config
}

type LoadSuite struct {
	suite.Suite
}

type ExistsSuite struct {
	suite.Suite
}

type FallbackSuite struct {
	suite.Suite
}

type MergeSuite struct {
	suite.Suite
}

type FormatKeysSuite struct {
	suite.Suite
}

type CheckKeySuite struct {
	suite.Suite
}

// --- Setup Methods ---

func (s *ConfigSuite) SetupTest() {
	s.config = New()
}

// --- Test Methods for ConfigSuite ---

func (s *ConfigSuite) TestString() {
	key := Variable[string]("TEST_STRING")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), "", s.config.String(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[string]]string{key: "hello"})
		assert.Equal(s.T(), "hello", s.config.String(key))
	})
}

func (s *ConfigSuite) TestInt() {
	key := Variable[int]("TEST_INT")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), 0, s.config.Int(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[int]]int{key: 123})
		assert.Equal(s.T(), 123, s.config.Int(key))
	})
}

func (s *ConfigSuite) TestInt8() {
	key := Variable[int8]("TEST_INT8")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), int8(0), s.config.Int8(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[int8]]int8{key: 12})
		assert.Equal(s.T(), int8(12), s.config.Int8(key))
	})
}

func (s *ConfigSuite) TestInt16() {
	key := Variable[int16]("TEST_INT16")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), int16(0), s.config.Int16(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[int16]]int16{key: 1234})
		assert.Equal(s.T(), int16(1234), s.config.Int16(key))
	})
}

func (s *ConfigSuite) TestInt32() {
	key := Variable[int32]("TEST_INT32")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), int32(0), s.config.Int32(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[int32]]int32{key: 123456})
		assert.Equal(s.T(), int32(123456), s.config.Int32(key))
	})
}

func (s *ConfigSuite) TestInt64() {
	key := Variable[int64]("TEST_INT64")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), int64(0), s.config.Int64(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[int64]]int64{key: 1234567890})
		assert.Equal(s.T(), int64(1234567890), s.config.Int64(key))
	})
}

func (s *ConfigSuite) TestUint() {
	key := Variable[uint]("TEST_UINT")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint(0), s.config.Uint(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[uint]]uint{key: 123})
		assert.Equal(s.T(), uint(123), s.config.Uint(key))
	})
}

func (s *ConfigSuite) TestUint8() {
	key := Variable[uint8]("TEST_UINT8")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint8(0), s.config.Uint8(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[uint8]]uint8{key: 12})
		assert.Equal(s.T(), uint8(12), s.config.Uint8(key))
	})
}

func (s *ConfigSuite) TestUint16() {
	key := Variable[uint16]("TEST_UINT16")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint16(0), s.config.Uint16(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[uint16]]uint16{key: 1234})
		assert.Equal(s.T(), uint16(1234), s.config.Uint16(key))
	})
}

func (s *ConfigSuite) TestUint32() {
	key := Variable[uint32]("TEST_UINT32")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint32(0), s.config.Uint32(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[uint32]]uint32{key: 123456})
		assert.Equal(s.T(), uint32(123456), s.config.Uint32(key))
	})
}

func (s *ConfigSuite) TestUint64() {
	key := Variable[uint64]("TEST_UINT64")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint64(0), s.config.Uint64(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[uint64]]uint64{key: 1234567890})
		assert.Equal(s.T(), uint64(1234567890), s.config.Uint64(key))
	})
}

func (s *ConfigSuite) TestUintptr() {
	key := Variable[uintptr]("TEST_UINTPTR")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uintptr(0), s.config.Uintptr(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[uintptr]]uintptr{key: 0xdeadbeef})
		assert.Equal(s.T(), uintptr(0xdeadbeef), s.config.Uintptr(key))
	})
}

func (s *ConfigSuite) TestBytes() {
	key := Variable[[]byte]("TEST_BYTES")
	s.Run("KeyNotExists", func() {
		assert.Nil(s.T(), s.config.Bytes(key))
	})
	s.Run("KeyExists", func() {
		val := []byte("hello")
		Write(s.config, map[Variable[[]byte]][]byte{key: val})
		assert.Equal(s.T(), val, s.config.Bytes(key))
	})
}

func (s *ConfigSuite) TestRunes() {
	key := Variable[[]rune]("TEST_RUNES")
	s.Run("KeyNotExists", func() {
		assert.Nil(s.T(), s.config.Runes(key))
	})
	s.Run("KeyExists", func() {
		val := []rune("hello😊")
		Write(s.config, map[Variable[[]rune]][]rune{key: val})
		assert.Equal(s.T(), val, s.config.Runes(key))
	})
}

func (s *ConfigSuite) TestFloat32() {
	key := Variable[float32]("TEST_FLOAT32")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), float32(0.0), s.config.Float32(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[float32]]float32{key: 3.14})
		assert.Equal(s.T(), float32(3.14), s.config.Float32(key))
	})
}

func (s *ConfigSuite) TestFloat64() {
	key := Variable[float64]("TEST_FLOAT64")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), float64(0.0), s.config.Float64(key))
	})
	s.Run("KeyExists", func() {
		Write(s.config, map[Variable[float64]]float64{key: 3.14159})
		assert.Equal(s.T(), float64(3.14159), s.config.Float64(key))
	})
}

func (s *ConfigSuite) TestBool() {
	key := Variable[bool]("TEST_BOOL")
	s.Run("KeyNotExists", func() {
		assert.False(s.T(), s.config.Bool(key))
	})
	s.Run("KeyExistsTrue", func() {
		Write(s.config, map[Variable[bool]]bool{key: true})
		assert.True(s.T(), s.config.Bool(key))
	})
	s.Run("KeyExistsFalse", func() {
		falseKey := Variable[bool]("TEST_BOOL_FALSE_VAL")
		Write(s.config, map[Variable[bool]]bool{falseKey: false})
		assert.False(s.T(), s.config.Bool(falseKey))
	})
}

// --- Test Methods for LoadSuite ---

func (s *LoadSuite) setEnvVar(key string, value string) {
	err := os.Setenv(key, value)
	s.Require().NoError(err)
	s.T().Cleanup(func() {
		os.Unsetenv(key)
	})
}

func (s *LoadSuite) unsetEnvVar(key string) {
	err := os.Unsetenv(key)
	// Allow unset to fail if var doesn't exist, as that's fine for test setup.
	if err != nil && !errors.Is(err, os.ErrNotExist) && !strings.Contains(err.Error(), "unsetenv: EINVAL: Invalid argument") { // macOS specific error for empty key
		s.Require().NoError(err) // Fail for other errors
	}
}

func (s *LoadSuite) TestLoadString() {
	cfg := New()
	key := Variable[string]("ENV_STR")
	fallback := "fb_str"
	envVal := "env_str_val"

	s.Run("EnvVarSet", func() {
		s.setEnvVar(string(key), envVal)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envVal, cfg.String(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))                     // Ensure it's unset for this specific sub-test
		Load(cfg, key, fallback)                       // Use fresh config or reset
		assert.Equal(s.T(), fallback, cfg.String(key)) // This line would fail if cfg is not reset or key re-added

		// Corrected approach for EnvVarNotSet
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.String(key))
	})
}

func (s *LoadSuite) TestLoadInt() {
	cfg := New()
	key := Variable[int]("ENV_INT")
	fallback := 100
	envValStr := "200"
	envValInt := 200

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValInt, cfg.Int(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int(key))
	})
}

func (s *LoadSuite) TestLoadInt8() {
	cfg := New()
	key := Variable[int8]("ENV_INT8")
	fallback := int8(10)
	envValStr := "20"
	envValInt8 := int8(20)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValInt8, cfg.Int8(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int8(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int8")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int8(key))
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "129") // Out of bounds for int8
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int8(key))
	})
}

func (s *LoadSuite) TestLoadInt16() {
	cfg := New()
	key := Variable[int16]("ENV_INT16")
	fallback := int16(1000)
	envValStr := "2000"
	envValInt16 := int16(2000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValInt16, cfg.Int16(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int16(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int16")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int16(key))
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "32768") // Out of bounds for int16
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int16(key))
	})
}

func (s *LoadSuite) TestLoadInt32() {
	cfg := New()
	key := Variable[int32]("ENV_INT32")
	fallback := int32(100000)
	envValStr := "200000"
	envValInt32 := int32(200000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValInt32, cfg.Int32(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int32(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int32")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int32(key))
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "2147483648") // Out of bounds for int32
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int32(key))
	})
}

func (s *LoadSuite) TestLoadInt64() {
	cfg := New()
	key := Variable[int64]("ENV_INT64")
	fallback := int64(1000000000)
	envValStr := "2000000000"
	envValInt64 := int64(2000000000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValInt64, cfg.Int64(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int64(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int64")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int64(key))
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "9223372036854775808") // Out of bounds for int64
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Int64(key))
	})
}

func (s *LoadSuite) TestLoadUint() {
	cfg := New()
	key := Variable[uint]("ENV_UINT")
	fallback := uint(100)
	envValStr := "200"
	envValUint := uint(200)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValUint, cfg.Uint(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint(key))
	})
	s.Run("EnvVarSetNegative", func() {
		s.setEnvVar(string(key), "-1") // Negative, invalid for uint
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint(key))
	})
}

func (s *LoadSuite) TestLoadUint8() {
	cfg := New()
	key := Variable[uint8]("ENV_UINT8")
	fallback := uint8(10)
	envValStr := "20"
	envValUint8 := uint8(20)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValUint8, cfg.Uint8(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint8(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint8")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint8(key))
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "256") // Out of bounds for uint8
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint8(key))
	})
}

func (s *LoadSuite) TestLoadUint16() {
	cfg := New()
	key := Variable[uint16]("ENV_UINT16")
	fallback := uint16(1000)
	envValStr := "2000"
	envValUint16 := uint16(2000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValUint16, cfg.Uint16(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint16(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint16")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint16(key))
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "65536") // Out of bounds for uint16
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint16(key))
	})
}

func (s *LoadSuite) TestLoadUint32() {
	cfg := New()
	key := Variable[uint32]("ENV_UINT32")
	fallback := uint32(100000)
	envValStr := "200000"
	envValUint32 := uint32(200000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValUint32, cfg.Uint32(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint32(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint32")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint32(key))
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "4294967296") // Out of bounds for uint32
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint32(key))
	})
}

func (s *LoadSuite) TestLoadUint64() {
	cfg := New()
	key := Variable[uint64]("ENV_UINT64")
	fallback := uint64(1000000000)
	envValStr := "2000000000"
	envValUint64 := uint64(2000000000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValUint64, cfg.Uint64(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint64(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint64")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint64(key))
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "18446744073709551616") // Out of bounds for uint64
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uint64(key))
	})
}

func (s *LoadSuite) TestLoadUintptr() {
	cfg := New()
	key := Variable[uintptr]("ENV_UINTPTR")
	fallback := uintptr(0x1000)
	envValStr := "0x2000" // Using hex for variety
	var envValUintptr uintptr
	_, err := fmt.Sscan(envValStr, &envValUintptr) // Parse hex string to uintptr
	s.Require().NoError(err)

	s.Run("EnvVarSetValid", func() {
		// strconv.ParseUint expects decimal unless base is specified (e.g. 0 for auto-detect prefix, 16 for hex)
		// For simplicity, let's use decimal strings for env vars if uintptr helper uses ParseUint(str, 10, ...)
		// Or adjust the helper stub to handle "0x"
		decimalEnvValStr := strconv.FormatUint(uint64(envValUintptr), 10)
		s.setEnvVar(string(key), decimalEnvValStr)
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envValUintptr, cfg.Uintptr(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uintptr(key))
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uintptr")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Uintptr(key))
	})
}

func (s *LoadSuite) TestLoadFloat32() {
	cfg := New()
	key := Variable[float32]("ENV_FLOAT32")
	fallback := float32(1.23)
	envValStr := "4.56"
	envValFloat32 := float32(4.56)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.InDelta(s.T(), envValFloat32, cfg.Float32(key), 0.0001)
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.InDelta(s.T(), fallback, freshCfg.Float32(key), 0.0001)
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-float")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.InDelta(s.T(), fallback, freshCfg.Float32(key), 0.0001)
	})
}

func (s *LoadSuite) TestLoadFloat64() {
	cfg := New()
	key := Variable[float64]("ENV_FLOAT64")
	fallback := float64(1.23456)
	envValStr := "7.89012"
	envValFloat64 := float64(7.89012)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		Load(cfg, key, fallback)
		assert.InDelta(s.T(), envValFloat64, cfg.Float64(key), 0.0000001)
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.InDelta(s.T(), fallback, freshCfg.Float64(key), 0.0000001)
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-float64")
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.InDelta(s.T(), fallback, freshCfg.Float64(key), 0.0000001)
	})
}

func (s *LoadSuite) TestLoadBool() {
	key := Variable[bool]("ENV_BOOL")

	testCases := []struct {
		name        string
		envValue    *string
		fallback    bool
		expectedReg bool
	}{
		{"EnvTrueFallbackF", func(s string) *string { return &s }("true"), false, true},
		{"EnvFalseFallbackT", func(s string) *string { return &s }("false"), true, false},
		{"Env1FallbackF", func(s string) *string { return &s }("1"), false, true},
		{"Env0FallbackT", func(s string) *string { return &s }("0"), true, false},
		{"EnvTFallbackF", func(s string) *string { return &s }("t"), false, true},
		{"EnvFFallbackT", func(s string) *string { return &s }("f"), true, false},
		{"EnvInvalidFallbackF", func(s string) *string { return &s }("invalid"), false, false},
		{"EnvInvalidFallbackT", func(s string) *string { return &s }("invalid"), true, true},
		{"EnvNotSetFallbackF", nil, false, false},
		{"EnvNotSetFallbackT", nil, true, true},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			currentCfg := New()
			if tc.envValue != nil {
				s.setEnvVar(string(key), *tc.envValue)
			} else {
				s.unsetEnvVar(string(key))
			}
			Load(currentCfg, key, tc.fallback)
			assert.Equal(s.T(), tc.expectedReg, currentCfg.Bool(key), "Mismatch in registered bool value")
		})
	}
}

func (s *LoadSuite) TestLoadBytes() {
	cfg := New()
	key := Variable[[]byte]("ENV_BYTES")
	fallback := []byte("fb_bytes")
	envVal := []byte("env_bytes_val")

	s.Run("EnvVarSet", func() {
		s.setEnvVar(string(key), string(envVal))
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envVal, cfg.Bytes(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Bytes(key))
	})
}

func (s *LoadSuite) TestLoadRunes() {
	cfg := New()
	key := Variable[[]rune]("ENV_RUNES")
	fallback := []rune("fb_runes")
	envValStr := "env_runes_😊"
	envVal := []rune(envValStr)

	s.Run("EnvVarSet", func() {
		s.setEnvVar(string(key), string(envValStr)) // Store string in env
		Load(cfg, key, fallback)
		assert.Equal(s.T(), envVal, cfg.Runes(key))
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := New()
		Load(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.Runes(key))
	})
}

// --- Test Methods for ExistsSuite ---

func (s *ExistsSuite) TestExists() {
	cfg := New()

	strKey1 := Variable[string]("STR_KEY_1")
	strKey2Missing := Variable[string]("STR_KEY_2_MISSING")
	intKey1 := Variable[int]("INT_KEY_1")
	floatKeyMissing := Variable[float32]("FLOAT_KEY_MISSING")

	Write(cfg, map[Variable[string]]string{strKey1: "val1"})
	Write(cfg, map[Variable[int]]int{intKey1: 100})

	s.Run("AllCheckedKeysExist", func() {
		err := cfg.Exists(strKey1, intKey1)
		assert.NoError(s.T(), err)
	})

	s.Run("SomeKeysMissing", func() {
		err := cfg.Exists(strKey1, strKey2Missing, intKey1, floatKeyMissing)
		s.Require().Error(err)

		var missingErr MissingVariableError
		s.Require().True(errors.As(err, &missingErr), "Error should be of type MissingVariableError")

		s.Require().ErrorIs(err, ErrMissingVariable, "Error should unwrap to ErrMissingVariable")

		assert.ElementsMatch(s.T(), []string{string(strKey2Missing), string(floatKeyMissing)}, missingErr.Keys)
		assert.Contains(s.T(), err.Error(), "missing configuration variables:")
		assert.Contains(s.T(), err.Error(), string(strKey2Missing))
		assert.Contains(s.T(), err.Error(), string(floatKeyMissing))
	})

	s.Run("AllCheckedKeysMissing", func() {
		missingStr := Variable[string]("COMPLETELY_MISSING_S")
		missingInt := Variable[int]("COMPLETELY_MISSING_I")
		err := cfg.Exists(missingStr, missingInt)
		s.Require().Error(err)
		var missingErr MissingVariableError
		s.Require().True(errors.As(err, &missingErr))
		assert.ElementsMatch(s.T(), []string{string(missingStr), string(missingInt)}, missingErr.Keys)
		s.Require().ErrorIs(err, ErrMissingVariable)
	})

	s.Run("NoKeysToCheck", func() {
		err := cfg.Exists()
		assert.NoError(s.T(), err)
	})

	s.Run("ErrorTypeAndUnwrap", func() {
		err := cfg.Exists(Variable[string]("ANY_MISSING_KEY"))
		s.Require().Error(err)

		_, ok := err.(MissingVariableError)
		assert.True(s.T(), ok, "Error should be MissingVariableError type")

		assert.ErrorIs(s.T(), err, ErrMissingVariable, "Error should unwrap to ErrMissingVariable via errors.Is")

		unwrappedErr := errors.Unwrap(err)
		assert.Equal(s.T(), ErrMissingVariable, unwrappedErr, "Unwrapped error should be exactly ErrMissingVariable")
	})
}

// --- Test Methods for FormatKeysSuite ---

func (s *FormatKeysSuite) TestFormatKeys() {
	testCases := []struct {
		name     string
		keys     []string
		expected string
	}{
		{"Empty", []string{}, "none"},
		{"Single", []string{"KEY1"}, "KEY1"},
		{"Multiple", []string{"KEY1", "KEY2", "KEY3"}, "KEY1, KEY2, KEY3"},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			assert.Equal(s.T(), tc.expected, formatKeys(tc.keys))
		})
	}
}

func (s *FallbackSuite) TestFallbackString() {
	tests := []struct {
		name     string
		value    string
		fallback string
		expected string
	}{
		{
			name:     "Value is not empty",
			value:    "actualValue",
			fallback: "fallbackValue",
			expected: "actualValue",
		},
		{
			name:     "Value is empty",
			value:    "",
			fallback: "fallbackValue",
			expected: "fallbackValue",
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			result := Fallback(tc.value, tc.fallback)
			s.Equal(tc.expected, result)
		})
	}
}

func (s *FallbackSuite) TestFallbackInt() {
	tests := []struct {
		name     string
		value    int
		fallback int
		expected int
	}{
		{
			name:     "Value is not zero",
			value:    10,
			fallback: 20,
			expected: 10,
		},
		{
			name:     "Value is zero",
			value:    0,
			fallback: 20,
			expected: 20,
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			result := Fallback(tc.value, tc.fallback)
			s.Equal(tc.expected, result)
		})
	}
}

func (s *FallbackSuite) TestFallbackBool() {
	tests := []struct {
		name     string
		value    bool
		fallback bool
		expected bool
	}{
		{
			name:     "Value is true",
			value:    true,
			fallback: false,
			expected: true,
		},
		{
			name:     "Value is false, fallback true",
			value:    false,
			fallback: true,
			expected: true,
		},
		{
			name:     "Value is false, fallback false",
			value:    false,
			fallback: false,
			expected: false,
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			result := Fallback(tc.value, tc.fallback)
			s.Equal(tc.expected, result)
		})
	}
}

// --- Test Methods for MergeSuite ---

func (s *MergeSuite) TestMergeEmpty() {
	cfg := Merge()
	s.Require().NotNil(cfg, "Merged config should not be nil")

	s.Empty(cfg.regString, "RegString should be empty")
	s.Empty(cfg.regInt, "RegInt should be empty")
	s.Empty(cfg.regInt8, "RegInt8 should be empty")
	s.Empty(cfg.regInt16, "RegInt16 should be empty")
	s.Empty(cfg.regInt32, "RegInt32 should be empty")
	s.Empty(cfg.regInt64, "RegInt64 should be empty")
	s.Empty(cfg.regUint, "RegUint should be empty")
	s.Empty(cfg.regUint8, "RegUint8 should be empty")
	s.Empty(cfg.regUint16, "RegUint16 should be empty")
	s.Empty(cfg.regUint32, "RegUint32 should be empty")
	s.Empty(cfg.regUint64, "RegUint64 should be empty")
	s.Empty(cfg.regUintptr, "RegUintptr should be empty")
	s.Empty(cfg.regBytes, "RegBytes should be empty")
	s.Empty(cfg.regRunes, "RegRunes should be empty")
	s.Empty(cfg.regFloat32, "RegFloat32 should be empty")
	s.Empty(cfg.regFloat64, "RegFloat64 should be empty")
	s.Empty(cfg.regBool, "RegBool should be empty")
}

func (s *MergeSuite) TestMergeSingle() {
	cfg1 := New()
	keyStr := Variable[string]("TEST_STR")
	Load(cfg1, keyStr, "value1")

	keyInt := Variable[int]("TEST_INT")
	Load(cfg1, keyInt, 123)

	mergedCfg := Merge(cfg1)
	s.Require().NotNil(mergedCfg, "Merged config should not be nil")

	s.Equal("value1", mergedCfg.String(keyStr))
	s.Equal(123, mergedCfg.Int(keyInt))
	s.Len(mergedCfg.regString, 1)
	s.Len(mergedCfg.regInt, 1)
}

func (s *MergeSuite) TestMergeTwoDistinct() {
	cfg1 := New()
	keyStr1 := Variable[string]("STR_KEY_1")
	Load(cfg1, keyStr1, "value1")

	cfg2 := New()
	keyInt1 := Variable[int]("INT_KEY_1")
	Load(cfg2, keyInt1, 100)

	mergedCfg := Merge(cfg1, cfg2)
	s.Require().NotNil(mergedCfg, "Merged config should not be nil")

	s.Equal("value1", mergedCfg.String(keyStr1))
	s.Equal(100, mergedCfg.Int(keyInt1))

	s.Len(mergedCfg.regString, 1, "RegString should have 1 entry")
	s.Equal("value1", mergedCfg.regString[keyStr1])
	s.Len(mergedCfg.regInt, 1, "RegInt should have 1 entry")
	s.Equal(100, mergedCfg.regInt[keyInt1])
}

func (s *MergeSuite) TestMergeTwoOverride() {
	cfg1 := New()
	keyStr := Variable[string]("OVERRIDE_STR")
	Load(cfg1, keyStr, "original_value")
	keyInt := Variable[int]("SHARED_INT")
	Load(cfg1, keyInt, 111) // This key is only in cfg1

	cfg2 := New()
	Load(cfg2, keyStr, "overridden_value") // Override
	keyBool := Variable[bool]("NEW_BOOL")
	Load(cfg2, keyBool, true) // This key is only in cfg2

	mergedCfg := Merge(cfg1, cfg2)
	s.Require().NotNil(mergedCfg, "Merged config should not be nil")

	s.Equal("overridden_value", mergedCfg.String(keyStr)) // Overridden
	s.Equal(111, mergedCfg.Int(keyInt))                   // From cfg1
	s.True(mergedCfg.Bool(keyBool))                       // From cfg2

	s.Len(mergedCfg.regString, 1)
	s.Equal("overridden_value", mergedCfg.regString[keyStr])
	s.Len(mergedCfg.regInt, 1)
	s.Equal(111, mergedCfg.regInt[keyInt])
	s.Len(mergedCfg.regBool, 1)
	s.True(mergedCfg.regBool[keyBool])
}

func (s *MergeSuite) TestMergeMultiple() {
	cfg1 := New()
	keyStr1 := Variable[string]("S1")
	Load(cfg1, keyStr1, "val_s1_cfg1") // In cfg1
	keyShared := Variable[string]("SHARED_KEY")
	Load(cfg1, keyShared, "shared_cfg1") // In cfg1, overridden by cfg2, then by cfg3

	cfg2 := New()
	keyInt1 := Variable[int]("I1")
	Load(cfg2, keyInt1, 222)             // In cfg2
	Load(cfg2, keyShared, "shared_cfg2") // Override from cfg1

	cfg3 := New()
	keyBool1 := Variable[bool]("B1")
	Load(cfg3, keyBool1, true)           // In cfg3
	Load(cfg3, keyShared, "shared_cfg3") // Override from cfg2

	mergedCfg := Merge(cfg1, cfg2, cfg3)
	s.Require().NotNil(mergedCfg)

	s.Equal("val_s1_cfg1", mergedCfg.String(keyStr1))
	s.Equal("shared_cfg3", mergedCfg.String(keyShared)) // Final override from cfg3
	s.Equal(222, mergedCfg.Int(keyInt1))
	s.True(mergedCfg.Bool(keyBool1))

	s.Len(mergedCfg.regString, 2) // S1, SHARED_KEY
	s.Len(mergedCfg.regInt, 1)    // I1
	s.Len(mergedCfg.regBool, 1)   // B1
}

func (s *MergeSuite) TestMergeAllTypes() {
	cfg1 := New()
	cfg2 := New()

	// Define keys and values
	kStr := Variable[string]("TYPE_STR")
	vStr1, vStr2 := "string_val1", "string_val2"
	kInt := Variable[int]("TYPE_INT")
	vInt1, vInt2 := 10, 20
	kInt8 := Variable[int8]("TYPE_INT8")
	vInt8_1, vInt8_2 := int8(1), int8(2)
	kInt16 := Variable[int16]("TYPE_INT16")
	vInt16_1, vInt16_2 := int16(100), int16(200)
	kInt32 := Variable[int32]("TYPE_INT32")
	vInt32_1, vInt32_2 := int32(1000), int32(2000)
	kInt64 := Variable[int64]("TYPE_INT64")
	vInt64_1, vInt64_2 := int64(10000), int64(20000)
	kUint := Variable[uint]("TYPE_UINT")
	vUint1 := uint(1)
	kUint8 := Variable[uint8]("TYPE_UINT8")
	vUint8_1 := uint8(10)
	kUint16 := Variable[uint16]("TYPE_UINT16")
	vUint16_1 := uint16(100)
	kUint32 := Variable[uint32]("TYPE_UINT32")
	vUint32_1 := uint32(1000)
	kUint64 := Variable[uint64]("TYPE_UINT64")
	vUint64_1 := uint64(10000)
	kUintptr := Variable[uintptr]("TYPE_UINTPTR")
	vUintptr1 := uintptr(0x1)
	kBytes := Variable[[]byte]("TYPE_BYTES")
	vBytes2 := []byte("bytes2")
	kRunes := Variable[[]rune]("TYPE_RUNES")
	vRunes2 := []rune("runes2")
	kFloat32 := Variable[float32]("TYPE_FLOAT32")
	vFloat32_2 := float32(2.5)
	kFloat64 := Variable[float64]("TYPE_FLOAT64")
	vFloat64_2 := 20.5
	kBool := Variable[bool]("TYPE_BOOL")
	vBool1, vBool2 := false, true // Bool1 is in cfg1, Bool2 will override

	// Load into cfg1 (these will be overridden or kept if not in cfg2)
	Load(cfg1, kStr, vStr1)
	Load(cfg1, kInt, vInt1)
	Load(cfg1, kInt8, vInt8_1)
	Load(cfg1, kInt16, vInt16_1)
	Load(cfg1, kInt32, vInt32_1)
	Load(cfg1, kInt64, vInt64_1)
	// Uint types only in cfg1
	Load(cfg1, kUint, vUint1)
	Load(cfg1, kUint8, vUint8_1)
	Load(cfg1, kUint16, vUint16_1)
	Load(cfg1, kUint32, vUint32_1)
	Load(cfg1, kUint64, vUint64_1)
	Load(cfg1, kUintptr, vUintptr1)
	Load(cfg1, kBool, vBool1) // Will be overridden

	// Load into cfg2 (these will override cfg1 or be new)
	Load(cfg2, kStr, vStr2)      // Override
	Load(cfg2, kInt, vInt2)      // Override
	Load(cfg2, kInt8, vInt8_2)   // Override
	Load(cfg2, kInt16, vInt16_2) // Override
	Load(cfg2, kInt32, vInt32_2) // Override
	Load(cfg2, kInt64, vInt64_2) // Override
	// Slice and float types only in cfg2
	Load(cfg2, kBytes, vBytes2)
	Load(cfg2, kRunes, vRunes2)
	Load(cfg2, kFloat32, vFloat32_2)
	Load(cfg2, kFloat64, vFloat64_2)
	Load(cfg2, kBool, vBool2) // Override

	mergedCfg := Merge(cfg1, cfg2)
	s.Require().NotNil(mergedCfg)

	// Assertions for overridden values (from cfg2)
	s.Equal(vStr2, mergedCfg.String(kStr))
	s.Equal(vInt2, mergedCfg.Int(kInt))
	s.Equal(vInt8_2, mergedCfg.Int8(kInt8))
	s.Equal(vInt16_2, mergedCfg.Int16(kInt16))
	s.Equal(vInt32_2, mergedCfg.Int32(kInt32))
	s.Equal(vInt64_2, mergedCfg.Int64(kInt64))
	s.Equal(vBool2, mergedCfg.Bool(kBool))

	// Assertions for values only in cfg1
	s.Equal(vUint1, mergedCfg.Uint(kUint))
	s.Equal(vUint8_1, mergedCfg.Uint8(kUint8))
	s.Equal(vUint16_1, mergedCfg.Uint16(kUint16))
	s.Equal(vUint32_1, mergedCfg.Uint32(kUint32))
	s.Equal(vUint64_1, mergedCfg.Uint64(kUint64))
	s.Equal(vUintptr1, mergedCfg.Uintptr(kUintptr))

	// Assertions for values only in cfg2
	s.Equal(vBytes2, mergedCfg.Bytes(kBytes))
	s.Equal(vRunes2, mergedCfg.Runes(kRunes))
	s.Equal(vFloat32_2, mergedCfg.Float32(kFloat32))
	s.Equal(vFloat64_2, mergedCfg.Float64(kFloat64))

	s.Len(mergedCfg.regString, 1)
	s.Len(mergedCfg.regInt, 1)
	s.Len(mergedCfg.regInt8, 1)
	s.Len(mergedCfg.regInt16, 1)
	s.Len(mergedCfg.regInt32, 1)
	s.Len(mergedCfg.regInt64, 1)
	s.Len(mergedCfg.regUint, 1)
	s.Len(mergedCfg.regUint8, 1)
	s.Len(mergedCfg.regUint16, 1)
	s.Len(mergedCfg.regUint32, 1)
	s.Len(mergedCfg.regUint64, 1)
	s.Len(mergedCfg.regUintptr, 1)
	s.Len(mergedCfg.regBytes, 1)
	s.Len(mergedCfg.regRunes, 1)
	s.Len(mergedCfg.regFloat32, 1)
	s.Len(mergedCfg.regFloat64, 1)
	s.Len(mergedCfg.regBool, 1)
}

// --- Test Methods for CheckKeySuite ---

func (s *CheckKeySuite) TestCheckKey() {
	cfg := New()

	strKey := Variable[string]("MY_STRING")
	intKey := Variable[int]("MY_INT")
	boolKey := Variable[bool]("MY_BOOL")
	float32Key := Variable[float32]("MY_FLOAT32")

	missingStrKey := Variable[string]("MISSING_STRING")
	missingIntKey := Variable[int]("MISSING_INT")
	uintptrKey := Variable[uintptr]("MY_UINTPTR_UNINIT_MAP_SCENARIO")

	cfg.regString[strKey] = "value"
	cfg.regInt[intKey] = 123
	cfg.regBool[boolKey] = true
	cfg.regFloat32[float32Key] = 3.14

	s.Run("ExistingKeys", func() {
		name, exists := cfg.checkKey(strKey)
		assert.True(s.T(), exists)
		assert.Equal(s.T(), string(strKey), name)

		name, exists = cfg.checkKey(intKey)
		assert.True(s.T(), exists)
		assert.Equal(s.T(), string(intKey), name)

		name, exists = cfg.checkKey(boolKey)
		assert.True(s.T(), exists)
		assert.Equal(s.T(), string(boolKey), name)

		name, exists = cfg.checkKey(float32Key)
		assert.True(s.T(), exists)
		assert.Equal(s.T(), string(float32Key), name)
	})

	s.Run("MissingKeys", func() {
		name, exists := cfg.checkKey(missingStrKey)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(missingStrKey), name)

		name, exists = cfg.checkKey(missingIntKey)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(missingIntKey), name)

		name, exists = cfg.checkKey(uintptrKey)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(uintptrKey), name)
	})

	s.Run("DifferentKeyTypeSameName", func() {
		diffTypeSameNameKey := Variable[string]("MY_INT")
		name, exists := cfg.checkKey(diffTypeSameNameKey)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(diffTypeSameNameKey), name)

		diffTypeSameNameKey2 := Variable[int]("MY_STRING")
		name, exists = cfg.checkKey(diffTypeSameNameKey2)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(diffTypeSameNameKey2), name)
	})
}

// --- Main Test Runner ---

func TestConfiguraSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
	suite.Run(t, new(LoadSuite))
	suite.Run(t, new(FormatKeysSuite))
	suite.Run(t, new(CheckKeySuite))
	suite.Run(t, new(ExistsSuite))
	suite.Run(t, new(FallbackSuite))
	suite.Run(t, new(MergeSuite))
}
