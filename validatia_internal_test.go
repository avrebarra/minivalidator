package validatia

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var (
	Validatia_SampleValid = SampleStruct{
		Value1: "abc",
		Value2: 120,
	}
	Validatia_SampleInvalid = SampleStruct{
		Value1: "",
		Value2: 1,
	}
)

func TestGetGlobal(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// arrange
		global = nil

		// act
		out := GetGlobal()

		// assert
		assert.NotNil(t, out)
	})

	t.Run("preset", func(t *testing.T) {
		// arrange
		global = &Validator{Core: validator.New()}

		// act
		out := GetGlobal()

		// assert
		assert.Equal(t, global, out)
	})
}

func TestSetGlobal(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// arrange
		global = nil

		// act
		SetGlobal(&Validator{Core: validator.New()})

		// assert
		assert.NotNil(t, global)
	})
}

func TestValidate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// arrange
		global = nil

		// act
		err := Validate(Validatia_SampleValid)

		// assert
		assert.NoError(t, err)
	})

	t.Run("ok invalid data", func(t *testing.T) {
		// arrange
		global = nil

		// act
		err := Validate(Validatia_SampleInvalid)

		// assert
		assert.Error(t, err)
		assert.Equal(t, "SampleStruct.Value1 must be required; SampleStruct.Value2 must be gte{12}, actual value is 1", err.Error())
	})
}

func TestValidateWithOpts(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// arrange
		global = nil

		// act
		err := ValidateWithOpts(SampleStruct{Value1: "abc", Value2: 120}, ValidateOptions{Mode: ModeVerbose})

		// assert
		assert.NoError(t, err)
	})

	t.Run("ok invalid data", func(t *testing.T) {
		// arrange
		global = nil

		// act
		err := ValidateWithOpts(SampleStruct{Value1: "", Value2: 1}, ValidateOptions{Mode: ModeVerbose})

		// assert
		assert.Error(t, err)
		assert.True(t, strings.HasPrefix(err.Error(), "invalid values at "))
		assert.True(t, strings.HasSuffix(err.Error(), ": SampleStruct.Value1 must be required; SampleStruct.Value2 must be gte{12}, actual value is 1"))
	})
}

func Test_errvalidatia_Error(t *testing.T) {
	t.Run("ok non validator origin", func(t *testing.T) {
		// arrange
		orig := fmt.Errorf("random non validation error that possibly wouldnt actually happened")
		err := errvalidatia{original: orig, metamode: ModeDefault, metaerrloc: "-"}

		// act
		msg := err.Error()

		// assert
		assert.Equal(t, orig.Error(), msg)
	})
}

func Test_errvalidatia_Unwrap(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// arrange
		orig := fmt.Errorf("random non validation error that possibly wouldnt actually happened")
		err := errvalidatia{original: orig, metamode: ModeDefault, metaerrloc: "-"}

		// act
		msg := err.Unwrap()

		// assert
		assert.Equal(t, orig, msg)
	})
}

// ***

type SampleStruct struct {
	Value1 string `validate:"required"`
	Value2 int    `validate:"required,gte=12"`
}
