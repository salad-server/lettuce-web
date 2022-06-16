package scores

var Modes = map[string]int{
	"vn!std":   0,
	"vn!taiko": 1,
	"vn!catch": 2,
	"vn!mania": 3,
	"rx!std":   4,
	"rx!taiko": 5,
	"rx!catch": 6,
	"ap!std":   7,
}

func ValidGamemode(m string) bool {
	for k := range Modes {
		if k == m {
			return true
		}
	}

	return false
}
