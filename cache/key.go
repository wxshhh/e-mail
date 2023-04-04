package cache

import "strconv"

const (
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return "view:product:" + strconv.Itoa(int(id))
}
