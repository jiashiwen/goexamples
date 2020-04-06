package context_sample

import "testing"

func ExampleContextWithTimeout() {
	ContextWithTimeout()

	//out:
	//process request with 500ms
	//main context deadline exceeded
}

//func TestContextWithCancel(t *testing.T) {
//	ContextWithCancel()
//}

func TestContextWithValue(t *testing.T) {
	ContextWithValue()
}
