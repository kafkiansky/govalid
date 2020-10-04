package validator

import (
	"testing"
)

func TestValidatorFromArrayIsEmpty(t *testing.T)  {
	data := map[string]interface{} {
		"age": 18,
	}

	rules := make(map[string][]Rule)

	rules["age"] = []Rule{
		Min{10},
		Max{19},
	}

	validator := Validator{FromArray(rules), data}

	errors := validator.validate()

	if !errors.isEmpty() {
		t.Error("Expected true, got", errors.isEmpty())
	}
}

type UserDto struct {
	age int
}

func (userDto UserDto) rules() map[string][]Rule  {
	rules := make(map[string][]Rule)

	rules["age"] = []Rule{
		Min{10},
		Max{19},
	}

	return rules
}

func TestValidatorFromStructIsEmpty(t *testing.T) {
	user := UserDto{}

	data := map[string]interface{} {
		"age": 18,
	}

	validator := Validator{FromStruct(user), data}

	errors := validator.validate()

	if !errors.isEmpty() {
		t.Error("Expected true, got", errors.isEmpty())
	}
}

func TestValidatorFromArrayHasErrors(t *testing.T) {
	data := map[string]interface{} {
		"age": 20,
	}

	rules := make(map[string][]Rule)

	rules["age"] = []Rule{
		Min{10},
		Max{19},
	}

	validator := Validator{FromArray(rules), data}

	errors := validator.validate()

	if errors.isEmpty() {
		t.Error("Expected false, got", errors.isEmpty())
	}

	message := errors.messages()[0]

	if message != "Value is greater than 19" {
		t.Error("Expected message 'Value is greater than 19', got", message)
	}
}