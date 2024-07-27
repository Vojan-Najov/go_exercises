package main

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type (
	User struct {
		ID      int64
		Name    string
		Balance int64
	}

	Expression struct {
		ID         int64
		Expression string
		UserID     int64
	}
)

func (u User) Print() string {
	id := strconv.FormatInt(u.ID, 10)
	balance := strconv.FormatInt(u.Balance, 10)
	return "ID: " + id + " Name: " + u.Name + " Balance: " + balance
}

func (e Expression) Print() string {
	id := strconv.FormatInt(e.ID, 10)
	userID := strconv.FormatInt(e.UserID, 10)
	return "ID: " + id + " Expression: " + e.Expression + " UserID: " + userID
}

func createTables(ctx context.Context, db *sql.DB) error {
	const (
		userTable = `
			CREATE TABLE IF NOT EXISTS users(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT,
				balance INTEGER NOT NULL CHECK(balance >= 0)
			);`

		expressionTable = `
			CREATE TABLE IF NOT EXISTS expressions(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				expression TEXT NOT NULL,
				user_id INTEGER NOT NULL,

				FOREIGN KEY (user_id) REFERENCES expressions (id)
			);`
	)

	if _, err := db.ExecContext(ctx, userTable); err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, expressionTable); err != nil {
		return err
	}

	return nil
}

func insertUser(ctx context.Context, db *sql.DB, user *User) (int64, error) {
	var q = `
		INSERT INTO users (name, balance) values ($1, $2)
	`

	result, err := db.ExecContext(ctx, q, user.Name, user.Balance)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func insertExpression(
	ctx context.Context,
	db *sql.DB,
	expression *Expression,
) (int64, error) {
	var q = `
		INSERT INTO expressions (expression, user_id) values ($1, $2)
	`
	result, err := db.ExecContext(ctx, q, expression.Expression, expression.UserID)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func selectUsers(ctx context.Context, db *sql.DB) ([]User, error) {
	var users []User
	var q = "SELECT id, name, balance FROM users"
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Balance)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func selectExpressions(ctx context.Context, db *sql.DB) ([]Expression, error) {
	var expressions []Expression
	var q = "SELECT id, expression, user_id FROM expressions"

	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		e := Expression{}
		err := rows.Scan(&e.ID, &e.Expression, &e.UserID)
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, e)
	}

	return expressions, nil
}

func main() {
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	if err = createTables(ctx, db); err != nil {
		panic(err)
	}

	user := &User{
		Name:    "Petr",
		Balance: 200,
	}
	userID, err := insertUser(ctx, db, user)
	if err != nil {
		panic(err)
	}

	expression := &Expression{
		Expression: "2+2",
		UserID:     userID,
	}
	expressionID, err := insertExpression(ctx, db, expression)
	if err != nil {
		panic(err)
	}
	expression.ID = expressionID

	users, err := selectUsers(ctx, db)
	if err != nil {
		panic(err)
	}
	for i := range users {
		log.Println(users[i].Print())
	}

	expressions, err := selectExpressions(ctx, db)
	if err != nil {
		panic(err)
	}
	for i := range expressions {
		log.Println(expressions[i].Print())
	}
}