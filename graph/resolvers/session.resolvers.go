package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"errors"
	"event_management/auth"
	"event_management/graph/model"
	service "event_management/graph/services"
	"fmt"
)

// CreateSession is the resolver for the createSession field.
func (r *mutationResolver) CreateSession(ctx context.Context, input model.NewSession) (*model.Session, error) {
	userToken := fmt.Sprint(ctx.Value("Token"))

	userId, authErr := auth.CheckLogin(userToken)
	if authErr != nil {
		return nil, authErr

	}
	_, err := auth.IsAdmin((userId), input.EventID, true, false)
	if err != nil {
		return nil, err
	}
	data, err := service.CreateSessionForEvent(input.Name, input.StartTime, *input.EndTime, input.EventID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// EditEventSession is the resolver for the editEventSession field.
func (r *mutationResolver) EditEventSession(ctx context.Context, input model.ScheduleSession) (*model.Session, error) {
	userToken := fmt.Sprint(ctx.Value("Token"))

	userId, authErr := auth.CheckLogin(userToken)
	if authErr != nil {
		return nil, authErr

	}

	eventId, err := service.FetchEventIdFromSession(input.ID)
	if err != nil {
		return nil, err
	}

	if eventId == "" {
		return nil, errors.New("session doesnt exists")
	}
	_, err = auth.IsAdmin(userId, eventId, true, false)
	if err != nil {
		return nil, err
	}

	data, err := service.ScheduleSession(input.Name, input.StartTime, input.EndTime, input.ID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetSessionsByEventID is the resolver for the getSessionsByEventId field.
func (r *queryResolver) GetSessionsByEventID(ctx context.Context, eventID string) ([]*model.Session, error) {
	userToken := fmt.Sprint(ctx.Value("Token"))

	user_id, authErr := auth.CheckLogin(userToken)
	if authErr != nil {
		return nil, authErr

	}

	_, err := service.IsUserRelatedToEvent((user_id), eventID)

	if err != nil {
		return nil, err
	}

	data, err := service.GetSessionsByEventId(eventID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
