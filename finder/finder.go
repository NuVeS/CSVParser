package finder

import (
	"github.com/NuVeS/CSVParser/evaluator"
	"github.com/NuVeS/CSVParser/parser"
)

func Find(records [][]string, expr *parser.Expression) [][]string {
	var rows [][]string = make([][]string, 0)
	column := defineColumnIndex(expr.ColumnName, records[0])
	if column == -1 {
		return make([][]string, 0)
	}

	for _, row := range records[1:] {
		res, err := evaluator.Evaluate(row[column], expr)
		if err != nil {
			return make([][]string, 0)
		}

		if res {
			rows = append(rows, row)
		}
	}
	return rows
}

func defineColumnIndex(searchingColumn string, columns []string) int {
	for i, column := range columns {
		if searchingColumn == column {
			return i
		}
	}
	return -1
}
