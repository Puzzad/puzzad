// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/progress"
	"github.com/greboid/puzzad/ent/question"
)

// ProgressCreate is the builder for creating a Progress entity.
type ProgressCreate struct {
	config
	mutation *ProgressMutation
	hooks    []Hook
}

// SetAdventureID sets the "adventure" edge to the Adventure entity by ID.
func (pc *ProgressCreate) SetAdventureID(id int) *ProgressCreate {
	pc.mutation.SetAdventureID(id)
	return pc
}

// SetNillableAdventureID sets the "adventure" edge to the Adventure entity by ID if the given value is not nil.
func (pc *ProgressCreate) SetNillableAdventureID(id *int) *ProgressCreate {
	if id != nil {
		pc = pc.SetAdventureID(*id)
	}
	return pc
}

// SetAdventure sets the "adventure" edge to the Adventure entity.
func (pc *ProgressCreate) SetAdventure(a *Adventure) *ProgressCreate {
	return pc.SetAdventureID(a.ID)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (pc *ProgressCreate) SetQuestionID(id int) *ProgressCreate {
	pc.mutation.SetQuestionID(id)
	return pc
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (pc *ProgressCreate) SetNillableQuestionID(id *int) *ProgressCreate {
	if id != nil {
		pc = pc.SetQuestionID(*id)
	}
	return pc
}

// SetQuestion sets the "question" edge to the Question entity.
func (pc *ProgressCreate) SetQuestion(q *Question) *ProgressCreate {
	return pc.SetQuestionID(q.ID)
}

// Mutation returns the ProgressMutation object of the builder.
func (pc *ProgressCreate) Mutation() *ProgressMutation {
	return pc.mutation
}

// Save creates the Progress in the database.
func (pc *ProgressCreate) Save(ctx context.Context) (*Progress, error) {
	var (
		err  error
		node *Progress
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProgressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Progress)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProgressMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProgressCreate) SaveX(ctx context.Context) *Progress {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProgressCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProgressCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProgressCreate) check() error {
	return nil
}

func (pc *ProgressCreate) sqlSave(ctx context.Context) (*Progress, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ProgressCreate) createSpec() (*Progress, *sqlgraph.CreateSpec) {
	var (
		_node = &Progress{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: progress.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: progress.FieldID,
			},
		}
	)
	if nodes := pc.mutation.AdventureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   progress.AdventureTable,
			Columns: []string{progress.AdventureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adventure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.progress_adventure = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   progress.QuestionTable,
			Columns: []string{progress.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.progress_question = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProgressCreateBulk is the builder for creating many Progress entities in bulk.
type ProgressCreateBulk struct {
	config
	builders []*ProgressCreate
}

// Save creates the Progress entities in the database.
func (pcb *ProgressCreateBulk) Save(ctx context.Context) ([]*Progress, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Progress, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProgressMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProgressCreateBulk) SaveX(ctx context.Context) []*Progress {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProgressCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProgressCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
