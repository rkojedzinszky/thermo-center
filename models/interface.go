/*
MIT License

Copyright (c) 2020 Richard Kojedzinszky

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package models

import (
	"database/sql"
	"fmt"
	"strings"
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

// AndFragment combines sub-fragments with AND operator
type AndFragment []ConditionFragment

// GetConditionFragment returns fragment with its parameter
func (a AndFragment) GetConditionFragment(c *PositionalCounter) (string, []interface{}) {
	var conds []string
	var condp []interface{}

	for _, cond := range a {
		s, p := cond.GetConditionFragment(c)

		conds = append(conds, s)
		condp = append(condp, p...)
	}

	return strings.Join(conds, " AND "), condp
}

// OrFragment combines sub-fragments with OR operator
type OrFragment []ConditionFragment

// GetConditionFragment returns fragment with its parameter
func (o OrFragment) GetConditionFragment(c *PositionalCounter) (string, []interface{}) {
	var conds []string
	var condp []interface{}

	if len(o) == 0 {
		return "false", nil
	}

	for _, cond := range o {
		s, p := cond.GetConditionFragment(c)

		conds = append(conds, s)
		condp = append(condp, p...)
	}

	return "(" + strings.Join(conds, " OR ") + ")", condp
}
