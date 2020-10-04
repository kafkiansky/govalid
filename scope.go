package validator

type UnderValidation interface {
	rules() map[string][]Rule
}

type Scope struct {
	rules map[string][]Rule
}

func FromArray(rules map[string][]Rule) Scope {
	return Scope{rules}
}

func FromStruct(validation UnderValidation) Scope {
	return Scope{validation.rules()}
}
