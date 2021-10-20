package validatia

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Mode int

const (
	ModeDefault Mode = iota
	ModeCompact      = iota
	ModeVerbose      = iota
)

var (
	global *Validator
)

func GetGlobal() *Validator {
	if global == nil {
		global = &Validator{Core: validator.New()}
	}
	return global
}

func SetGlobal(in *Validator) {
	global = in
}

func Validate(in interface{}) (err error) {
	return GetGlobal().Validate(in)
}

func ValidateWithOpts(in interface{}, opt ValidateOptions) (err error) {
	return GetGlobal().ValidateWithOpts(in, opt)
}

// ***

type errvalidatia struct {
	original error

	metamode   Mode
	metaerrloc string
}

func (err errvalidatia) Error() (out string) {
	// check if castable to validation error
	errValidator, ok := err.original.(validator.ValidationErrors)
	if !ok {
		return err.original.Error()
	}

	// build error message
	switch err.metamode {
	case ModeCompact:
		fields := []string{}
		for _, q := range errValidator {
			fields = append(fields, q.Field())
		}
		out += fmt.Sprintf("invalid values: %s", strings.Join(fields, ", "))

	case ModeVerbose:
		out += fmt.Sprintf("invalid values at %s: ", err.metaerrloc)
		fallthrough

	default:
		errs := []string{}
		for _, errfield := range errValidator {
			msg := ""

			msg += errfield.StructNamespace()
			msg += " must be " + errfield.ActualTag()

			// print condition parameters, e.g. oneof=red blue -> { red blue }
			if errfield.Param() != "" {
				msg += "{" + errfield.Param() + "}"
			}

			// print actual value
			if errfield.Value() != nil && errfield.Value() != "" {
				msg += fmt.Sprintf(", actual value is %v", errfield.Value())
			}

			errs = append(errs, msg)
		}

		out += strings.Join(errs, "; ")
	}

	return
}

func (e errvalidatia) Unwrap() error {
	return e.original
}

// ***

type Validator struct {
	Core *validator.Validate
}

func (v Validator) Validate(in interface{}) (err error) {
	errv := v.Core.Struct(in)
	if errv == nil {
		return
	}
	return errvalidatia{original: errv, metamode: ModeDefault, metaerrloc: getCallerFuncName()}
}

type ValidateOptions struct {
	Mode Mode
}

func (v Validator) ValidateWithOpts(in interface{}, opt ValidateOptions) (err error) {
	errv := v.Core.Struct(in)
	if errv == nil {
		return
	}
	return errvalidatia{original: errv, metamode: opt.Mode, metaerrloc: getCallerFuncName()}
}

// ***

func getCallerFuncName() (fname string) {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		f, l := details.FileLine(pc)
		fname = fmt.Sprintf("%s:%d", f, l)
	}
	return
}
