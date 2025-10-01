package configura_test

import (
	"os"
	"testing"

	"github.com/Kansuler/configura"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Bool
	_ = os.Setenv("valid_bool_1", "true")
	_ = os.Setenv("valid_bool_2", "false")
	_ = os.Setenv("invalid_bool_1", "")
	_ = os.Setenv("invalid_bool_2", "123")
	_ = os.Setenv("invalid_bool_3", "test")

	// Int
	_ = os.Setenv("valid_int_1", "2")
	_ = os.Setenv("invalid_int_1", "")
	_ = os.Setenv("invalid_int_2", "test")

	// Int8
	_ = os.Setenv("valid_int8_1", "2")
	_ = os.Setenv("invalid_int8_1", "")
	_ = os.Setenv("invalid_int8_2", "test")
	_ = os.Setenv("invalid_int8_3", "300")

	// Int16
	_ = os.Setenv("valid_int16_1", "2")
	_ = os.Setenv("invalid_int16_1", "")
	_ = os.Setenv("invalid_int16_2", "test")
	_ = os.Setenv("invalid_int16_3", "32768")

	// Int32
	_ = os.Setenv("valid_int32_1", "2")
	_ = os.Setenv("invalid_int32_1", "")
	_ = os.Setenv("invalid_int32_2", "test")
	_ = os.Setenv("invalid_int32_3", "2147483648")

	// Int64
	_ = os.Setenv("valid_int64_1", "2")
	_ = os.Setenv("invalid_int64_1", "")
	_ = os.Setenv("invalid_int64_2", "test")
	_ = os.Setenv("invalid_int64_3", "9223372036854775808")

	// Uint
	_ = os.Setenv("valid_uint_1", "2")
	_ = os.Setenv("invalid_uint_1", "")
	_ = os.Setenv("invalid_uint_2", "test")
	_ = os.Setenv("invalid_uint_3", "-1")
	_ = os.Setenv("invalid_uint_4", "18446744073709551616")

	// Uint8
	_ = os.Setenv("valid_uint8_1", "2")
	_ = os.Setenv("invalid_uint8_1", "")
	_ = os.Setenv("invalid_uint8_2", "test")
	_ = os.Setenv("invalid_uint8_3", "-1")
	_ = os.Setenv("invalid_uint8_4", "256")

	// Uint16
	_ = os.Setenv("valid_uint16_1", "2")
	_ = os.Setenv("invalid_uint16_1", "")
	_ = os.Setenv("invalid_uint16_2", "test")
	_ = os.Setenv("invalid_uint16_3", "-1")
	_ = os.Setenv("invalid_uint16_4", "65536")

	// Uint32
	_ = os.Setenv("valid_uint32_1", "2")
	_ = os.Setenv("invalid_uint32_1", "")
	_ = os.Setenv("invalid_uint32_2", "test")
	_ = os.Setenv("invalid_uint32_3", "-1")
	_ = os.Setenv("invalid_uint32_4", "4294967296")

	// Uint64
	_ = os.Setenv("valid_uint64_1", "2")
	_ = os.Setenv("invalid_uint64_1", "")
	_ = os.Setenv("invalid_uint64_2", "test")
	_ = os.Setenv("invalid_uint64_3", "-1")
	_ = os.Setenv("invalid_uint64_4", "18446744073709551616")

	// Uintptr
	_ = os.Setenv("valid_uintptr_1", "2")
	_ = os.Setenv("invalid_uintptr_1", "")
	_ = os.Setenv("invalid_uintptr_2", "test")
	_ = os.Setenv("invalid_uintptr_3", "-1")
	_ = os.Setenv("invalid_uintptr_4", "18446744073709551616")

	// Bytes
	_ = os.Setenv("valid_bytes_1", "2")
	_ = os.Setenv("valid_bytes_2", "testing")
	_ = os.Setenv("invalid_bytes_1", "foo")

	// String
	_ = os.Setenv("valid_string_1", "2")
	_ = os.Setenv("valid_string_2", "testing")
	_ = os.Setenv("invalid_string_1", "")

	// Runes
	_ = os.Setenv("valid_runes_1", "2")
	_ = os.Setenv("valid_runes_2", "testing")
	_ = os.Setenv("invalid_runes_1", "foo")

	// Float32
	_ = os.Setenv("valid_float32_1", "2")
	_ = os.Setenv("valid_float32_2", "12.134")
	_ = os.Setenv("valid_float32_3", "-12.642322")
	_ = os.Setenv("valid_float32_4", "9223372036854775808")
	_ = os.Setenv("invalid_float32_1", "")
	_ = os.Setenv("invalid_float32_2", "test")

	// Float64
	_ = os.Setenv("valid_float64_1", "2")
	_ = os.Setenv("valid_float64_2", "12.134")
	_ = os.Setenv("valid_float64_3", "-12.642322")
	_ = os.Setenv("valid_float64_4", "9223372036854775808")
	_ = os.Setenv("invalid_float64_1", "")
	_ = os.Setenv("invalid_float64_2", "test")
}

