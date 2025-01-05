package expense

import (
	"context"
	"fmt"
)

func (m *Expense) DeleteExpense(ctx context.Context, id int64) error {
	log := m.logger.With("method", "DeleteExpense", "id", id)

	stmt := `
		DELETE FROM expense
		WHERE id=$1
		`

	res, err := m.db.ExecContext(ctx, stmt, id)
	if err != nil {
		log.ErrorContext(ctx, "fail to delete from the table expense error", "error", err, "id", id)
		return err

	}

	num, err := res.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "fail to delete from the table expense error", "error", err, "id", id)
		return err
	}

	if num == 0 {
		log.WarnContext(ctx, "expense with id was not found", "id", id)
		return fmt.Errorf("expense with id was not found")
	}
	log.InfoContext(ctx, "success delete from the table expense", "id", id)
	return nil

}
