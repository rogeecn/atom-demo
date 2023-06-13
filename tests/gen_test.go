package main

import (
	"regexp"
	"testing"
)

func Test_Gen(t *testing.T) {
	pattern := regexp.MustCompile(`(?mi)\{(.*?)\}`)

	data := map[string]string{
		"/user/{id}":                       "/user/:id",
		"/user/{id}/name":                  "/user/:id/name",
		"/user/{id}/name/{name}":           "/user/:id/name/:name",
		"/user/{id}/name/{name}/age/{age}": "/user/:id/name/:name/age/:age",
	}

	for k, v := range data {
		if pattern.MatchString(k) {
			t.Log(k, "match")
		} else {
			t.Error(k, "not match")
		}

		if pattern.ReplaceAllString(k, ":$1") == v {
			t.Log(k, "replace")
		} else {
			t.Error(k, "not replace")
		}
	}

}
