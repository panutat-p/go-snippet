package logger

import (
	"encoding/json"
	"os"
	"testing"

  "github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	myLogger  *zap.Logger
	myEncoder CustomEncoder
)

func init() {
	m := map[string]func(string) string{
		"password": MaskString,
		"email":    MaskEmail,
	}
	myEncoder = *NewCustomEncoder(zapcore.EncoderConfig{}, m)
	core := zapcore.NewCore(NewCustomEncoder(zapcore.EncoderConfig{}, m), zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	myLogger = zap.New(core)
}

func TestCustomEncoder_EncodeEntry_types(t *testing.T) {
	myLogger.Info("Monkey",
		zap.String("animal", "🐵"),
		zap.Int("weight", 42),
		zap.String("email", "monkey@gmail.com"),
		zap.String("password", "1234"),
	)
}

func TestCustomEncoder_EncodeEntry_nil(t *testing.T) {
	myLogger.Info("nil",
		zap.Any("name", nil),
	)
}

func TestCustomEncoder_EncodeEntry_string(t *testing.T) {
	myLogger.Info(
		"John",
		zap.Any("passworD", "John"),
	)
}

func TestCustomEncoder_EncodeEntry_integer(t *testing.T) {
	myLogger.Info("nil",
		zap.Any("number", 123),
	)
}

func TestCustomEncoder_EncodeEntry_nil_in_map(t *testing.T) {
	myLogger.Info("nil",
		zap.Any("map", map[string]any{
			"key1": "hello",
			"key2": nil,
		}),
	)
}

func TestCustomEncoder_MaskFields_map(t *testing.T) {
	given := map[string]any{
		"animal":   "🐵",
		"age":      42,
		"email":    "monkey@gmail.com",
		"password": "secret",
	}
	assert.Equal(
		t,
		map[string]any{"animal": "🐵", "age": 42, "email": "monkey@gmail.com", "password": "***"},
		myEncoder.MaskFields(given),
	)
	myLogger.Info(
		"Monkey",
		zap.Any("monkey", given),
	)
}
