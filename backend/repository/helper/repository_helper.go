package helper

import (
	"log"

	"gorm.io/gorm"
)

func ClearPaginateParams(conds ...interface{}) []interface{} {
	var result []interface{}
	var skipNext bool = false
	for _, cond := range conds {
		if cond == "offset" || cond == "page_size" {
			skipNext = true
		} else if !skipNext {
			result = append(result, cond)
		} else {
			skipNext = false
		}
	}

	return result
}

func Paginate(conds ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var offset int
		var pageSize int

		for i, cond := range conds {
			if cond == "offset" {
				if i+1 < len(conds) {
					offset = conds[i+1].(int)
				} else {
					log.Println("Invalid  parameters found in Pagination when looking for offset value")
				}
			} else if cond == "page_size" {
				if i+1 < len(conds) {
					pageSize = conds[i+1].(int)
				} else {
					log.Println("Invalid parameters found in Pagination when looking for page_size value")
				}
			}
		}

		if offset >= 0 && pageSize > 0 {
			return db.Offset(offset).Limit(pageSize)
		} else {
			return db
		}

	}
}
