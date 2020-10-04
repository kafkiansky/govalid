package validator

import "fmt"

type Rule interface {
	try(value interface{}, data map[string]interface{}) Result
}

type Min struct {
	value int
}

func (min Min) try(value interface{}, data map[string]interface{}) Result {
	return either(value.(int) > min.value, fmt.Sprint("Value is less than ", min.value))
}

type Max struct {
	value int
}

func (max Max) try(value interface{}, data map[string]interface{}) Result {
	return either(value.(int) < max.value, fmt.Sprint("Value is greater than ", max.value))
}