package base64Captcha

import (
	"strings"
	"testing"
)

func TestDefaultDigitCaptcha(t *testing.T) {
	impl := NewCaptcha(WithDigitCaptcha(DefaultDigitCaptcha()))
	id, data := impl.Generate()
	if !strings.Contains(data, "data:image/png;base64") {
		t.Error(id, data)
	}
}

func TestDefaultAudioCaptcha(t *testing.T) {
	impl := NewCaptcha(WithAudioCaptcha(DefaultAudioCaptcha()))
	id, data := impl.Generate()
	if !strings.Contains(data, "data:audio/wav;base64") {
		t.Error(id, data)
	}
}

func TestCharacterCaptcha(t *testing.T) {
	impl := NewCaptcha(WithCharacterCaptcha(DefaultCharacterCaptcha()))
	id, data := impl.Generate()
	if !strings.Contains(data, "data:image/png;base64") {
		t.Error(id, data)
	}
}
