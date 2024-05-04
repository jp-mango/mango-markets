CREATE TABLE IF NOT EXISTS daily_stock_data(id SERIAL PRIMARY KEY,
																																													ticker VARCHAR(10) NOT NULL,
																																													price_date DATE, open NUMERIC(10,2),
																																													high NUMERIC(10,2),
																																													low NUMERIC(10,2),
																																													close NUMERIC (10,2),
																																													volume BIGINT);