package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const path = "test.csv"

func TestExport(t *testing.T) {
	assert.Nil(t, exportProducts(generatorProducts(2), path))
}

func TestImport(t *testing.T) {
	p, errImport := importProducts(path)
	assert.Nil(t, errImport)
	assert.Equal(t, len(p), 2)

	for _, v := range p {
		assert.Nil(t, validateProduct(v), v)
	}
}
