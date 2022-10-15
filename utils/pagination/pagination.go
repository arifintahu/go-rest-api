package pagination

func OffsetAndLimit(page int, limit int) (int, int) {
	if page < 1 {
		page = 1
	}

	if limit < 10 || limit > 200 {
		limit = 50
	}

	offset := (page - 1) * limit
	return offset, limit
}
