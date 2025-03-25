package repositories

import "fmt"

var Where = func(column string) string {
	return fmt.Sprintf("%v = ?", column)
}

var Sum = func(column string) string {
	return fmt.Sprintf("SUM(%v)", column)
}
