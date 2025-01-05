package db

import "context"

func (db *DB) Init(ctx context.Context) error {
	log := db.logger.With("method", "init")
	stmt := `
	CREATE TABLE IF NOT EXISTS expense (
		id SERIAL PRIMARY KEY,
		description TEXT NOT NULL,
		amount DECIMAL(10,2) NOT NULL,
	    category TEXT NOT NULL,
	    recurring BOOLEAN NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
		`

	if _, err := db.pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "create table expense error", "error", err)
		return err
	}

	seedStmt := `
	INSERT INTO expense (description, amount, category,recurring)
	VALUES ('Netflix',7.99,'Entertainment', TRUE),
	       ('Ice Cream',2.00,'Food', FALSE),
		   ('Book',10.00,'Education', FALSE),
	       `

	if _, err := db.pg.Exec(seedStmt); err != nil {
		log.ErrorContext(ctx, "fail seed table expense error", "error", err)
		return err
	}

	log.InfoContext(ctx, "success expense table created")
	return nil

}
