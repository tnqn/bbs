package helpers

import (
	"database/sql"
	"fmt"
	"strings"

	"code.cloudfoundry.org/lager"
)

// SELECT <columns> FROM <table> WHERE ... [FOR UPDATE]
func (h *sqlHelper) All(
	logger lager.Logger,
	q Queryable,
	table string,
	columns ColumnList,
	lockRow RowLock,
	wheres string,
	whereBindings ...interface{},
) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s\n", strings.Join(columns, ", "), table)

	if len(wheres) > 0 {
		query += "WHERE " + wheres
	}

	if lockRow {
		query += "\nFOR UPDATE"
	}

	return q.Query(h.Rebind(query), whereBindings...)
}
