package valeed_test

import (
	"strings"
	"testing"

	"github.com/avrebarra/valeed"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	valeed_SampleValid = SampleStruct{
		Value1: "ok",
		Value2: 100,
	}
	valeed_SampleInvalid = SampleStruct{
		Value1: "",
		Value2: 2,
	}
)

func TestValidator_Validate(t *testing.T) {
	t.Run("ok valid", func(t *testing.T) {
		// arrange
		v := valeed.Validator{Core: validator.New()}

		// act
		err := v.Validate(valeed_SampleValid)

		// assert
		assert.NoError(t, err)
	})

	t.Run("ok invalid", func(t *testing.T) {
		// arrange
		v := valeed.Validator{Core: validator.New()}

		// act
		err := v.Validate(valeed_SampleInvalid)

		// assert
		assert.Error(t, err)
		assert.Equal(t, "SampleStruct.Value1 must be required; SampleStruct.Value2 must be gte{12}, actual value is 2", err.Error())
	})
}

func TestValidator_ValidateWithOpts(t *testing.T) {
	t.Run("ok valid", func(t *testing.T) {
		// arrange
		v := valeed.Validator{Core: validator.New()}

		// act
		err := v.ValidateWithOpts(valeed_SampleValid, valeed.ValidateOptions{Mode: valeed.ModeVerbose})

		// assert
		assert.NoError(t, err)
	})

	t.Run("ok invalid", func(t *testing.T) {
		t.Skip() // already tested in Test_errvaleed_Error
	})
}

func Test_errvaleed_Error(t *testing.T) {
	t.Run("ok mode compact", func(t *testing.T) {
		// arrange
		err := valeed.ValidateWithOpts(valeed_SampleInvalid, valeed.ValidateOptions{Mode: valeed.ModeCompact})
		require.Error(t, err)

		// act
		msg := err.Error()

		// assert
		assert.Equal(t, "invalid values: Value1, Value2", msg)
	})

	t.Run("ok mode verbose", func(t *testing.T) {
		// arrange
		err := valeed.ValidateWithOpts(valeed_SampleInvalid, valeed.ValidateOptions{Mode: valeed.ModeVerbose})
		require.Error(t, err)

		// act
		msg := err.Error()

		// assert
		assert.Error(t, err)
		assert.True(t, strings.HasPrefix(msg, "invalid values at "))
		assert.True(t, strings.HasSuffix(msg, ": SampleStruct.Value1 must be required; SampleStruct.Value2 must be gte{12}, actual value is 2"))
	})
}

// ***

type SampleStruct struct {
	Value1 string `validate:"required"`
	Value2 int    `validate:"required,gte=12"`
}
