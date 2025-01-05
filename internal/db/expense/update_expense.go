package expense

import (
	"context"
	"fmt"
)

func (m *Expense) UpdateExpense(ctx context.Context, id int64, insertData *ModelExpense) error {
	log := m.logger.With("method", "UpdateExpense", "id", id)

	stmt := `
		UPDATE expense
		SET description = $2, amount=$3, category=$4, recurring=$5
		WHERE id=$1
		`

	res, err := m.db.ExecContext(ctx, stmt, id, insertData.Description, insertData.Amount, insertData.Category, insertData.Recurring)
	if err != nil {
		log.ErrorContext(ctx, "fail to update table expense error", "error", err, "id", id)
		return err

	}
	num, err := res.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "fail to update table expense error", "error", err, "id", id)
		return err
	}

	if num == 0 {
		log.WarnContext(ctx, "expense with id was not found", "id", id)
		return fmt.Errorf("expense with id was not found")
	}

	log.InfoContext(ctx, "success update table created", "id", id)
	return nil

}
