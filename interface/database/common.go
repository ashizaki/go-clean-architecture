package database

const checkNum = 1

func cursorToOffset(cursor uint) uint {
	return cursor - 1
}

// limitForHasNext sets limit for hasNext.
func limitForHasNext(limit int) int {
	return limit + checkNum // DBで次が存在するかを確認するために、limitで指定された数に+1を行う
}

// checkHasNext check whether the data has already existed or not.
func checkHasNext(length, limit int) bool {
	return length >= limit+checkNum
}
