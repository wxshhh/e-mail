package cache

import "strconv"

const (
	RankKey      = "rank"
	OrderTimeKey = "OrderTime"
)

func ProductViewKey(id uint) string {
	return "view:product:" + strconv.Itoa(int(id))
}
