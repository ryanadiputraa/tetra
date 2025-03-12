package pagination

import (
	"errors"
	"strconv"
)

const (
	defaultPage = 1
	defaultSize = 50
)

type Pagination interface {
	ValidateParam(pageParam, sizeParam string) (page, size int, errDetails map[string]string, err error)
}

type pagination struct{}

func New() Pagination {
	return &pagination{}
}

type Meta struct {
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	Size        int   `json:"size"`
	TotalData   int64 `json:"total_data"`
}

func (p *pagination) ValidateParam(pageParam, sizeParam string) (page, size int, errDetails map[string]string, err error) {
	errDetails = make(map[string]string)

	if len(pageParam) == 0 {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			errDetails["page"] = "invalid 'page' param expecting int"
		}
	}
	if len(sizeParam) == 0 {
		size = defaultSize
	} else {
		size, err = strconv.Atoi(sizeParam)
		if err != nil {
			errDetails["size"] = "invalid 'size' param expecting int"
		}
	}

	if len(errDetails) > 0 {
		err = errors.New("invalid param")
		return
	}

	return
}