func TestBool(t *testing.T) {
	valid1 := configura.Bool("valid_bool_1", false)
	assert.True(t, valid1)
	valid2 := configura.Bool("valid_bool_2", true)
	assert.False(t, valid2)
	invalid1 := configura.Bool("invalid_bool_1", true)
	assert.True(t, invalid1)
	invalid2 := configura.Bool("invalid_bool_2", true)
	assert.True(t, invalid2)
	invalid3 := configura.Bool("invalid_bool_3", true)
	assert.True(t, invalid3)
}

func TestInt(t *testing.T) {
	valid1 := configura.Int("valid_int_1", 0)
	assert.Equal(t, 2, valid1)
	invalid1 := configura.Int("invalid_int_1", 1)
	assert.Equal(t, 1, invalid1)
	invalid2 := configura.Int("invalid_int_2", 300)
	assert.Equal(t, 300, invalid2)
}

func TestInt8(t *testing.T) {
	valid1 := configura.Int8("valid_int8_1", 0)
	assert.Equal(t, int8(2), valid1)
	invalid1 := configura.Int8("invalid_int8_1", 1)
	assert.Equal(t, int8(1), invalid1)
	invalid2 := configura.Int8("invalid_int8_2", 127)
	assert.Equal(t, int8(127), invalid2)
	invalid3 := configura.Int8("invalid_int8_3", 0)
	assert.Equal(t, int8(0), invalid3)
}

func TestInt16(t *testing.T) {
	valid1 := configura.Int16("valid_int16_1", 0)
	assert.Equal(t, int16(2), valid1)
	invalid1 := configura.Int16("invalid_int16_1", 1)
	assert.Equal(t, int16(1), invalid1)
	invalid2 := configura.Int16("invalid_int16_2", 55)
	assert.Equal(t, int16(55), invalid2)
	invalid3 := configura.Int16("invalid_int16_3", 127)
	assert.Equal(t, int16(127), invalid3)
}

func TestInt32(t *testing.T) {
	valid1 := configura.Int32("valid_int32_1", 0)
	assert.Equal(t, int32(2), valid1)
	invalid1 := configura.Int32("invalid_int32_1", 1)
	assert.Equal(t, int32(1), invalid1)
	invalid2 := configura.Int32("invalid_int32_2", 55)
	assert.Equal(t, int32(55), invalid2)
	invalid3 := configura.Int32("invalid_int32_3", 127)
	assert.Equal(t, int32(127), invalid3)
}

func TestInt64(t *testing.T) {
	valid1 := configura.Int64("valid_int64_1", 0)
	assert.Equal(t, int64(2), valid1)
	invalid1 := configura.Int64("invalid_int64_1", 1)
	assert.Equal(t, int64(1), invalid1)
	invalid2 := configura.Int64("invalid_int64_2", 55)
	assert.Equal(t, int64(55), invalid2)
	invalid3 := configura.Int64("invalid_int64_3", 127)
	assert.Equal(t, int64(127), invalid3)
}

func TestUInt(t *testing.T) {
	valid1 := configura.Uint("valid_uint_1", 0)
	assert.Equal(t, uint(2), valid1)
	invalid1 := configura.Uint("invalid_uint_1", 1)
	assert.Equal(t, uint(1), invalid1)
	invalid2 := configura.Uint("invalid_uint_2", 55)
	assert.Equal(t, uint(55), invalid2)
	invalid3 := configura.Uint("invalid_uint_3", 127)
	assert.Equal(t, uint(127), invalid3)
	invalid4 := configura.Uint("invalid_uint_4", 22)
	assert.Equal(t, uint(22), invalid4)
}

func TestUInt8(t *testing.T) {
	valid1 := configura.Uint8("valid_uint8_1", 0)
	assert.Equal(t, uint8(2), valid1)
	invalid1 := configura.Uint8("invalid_uint8_1", 1)
	assert.Equal(t, uint8(1), invalid1)
	invalid2 := configura.Uint8("invalid_uint8_2", 55)
	assert.Equal(t, uint8(55), invalid2)
	invalid3 := configura.Uint8("invalid_uint8_3", 127)
	assert.Equal(t, uint8(127), invalid3)
	invalid4 := configura.Uint8("invalid_uint8_4", 22)
	assert.Equal(t, uint8(22), invalid4)
}

func TestUInt16(t *testing.T) {
	valid1 := configura.Uint16("valid_uint16_1", 0)
	assert.Equal(t, uint16(2), valid1)
	invalid1 := configura.Uint16("invalid_uint16_1", 1)
	assert.Equal(t, uint16(1), invalid1)
	invalid2 := configura.Uint16("invalid_uint16_2", 55)
	assert.Equal(t, uint16(55), invalid2)
	invalid3 := configura.Uint16("invalid_uint16_3", 127)
	assert.Equal(t, uint16(127), invalid3)
	invalid4 := configura.Uint16("invalid_uint16_4", 22)
	assert.Equal(t, uint16(22), invalid4)
}

