package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForms_Valid(t *testing.T){
  r := httptest.NewRequest("POST","/whatever", nil)
  form := New(r.PostForm)
  isValid := form.Valid()
  if !isValid {
    t.Error("got invalid when it should be valid")
  }
}

func TestForms_Required(t *testing.T){
  r := httptest.NewRequest("POST","/whatever", nil)
  form := New(r.PostForm)
  form.Required("a","b","c")
  if form.Valid() {
    t.Error("form shows valid when required fields is missing")
  }

  postedData := url.Values{}
  postedData.Add("a","a")
  postedData.Add("b","b")
  postedData.Add("c","c")

  r, _ = http.NewRequest("POST","/whatever",nil)
  r.PostForm = postedData
  form = New(r.PostForm)
  form.Required("a","b","c")
  if !form.Valid() {
    t.Error("show does not have required fields when it does")
  }
}
