package hw09structvalidator

import (
	"errors"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ExpectedStructErr        = errors.New("expected struct, given")
	MinimumValueViolationErr = errors.New("minimum constraint violation")
	MaxValueViolationErr     = errors.New("maximum constraint violation")
	RegexpViolationErr       = errors.New("regexp does not match")
	LenViolationErr          = errors.New("length violation")
	NotInRangeViolationErr   = errors.New("value is not in range")
	validationErrors         ValidationErrors
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var strB strings.Builder

	for _, validationError := range v {
		strB.WriteString("validation error for field: ")
		strB.Write([]byte(validationError.Err.Error()))
		strB.Write([]byte(validationError.Field))
		strB.WriteByte('\n')
		strB.WriteString("error: ")
		strB.WriteByte('\n')
	}

	return strB.String()
}

func Validate(v interface{}) error {
	iv := reflect.ValueOf(v)

	if iv.Kind() != reflect.Struct {
		return ExpectedStructErr
	}

	t := iv.Type()

	structMap := make(map[string]map[string]interface{}, iv.NumField())

	for i := 0; i < iv.NumField(); i++ {
		field := t.Field(i)
		tagVal := field.Tag.Get("validate")

		if tagVal == "" {
			continue
		}

		val := iv.Field(i)

		if val.CanInterface() {
			existing, ok := structMap[field.Name]

			if !ok {
				structMap[field.Name] = make(map[string]interface{})
				structMap[field.Name]["rules"] = tagVal
				structMap[field.Name]["values"] = val.Interface()
			} else {
				existing["rules"] = tagVal
				existing["values"] = val.Interface()
			}
		}
	}

	validationErrors = ValidationErrors{}

	handleValidations(structMap)

	if len(validationErrors) > 0 {
		return validationErrors
	} else {
		return nil
	}
}

func handleValidations(m map[string]map[string]interface{}) {
	for fieldName, val := range m {
		rules := val["rules"]
		values := val["values"]

		resolveValidationType(rules.(string), values, fieldName)
	}
}

func resolveValidationType(rulesStr string, values interface{}, fName string) {
	rules := strings.Split(rulesStr, "|")
	var err error

	for _, rule := range rules {
		r := strings.Split(rule, ":")
		rType := r[0]
		rVal := r[1]

		switch rType {
		case "min":
			value, _ := strconv.Atoi(rVal)
			err = validateMin(value, values)
		case "max":
			value, _ := strconv.Atoi(rVal)
			err = validateMax(value, values)
		case "len":
			value, _ := strconv.Atoi(rVal)
			err = validateLen(value, values)
		case "regexp":
			err = validateRegexp(rVal, values)
		case "in":
			err = validateIn(rVal, values)
		default:
			log.Fatalf("no suitable validation handler found for type %s", rType)
		}
		if err != nil {
			newErr := NewError(fName, err)
			validationErrors = append(validationErrors, *newErr)
		}
	}
}

func validateMin(must int, have interface{}) error {
	switch x := have.(type) {
	case []int:
		for _, val := range x {
			if must > val {
				return MinimumValueViolationErr
			}
		}
	case int:
		if must > x {
			return MinimumValueViolationErr
		}
	}

	return nil
}

func validateMax(must int, have interface{}) error {
	switch x := have.(type) {
	case []int:
		for _, val := range x {
			if must < val {
				return MaxValueViolationErr
			}
		}
	case int:
		if must < x {
			return MaxValueViolationErr
		}
	}

	return nil
}

func validateRegexp(reg string, have interface{}) error {
	switch x := have.(type) {
	case []string:
		for _, val := range x {
			match, _ := regexp.Match(reg, []byte(val))
			if !match {
				return RegexpViolationErr
			}
		}
	case string:
		match, _ := regexp.Match(reg, []byte(x))
		if !match {
			return RegexpViolationErr
		}
	}

	return nil
}

func validateLen(must int, have interface{}) error {
	switch x := have.(type) {
	case []string:
		for _, val := range x {
			if must != len(val) {
				return LenViolationErr
			}
		}
	case string:
		if must != len(x) {
			return LenViolationErr
		}
	}

	return nil
}

func validateIn(must string, have interface{}) error {
	switch x := have.(type) {
	case []string:
		for _, val := range x {
			if must != val {
				return NotInRangeViolationErr
			}
		}
	case string:
		if must != x {
			return NotInRangeViolationErr
		}
	}

	return nil
}

func NewError(field string, err error) *ValidationError {
	return &ValidationError{
		Field: field,
		Err:   err,
	}
}
