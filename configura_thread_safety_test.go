package configura

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ThreadSafetySuite struct {
	suite.Suite
	config Config
}

func (s *ThreadSafetySuite) SetupTest() {
	s.config = NewConfigImpl()
}

func (s *ThreadSafetySuite) runConcurrently(goroutines int, f func(i int)) {
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()
			f(i)
		}(i)
	}

	wg.Wait()
}

func (s *ThreadSafetySuite) TestConcurrentAccess() {
	type testCase struct {
		name   string
		action func(i int)
	}

	testCases := []testCase{
		{
			name: "String",
			action: func(i int) {
				key := Variable[string](fmt.Sprintf("KEY_%d", i))
				_ = s.config.String(key)
			},
		},
		{
			name: "Int",
			action: func(i int) {
				key := Variable[int](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Int(key)
			},
		},
		{
			name: "Int8",
			action: func(i int) {
				key := Variable[int8](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Int8(key)
			},
		},
		{
			name: "Int16",
			action: func(i int) {
				key := Variable[int16](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Int16(key)
			},
		},
		{
			name: "Int32",
			action: func(i int) {
				key := Variable[int32](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Int32(key)
			},
		},
		{
			name: "Int64",
			action: func(i int) {
				key := Variable[int64](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Int64(key)
			},
		},
		{
			name: "Uint",
			action: func(i int) {
				key := Variable[uint](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Uint(key)
			},
		},
		{
			name: "Uint8",
			action: func(i int) {
				key := Variable[uint8](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Uint8(key)
			},
		},
		{
			name: "Uint16",
			action: func(i int) {
				key := Variable[uint16](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Uint16(key)
			},
		},
		{
			name: "Uint32",
			action: func(i int) {
				key := Variable[uint32](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Uint32(key)
			},
		},
		{
			name: "Uint64",
			action: func(i int) {
				key := Variable[uint64](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Uint64(key)
			},
		},
		{
			name: "Uintptr",
			action: func(i int) {
				key := Variable[uintptr](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Uintptr(key)
			},
		},
		{
			name: "Bytes",
			action: func(i int) {
				key := Variable[[]byte](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Bytes(key)
			},
		},
		{
			name: "Runes",
			action: func(i int) {
				key := Variable[[]rune](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Runes(key)
			},
		},
		{
			name: "Float32",
			action: func(i int) {
				key := Variable[float32](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Float32(key)
			},
		},
		{
			name: "Float64",
			action: func(i int) {
				key := Variable[float64](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Float64(key)
			},
		},
		{
			name: "Bool",
			action: func(i int) {
				key := Variable[bool](fmt.Sprintf("KEY_%d", i))
				_ = s.config.Bool(key)
			},
		},
	}

	const numGoroutines = 100

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()
			s.runConcurrently(numGoroutines, tc.action)
		})
	}
}

func TestThreadSafety(t *testing.T) {
	suite.Run(t, new(ThreadSafetySuite))
}
