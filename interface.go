package captcha

// Captcha interface.
type Captcha interface {
	// Generate captcha return this id and base64 data.
	Generate() (id string, data string)
	// Use id and value verify captcha when return true than succeed.
	Verify(id string, value string) bool
}
