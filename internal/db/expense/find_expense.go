package expense

import (
	"context"
)

func (m *Expense) FindExpense(ctx context.Context, id int64) (*ModelExpense, error) {
	log := m.logger.With("method", "FindExpense")

	stmt := `
		SELECT id, description, amount, category, recurring, created_at, updated_at 
		FROM expense	
		WHERE id = $1
		`

	row := m.db.QueryRowContext(ctx, stmt, id)

	if row.Err() != nil {
		log.ErrorContext(ctx, "fail to query table expense error", "error", row.Err())
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
		log.ErrorContext(ctx, "fail to scan table expense error", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success query table created")
	return &expense, nil

}
