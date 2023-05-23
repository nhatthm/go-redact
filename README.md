# Redact

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/go-redact)](https://github.com/nhatthm/go-redact/releases/latest)
[![Build Status](https://github.com/nhatthm/go-redact/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/go-redact/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/go-redact/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/go-redact)
[![Go Report Card](https://goreportcard.com/badge/go.nhat.io/redact)](https://goreportcard.com/report/go.nhat.io/redact)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/go.nhat.io/redact)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

A simple library to redact sensitive data.

## Prerequisites

- `Go >= 1.18`

## Install

```bash
go get go.nhat.io/redact
```

## Usage

```go
package main

import (
	"fmt"

	"go.nhat.io/redact"
)

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
```

## Donation

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
