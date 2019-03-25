
func sentinel(sess *xorm.Session) (amount, balance, lastUpdateAt int) {
	rules := make([]model.LotteryRule, 0)

	err := sess.Find(&rules)
	if err != nil {
		return
	}

	for _, r := range rules {
		if r.LastUpdateAt > lastUpdateAt {
			lastUpdateAt = r.LastUpdateAt
		}
		balance += r.Balance
		amount += r.Amount
	}

	return

}
func drawHelper(sess *xorm.Session, beginAt, endAt int) *model.LotteryRule {
	r := randomPickUp(sess)

	//奖品已经发完
	if r == nil {
		return nil
	}

	now := time.Now().Unix()

	amount, balance, lastUpdateAt := sentinel(sess)
	delta := (endAt - beginAt) / amount

	//使用lastUpdateAt作为种子，保证下一个奖品的随机时间对每个抽奖者一样
	rand.Seed(int64(lastUpdateAt))

	//计算下一个奖品的释放时间点
	releaseAt := beginAt + (amount-balance)*delta + int(math.Abs(float64(rand.Int63()%int64(delta))))
	if releaseAt > endAt {
		releaseAt = endAt
	}
	if int(now) < releaseAt {
		return nil
	}
	return r

}

//从所有等级中选中一类奖品,剩余越多的越容易中
func randomPickUp(sess *xorm.Session) *model.LotteryRule {

	rules := make([]model.LotteryRule, 0)

	err := sess.Find(&rules)
	if err != nil {
		return nil
	}

	weight := 0

	for _, r := range rules {
		weight += r.Balance
	}

	//所有商品已经被领完
	if weight == 0 {
		return nil
	}

	rand.Seed(int64(weight))
	num := rand.Intn(weight)

	for _, r := range rules {

		num -= r.Balance

		if num < 0 {
			return &r
		}

	}

	return nil
}

/*
type LotteryRule struct {
	Amount       int    `xorm:"not null comment('奖品数量(-1:无限)') INT(255)"`
	Balance      int    `xorm:"not null comment('剩余奖品数') INT(255)"`
	Desc         string `xorm:"comment('描述信息') TEXT"`
	Id           int    `xorm:"not null pk autoincr INT(10)"`
	LastUpdateAt int    `xorm:"not null default 0 comment('上次中奖时间') INT(11)"`
	Level        string `xorm:"not null comment('等级') VARCHAR(255)"`
	LotteryId    int    `xorm:"not null comment('活动id') INT(11)"`
	Name         string `xorm:"not null comment('规则名') VARCHAR(255)"`
	Rate         int    `xorm:"not null comment('中奖概率(以10亿计)') INT(255)"`
}

type LotteryLog struct {
	CreatedAt     int   `xorm:"comment('中奖时间') INT(255)"`
	Id            int   `xorm:"not null pk autoincr INT(11)"`
	LotteryRuleId int   `xorm:"comment('中奖类型') INT(11)"`
	PickUpAt      int   `xorm:"comment('领取时间') INT(255)"`
	PickUpStatus  int   `xorm:"comment('领取状态: (1:未领取 2-已领取)') INT(255)"`
	Uid           uint64 `xorm:"not null comment('中奖用户') BIGINT(20)"`
}

type Lottery struct {
	Audience  int    `xorm:"default -1 comment('受众(-1:所有用户)') INT(255)"`
	BeginAt   int    `xorm:"not null comment('开启时间') INT(255)"`
	CreatedAt int    `xorm:"not null comment('活动创建时间') INT(255)"`
	EndAt     int    `xorm:"not null default -1 comment('结束时间 (-1:永久)') INT(255)"`
	Id        int    `xorm:"not null pk autoincr INT(10)"`
	Limit     int    `xorm:"not null default 1 comment('次数限制') INT(255)"`
	Name      string `xorm:"default '' comment('活动名称') VARCHAR(255)"`
	Operator  int64  `xorm:"not null default 0 comment('操作者(目前未启用)') BIGINT(255)"`
}

*/
