// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/examples/o2o2types/ent/card"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/predicate"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/user"

	"github.com/facebookincubator/ent/dialect/sql"
)

// CardQuery is the builder for querying Card entities.
type CardQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Card
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (cq *CardQuery) Where(ps ...predicate.Card) *CardQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CardQuery) Limit(limit int) *CardQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CardQuery) Offset(offset int) *CardQuery {
	cq.offset = &offset
	return cq
}

// Order adds an order step to the query.
func (cq *CardQuery) Order(o ...Order) *CardQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryOwner chains the current query on the owner edge.
func (cq *CardQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: cq.config}
	t1 := sql.Table(user.Table)
	t2 := cq.sqlQuery()
	t2.Select(t2.C(card.OwnerColumn))
	query.sql = sql.Select(t1.Columns(user.Columns...)...).
		From(t1).
		Join(t2).
		On(t1.C(user.FieldID), t2.C(card.OwnerColumn))
	return query
}

// Get returns a Card entity by its id.
func (cq *CardQuery) Get(ctx context.Context, id int) (*Card, error) {
	return cq.Where(card.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (cq *CardQuery) GetX(ctx context.Context, id int) *Card {
	c, err := cq.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return c
}

// First returns the first Card entity in the query. Returns *ErrNotFound when no card was found.
func (cq *CardQuery) First(ctx context.Context) (*Card, error) {
	cs, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(cs) == 0 {
		return nil, &ErrNotFound{card.Label}
	}
	return cs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CardQuery) FirstX(ctx context.Context) *Card {
	c, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return c
}

// FirstID returns the first Card id in the query. Returns *ErrNotFound when no id was found.
func (cq *CardQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{card.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (cq *CardQuery) FirstXID(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Card entity in the query, returns an error if not exactly one entity was returned.
func (cq *CardQuery) Only(ctx context.Context) (*Card, error) {
	cs, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(cs) {
	case 1:
		return cs[0], nil
	case 0:
		return nil, &ErrNotFound{card.Label}
	default:
		return nil, &ErrNotSingular{card.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CardQuery) OnlyX(ctx context.Context) *Card {
	c, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// OnlyID returns the only Card id in the query, returns an error if not exactly one id was returned.
func (cq *CardQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{card.Label}
	default:
		err = &ErrNotSingular{card.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (cq *CardQuery) OnlyXID(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cards.
func (cq *CardQuery) All(ctx context.Context) ([]*Card, error) {
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CardQuery) AllX(ctx context.Context) []*Card {
	cs, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return cs
}

// IDs executes the query and returns a list of Card ids.
func (cq *CardQuery) IDs(ctx context.Context) ([]int, error) {
	return cq.sqlIDs(ctx)
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CardQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CardQuery) Count(ctx context.Context) (int, error) {
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CardQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CardQuery) Exist(ctx context.Context) (bool, error) {
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CardQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CardQuery) Clone() *CardQuery {
	return &CardQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]Order{}, cq.order...),
		unique:     append([]string{}, cq.unique...),
		predicates: append([]predicate.Card{}, cq.predicates...),
		// clone intermediate queries.
		sql: cq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Expired time.Time `json:"expired,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Card.Query().
//		GroupBy(card.FieldExpired).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *CardQuery) GroupBy(field string, fields ...string) *CardGroupBy {
	group := &CardGroupBy{config: cq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = cq.sqlQuery()
	return group
}

func (cq *CardQuery) sqlAll(ctx context.Context) ([]*Card, error) {
	rows := &sql.Rows{}
	selector := cq.sqlQuery()
	if unique := cq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := cq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var cs Cards
	if err := cs.FromRows(rows); err != nil {
		return nil, err
	}
	cs.config(cq.config)
	return cs, nil
}

func (cq *CardQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := cq.sqlQuery()
	unique := []string{card.FieldID}
	if len(cq.unique) > 0 {
		unique = cq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := cq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (cq *CardQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cq *CardQuery) sqlIDs(ctx context.Context) ([]int, error) {
	vs, err := cq.sqlAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []int
	for _, v := range vs {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

func (cq *CardQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(card.Table)
	selector := sql.Select(t1.Columns(card.Columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(card.Columns...)...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt64)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CardQuery is the builder for group-by Card entities.
type CardGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CardGroupBy) Aggregate(fns ...Aggregate) *CardGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scan the result into the given value.
func (cgb *CardGroupBy) Scan(ctx context.Context, v interface{}) error {
	return cgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cgb *CardGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cgb *CardGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CardGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cgb *CardGroupBy) StringsX(ctx context.Context) []string {
	v, err := cgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cgb *CardGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CardGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cgb *CardGroupBy) IntsX(ctx context.Context) []int {
	v, err := cgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cgb *CardGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CardGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cgb *CardGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cgb *CardGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CardGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cgb *CardGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cgb *CardGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cgb.sqlQuery().Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CardGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql
	columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
	columns = append(columns, cgb.fields...)
	for _, fn := range cgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(cgb.fields...)
}