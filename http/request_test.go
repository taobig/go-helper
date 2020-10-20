package http

import (
	"net/http"
	"testing"
)

var ContextKeyTest = HttpRequestContextKey("test")

func TestSetContext(t *testing.T) {
	t.Parallel()

	str := "hello world"
	var (
		expected = str
	)

	r := &http.Request{}
	r = SetStringInContext(r, ContextKeyTest, str)
	actual := GetStringFromContext(r, ContextKeyTest)
	if actual != expected {
		t.Errorf("GetStringFromContext():%s; expected %s", actual, expected)
	}
}
