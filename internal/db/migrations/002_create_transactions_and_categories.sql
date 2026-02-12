CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL UNIQUE,
    type VARCHAR(10) NOT NULL CHECK (type IN ('income', 'expense')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Create enum type for transaction types
CREATE TYPE transaction_type AS ENUM ('income', 'expense', 'transfer');

-- Create enum type for transaction status
CREATE TYPE transaction_status AS ENUM ('pending', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type transaction_type NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id),
    amount NUMERIC(12, 2) NOT NULL CHECK (amount > 0),
    currency VARCHAR(3) NOT NULL DEFAULT 'IDR',
    description TEXT,
    status transaction_status NOT NULL DEFAULT 'completed',
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX idx_transactions_type ON transactions (type);
CREATE INDEX idx_transactions_category_id ON transactions (category_id);
CREATE INDEX idx_transactions_status ON transactions (status);
CREATE INDEX idx_transactions_date ON transactions (date);

-- Seed default categories
INSERT INTO categories (name, type) VALUES
    ('Gaji', 'income'),
    ('Freelance', 'income'),
    ('Investasi', 'income'),
    ('Lainnya (Pemasukan)', 'income'),
    ('Makanan & Minuman', 'expense'),
    ('Transportasi', 'expense'),
    ('Belanja', 'expense'),
    ('Tagihan', 'expense'),
    ('Hiburan', 'expense'),
    ('Kesehatan', 'expense'),
    ('Pendidikan', 'expense'),
    ('Lainnya (Pengeluaran)', 'expense')
ON CONFLICT (name) DO NOTHING;
