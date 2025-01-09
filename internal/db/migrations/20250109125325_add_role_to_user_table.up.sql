DO
$$
BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_type WHERE typname = 'user_role') THEN
CREATE TYPE user_role AS ENUM ('admin', 'authorized');
END IF;
END
$$;

ALTER TABLE user_
ADD COLUMN role user_role NOT NULL DEFAULT 'authorized';