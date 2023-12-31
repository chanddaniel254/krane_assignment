package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"event_management/auth"
	"event_management/graph"
	"event_management/graph/model"
	service "event_management/graph/services"
	"fmt"
)

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	tokenString := fmt.Sprint(ctx.Value("Token"))
	userId, err := auth.CheckLogin(tokenString)

	if err != nil {
		return nil, err
	}
	event, err := service.CreateEvent(input.Name, input.StartDate, input.EndDate, input.Location, userId)
	if err != nil {
		return nil, err
	}
	return event, nil

}

// EditEventSchedule is the resolver for the editEventSchedule field.
func (r *mutationResolver) EditEventSchedule(ctx context.Context, input model.ScheduleEvent) (*model.Event, error) {
	userToken := fmt.Sprint(ctx.Value("Token"))

	userId, err := auth.CheckLogin(userToken)
	if err != nil {
		return nil, err

	}
	_, err = auth.IsAdmin(fmt.Sprint(userId), input.EventID, false, false)
	if err != nil {
		return nil, err

	}
	data, err := service.EditEventSchedule(input.StartDate, input.EndDate, input.Location, input.EventID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	tokenString := fmt.Sprint(ctx.Value("Token"))
	userId, err := auth.CheckLogin(tokenString)

	if err != nil {
		return nil, err
	}
	events, err := service.GetEvents(userId)
	if err != nil {
		return nil, err
	}
	return events, err
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
