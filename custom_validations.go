/**
Custom Validator Package
Define Your Own Tags For validation payloads At Run time
Author shubham.dixit@arya.ag
*/

package validations

import (
	"fmt"
	"reflect"
	"strings"
)

const tag_name = "validate"

type Validator interface {
	Validate(interface{}) (bool, error)
}

type StringValidator struct {
	Min int
	Max int
}

func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))
	if l == 0 {
		return false, fmt.Errorf("cannot be blank")
	}
	if l < v.Min {
		return false, fmt.Errorf("should be at least %v chars long", v.Min)
	}
	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("should be less than %v chars long", v.Max)
	}
	return true, nil
}

// DefaultValidator only for default case for switch statement
type DefaultValidator struct {
}

func (d DefaultValidator) Validate(i interface{}) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")
	switch args[0] {
	case "string":
		validator := StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator

	}
	return DefaultValidator{}
}

func ValidateStruct(s interface{}) []error {
	errs := []error{}
	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tag_name)
		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}
		// Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)
		// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface())
		// Append error to results
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}
	return errs
}
