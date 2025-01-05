package expense

import (
	"context"
	"fmt"
)

func (m *Expense) FindExpenses(ctx context.Context, offset, limit int) ([]ModelExpense, error) {
	log := m.logger.With("method", "FindExpenses")

	expenses := make([]ModelExpense, 0)

	stmt := `
		SELECT id, description, amount, category, recurring, created_at, updated_at 
		FROM expense	
		OFFSET $1
		LIMIT $2
		`

	rows, err := m.db.QueryContext(ctx, stmt, offset, limit)
	if err != nil {
		log.ErrorContext(ctx, "fail to query table expense error", "error", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		expense := ModelExpense{}

		if err := rows.Scan(
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

		fmt.Println("expense", expense)

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success expenses query table created")
	return expenses, nil

}
