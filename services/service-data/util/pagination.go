package util

import (
	"net/http"
	"strconv"
)

func GetPaginationParams(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		page = 1
	}

	if page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("size"))

	if err != nil {
		pageSize = 25
	}

	return page, pageSize
}
