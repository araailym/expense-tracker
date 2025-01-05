package expense

import (
	"context"
	"database/sql"
	"errors"
)

func (m *Expense) CreateExpense(ctx context.Context, insertData *ModelExpense) (*ModelExpense, error) {
	log := m.logger.With("method", "CreateExpense")

	stmt := `
		INSERT INTO expense ( description, amount, category, recurring)
		VALUES ($1, $2, $3, $4)
		RETURNING id, description, amount, category, recurring, created_at, updated_at 
		`

	row := m.db.QueryRowContext(ctx, stmt, insertData.Description, insertData.Amount, insertData.Category, insertData.Recurring)

	if row.Err() != nil {
		log.ErrorContext(ctx, "fail to insert into table expense error", "error", row.Err())
		return nil, row.Err()
	}

	expense := ModelExpense{}

	if err := row.Scan(
		&expense.ID,
		&expense.Description,
		&expense.Amount,
		&expense.Category,
		&expense.Recurring,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			log.ErrorContext(ctx, "no values found", "error", err)
			return nil, nil

		}
		log.ErrorContext(ctx, "fail to scan table expense error", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success insert into table created")
	return &expense, nil

}
