// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/greboid/puzzad/ent/access"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/predicate"
	"github.com/greboid/puzzad/ent/question"
	"github.com/greboid/puzzad/ent/team"
)

// AdventureQuery is the builder for querying Adventure entities.
type AdventureQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.Adventure
	withTeam      *TeamQuery
	withQuestions *QuestionQuery
	withAccess    *AccessQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdventureQuery builder.
func (aq *AdventureQuery) Where(ps ...predicate.Adventure) *AdventureQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AdventureQuery) Limit(limit int) *AdventureQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AdventureQuery) Offset(offset int) *AdventureQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AdventureQuery) Unique(unique bool) *AdventureQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AdventureQuery) Order(o ...OrderFunc) *AdventureQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryTeam chains the current query on the "team" edge.
func (aq *AdventureQuery) QueryTeam() *TeamQuery {
	query := &TeamQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adventure.Table, adventure.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, adventure.TeamTable, adventure.TeamPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryQuestions chains the current query on the "questions" edge.
func (aq *AdventureQuery) QueryQuestions() *QuestionQuery {
	query := &QuestionQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adventure.Table, adventure.FieldID, selector),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adventure.QuestionsTable, adventure.QuestionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccess chains the current query on the "access" edge.
func (aq *AdventureQuery) QueryAccess() *AccessQuery {
	query := &AccessQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adventure.Table, adventure.FieldID, selector),
			sqlgraph.To(access.Table, access.AdventuresColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, adventure.AccessTable, adventure.AccessColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Adventure entity from the query.
// Returns a *NotFoundError when no Adventure was found.
func (aq *AdventureQuery) First(ctx context.Context) (*Adventure, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adventure.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AdventureQuery) FirstX(ctx context.Context) *Adventure {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Adventure ID from the query.
// Returns a *NotFoundError when no Adventure ID was found.
func (aq *AdventureQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adventure.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AdventureQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Adventure entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Adventure entity is found.
// Returns a *NotFoundError when no Adventure entities are found.
func (aq *AdventureQuery) Only(ctx context.Context) (*Adventure, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adventure.Label}
	default:
		return nil, &NotSingularError{adventure.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AdventureQuery) OnlyX(ctx context.Context) *Adventure {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Adventure ID in the query.
// Returns a *NotSingularError when more than one Adventure ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AdventureQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adventure.Label}
	default:
		err = &NotSingularError{adventure.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AdventureQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Adventures.
func (aq *AdventureQuery) All(ctx context.Context) ([]*Adventure, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AdventureQuery) AllX(ctx context.Context) []*Adventure {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Adventure IDs.
func (aq *AdventureQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(adventure.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AdventureQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AdventureQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AdventureQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AdventureQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AdventureQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdventureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AdventureQuery) Clone() *AdventureQuery {
	if aq == nil {
		return nil
	}
	return &AdventureQuery{
		config:        aq.config,
		limit:         aq.limit,
		offset:        aq.offset,
		order:         append([]OrderFunc{}, aq.order...),
		predicates:    append([]predicate.Adventure{}, aq.predicates...),
		withTeam:      aq.withTeam.Clone(),
		withQuestions: aq.withQuestions.Clone(),
		withAccess:    aq.withAccess.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithTeam tells the query-builder to eager-load the nodes that are connected to
// the "team" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AdventureQuery) WithTeam(opts ...func(*TeamQuery)) *AdventureQuery {
	query := &TeamQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withTeam = query
	return aq
}

// WithQuestions tells the query-builder to eager-load the nodes that are connected to
// the "questions" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AdventureQuery) WithQuestions(opts ...func(*QuestionQuery)) *AdventureQuery {
	query := &QuestionQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withQuestions = query
	return aq
}

// WithAccess tells the query-builder to eager-load the nodes that are connected to
// the "access" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AdventureQuery) WithAccess(opts ...func(*AccessQuery)) *AdventureQuery {
	query := &AccessQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withAccess = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Adventure.Query().
//		GroupBy(adventure.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AdventureQuery) GroupBy(field string, fields ...string) *AdventureGroupBy {
	grbuild := &AdventureGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = adventure.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Adventure.Query().
//		Select(adventure.FieldName).
//		Scan(ctx, &v)
func (aq *AdventureQuery) Select(fields ...string) *AdventureSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AdventureSelect{AdventureQuery: aq}
	selbuild.label = adventure.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

func (aq *AdventureQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !adventure.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AdventureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Adventure, error) {
	var (
		nodes       = []*Adventure{}
		_spec       = aq.querySpec()
		loadedTypes = [3]bool{
			aq.withTeam != nil,
			aq.withQuestions != nil,
			aq.withAccess != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Adventure).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Adventure{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withTeam; query != nil {
		if err := aq.loadTeam(ctx, query, nodes,
			func(n *Adventure) { n.Edges.Team = []*Team{} },
			func(n *Adventure, e *Team) { n.Edges.Team = append(n.Edges.Team, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withQuestions; query != nil {
		if err := aq.loadQuestions(ctx, query, nodes,
			func(n *Adventure) { n.Edges.Questions = []*Question{} },
			func(n *Adventure, e *Question) { n.Edges.Questions = append(n.Edges.Questions, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAccess; query != nil {
		if err := aq.loadAccess(ctx, query, nodes,
			func(n *Adventure) { n.Edges.Access = []*Access{} },
			func(n *Adventure, e *Access) { n.Edges.Access = append(n.Edges.Access, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AdventureQuery) loadTeam(ctx context.Context, query *TeamQuery, nodes []*Adventure, init func(*Adventure), assign func(*Adventure, *Team)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Adventure)
	nids := make(map[int]map[*Adventure]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(adventure.TeamTable)
		s.Join(joinT).On(s.C(team.FieldID), joinT.C(adventure.TeamPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(adventure.TeamPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(adventure.TeamPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Adventure]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "team" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aq *AdventureQuery) loadQuestions(ctx context.Context, query *QuestionQuery, nodes []*Adventure, init func(*Adventure), assign func(*Adventure, *Question)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Adventure)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Question(func(s *sql.Selector) {
		s.Where(sql.InValues(adventure.QuestionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.adventure_questions
		if fk == nil {
			return fmt.Errorf(`foreign-key "adventure_questions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "adventure_questions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AdventureQuery) loadAccess(ctx context.Context, query *AccessQuery, nodes []*Adventure, init func(*Adventure), assign func(*Adventure, *Access)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Adventure)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Access(func(s *sql.Selector) {
		s.Where(sql.InValues(adventure.AccessColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AdventureID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "adventure_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AdventureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AdventureQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *AdventureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adventure.Table,
			Columns: adventure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adventure.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adventure.FieldID)
		for i := range fields {
			if fields[i] != adventure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AdventureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(adventure.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = adventure.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdventureGroupBy is the group-by builder for Adventure entities.
type AdventureGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AdventureGroupBy) Aggregate(fns ...AggregateFunc) *AdventureGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AdventureGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AdventureGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !adventure.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AdventureGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AdventureSelect is the builder for selecting fields of Adventure entities.
type AdventureSelect struct {
	*AdventureQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AdventureSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AdventureQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AdventureSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
