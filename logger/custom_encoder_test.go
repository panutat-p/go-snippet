package logger

import (
	"encoding/json"
	"fmt"
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
	c := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	myEncoder = *NewCustomEncoder(c)
	core := zapcore.NewCore(NewCustomEncoder(c), zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	myLogger = zap.New(core)
}

func TestCustomEncoder_EncodeEntry_simple(t *testing.T) {
	given := []zapcore.Field{
		{"animal", zapcore.StringType, 0, "🐵", nil},
		{"name", zapcore.StringType, 0, "John Doe", nil},
		{"weight", zapcore.Int64Type, 42, "", nil},
		{"email", zapcore.StringType, 0, "monkey@gmail.com", nil},
		{"password", zapcore.StringType, 0, "1234", nil},
		{"uuid", zapcore.ReflectType, 0, "", nil},
	}
	b, err := myEncoder.EncodeEntry(
		zapcore.Entry{},
		given,
	)
	if err != nil {
		t.Error("Failed to encode log, err:", err)
	}
	assert.Equal(
		t,
		"{\"animal\":\"🐵\",\"name\":\"John Doe\",\"weight\":42,\"email\":\"monkey@gmail.com\",\"password\":\"***\",\"uuid\":null}\n",
		b.String(),
	)
}

func TestCustomEncoder_EncodeEntry_HTTP_body(t *testing.T) {
	given := `
	{
		"animal": "🐵",
		"profile": {
			"name": "John Doe",
			"address": "111/222 banana road",
			"uuid": null
		},
		"credentials": {
			"email": "monkey@gmail.com",
			"password": "secret"
		}
	}`
	var m map[string]any
	err := json.Unmarshal([]byte(given), &m)
	if err != nil {
		t.Error("Failed to Unmarshal JSON, err:", err)
	}
	b, err := myEncoder.EncodeEntry(
		zapcore.Entry{},
		[]zapcore.Field{{"monkey", zapcore.ReflectType, 0, "", m}},
	)
	if err != nil {
		t.Error("Failed to encode log, err:", err)
	}
	fmt.Println(b)
	assert.Equal(
		t,
		"{\"my_map\":{\"animal\":\"🐵\",\"credentials\":{\"email\":\"monkey@gmail.com\",\"password\":\"***\"},\"profile\":{\"address\":\"111/222 banana road\",\"name\":\"John Doe\",\"uuid\":null}}}\n",
		b.String(),
	)
}

func TestCustomEncoder_EncodeEntry_nil_in_map(t *testing.T) {
	given := map[string]any{
		"key1":     "hello",
		"key2":     nil,
		"key3":     nil,
		"key4":     nil,
		"key5":     nil,
		"password": "secret",
	}
	b, err := myEncoder.EncodeEntry(
		zapcore.Entry{},
		[]zapcore.Field{{"my_map", zapcore.ReflectType, 0, "", given}},
	)
	if err != nil {
		t.Error("Failed to encode log, err:", err)
	}
	assert.Equal(
		t,
		"{\"my_map\":{\"key1\":\"hello\",\"key2\":null,\"key3\":null,\"key4\":null,\"key5\":null,\"password\":\"***\"}}\n",
		b.String(),
	)
}
