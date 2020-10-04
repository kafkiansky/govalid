package validator

type Validator struct {
	scope Scope
	data map[string]interface{}
}

func (validator *Validator) validate() ErrorCollection {
	var errors []Result

	catch:
	for field, rules := range validator.scope.rules {

		value := validator.data[field]

		for _, rule := range rules {
			if result := rule.try(value, validator.data); !result.isOk {
				errors = append(errors, result)

				continue catch
			}
		}
	}

	return ErrorCollection{
		errors,
	}
}
