package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"event_management/auth"
	"event_management/graph/model"
	service "event_management/graph/services"
	"fmt"
)

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, input model.NewExpense) (*model.Expense, error) {
	userToken := fmt.Sprint(ctx.Value("Token"))

	userId, authErr := auth.CheckLogin(userToken)
	if authErr != nil {
		return nil, authErr

	}
	_, err := auth.IsAdmin((userId), input.EventID, false, false)
	if err != nil {
		return nil, err
	}
	data, err := service.CreateExpense(input.ItemName, input.Cost, input.Description, input.Category, input.EventID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetExpenseByID is the resolver for the getExpenseById field.
func (r *queryResolver) GetExpenseByID(ctx context.Context, id string) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: GetExpenseByID - getExpenseById"))
}

// GetExpensesByEventID is the resolver for the getExpensesByEventId field.
func (r *queryResolver) GetExpensesByEventID(ctx context.Context, eventID string) ([]*model.Expense, error) {
	userToken := fmt.Sprint(ctx.Value("Token"))

	userId, err := auth.CheckLogin(userToken)
	if err != nil {
		return nil, err

	}

	_, err = auth.IsAdmin((userId), eventID, true, true)
	if err != nil {
		return nil, err

	}

	data, err := service.GetExpensesByEventId(eventID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
