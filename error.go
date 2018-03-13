package orwell

type (
	// ValidationError struct
	validationError struct {
		error
		fieldName string
		jsonName  string
		msg       string
	}
	// StructValidationError struct
	structValidationError struct {
		error
		errors []error
	}
	// InternalError struct
	internalError struct {
		err error
	}
)

// NewValidationError func
func NewValidationError(fieldName, jsonName, msg string) *validationError {
	return &validationError{fieldName: fieldName, jsonName: jsonName, msg: msg}
}

func (e validationError) Error() string {
	return e.msg
}

func (e validationError) FieldName() string {
	return e.fieldName
}

func (e validationError) JSONName() string {
	return e.jsonName
}

// SetFieldName func
func (e *validationError) SetFieldName(str string) {
	e.fieldName = str
}

// SetJSONName func
func (e *validationError) SetJSONName(str string) {
	e.jsonName = str
}

func NewStructValidationError() *structValidationError {
	return &structValidationError{}
}

func (e structValidationError) Error() string {
	return "Validation of struct returned erros"
}

func (e structValidationError) Len() int {
	return len(e.errors)
}

func (e structValidationError) ValueAt(i int) error {
	return e.errors[i]
}

func NewInternalError(err error) *internalError {
	return &internalError{err: err}
}

func (e internalError) Error() string {
	return e.err.Error()
}

func (e internalError) Throw() {
}
