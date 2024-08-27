package mongo

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/squirrel"
)

// Squirrel provides wrapper around squirrel package
type Squirrel struct {
	Builder squirrel.StatementBuilderType
}

func NewSquirrel() *Squirrel {
	return &Squirrel{squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
}

func (s *Squirrel) Equal(key string, value interface{}) squirrel.Eq {
	return squirrel.Eq{key: value}
}

func (s *Squirrel) EqualStr(key string) EqualStr {
	return EqualStr(key)
}

func (s *Squirrel) ILike(key string, value interface{}) squirrel.ILike {
	return squirrel.ILike{key: value}
}

func (s *Squirrel) NotEqual(key string, value interface{}) squirrel.NotEq {
	return squirrel.NotEq{key: value}
}

func (s *Squirrel) Or(cond ...squirrel.Sqlizer) squirrel.Or {
	sl := make([]squirrel.Sqlizer, 0, len(cond))
	sl = append(sl, cond...)
	return sl
}

func (s *Squirrel) And(cond ...squirrel.Sqlizer) squirrel.And {
	sl := make([]squirrel.Sqlizer, 0, len(cond))
	sl = append(sl, cond...)
	return sl
}

func (s *Squirrel) Alias(expr squirrel.Sqlizer, alias string) squirrel.Sqlizer {
	return squirrel.Alias(expr, alias)
}

func (s *Squirrel) EqualMany(clauses map[string]interface{}) squirrel.Eq {
	eqMany := make(squirrel.Eq, len(clauses))
	for key, value := range clauses {
		eqMany[key] = value
	}
	return eqMany
}

func (s *Squirrel) Gt(key string, value interface{}) squirrel.Gt {
	return squirrel.Gt{key: value}
}

func (s *Squirrel) Lt(key string, value interface{}) squirrel.Lt {
	return squirrel.Lt{key: value}
}

func (s *Squirrel) Expr(sql string, args ...interface{}) squirrel.Sqlizer {
	return squirrel.Expr(sql, args)
}

func (s *Squirrel) JSONPathWhere(fieldName, jsonbOp, searchField, value string) (string, error) {
	var b strings.Builder
	value = template.HTMLEscapeString(value)
	_, err := fmt.Fprintf(&b, "%s %s? '$.%s ?? (@ == %q)'", fieldName, jsonbOp, searchField, value)
	if err != nil {
		return "", fmt.Errorf("in JSONPathWhere squirrel method: %w", err)
	}

	return b.String(), nil
}

type EqualStr string

func (e EqualStr) ToSql() (sql string, args []interface{}, err error) {
	sql = string(e)
	return
}
