package scores

import "strconv"

var MapStatus = map[int]string{
	-1: "NotSubmitted",
	0:  "Pending",
	1:  "UpdateAvailable",
	2:  "Ranked",
	3:  "Approved",
	4:  "Qualified",
	5:  "Loved",
}

func ConvertStatus(s string) string {
	status, serr := strconv.Atoi(s)
	if serr != nil {
		status = 0
	}

	return MapStatus[status]
}
