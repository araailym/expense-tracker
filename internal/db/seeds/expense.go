package seeds

import (
	"database/sql"
	"time"
)

type expense struct {
	id          int
	description string
	amount      float64
	category    string
	recurring   bool
	createdAt   *time.Time
	updatedAt   *time.Time
}

var expenses []*expense

func (s *seeder) expense(tx *sql.Tx) error {
	expenses = []*expense{
		{
			description: "Netflix",
			amount:      7.99,
			category:    "Entertainment",
			recurring:   true,
		},
		{
			description: "Ice Cream",
			amount:      2.00,
			category:    "Food",
			recurring:   false,
		},
		{
			description: "Book",
			amount:      10.00,
			category:    "Education",
			recurring:   false,
		},
	}
	sqlQuery := `
		INSERT INTO expense (description, amount, category, recurring)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	sqlStmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return err
	}
	s.cleanups = append(
		s.cleanups, func() error {
			return sqlStmt.Close()
		},
	)
	for _, g := range expenses {
		err := sqlStmt.
			QueryRow(g.description, g.amount, g.category, g.recurring).
			Scan(&g.id)
		if err != nil {
			return err
		}
	}
	return nil
}
