CREATE TABLE IF NOT EXISTS expenses (
                          id SERIAL PRIMARY KEY,
                          user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
                          amount NUMERIC NOT NULL CHECK (amount > 0),
                          currency CHAR(3) NOT NULL,
                          spent_at TIMESTAMP NOT NULL,
                          created_at TIMESTAMP DEFAULT NOW(),
                          note TEXT
);

CREATE INDEX IF NOT EXISTS idx_expenses_user_id ON expenses(user_id);