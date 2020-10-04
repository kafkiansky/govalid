package validator

type Result struct {
	isOk bool
	errorMessage string
}

func ok() Result {
	return Result{true, ""}
}

func err(errorMessage string) Result {
	return Result{false, errorMessage}
}

func either(isOk bool, errorMessage string) Result {
	if isOk { return ok() }

	return err(errorMessage)
}

type ErrorCollection struct {
	errors []Result
}

func (collection *ErrorCollection) isEmpty() bool  {
	return 0 == len(collection.errors)
}

func (collection *ErrorCollection) messages() []string  {
	var messages []string

	for _, result := range collection.errors {
		if !result.isOk {
			messages = append(messages, result.errorMessage)
		}
	}

	return messages
}