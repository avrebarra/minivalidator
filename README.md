# 🙅 minivalidator

Quick opinionated validation library wrapper over [go-playground/validator](https://github.com/go-playground/validator). Struct tag-based.

## Usage

```go
func main() {
	type UserForm struct {
		Name     string    `validate:"required,gte=7"`
		Email    string    `validate:"email"`
		Age      int       `validate:"required,numeric,min=1,max=99"`
		CreateAt int       `validate:"gte=1"`
		Safe     int       `validate:"-"`
		UpdateAt time.Time `validate:"required"`
		Code     string    `validate:"required"`
	}

	err := minivalidator.Validate(UserForm{})
	fmt.Println(err.Error()) // UserForm.Name must be required; UserForm.Email must be email; UserForm.Age must be required, actual value is 0; UserForm.CreateAt must be gte{1}, actual value is 0; UserForm.UpdateAt must be required, actual value is 0001-01-01 00:00:00 +0000 UTC; UserForm.Code must be required

	err = minivalidator.ValidateWithOpts(UserForm{}, minivalidator.ValidateOptions{Mode: minivalidator.ModeVerbose})
	fmt.Println(err.Error()) // invalid fields at home/projects/myproject/main.go:24: UserForm.Name must be required; UserForm.Email must be email; UserForm.Age must be required, actual value is 0; UserForm.CreateAt must be gte{1}, actual value is 0; UserForm.UpdateAt must be required, actual value is 0001-01-01 00:00:00 +0000 UTC; UserForm.Code must be required

	err = minivalidator.ValidateWithOpts(UserForm{}, minivalidator.ValidateOptions{Mode: minivalidator.ModeCompact})
	fmt.Println(err.Error()) // invalid fields: Name, Email, Age, CreateAt, UpdateAt, Code
}

```

### Supported Struct Tags

This lib is basically a wrapper over github.com/go-playground/validator (v10). Which means all supported validation structs would follow what originally defined by the lib (SPOILER: there are tons).

Check here for available tags or how to add your custom validation tag:

- https://github.com/go-playground/validator
- https://github.com/go-playground/validator#usage-and-documentation
- https://pkg.go.dev/github.com/go-playground/validator

## FAQ

##### Why not just use go-playground/validator?

Because the default error message format from go-playground/validator is imo overly verbose and includes characters like `\n` like this:

`Key: 'UserForm.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'UserForm.Email' Error:Field validation for 'Email' failed on the 'email' tag\nKey: 'UserForm.Age' Error:Field validation for 'Age' failed on the 'required' tag\nKey: 'UserForm.CreateAt' Error:Field validation for 'CreateAt' failed on the 'gte' tag\nKey: 'UserForm.UpdateAt' Error:Field validation for 'UpdateAt' failed on the 'required' tag\nKey: 'UserForm.Code' Error:Field validation for 'Code' failed on the 'required' tag\n`

Using minivalidator, will result to this message:

`UserForm.Name must be required; UserForm.Email must be email; UserForm.Age must be required, actual value is 0; UserForm.CreateAt must be gte{1}, actual value is 0; UserForm.UpdateAt must be required, actual value is 0001-01-01 00:00:00 +0000 UTC; UserForm.Code must be required`
