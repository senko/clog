package clog

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSetup(t *testing.T) {
	Setup(INFO, true)
	if cfg.level != INFO || cfg.useColor != true {
		t.Errorf("Setup() doesn't set up config correctly")
	}
}

func TestSetOutput(t *testing.T) {
	out := bytes.Buffer{}

	Setup(DEBUG, false)
	SetOutput(&out)

	Log(DEBUG, "test set output")
	if !strings.Contains(out.String(), "test set output") {
		t.Error("output not redirected correctly")
	}
}

func TestSetupFromEnv(t *testing.T) {
	os.Setenv("LOG_LEVEL", "WARNING")
	os.Setenv("LOG_COLOR", "true")

	SetupFromEnv()
	if cfg.level != WARNING || cfg.useColor != true {
		t.Errorf("SetupFromEnv() doesn't set up config correctly")
	}
}

func TestColorOutput(t *testing.T) {
	out := bytes.Buffer{}

	Setup(DEBUG, true)
	SetOutput(&out)

	Log(WARNING, "test color")
	if !strings.Contains(out.String(), colorCodes[WARNING-DEBUG]) {
		t.Errorf("Correct color output not detected")
	}
}

func TestMessageFormatting(t *testing.T) {
	out := bytes.Buffer{}

	Setup(DEBUG, false)
	SetOutput(&out)

	Logf(DEBUG, "Hello %s world %d!", "happy", 42)
	if !strings.Contains(out.String(), "Hello happy world 42!") {
		t.Errorf("Log message not formatted as expected")
	}
}

func TestLogFormat(t *testing.T) {
	out := bytes.Buffer{}

	Setup(DEBUG, false)
	SetOutput(&out)

	Log(WARNING, "message")

	parts := strings.Split(out.String(), " ")

	if len(parts) != 3 {
		t.Errorf("Incorrect log output format: %s", out.String())
	}

	if parts[1] != "WARNING" {
		t.Errorf("Incorrect log level string: %s", parts[1])
	}

	if parts[2] != "message\n" {
		t.Errorf("Incorrect log message: %s", parts[2])
	}

	_, err := time.Parse(time.RFC3339, parts[0])
	if err != nil {
		t.Errorf("Incorrect log timestamp: %s", err)
	}
}

func TestPanic(t *testing.T) {
	out := bytes.Buffer{}
	Setup(DEBUG, false)
	SetOutput(&out)

	defer func() {
		if p := recover(); p == nil {
			t.Errorf("Expected panic, but it didn't happen")
		}
	}()

	Log(PANIC, "omg!")
}
