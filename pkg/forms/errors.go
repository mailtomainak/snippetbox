package forms

// Define a formErrors type which will be used to hold the error messages for
// the forms.
type formErrors map[string][]string

func (fe formErrors) Add(field, message string) {
	fe[field] = append(fe[field], message)
}
func (fe formErrors) Get(field string) string {
	messageList := fe[field]
	if len(messageList) == 0 {
		return ""
	}
	return messageList[0]
}
