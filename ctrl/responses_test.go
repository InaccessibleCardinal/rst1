package ctrl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSuccessResponse(t *testing.T) {

	rec := httptest.NewRecorder()
	mockResp := map[string]int{"problems": 99}
	expected := []byte(`{"problems":99}`)

	SendResponse(rec, mockResp, http.StatusOK)
	if status := rec.Result().StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d",
			status, http.StatusOK)
	}
	defer rec.Result().Body.Close()
	actual, _ := io.ReadAll(rec.Result().Body)
	if strings.Compare(string(actual), string(expected)) != 1 {
		t.Errorf("expected %s but got %s", string(actual), string(expected))
	}
}
