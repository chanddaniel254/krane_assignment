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

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := service.RegisterUser(input.Name, input.Email, input.Password, input.Phoneno)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (*model.LoginResponse, error) {
	token, err := service.LoginUser(input.Email, input.Password)

	if err != nil {
		return nil, err
	}
	return token, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {

	tokenString := fmt.Sprint(ctx.Value("Token"))
	_, err := auth.CheckLogin(tokenString)

	if err != nil {
		return nil, err
	}

	users, err := service.GetUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {

	tokenString := fmt.Sprint(ctx.Value("Token"))
	_, err := auth.CheckLogin(tokenString)

	if err != nil {
		return nil, err
	}
	user, err := service.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
