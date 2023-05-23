package redact_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/redact"
)

func Test_NoRedact(t *testing.T) {
	t.Parallel()

	values := []string{
		"hello",
		"world",
	}

	actual := redact.NoRedact.Redact(values...)

	assert.Equal(t, values, actual)
}

func Test_Environ(t *testing.T) { //nolint: paralleltest
	t.Setenv("ENV_TEST_ENVIRON", t.Name())

	originalValues := []string{
		"hello",
		t.Name(),
		"world",
	}

	values := append([]string(nil), originalValues...)

	actual := redact.Environ.Redact(values...)
	expected := []string{
		"hello",
		"******",
		"world",
	}

	assert.Equal(t, originalValues, values)
	assert.Equal(t, expected, actual)
}

func Test_Values(t *testing.T) {
	t.Parallel()

	originalValues := []string{
		"hello",
		"world",
	}

	values := append([]string(nil), originalValues...)

	actual := redact.Values("hello").Redact(values...)
	expected := []string{
		"******",
		"world",
	}

	assert.Equal(t, originalValues, values)
	assert.Equal(t, expected, actual)
}

func Test_Use(t *testing.T) { //nolint: paralleltest
	t.Setenv("ENV_TEST_USE", t.Name())

	originalValues := []string{
		"hello world",
		t.Name(),
		"unknown",
	}

	values := append([]string(nil), originalValues...)

	r := redact.Use(
		redact.Use(redact.Environ, nil),
		redact.Values("hello", ""),
	)

	actual := r.Redact(values...)
	expected := []string{
		"****** world",
		"******",
		"unknown",
	}

	assert.Equal(t, originalValues, values)
	assert.Equal(t, expected, actual)
}

func Test_Use_NoRedact(t *testing.T) {
	t.Parallel()

	values := []string{
		"hello",
		"world",
	}

	r := redact.Use(redact.NoRedact)

	actual := r.Redact(values...)

	assert.Equal(t, values, actual)
}
