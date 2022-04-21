package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	assert := assert.New(t)
	expr, err := Parse("age = 10")
	assert.Nil(err)
	assert.Equal(expr.ColumnName, "age")
	assert.Equal(expr.Operation, Equal)
	assert.Equal(expr.Value, "10")

	expr, err = Parse("name != Maxud")
	assert.Nil(err)
	assert.Equal(expr.ColumnName, "name")
	assert.Equal(expr.Operation, NotEqual)
	assert.Equal(expr.Value, "Maxud")
}
