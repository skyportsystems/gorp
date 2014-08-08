//
package gorp

import (
	"database/sql"
)

// Starts a gorp Transaction using an existing Sql transaction
//
func (m *DbMap) SetSqlTransaction(tx *sql.Tx) *Transaction {
	return &Transaction{m, tx, false}
}

// Return underlying sql.Tx object
//
func (t *Transaction) GetSqlTransaction() *sql.Tx {
	return t.tx
}

// SetReadOnly allows you to mark the column as read-only. If true
// this column will be skipped when UPDATE SQL statement is generated
//
func (c *ColumnMap) SetReadOnly(b bool) *ColumnMap {
	c.ReadOnly = b
	return c
}
