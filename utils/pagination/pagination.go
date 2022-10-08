//go:generate rm -fr mocks
//go:generate mockery --all

package pagination

import (
	"math"
)

type Pagination struct {
	PreviousPage interface{} `json:"previous_page"`
	CurrentPage  int         `json:"current_page"`
	NextPage     interface{} `json:"next_page"`
	Total        int         `json:"total"`
	PerPage      int         `json:"per_page"`
}

func Paginate(page int, perPage int, count int) Pagination {
	var nextPage interface{}
	var previousPage interface{}
	if math.Ceil(float64(count)/float64(perPage)) > float64(page) {
		nextPage = page + 1
	}
	if page > 1 {
		previousPage = page - 1
	}
	return Pagination{
		PreviousPage: previousPage,
		CurrentPage:  page,
		NextPage:     nextPage,
		Total:        count,
		PerPage:      perPage,
	}
}

func CountLimitAndOffset(page int, dataPerPage int) (int, int, int64) {
	var offset, limit int
	var count int64

	offset = -1
	if page > 0 {
		offset = (page - 1) * dataPerPage
	}

	limit = -1
	if dataPerPage > 0 {
		limit = dataPerPage
	}
	return offset, limit, count
}
