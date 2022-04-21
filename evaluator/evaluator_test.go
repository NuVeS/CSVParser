package evaluator

import (
	"testing"

	"github.com/NuVeS/CSVParser/parser"
	"github.com/stretchr/testify/assert"
)

func TestEvaluator(t *testing.T) {
	assert := assert.New(t)

	expr := &parser.Expression{ColumnName: "age", Operation: parser.Equal, Value: "10"}

	res, err := Evaluate("100", expr)
	assert.Nil(err)
	assert.Equal(res, false)

	res, err = Evaluate("10", expr)
	assert.Nil(err)
	assert.Equal(res, true)

	res, err = Evaluate("StringCase", expr)
	assert.NotNil(err)
}
