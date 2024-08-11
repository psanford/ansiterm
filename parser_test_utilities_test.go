package ansiterm

import (
	"context"
	"log"
	"testing"
)

var debugLogOpt = WithLogf(func(s string, i ...interface{}) {
	log.Printf(s, i...)
})

func createTestParser(ctx context.Context, s string) (*AnsiParser, *TestAnsiEventHandler) {
	evtHandler := CreateTestAnsiEventHandler()
	parser := createParserEventHandler(ctx, evtHandler, WithInitialState(s))

	return parser, evtHandler
}

func validateState(t *testing.T, start string, actualState state, expectedStateName string) {
	actualName := "Nil"

	if actualState != nil {
		actualName = actualState.Name()
	}

	if start == "Utf8" {
		return
	}

	if actualName != expectedStateName {
		if start == "Ground" && actualName == "Utf8" && expectedStateName == "Ground" {
			// for our purposes, utf8 is a sub state within ground
			return
		}
		t.Errorf("Invalid state: '%s' != '%s'", actualName, expectedStateName)
	}
}

func validateFuncCalls(t *testing.T, actualCalls []string, expectedCalls []string) {
	actualCount := len(actualCalls)
	expectedCount := len(expectedCalls)

	if actualCount != expectedCount {
		t.Errorf("Actual   calls: %v", actualCalls)
		t.Errorf("Expected calls: %v", expectedCalls)
		t.Errorf("Call count error: %d != %d", actualCount, expectedCount)
		return
	}

	for i, v := range actualCalls {
		if v != expectedCalls[i] {
			t.Errorf("Actual   calls: %v", actualCalls)
			t.Errorf("Expected calls: %v", expectedCalls)
			t.Errorf("Mismatched calls: %s != %s with lengths %d and %d", v, expectedCalls[i], len(v), len(expectedCalls[i]))
		}
	}
}

func fillContext(context *ansiContext) {
	context.CollectCurrentChar('C')
	context.CurCharToParamBuffer()
	context.CollectCurrentChar('D')
	context.CurCharToParamBuffer()
	context.CollectCurrentChar('E')
	context.CurCharToParamBuffer()
	context.CollectCurrentChar('A')
}

func validateEmptyContext(t *testing.T, context *ansiContext) {
	var expectedCurrChar byte = 0x0
	if context.CurrentChar() != expectedCurrChar {
		t.Errorf("Currentchar mismatch '%#x' != '%#x'", context.CurrentChar(), expectedCurrChar)
	}

	if len(context.ParamBuffer()) != 0 {
		t.Errorf("Non-empty parameter buffer: %v", context.ParamBuffer())
	}
}
