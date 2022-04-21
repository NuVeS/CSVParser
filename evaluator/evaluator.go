package evaluator

import (
	"errors"
	"strconv"

	"github.com/NuVeS/CSVParser/parser"
)

func Evaluate(value string, expr *parser.Expression) (bool, error) {
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return false, err
	}

	exprVal, err := strconv.Atoi(expr.Value)
	if err != nil {
		return false, err
	}

	return perform(expr.Operation, valueInt, exprVal)
}

func perform(op parser.Operation, lhs int, rhs int) (bool, error) {
	switch op {
	case parser.Equal:
		return lhs == rhs, nil
	case parser.Less:
		return lhs < rhs, nil
	case parser.More:
		return lhs > rhs, nil
	case parser.LessOrEqual:
		return lhs <= rhs, nil
	case parser.MoreOrEqual:
		return lhs >= rhs, nil
	case parser.NotEqual:
		return lhs != rhs, nil
	}
	return false, errors.New("wrong operation")
}
