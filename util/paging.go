package util

// Parse pagination
func Parse(page *int, perPage *int) (int, int) {
	offset := 0 // no. of records to skip
	limit := 10 // limit

	if page == nil || perPage == nil {
		return offset, limit
	}

	if *perPage > 0 && *perPage <= 20 {
		limit = *perPage
	}

	if *page > 1 {
		offset = (*page - 1) * limit
	}

	return offset, limit
}
