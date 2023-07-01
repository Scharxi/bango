package util

func CalculateOffset(page, limit int) int32 {
	offset := (page - 1) * limit
	return int32(offset)
}
