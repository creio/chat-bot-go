// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chathistory"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/predicate"
)

// ChatHistoryDelete is the builder for deleting a ChatHistory entity.
type ChatHistoryDelete struct {
	config
	hooks    []Hook
	mutation *ChatHistoryMutation
}

// Where appends a list predicates to the ChatHistoryDelete builder.
func (chd *ChatHistoryDelete) Where(ps ...predicate.ChatHistory) *ChatHistoryDelete {
	chd.mutation.Where(ps...)
	return chd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (chd *ChatHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ChatHistoryMutation](ctx, chd.sqlExec, chd.mutation, chd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (chd *ChatHistoryDelete) ExecX(ctx context.Context) int {
	n, err := chd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (chd *ChatHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(chathistory.Table, sqlgraph.NewFieldSpec(chathistory.FieldID, field.TypeInt))
	if ps := chd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, chd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	chd.mutation.done = true
	return affected, err
}

// ChatHistoryDeleteOne is the builder for deleting a single ChatHistory entity.
type ChatHistoryDeleteOne struct {
	chd *ChatHistoryDelete
}

// Where appends a list predicates to the ChatHistoryDelete builder.
func (chdo *ChatHistoryDeleteOne) Where(ps ...predicate.ChatHistory) *ChatHistoryDeleteOne {
	chdo.chd.mutation.Where(ps...)
	return chdo
}

// Exec executes the deletion query.
func (chdo *ChatHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := chdo.chd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{chathistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (chdo *ChatHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := chdo.Exec(ctx); err != nil {
		panic(err)
	}
}