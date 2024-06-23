package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGeneratedHandleFunc(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	generateHandleFunc(Endpoint{
		Path:   "v1/test",
		Method: "GET",
		Status: 200,
		Delay:  0,
		Body: map[string]interface{}{
			"hello": "world",
		},
	})(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if string(data) != `{"hello":"world"}` {
		t.Errorf(`expected {"hello":"world"} got %v`, string(data))
	}

	if res.StatusCode != 200 {
		t.Errorf(`expected 200 got %v`, res.Status)
	}
}
