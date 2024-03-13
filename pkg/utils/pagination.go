package utils

import (
	"fmt"
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	defaultSize = 10
)

// Pagination query params
type PaginationQuery struct {
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
	Search string `json:"search,omitempty"`
}

// Set page size
func (q *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Limit = defaultSize
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	q.Limit = n

	return nil
}

// Set page number
func (q *PaginationQuery) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Limit = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	q.Page = n

	return nil
}

// // Set order by
// func (q *PaginationQuery) SetOrderBy(orderByQuery string) {
// 	q.OrderBy = orderByQuery
// }

// Get offset
func (q *PaginationQuery) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Limit
}

// Get limit
func (q *PaginationQuery) GetLimit() int {
	return q.Page
}

// // Get OrderBy
// func (q *PaginationQuery) GetOrderBy() string {
// 	return q.OrderBy
// }

// Get OrderBy
func (q *PaginationQuery) GetPage() int {
	return q.Page
}

// Get OrderBy
func (q *PaginationQuery) GetSize() int {
	return q.Limit
}

func (q *PaginationQuery) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", q.GetPage(), q.GetSize())
}

// Get pagination query struct from
func GetPaginationFromCtx(c echo.Context) (*PaginationQuery, error) {
	q := &PaginationQuery{}
	if err := q.SetPage(c.QueryParam("page")); err != nil {
		return nil, err
	}
	if err := q.SetSize(c.QueryParam("limit")); err != nil {
		return nil, err
	}
	q.Search = c.QueryParam("search")
	fmt.Println(c.QueryParam("search"))

	return q, nil
}

// Get total pages int
func GetTotalPages(totalCount int, pageSize int) int {
	d := float64(totalCount) / float64(pageSize)
	return int(math.Ceil(d))
}

// Get has more
func GetHasMore(currentPage int, totalCount int, pageSize int) bool {
	return currentPage < totalCount/pageSize
}
