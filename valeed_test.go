package minivalidator_test

import (
	"strings"
	"testing"

	"github.com/avrebarra/minivalidator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	minivalidator_SampleValid = SampleStruct{
		Value1: "ok",
		Value2: 100,
	}
	minivalidator_SampleInvalid = SampleStruct{
		Value1: "",
		Value2: 2,
	}
)

func TestValidator_Validate(t *testing.T) {
	t.Run("ok valid", func(t *testing.T) {
		// arrange
		v := minivalidator.Validator{Core: validator.New()}

		// act
		err := v.Validate(minivalidator_SampleValid)

		// assert
		assert.NoError(t, err)
	})

	t.Run("ok invalid", func(t *testing.T) {
		// arrange
		v := minivalidator.Validator{Core: validator.New()}

		// act
		err := v.Validate(minivalidator_SampleInvalid)

		// assert
		assert.Error(t, err)
		assert.Equal(t, "SampleStruct.Value1 must be required; SampleStruct.Value2 must be gte{12}, actual value is 2", err.Error())
	})
}

func TestValidator_ValidateWithOpts(t *testing.T) {
	t.Run("ok valid", func(t *testing.T) {
		// arrange
		v := minivalidator.Validator{Core: validator.New()}

		// act
		err := v.ValidateWithOpts(minivalidator_SampleValid, minivalidator.ValidateOptions{Mode: minivalidator.ModeVerbose})

		// assert
		assert.NoError(t, err)
	})

	t.Run("ok invalid", func(t *testing.T) {
		t.Skip() // already tested in Test_errminivalidator_Error
	})
}

func Test_errminivalidator_Error(t *testing.T) {
	t.Run("ok mode compact", func(t *testing.T) {
		// arrange
		err := minivalidator.ValidateWithOpts(minivalidator_SampleInvalid, minivalidator.ValidateOptions{Mode: minivalidator.ModeCompact})
		require.Error(t, err)

		// act
		msg := err.Error()

		// assert
		assert.Equal(t, "invalid values: Value1, Value2", msg)
	})

	t.Run("ok mode verbose", func(t *testing.T) {
		// arrange
		err := minivalidator.ValidateWithOpts(minivalidator_SampleInvalid, minivalidator.ValidateOptions{Mode: minivalidator.ModeVerbose})
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
