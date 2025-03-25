package nuon

func handlePagination[T any](items []T, offset, limit int64) ([]T, bool) {
	hasMore := false

	if limit == 0 {
		limit = 10
	}

	// reduce if over limit
	if len(items) > int(limit) {
		items = items[:len(items)-1]
		hasMore = true
	}

	return items, hasMore
}
