package parser

import (
	"errors"
	"strings"
)

func Parse(s string) (*Expression, error) {
	values := strings.Split(s, " ")
	operation, err := getOperation(values[1])
	if err != nil {
		return nil, err
	}

	expr := &Expression{ColumnName: values[0], Operation: operation, Value: values[2]}
	return expr, nil
}

func getOperation(s string) (Operation, error) {
	switch s {
	case ">":
		return More, nil
	case "<":
		return Less, nil
	case "=":
		return Equal, nil
	case "!=":
		return NotEqual, nil
	case "<=":
		return LessOrEqual, nil
	case ">=":
		return MoreOrEqual, nil
	default:
		return none, errors.New("wrong operation")

	}
}
