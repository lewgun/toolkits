package sql

import (
	"fmt"
	"strings"
)

// 给定列, 返回起始时间条件SQL语句, [begin, end)
func RangeCondition(column string, begin, end int64) string {
	return fmt.Sprintf("(`%s` BETWEEN %d AND %d)", column, begin, end)
}


func EqIntCondition(col string, v int) string {
	return fmt.Sprintf("`%s`=%d", col, v)
}

func EqInt64Condition(col string, v int64) string {
	return fmt.Sprintf("`%s`=%d", col, v)
}

func LtInt64Condition(col string, v int64) string {
	return fmt.Sprintf("`%s`<%d", col, v)
}

func Combined(cond ...string) string {
	return strings.Join(cond, " AND ")
}
