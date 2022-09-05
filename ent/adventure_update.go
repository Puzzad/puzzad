// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/game"
	"github.com/greboid/puzzad/ent/predicate"
	"github.com/greboid/puzzad/ent/puzzle"
)

// AdventureUpdate is the builder for updating Adventure entities.
type AdventureUpdate struct {
	config
	hooks    []Hook
	mutation *AdventureMutation
}

// Where appends a list predicates to the AdventureUpdate builder.
func (au *AdventureUpdate) Where(ps ...predicate.Adventure) *AdventureUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AdventureUpdate) SetName(s string) *AdventureUpdate {
	au.mutation.SetName(s)
	return au
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (au *AdventureUpdate) AddGameIDs(ids ...int) *AdventureUpdate {
	au.mutation.AddGameIDs(ids...)
	return au
}

// AddGame adds the "game" edges to the Game entity.
func (au *AdventureUpdate) AddGame(g ...*Game) *AdventureUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return au.AddGameIDs(ids...)
}

// AddPuzzleIDs adds the "puzzles" edge to the Puzzle entity by IDs.
func (au *AdventureUpdate) AddPuzzleIDs(ids ...int) *AdventureUpdate {
	au.mutation.AddPuzzleIDs(ids...)
	return au
}

// AddPuzzles adds the "puzzles" edges to the Puzzle entity.
func (au *AdventureUpdate) AddPuzzles(p ...*Puzzle) *AdventureUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.AddPuzzleIDs(ids...)
}

// Mutation returns the AdventureMutation object of the builder.
func (au *AdventureUpdate) Mutation() *AdventureMutation {
	return au.mutation
}

// ClearGame clears all "game" edges to the Game entity.
func (au *AdventureUpdate) ClearGame() *AdventureUpdate {
	au.mutation.ClearGame()
	return au
}

// RemoveGameIDs removes the "game" edge to Game entities by IDs.
func (au *AdventureUpdate) RemoveGameIDs(ids ...int) *AdventureUpdate {
	au.mutation.RemoveGameIDs(ids...)
	return au
}

// RemoveGame removes "game" edges to Game entities.
func (au *AdventureUpdate) RemoveGame(g ...*Game) *AdventureUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return au.RemoveGameIDs(ids...)
}

// ClearPuzzles clears all "puzzles" edges to the Puzzle entity.
func (au *AdventureUpdate) ClearPuzzles() *AdventureUpdate {
	au.mutation.ClearPuzzles()
	return au
}

// RemovePuzzleIDs removes the "puzzles" edge to Puzzle entities by IDs.
func (au *AdventureUpdate) RemovePuzzleIDs(ids ...int) *AdventureUpdate {
	au.mutation.RemovePuzzleIDs(ids...)
	return au
}

// RemovePuzzles removes "puzzles" edges to Puzzle entities.
func (au *AdventureUpdate) RemovePuzzles(p ...*Puzzle) *AdventureUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.RemovePuzzleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdventureUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdventureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdventureUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdventureUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdventureUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AdventureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adventure.Table,
			Columns: adventure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adventure.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adventure.FieldName,
		})
	}
	if au.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adventure.GameTable,
			Columns: []string{adventure.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedGameIDs(); len(nodes) > 0 && !au.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adventure.GameTable,
			Columns: []string{adventure.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adventure.GameTable,
			Columns: []string{adventure.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.PuzzlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adventure.PuzzlesTable,
			Columns: []string{adventure.PuzzlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedPuzzlesIDs(); len(nodes) > 0 && !au.mutation.PuzzlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adventure.PuzzlesTable,
			Columns: []string{adventure.PuzzlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.PuzzlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adventure.PuzzlesTable,
			Columns: []string{adventure.PuzzlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adventure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AdventureUpdateOne is the builder for updating a single Adventure entity.
type AdventureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdventureMutation
}

// SetName sets the "name" field.
func (auo *AdventureUpdateOne) SetName(s string) *AdventureUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (auo *AdventureUpdateOne) AddGameIDs(ids ...int) *AdventureUpdateOne {
	auo.mutation.AddGameIDs(ids...)
	return auo
}

// AddGame adds the "game" edges to the Game entity.
func (auo *AdventureUpdateOne) AddGame(g ...*Game) *AdventureUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return auo.AddGameIDs(ids...)
}

// AddPuzzleIDs adds the "puzzles" edge to the Puzzle entity by IDs.
func (auo *AdventureUpdateOne) AddPuzzleIDs(ids ...int) *AdventureUpdateOne {
	auo.mutation.AddPuzzleIDs(ids...)
	return auo
}

// AddPuzzles adds the "puzzles" edges to the Puzzle entity.
func (auo *AdventureUpdateOne) AddPuzzles(p ...*Puzzle) *AdventureUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.AddPuzzleIDs(ids...)
}

// Mutation returns the AdventureMutation object of the builder.
func (auo *AdventureUpdateOne) Mutation() *AdventureMutation {
	return auo.mutation
}

// ClearGame clears all "game" edges to the Game entity.
func (auo *AdventureUpdateOne) ClearGame() *AdventureUpdateOne {
	auo.mutation.ClearGame()
	return auo
}

// RemoveGameIDs removes the "game" edge to Game entities by IDs.
func (auo *AdventureUpdateOne) RemoveGameIDs(ids ...int) *AdventureUpdateOne {
	auo.mutation.RemoveGameIDs(ids...)
	return auo
}

// RemoveGame removes "game" edges to Game entities.
func (auo *AdventureUpdateOne) RemoveGame(g ...*Game) *AdventureUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return auo.RemoveGameIDs(ids...)
}

// ClearPuzzles clears all "puzzles" edges to the Puzzle entity.
func (auo *AdventureUpdateOne) ClearPuzzles() *AdventureUpdateOne {
	auo.mutation.ClearPuzzles()
	return auo
}

// RemovePuzzleIDs removes the "puzzles" edge to Puzzle entities by IDs.
func (auo *AdventureUpdateOne) RemovePuzzleIDs(ids ...int) *AdventureUpdateOne {
	auo.mutation.RemovePuzzleIDs(ids...)
	return auo
}

// RemovePuzzles removes "puzzles" edges to Puzzle entities.
func (auo *AdventureUpdateOne) RemovePuzzles(p ...*Puzzle) *AdventureUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.RemovePuzzleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AdventureUpdateOne) Select(field string, fields ...string) *AdventureUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Adventure entity.
func (auo *AdventureUpdateOne) Save(ctx context.Context) (*Adventure, error) {
	var (
		err  error
		node *Adventure
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdventureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Adventure)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AdventureMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdventureUpdateOne) SaveX(ctx context.Context) *Adventure {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdventureUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdventureUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AdventureUpdateOne) sqlSave(ctx context.Context) (_node *Adventure, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adventure.Table,
			Columns: adventure.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adventure.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Adventure.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adventure.FieldID)
		for _, f := range fields {
			if !adventure.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adventure.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adventure.FieldName,
		})
	}
	if auo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adventure.GameTable,
			Columns: []string{adventure.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedGameIDs(); len(nodes) > 0 && !auo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adventure.GameTable,
			Columns: []string{adventure.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adventure.GameTable,
			Columns: []string{adventure.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.PuzzlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adventure.PuzzlesTable,
			Columns: []string{adventure.PuzzlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedPuzzlesIDs(); len(nodes) > 0 && !auo.mutation.PuzzlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adventure.PuzzlesTable,
			Columns: []string{adventure.PuzzlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.PuzzlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adventure.PuzzlesTable,
			Columns: []string{adventure.PuzzlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Adventure{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adventure.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
