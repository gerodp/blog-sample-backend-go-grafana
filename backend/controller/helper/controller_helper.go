package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParsePaginationParams(c *gin.Context) ([]interface{}, error) {

	var page int
	var pageSize int
	var err error

	pageParam, pageFound := c.GetQuery("page")

	if pageFound {
		page, err = strconv.Atoi(pageParam)

		if err != nil {
			return nil, err
		}

	}

	if page == 0 {
		page = 1
	}

	pageSizeParam, pageSizeFound := c.GetQuery("page_size")
	if pageSizeFound {
		pageSize, err = strconv.Atoi(pageSizeParam)

		if err != nil {
			return nil, err
		}
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	result := make([]interface{}, 4)

	result[0] = "offset"
	result[1] = offset

	result[2] = "page_size"
	result[3] = pageSize

	return result, nil

}
