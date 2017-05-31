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

/*
func QueryNewRoleCount(appId string, channels []string, areas []*protocol.Area, endAt int64) (int, error) {
	var (
		bean = StatsNewRole{AppId: appId}
		cond = ""
	)

	if len(channels) > 0 {
		cond = ChannelCondition(channels)
	} else {
		cond = AreaVectorCondition(areas)
	}

	n, err := defaultEngine.Where(Combined(cond, LtInt64Condition("start_at", endAt))).Sum(bean, "role_count")
	if err != nil {
		return 0, nil
	}

	return int(n), nil
}

*/