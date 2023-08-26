package models

import "github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/066-server-side-form-validation/internal/forms"

// TemplateData holds data send from handlers to templates.
type TemplateData struct {
	StringMap    map[string]string
	IntMap       map[string]int32
	FloatMap     map[string]float32
	Data         map[string]interface{}
	CSRFToken    string
	FlashMessage string
	WarnMessage  string
	ErrorMessage string
	Form         *forms.Form
}
