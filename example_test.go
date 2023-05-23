package redact_test

import (
	"fmt"
	"os"

	"go.nhat.io/redact"
)

func ExampleUse() {
	_ = os.Setenv("CUSTOM_ENV", "world") //nolint: errcheck

	r := redact.Use(
		redact.Environ,
		redact.Values("hello"),
	)

	result := r.Redact("hello world!", "hello there")

	for _, s := range result {
		fmt.Println(s)
	}

	// Output:
	// ****** ******!
	// ****** there
}

func ExampleValues() {
	r := redact.Values("hello")

	result := r.Redact("hello world!", "hello there")

	for _, s := range result {
		fmt.Println(s)
	}

	// Output:
	// ****** world!
	// ****** there
}
