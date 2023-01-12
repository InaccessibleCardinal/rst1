package ctrl

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessResponse(t *testing.T) {

	rec := httptest.NewRecorder()
	mockResp := map[string]int{"problems": 99}

	SendResponse(rec, mockResp, http.StatusOK)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
