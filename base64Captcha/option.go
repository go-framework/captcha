package base64Captcha

// With digit captcha option.
func WithDigitCaptcha(config *ConfigDigit) Option {
	return func(impl *captchaImpl) {
		impl.config = *config
	}
}

// With audio captcha option.
func WithAudioCaptcha(config *ConfigAudio) Option {
	return func(impl *captchaImpl) {
		impl.config = *config
	}
}

// With character captcha option.
func WithCharacterCaptcha(config *ConfigCharacter) Option {
	return func(impl *captchaImpl) {
		impl.config = *config
	}
}
