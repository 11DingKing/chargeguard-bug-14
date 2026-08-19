package httpapi

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTaskBehavior(t *testing.T) {
	rr := httptest.NewRecorder()
	TaskHTTPHandler(rr, httptest.NewRequest("POST", "/task", nil))
	if rr.Code != 200 || !strings.Contains(rr.Body.String(), "operator-a") {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}
}
