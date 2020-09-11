package bago1

import (
	"testing"
)

func TestHandleRequest(t *testing.T) {
	t.Log("Given the need to test something...")
	{
		t.Logf("\tShould expect something...")

		event := MyEvent{ Name: "MyTest"}
		actual := "Hello, MyName"

		res, err := HandleRequest( nil, event)
		if err != nil {
			t.Errorf("Should expect response but received error=%v", err)
		} else {
			if res.Message != actual {
				t.Errorf("Should expect response message \"%s\" but received \"%s\"", actual, res.Message)
			} else {
				t.Logf("Should expect response message \"%s\"", actual)
			}
		}
	}
}
