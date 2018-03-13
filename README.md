# Orwell
Orwell is a validator package for the Go programming language.

## Features
- Quick and simple to use
- Easy to extend with custom validation rules
- Some of the most common validation rules baked in
- Lots of cross-field validation rules (e.g. dependant struct fields)
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

o := orwell.NewOrwell()
````

Validate something:
````
i := 42

err := o.Validate(i, o.Required())

if err != nil {
    log.Print(err.Error())
}
````

Add more Rules:
````
err := o.Validate(i, o.Required(), o.Max(42), o.DivBy(21))
````

Rules after failing rules wil not be applied:
````
// o.DivBy() will not be called
err := o.Validate(i, o.Required(), o.Max(41), o.DivBy(20))
````

Nil or empty values are always valid except for the Required(), NotNil(), NotEmpty() and X-Rules:
````
var (
    i int
    err error
)

// valid
err = o.Validate(i, o.Min(42))

// not valid
err = o.Validate(i, o.NotEmpty(), o.Min(42))

i = 41

// not valid
err = o.Validate(i, o.NotEmpty(), o.Min(42))
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
err := o.ValidateStruct(&s,
    o.FieldRules(&s.StringField, o.Required()),
    o.FieldRules(&s.StructField.ChildIntField, o.XLt(s.IntField, 1000))
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
- `DivBy(arg int)` Must be dividable by arg
- `In(args ...interface{})` Must be equal to one of args
- `Email(arg bool)` Must be RFC 5322 compliant address (uses net/mail.ParseAddress() and optionally net.LookupMX())
- `Max(arg int)` Must not be greater than arg
- `Min(arg int)` Must not be lower than arg
- `NotEmpty()` Must not be empty
- `NotNil()` Must not be nil
- `Required()` Must neither be nil nor empty
- `XAnd(arg interface{})` Required when arg is neither nil nor empty
- `XAndOr(arg interface{}, args ...interface{})` Required when arg is neither nil nor empty and all args are nil or empty
- `XGt(arg interface{}, arg int)` Required when 1st arg is greater than 2nd arg
- `XLt(arg interface{}, arg int)` Required when 1st arg is lower than 2nd arg
- `XNor(args ...interface{})` Must be nil or empty when all args are nil or empty
- `XNot(arg interface{})` Must be nil or empty when arg is neither nil nor empty
- `XOr(args ...interface{})` Required when all args are nil or empty

Orwell is a new project and not stable, yet. Don't use it in production! Thankful respects to [asaskevich/govalidator](https://github.com/asaskevich/govalidator) and [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation) for inspiration. Feel free to contribute.

