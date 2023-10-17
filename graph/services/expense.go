package service

import (
	"event_management/database"
	"event_management/graph/model"
	"fmt"
)

func CreateExpense(itemName, cost, description, category, eventId string) (*model.Expense, error) {
	db := database.Db

	var expenseId int
	err := db.QueryRow("Insert into expense (item_name,cost,description,category,event_id) values ($1,$2,$3,$4,$5) returning id", itemName, cost, description, category, eventId).Scan(&expenseId)

	if err != nil {
		return nil, err
	}

	return &model.Expense{
		ID:          fmt.Sprint(expenseId),
		ItemName:    itemName,
		Cost:        cost,
		Description: description,
		Category:    category,
		EventID:     &model.Event{ID: eventId},
	}, nil

}

func GetExpensesByEventId(eventId string) ([]*model.Expense, error) {
	db := database.Db
	rows, err := db.Query("SELECT id, item_name, cost, description, category FROM expense WHERE event_id = $1", eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var expenses []*model.Expense
	for rows.Next() {
		var expense model.Expense
		err = rows.Scan(&expense.ID, &expense.ItemName, &expense.Cost, &expense.Description, &expense.Category)

		if err != nil {
			return nil, err
		}
		expenses = append(expenses, &expense)
	}

	return expenses, nil

}
