// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chatconfig"
)

// ChatConfigCreate is the builder for creating a ChatConfig entity.
type ChatConfigCreate struct {
	config
	mutation *ChatConfigMutation
	hooks    []Hook
}

// SetChatID sets the "chat_id" field.
func (ccc *ChatConfigCreate) SetChatID(i int64) *ChatConfigCreate {
	ccc.mutation.SetChatID(i)
	return ccc
}

// SetJSON sets the "json" field.
func (ccc *ChatConfigCreate) SetJSON(s string) *ChatConfigCreate {
	ccc.mutation.SetJSON(s)
	return ccc
}

// Mutation returns the ChatConfigMutation object of the builder.
func (ccc *ChatConfigCreate) Mutation() *ChatConfigMutation {
	return ccc.mutation
}

// Save creates the ChatConfig in the database.
func (ccc *ChatConfigCreate) Save(ctx context.Context) (*ChatConfig, error) {
	return withHooks[*ChatConfig, ChatConfigMutation](ctx, ccc.sqlSave, ccc.mutation, ccc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ccc *ChatConfigCreate) SaveX(ctx context.Context) *ChatConfig {
	v, err := ccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccc *ChatConfigCreate) Exec(ctx context.Context) error {
	_, err := ccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccc *ChatConfigCreate) ExecX(ctx context.Context) {
	if err := ccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccc *ChatConfigCreate) check() error {
	if _, ok := ccc.mutation.ChatID(); !ok {
		return &ValidationError{Name: "chat_id", err: errors.New(`ent: missing required field "ChatConfig.chat_id"`)}
	}
	if _, ok := ccc.mutation.JSON(); !ok {
		return &ValidationError{Name: "json", err: errors.New(`ent: missing required field "ChatConfig.json"`)}
	}
	return nil
}

func (ccc *ChatConfigCreate) sqlSave(ctx context.Context) (*ChatConfig, error) {
	if err := ccc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ccc.mutation.id = &_node.ID
	ccc.mutation.done = true
	return _node, nil
}

func (ccc *ChatConfigCreate) createSpec() (*ChatConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &ChatConfig{config: ccc.config}
		_spec = sqlgraph.NewCreateSpec(chatconfig.Table, sqlgraph.NewFieldSpec(chatconfig.FieldID, field.TypeInt))
	)
	if value, ok := ccc.mutation.ChatID(); ok {
		_spec.SetField(chatconfig.FieldChatID, field.TypeInt64, value)
		_node.ChatID = value
	}
	if value, ok := ccc.mutation.JSON(); ok {
		_spec.SetField(chatconfig.FieldJSON, field.TypeString, value)
		_node.JSON = value
	}
	return _node, _spec
}

// ChatConfigCreateBulk is the builder for creating many ChatConfig entities in bulk.
type ChatConfigCreateBulk struct {
	config
	builders []*ChatConfigCreate
}

// Save creates the ChatConfig entities in the database.
func (cccb *ChatConfigCreateBulk) Save(ctx context.Context) ([]*ChatConfig, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cccb.builders))
	nodes := make([]*ChatConfig, len(cccb.builders))
	mutators := make([]Mutator, len(cccb.builders))
	for i := range cccb.builders {
		func(i int, root context.Context) {
			builder := cccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChatConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cccb *ChatConfigCreateBulk) SaveX(ctx context.Context) []*ChatConfig {
	v, err := cccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cccb *ChatConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := cccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cccb *ChatConfigCreateBulk) ExecX(ctx context.Context) {
	if err := cccb.Exec(ctx); err != nil {
		panic(err)
	}
}