func TestUInt32(t *testing.T) {
	valid1 := configura.Uint32("valid_uint32_1", 0)
	assert.Equal(t, uint32(2), valid1)
	invalid1 := configura.Uint32("invalid_uint32_1", 1)
	assert.Equal(t, uint32(1), invalid1)
	invalid2 := configura.Uint32("invalid_uint32_2", 55)
	assert.Equal(t, uint32(55), invalid2)
	invalid3 := configura.Uint32("invalid_uint32_3", 127)
	assert.Equal(t, uint32(127), invalid3)
	invalid4 := configura.Uint32("invalid_uint32_4", 22)
	assert.Equal(t, uint32(22), invalid4)
}

func TestUInt64(t *testing.T) {
	valid1 := configura.Uint64("valid_uint64_1", 0)
	assert.Equal(t, uint64(2), valid1)
	invalid1 := configura.Uint64("invalid_uint64_1", 1)
	assert.Equal(t, uint64(1), invalid1)
	invalid2 := configura.Uint64("invalid_uint64_2", 55)
	assert.Equal(t, uint64(55), invalid2)
	invalid3 := configura.Uint64("invalid_uint64_3", 127)
	assert.Equal(t, uint64(127), invalid3)
	invalid4 := configura.Uint64("invalid_uint64_4", 22)
	assert.Equal(t, uint64(22), invalid4)
}

func TestUIntptr(t *testing.T) {
	valid1 := configura.Uintptr("valid_uintptr_1", 0)
	assert.Equal(t, uintptr(2), valid1)
	invalid1 := configura.Uintptr("invalid_uintptr_1", 1)
	assert.Equal(t, uintptr(1), invalid1)
	invalid2 := configura.Uintptr("invalid_uintptr_2", 55)
	assert.Equal(t, uintptr(55), invalid2)
	invalid3 := configura.Uintptr("invalid_uintptr_3", 127)
	assert.Equal(t, uintptr(127), invalid3)
	invalid4 := configura.Uintptr("invalid_uintptr_4", 22)
	assert.Equal(t, uintptr(22), invalid4)
}

func TestBytes(t *testing.T) {
	valid1 := configura.Bytes("valid_bytes_1", []byte{'f', 'o', 'o'})
	assert.Equal(t, []byte{'2'}, valid1)
	valid2 := configura.Bytes("valid_bytes_2", []byte{'f', 'o', 'o'})
	assert.Equal(t, []byte{'t', 'e', 's', 't', 'i', 'n', 'g'}, valid2)
	invalid1 := configura.Bytes("invalid_bytes_1", []byte{'f', 'o', 'o'})
	assert.Equal(t, []byte{'f', 'o', 'o'}, invalid1)
}

func TestString(t *testing.T) {
	valid1 := configura.String("valid_bytes_1", "foo")
	assert.Equal(t, "2", valid1)
	valid2 := configura.String("valid_bytes_2", "foo")
	assert.Equal(t, "testing", valid2)
	invalid1 := configura.String("invalid_bytes_1", "foo")
	assert.Equal(t, "foo", invalid1)
}

func TestRunes(t *testing.T) {
	valid1 := configura.Runes("valid_runes_1", []rune("foo"))
	assert.Equal(t, []rune("2"), valid1)
	valid2 := configura.Runes("valid_runes_2", []rune("foo"))
	assert.Equal(t, []rune("testing"), valid2)
	invalid1 := configura.Runes("invalid_runes_1", []rune("foo"))
	assert.Equal(t, []rune("foo"), invalid1)
}

func TestFloat32(t *testing.T) {
	valid1 := configura.Float32("valid_float32_1", 0)
	assert.Equal(t, float32(2), valid1)
	valid2 := configura.Float32("valid_float32_2", 0)
	assert.Equal(t, float32(12.134), valid2)
	valid3 := configura.Float32("valid_float32_3", 0)
	assert.Equal(t, float32(-12.642322), valid3)
	valid4 := configura.Float32("valid_float32_4", 22.44)
	assert.Equal(t, float32(9.223372e+18), valid4)
	invalid1 := configura.Float32("invalid_float32_1", 22.22)
	assert.Equal(t, float32(22.22), invalid1)
	invalid2 := configura.Float32("invalid_float32_2", 22.33)
	assert.Equal(t, float32(22.33), invalid2)
}

func TestFloat64(t *testing.T) {
	valid1 := configura.Float64("valid_float64_1", 0)
	assert.Equal(t, float64(2), valid1)
	valid2 := configura.Float64("valid_float64_2", 0)
	assert.Equal(t, 12.134, valid2)
	valid3 := configura.Float64("valid_float64_3", 0)
	assert.Equal(t, -12.642322, valid3)
	valid4 := configura.Float64("valid_float64_4", 22.44)
	assert.Equal(t, 9.223372036854776e+18, valid4)
	invalid1 := configura.Float64("invalid_float64_1", 22.22)
	assert.Equal(t, 22.22, invalid1)
	invalid2 := configura.Float64("invalid_float64_2", 22.33)
	assert.Equal(t, 22.33, invalid2)
}
