package cache

import (
	"fmt"
	"strconv"
)

var (
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
