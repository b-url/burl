//go:build tools
// +build tools

package main

import (
	_ "github.com/matryer/moq"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "golang.org/x/tools/cmd/stringer"
)
