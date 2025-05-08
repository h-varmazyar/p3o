package repositories

import (
	"fmt"
)

var Where = func(column string) string {
	return fmt.Sprintf("%v = ?", column)
}

var Sum = func(column string) string {
	return fmt.Sprintf("COALESCE(SUM(%v), 0)", column)
}

var Count = func(column string) string {
	return fmt.Sprintf("COUNT(%v)", column)
}

var Between = func(column string) string {
	return fmt.Sprintf("%v BETWEEN ? AND ?", column)
}
