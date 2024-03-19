package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHola(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := hola("Juan")

	// coincidencia exacta de msg con juan. si lo que devuelve matchString es False o hay un error manda el fatal
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello() `)
	}
	fmt.Sprint(want.MatchString(msg), msg)
}
