package models

import (
    "database/sql"
    "time"

    "backend-urubu-do-pix/utils"
)

type User struct {
    ID            string  `json:"id"`
    InitialAmount float64 `json:"initial_amount"`
    DepositDate   string  `json:"deposit_date"`
    Balance       float64 `json:"balance"`
}

func (u *User) Save() error {
    query := `
    INSERT INTO users (id, initial_amount, deposit_date, balance)
    VALUES (?, ?, ?, ?)`
    _, err := utils.DB.Exec(query, u.ID, u.InitialAmount, u.DepositDate, u.Balance)
    return err
}

func GetByID(id string) (*User, error) {
    query := `SELECT id, initial_amount, deposit_date, balance FROM users WHERE id = ?`
    row := utils.DB.QueryRow(query, id)

    var user User
    err := row.Scan(&user.ID, &user.InitialAmount, &user.DepositDate, &user.Balance)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &user, err
}

func (u *User) GetDepositDateAsTime() (time.Time, error) {
    return time.Parse(time.RFC3339, u.DepositDate)
}