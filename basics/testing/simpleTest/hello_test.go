package main

import "testing"

func TestHelloEmptyArg(t *testing.T)  {
	emptyResult := hello("")
	if emptyResult != "Hello Dude!"{
		t.Errorf("failed. Expected %v , got %v", "Hello Dude!", emptyResult)
	}else{
		t.Logf("success. Expected %v , got %v", "Hello Dude!", emptyResult)
	}
}

func TestHelloValidArg(t *testing.T)  {
	result := hello("Mike")
	if result != "Hello Mike!"{
		t.Errorf("failed. Expected %v , got %v", "Hello Mike!", result)
	}else{
		t.Logf("success. Expected %v , got %v", "Hello Mike!", result)
	}
}
/*
just to show cover percentage, "go test -cover"
to make cover.txt file , "go test -coverprofile=cover.txt"
to visualize , "go tool cover -html=cover.txt -o cover.html"
 */