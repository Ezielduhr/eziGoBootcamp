package forms

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form{
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) HasRequired(tagIDs ...string){
	for _, tagID := range tagIDs {
		value := f.Get(tagID)
		if strings.TrimSpace(value) == "" {
			f.Errors.AddError(tagID, "This field can't be blank")
		}
	}
}

func (f *Form) HasValue(tagID string, r *http.Request) bool {
	x := r.Form.Get(tagID)
	return x != ""
}

func (f *Form) MinLength(tagID string, length int, r *http.Request) bool {
	x := r.Form.Get(tagID)
	if len(x) < length {
		f.Errors.AddError(tagID, fmt.Sprintf("This field must be %d characters long or more"))
	}
	return true
}

func (f *Form) IsEmail(tagID string) {
	if !govalidator.IsEmail(f.Get(tagID)) {
		f.Errors.AddError(tagID, "Invalid Email")
	}
}


func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}