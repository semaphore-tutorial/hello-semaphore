package main

import (
    "io/ioutil"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    helloHandler(w, req)

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)

    expected := "Hello Go!\n"
    if string(body) != expected {
        t.Errorf("expected %q, got %q", expected, string(body))
    }
}
