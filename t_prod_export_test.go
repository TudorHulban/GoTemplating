package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {
	path := "test.csv"

	assert.Nil(t, exportProducts(generatorProducts(2), path))
}
