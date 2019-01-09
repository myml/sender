package share

// TemplateOption template option
// Template is templateID or templateName
// TemplateVariable is template variable
type TemplateOption struct {
	Template         string
	TemplateVariable map[string]interface{}
}

// SendSMSOption send message option
type SendSMSOption struct {
	Phone []string
	TemplateOption
}
