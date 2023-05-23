package redact

import (
	"os"
	"strings"
)

const defaultMask = "******"

// Redactor is an interface that redacts values.
type Redactor interface {
	Redact(values ...string) []string
}

// Redactors is a list of redactors.
type Redactors []Redactor

// Redact redacts values.
func (rs Redactors) Redact(values ...string) []string {
	for _, r := range rs {
		values = r.Redact(values...)
	}

	return values
}

// Use creates a new Redactor from a list of Redactors.
func Use(redactors ...Redactor) Redactor {
	rs := make(Redactors, 0, len(redactors))

	for _, r := range redactors {
		switch r := r.(type) {
		case nil:
		case Redactors:
			rs = append(rs, r...)

		case fn:
			rs = append(rs, r.Fn)

		default:
			rs = append(rs, r)
		}
	}

	return rs
}

// Fn is a function that redacts values.
type Fn func(values ...string) []string

// Redact redacts values.
func (f Fn) Redact(values ...string) []string {
	return f(values...)
}

type fn struct{ Fn }

// NoRedact does not redact values.
var NoRedact = fn{func(values ...string) []string {
	return values
}}

// Environ redacts values that contain environment variables.
var Environ = fn{func(values ...string) []string {
	env := os.Environ()
	oldNew := make([]string, 0, len(env)*2)

	for _, value := range env {
		name, _, _ := strings.Cut(value, "=")

		if v := os.Getenv(name); v != "" {
			oldNew = append(oldNew, v, defaultMask)
		}
	}

	return Redact(strings.NewReplacer(oldNew...).Replace, values)
}}

// Values redacts the values.
func Values(sensitiveValues ...string) Replacer {
	return NewReplacer(defaultMask, sensitiveValues)
}

// Replacer is a string replacer that redacts values.
type Replacer func(s string) string

// Redact redacts values.
func (r Replacer) Redact(values ...string) []string {
	return Redact(r, values)
}

// NewReplacer creates a new Replacer that replaces the sensitive values with the mask.
func NewReplacer(mask string, values []string) Replacer {
	oldNew := make([]string, 0, len(values)*2)

	for _, value := range values {
		if value != "" {
			oldNew = append(oldNew, value, mask)
		}
	}

	return strings.NewReplacer(oldNew...).Replace
}

// Redact redacts values.
func Redact(redact func(s string) string, values []string) []string {
	result := make([]string, len(values))

	for i, value := range values {
		result[i] = redact(value)
	}

	return result
}
