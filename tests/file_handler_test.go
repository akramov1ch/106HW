package tests

import (
	"106HW/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUploadFile(t *testing.T) {
	req, err := http.NewRequest("POST", "/upload", strings.NewReader("test file content"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "multipart/form-data")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.UploadFile)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, status)
	}
}

func TestDownloadFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/download/testfile.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DownloadFile)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, status)
	}
}
