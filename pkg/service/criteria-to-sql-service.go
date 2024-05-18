package service

import (
	"github.com/adriein/tibia-mkt/pkg/types"
	"strings"
)

type CriteriaToSqlService struct {
	table string
}

func NewCriteriaToSqlService(table string) *CriteriaToSqlService {
	return &CriteriaToSqlService{
		table: table,
	}
}

func (c *CriteriaToSqlService) Transform(criteria types.Criteria) (string, error) {
	var where []string

	sql := "SELECT * FROM" + " " + c.table + " WHERE "

	for _, filter := range criteria.Filters {
		clause := filter.Name + " " + filter.Operand + filter.Value

		where = append(where, clause)
	}

	completeSQL := sql + strings.Join(where, " AND ")

	return completeSQL, nil
}
