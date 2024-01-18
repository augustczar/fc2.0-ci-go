package main

import (
	"log"
	"os"
	"testing"
)

func Testmain(t *testing.T){
	result:= t.main()
	log.Println("Finishing all tests")
	os.Exit(result)
}

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Sum result is invalid: Result: %d. Expected: %d", total, 30)
	}	
}