package hello

import "testing"

func TestSayHello(t *testing.T) {
	test_string := say_hello()
	assert_string := "Hello, world!"

	if test_string != assert_string {
		t.Errorf("got %q want %q", test_string, assert_string)
	}
}

func TestHello(t *testing.T) {
	test_string := Hello()
	assert_string := "Ahoy, world"

	if test_string != assert_string {
		t.Errorf("got %q want %q", test_string, assert_string)
	}
}
