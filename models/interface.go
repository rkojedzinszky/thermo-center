package models

import (
	"database/sql"
	"fmt"
)

// DBInterface holds common operations for sql.DB and sql.Tx
type DBInterface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// PositionalCounter generates positional parameter markers
type PositionalCounter struct {
	position int
}

// Get generates next positional parameter marker
func (p *PositionalCounter) Get() string {
	p.position++

	return fmt.Sprintf("$%d", p.position)
}

// ConditionFragment represents a condition fragment
type ConditionFragment interface {
	// GetConditionFragment returns ConditionFragment stringified
	GetConditionFragment(*PositionalCounter) (string, []interface{})
}

// ConstantFragment holds constant expression, without parameters
type ConstantFragment struct {
	Constant string
}

// GetConditionFragment returns the constant fragment
func (c *ConstantFragment) GetConditionFragment(*PositionalCounter) (string, []interface{}) {
	return c.Constant, nil
}

// UnaryFragment holds fragment with one parameter
type UnaryFragment struct {
	Frag  string
	Param interface{}
}

// GetConditionFragment returns fragment with its parameter
func (u *UnaryFragment) GetConditionFragment(p *PositionalCounter) (string, []interface{}) {
	return u.Frag + " " + p.Get(), []interface{}{u.Param}
}
