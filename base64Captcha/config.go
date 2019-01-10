package base64Captcha

import "github.com/mojocn/base64Captcha"

// Default digit captcha.
func DefaultDigitCaptcha() *ConfigDigit {
	return &ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 4,
	}
}

// Default audio captcha.
func DefaultAudioCaptcha() *ConfigAudio {
	return &ConfigAudio{
		CaptchaLen: 4,
		Language:   "zh",
	}
}

// Default character captcha.
func DefaultCharacterCaptcha() *ConfigCharacter {
	return &ConfigCharacter{
		Height:             80,
		Width:              240,
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    true,
		IsShowSlimeLine:    true,
		IsShowSineLine:     true,
		CaptchaLen:         4,
	}
}
