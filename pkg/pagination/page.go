package pagination

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Paginator struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type Pagination struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	Total       int64 `json:"total"`
}

func (p Paginator) Offset() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return (p.Page - 1) * p.PageSize
}

func (p Paginator) Limit() int {
	if p.PageSize == 0 {
		p.Page = PageSize10
	}
	return p.PageSize
}

func (p Paginator) NextPage(total int64) Pagination {
	return Pagination{
		CurrentPage: p.Page + 1,
		PageSize:    p.PageSize,
		Total:       total,
	}
}

func GinPaginator(ctx *gin.Context) Paginator {
	pageStr := ctx.Query("page")
	sizeStr := ctx.Query("page_size")

	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.ParseUint(sizeStr, 10, 64)
	if err != nil {
		pageSize = PageSize10
	}

	return Paginator{
		Page:     int(page),
		PageSize: int(pageSize),
	}
}
