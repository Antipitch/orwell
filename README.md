# Orwell
Orwell is a validator package for the Go programming language.

## Features
- Quick and simple to use
- Easy to extend with custom validation rules
- Some of the most common validation rules baked in
- Many cross-field validation rules (e.g. dependant struct fields)
- Pure Go, no tags
- No third party dependencies
- Avoids complex regexes and reflection whenever possible


## Usage
### Simple data validation
Get a new validator instance:
````
import (
	"github.com/Antipitch/orwell"
)

v := orwell.NewValidator()
````

Validate something:
````
i := 42

err := v.Validate(i, v.Required())

if err != nil {
    log.Print(err.Error())
}
````

Add more Rules:
````
err := v.Validate(i, v.Required(), v.Max(42), v.DivBy(21))
````

Rules after failing rules wil not be applied:
````
// v.DivBy() will not be called
err := v.Validate(i, v.Required(), v.Max(41), v.DivBy(20))
````

Nil or empty values are always valid except for the Required(), NotNil(), NotEmpty() and X-Rules:
````
var (
    i int
    err error
)

// valid
err = v.Validate(i, v.Min(42))

// not valid
err = v.Validate(i, v.NotEmpty(), v.Min(42))

i = 41

// not valid
err = v.Validate(i, v.NotEmpty(), v.Min(42))
````

### Struct Validation
Use pointers both to the struct itself and to the fields to validate (fields passed as arguments to X-rules don't have to be pointers).
````
s := {
    StringField: "somestring",
    IntField: 1000,
    StructField: {
        ChildStringFiled: "",
        ChildIntField: 0
    } 
}

// will return orwell.IterableError interface (s.Intfield is not valid)
err := v.ValidateStruct(&s,
    v.FieldRules(&s.StringField, v.Required()),
    v.FieldRules(&s.StructField.ChildIntField, v.XLt(s.IntField, 1000))
)
````
### Error handling
Orwell exports three error interfaces: InternalError, IterableError and FieldableError (all of them implement Go's error interface).

Orwell's Validate method will -for now- return a standard Go error, nothing special here. The ValidateStruct method will either return an InternalError or an IterableError. Proceed according to your needs with type assertion:
````
err := v.ValidateStruct(
    // validate some struct
)

if err != nil {
    // don't react specifically, maybe log
    log.Print(err.Error())

    // internal error
    internalError, ok := err.(orwell.InternalError)
    if ok {
        // do something with internal error
    }

    // iterable error
    iterableError, ok := err.(orwell.IterableError)
    if ok {
        l := iterableError.Len()
        for i := 0; i < l; i++ {
            if fe, ok := iterableError.ValueAt(i).(orwell.FieldableError); ok {
                f := fe.FieldName
                j := fe.JSONName
                e := fe.Error()
            }
        }
    }
}
````

## Rules
### Simple rules
- `DivBy(arg int)` Must be dividable by arg
- `In(args ...interface{})` Must be equal to one of args
- `Email(arg bool)` Must be RFC 5322 compliant address (uses net/mail.ParseAddress and optionally net.LookupMX)
- `LengthMax(arg int)` Must be of length not greater than arg
- `LengthMin(arg int)` Must be of length not less than arg
- `Max(arg int)` Must not be greater than arg
- `Min(arg int)` Must not be lower than arg
- `Match(arg string)` Must match regular expression described by arg
- `NotNil()` Must not be nil
- `Required()` Must neither be nil nor empty
### X-Rules
- `XAnd(arg interface{})` Required when arg is neither nil nor empty
- `XAndOr(arg interface{}, args ...interface{})` Required when arg is neither nil nor empty and all args are nil or empty
- `XGt(arg1 interface{}, arg2 int)` Required when arg1 is greater than arg2
- `XGtEql(arg1 interface{}, arg2 int, arg3 interface{}, arg4 interface{})` Required when arg1 is greater than arg2 and arg3 is equal to arg4
- `XGtEqlOr(arg1 interface{}, arg2 int, arg3 interface{}, arg4 interface{}, args ...interface{})` Required when arg1 is greater than arg2, arg3 is equal to arg4 and all args are nil or empty
- `XGtAndOr(arg1 interface{}, arg2 int, arg3 interface{}, args ...interface{})` Required when arg1 is greater than arg2, arg3 is neither nil nor empty and all args are nil or empty
- `XGtOr(arg1 interface{}, arg2 int, args ...interface{})` Required when arg1 is greater than arg2 and all args are nil or empty
- `XLt(arg1 interface{}, arg2 int)` Required when arg1 is lower than arg2
- `XLtOr(arg1 interface{}, arg2 int, args ...interface{})` Required when arg1 is lower than arg2 and all args are nil or empty
- `XNor(args ...interface{})` Must be nil or empty when all args are nil or empty
- `XNot(arg interface{})` Must be nil or empty when arg is neither nil nor empty
- `XOr(args ...interface{})` Required when all args are nil or empty

Orwell is a new project and not stable, yet. Don't use it in production! Thankful respects to [asaskevich/govalidator](https://github.com/asaskevich/govalidator) and [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation) for inspiration. Feel free to contribute.

