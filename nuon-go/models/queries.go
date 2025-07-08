package models

type GetPaginatedQuery struct {
	Offset            int
	Limit             int
	PaginationEnabled bool
}
