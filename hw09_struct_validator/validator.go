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
	ErrExpectedStruct        = errors.New("expected struct, given")
	ErrMinimumValueViolation = errors.New("minimum constraint violation")
	ErrMaxValueViolation     = errors.New("maximum constraint violation")
	ErrRegexpViolation       = errors.New("regexp does not match")
	ErrLenViolation          = errors.New("length violation")
	ErrNotInRangeViolation   = errors.New("value is not in range")
	validationErrors         ValidationErrors
)

type ValidationError struct {
	Field string
	Err   error
}

func (v ValidationError) Error() string {
	var strB strings.Builder

	strB.Write([]byte(v.Err.Error()))

	return strB.String()
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var strB strings.Builder

	for _, validationError := range v {
		strB.WriteString(validationError.Error())
	}

	return strB.String()
}

func Validate(v interface{}) error {
	iv := reflect.ValueOf(v)

	if iv.Kind() != reflect.Struct {
		return ErrExpectedStruct
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
	}

	return nil
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
				return ErrMinimumValueViolation
			}
		}
	case int:
		if must > x {
			return ErrMinimumValueViolation
		}
	}

	return nil
}

func validateMax(must int, have interface{}) error {
	switch x := have.(type) {
	case []int:
		for _, val := range x {
			if must < val {
				return ErrMaxValueViolation
			}
		}
	case int:
		if must < x {
			return ErrMaxValueViolation
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
				return ErrRegexpViolation
			}
		}
	case string:
		match, _ := regexp.Match(reg, []byte(x))
		if !match {
			return ErrRegexpViolation
		}
	}

	return nil
}

func validateLen(must int, have interface{}) error {
	switch x := have.(type) {
	case []string:
		for _, val := range x {
			if must != len(val) {
				return ErrLenViolation
			}
		}
	case string:
		if must != len(x) {
			return ErrLenViolation
		}
	}

	return nil
}

func validateIn(must string, have interface{}) error {
	numRange := strings.Split(must, ",")

	switch x := have.(type) {
	case []string:
		for _, n := range numRange {
			for _, val := range x {
				if n != val {
					return ErrNotInRangeViolation
				}
			}
		}
	case string:
		for _, n := range numRange {
			if n != x {
				return ErrNotInRangeViolation
			}
		}
	case []int:
		found := false
		for _, n := range numRange {
			for _, val := range x {
				m, _ := strconv.Atoi(n)
				if m == val {
					found = true
					break
				}
			}
		}
		if !found {
			return ErrNotInRangeViolation
		}

	case int:
		found := false
		for _, n := range numRange {
			m, _ := strconv.Atoi(n)
			if m == x {
				found = true
				break
			}
		}

		if !found {
			return ErrNotInRangeViolation
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
