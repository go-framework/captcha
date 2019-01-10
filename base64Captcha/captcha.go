package base64Captcha

import (
	"gitee.com/mjob/account/captcha"
	"github.com/mojocn/base64Captcha"
)

// Captcha option.
type Option func(*captchaImpl)

// Type base64Captcha type.
type ConfigDigit = base64Captcha.ConfigDigit
type ConfigAudio = base64Captcha.ConfigAudio
type ConfigCharacter = base64Captcha.ConfigCharacter

// implementing captcha interface.
type captchaImpl struct {
	config interface{}
}

// New captcha with option, default is use character captcha.
func NewCaptcha(opt ...Option) captcha.Captcha {
	impl := &captchaImpl{}

	// just one option
	for _, o := range opt {
		o(impl)
	}

	// default is character captcha
	if impl.config == nil {
		impl.config = DefaultCharacterCaptcha()
	}

	return impl
}

// Generate captcha return this id and base64 data.
func (c *captchaImpl) Generate() (string, string) {
	id, instance := base64Captcha.GenerateCaptcha("", c.config)
	data := base64Captcha.CaptchaWriteToBase64Encoding(instance)

	return id, data
}

// Use id and value verify captcha when return true than succeed.
func (c *captchaImpl) Verify(id string, value string) bool {
	return base64Captcha.VerifyCaptcha(id, value)
}
