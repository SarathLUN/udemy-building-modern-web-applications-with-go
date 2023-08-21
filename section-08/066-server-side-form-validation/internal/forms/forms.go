package forms

import (
  "net/url"
  "net/http"
)

type Form struct {
  url.Values
  Errors errors
}

// New initializes a form struct.
func New(data url.Values) *Form {
  return &Form{
    data,
    errors(map[string][]string{}),
  }
}

// Has validate required form field and retrun wether it pass or fail.
func (f *Form) Has(field string, r *http.Request) bool {
  x := r.Form.Get(field)
  if x == "" {
    return false
  }
  return true
}
