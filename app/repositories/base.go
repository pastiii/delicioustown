package repositories

func GetOffsetAndPageSize(page int, pageSize int) (offset int, size int) {
	if pageSize < 1 {
		pageSize = 10
	}

	if page < 1 {
		page = 1
	}

	return pageSize * (page - 1), pageSize
}
