package yoctog

import (
	"bytes"
	"log"
	"os"
	"testing"
)

type TestCase struct {
	Level Level
	ExpectValue string
}

var output = new(bytes.Buffer)

func TestDebug(t *testing.T) {
	debugInput := "こんにちは"
	debugExpect := "[DEBUG] こんにちは\n"
	debugSuppressExpect := ""

	testCase := []TestCase{
		{
			Level: LevelDebug,
			ExpectValue: debugExpect,
		},
		{
			Level: LevelInfo,
			ExpectValue: debugSuppressExpect,
		},
		{
			Level: LevelWarn,
			ExpectValue: debugSuppressExpect,
		},
		{
			Level: LevelError,
			ExpectValue: debugSuppressExpect,
		},
		{
			Level: LevelAllSuppress,
			ExpectValue: debugSuppressExpect,
		},
	}
	for _, v := range testCase {
		SetLogLevel(v.Level)
		Debug(debugInput)
		if output.String() != v.ExpectValue {
			t.Error("unexpected value")
		}
		output.Reset()
	}
}

func TestInfo(t *testing.T) {
	infoInput := "こんにちは"
	infoExpect := "[INFO] こんにちは\n"
	infoSuppressExpect := ""

	testCase := []TestCase{
		{
			Level: LevelDebug,
			ExpectValue: infoExpect,
		},
		{
			Level: LevelInfo,
			ExpectValue: infoExpect,
		},
		{
			Level: LevelWarn,
			ExpectValue: infoSuppressExpect,
		},
		{
			Level: LevelError,
			ExpectValue: infoSuppressExpect,
		},
		{
			Level: LevelAllSuppress,
			ExpectValue: infoSuppressExpect,
		},
	}
	for _, v := range testCase {
		SetLogLevel(v.Level)
		Info(infoInput)
		if output.String() != v.ExpectValue {
			t.Error("unexpected value")
		}
		output.Reset()
	}
}

func TestWarn(t *testing.T) {
	warnInput := "こんにちは"
	warnExpect := "[WARN] こんにちは\n"
	warnSuppressExpect := ""

	testCase := []TestCase{
		{
			Level: LevelDebug,
			ExpectValue: warnExpect,
		},
		{
			Level: LevelInfo,
			ExpectValue: warnExpect,
		},
		{
			Level: LevelWarn,
			ExpectValue: warnExpect,
		},
		{
			Level: LevelError,
			ExpectValue: warnSuppressExpect,
		},
		{
			Level: LevelAllSuppress,
			ExpectValue: warnSuppressExpect,
		},
	}
	for _, v := range testCase {
		SetLogLevel(v.Level)
		Warn(warnInput)
		if output.String() != v.ExpectValue {
			t.Error("unexpected value")
		}
		output.Reset()
	}
}

func TestError(t *testing.T) {
	errorInput := "こんにちは"
	errorExpect := "[ERROR] こんにちは\n"
	errorSuppressExpect := ""

	testCase := []TestCase{
		{
			Level: LevelDebug,
			ExpectValue: errorExpect,
		},
		{
			Level: LevelInfo,
			ExpectValue: errorExpect,
		},
		{
			Level: LevelWarn,
			ExpectValue: errorExpect,
		},
		{
			Level: LevelError,
			ExpectValue: errorExpect,
		},
		{
			Level: LevelAllSuppress,
			ExpectValue: errorSuppressExpect,
		},
	}
	for _, v := range testCase {
		SetLogLevel(v.Level)
		Error(errorInput)
		if output.String() != v.ExpectValue {
			t.Error("unexpected value")
		}
		output.Reset()
	}
}

// Testing setup/teardown https://gist.github.com/atotto/88ed1be361976807c171

func setup() {
	log.SetFlags(0)
	log.SetOutput(output)
}

func teardown() {
	log.SetOutput(os.Stdout)
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}