package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors formErrors
}

func NewForm(values url.Values) *Form {
	return &Form{values, formErrors(map[string][]string{})}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		if strings.TrimSpace(f.Get(field)) == "" {
			f.Errors.Add(field, "The field can not be blank")
		}
	}
}

func (f *Form) MaxLength(field string, maxLength int) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > maxLength {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (max  %d characters),",
			maxLength))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
