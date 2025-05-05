package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"payments/config"
	"payments/internal/domain/entity"
	"payments/internal/domain/ports"
	"time"
)

type PostgresPaymentRepo struct {
	db *sql.DB
}

func NewPostgresPaymentRepo(db *sql.DB) ports.PaymentRepository {
	return &PostgresPaymentRepo{db: db}
}

func NewPostgresDB(cfg config.DBConfig) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(fmt.Sprintf("database ping failed: %v", err))
	}

	// Set connection pool parameters
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}

func (r *PostgresPaymentRepo) Save(p *entity.Payment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO payments (id, auction_id, amount, status, payment_link, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (auction_id) DO UPDATE
		SET status = EXCLUDED.status, 
			payment_link = EXCLUDED.payment_link,
			amount = EXCLUDED.amount`,
		p.ID, p.AuctionID, p.Amount, p.Status, p.PaymentLink, p.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save payment: %w", err)
	}
	return nil
}

func (r *PostgresPaymentRepo) UpdateStatus(auctionID string, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.db.ExecContext(ctx, `
		UPDATE payments 
		SET status = $1 
		WHERE auction_id = $2`,
		status, auctionID,
	)

	if err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no payment found with auction_id: %s", auctionID)
	}

	return nil
}

func (r *PostgresPaymentRepo) FindByAuctionID(auctionID string) (*entity.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var payment entity.Payment
	err := r.db.QueryRowContext(ctx, `
		SELECT id, auction_id, amount, status, payment_link, created_at
		FROM payments
		WHERE auction_id = $1`,
		auctionID,
	).Scan(
		&payment.ID,
		&payment.AuctionID,
		&payment.Amount,
		&payment.Status,
		&payment.PaymentLink,
		&payment.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("payment not found for auction_id: %s", auctionID)
		}
		return nil, fmt.Errorf("failed to find payment: %w", err)
	}

	return &payment, nil
}

// Additional useful method for expiration worker
func (r *PostgresPaymentRepo) FindExpired(threshold int64) ([]*entity.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, `
		SELECT id, auction_id, amount, status, payment_link, created_at
		FROM payments
		WHERE status = 'pending' 
		AND created_at < $1`,
		time.Now().Add(-time.Duration(threshold)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query expired payments: %w", err)
	}
	defer rows.Close()

	var payments []*entity.Payment
	for rows.Next() {
		var p entity.Payment
		if err := rows.Scan(
			&p.ID,
			&p.AuctionID,
			&p.Amount,
			&p.Status,
			&p.PaymentLink,
			&p.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan payment row: %w", err)
		}
		payments = append(payments, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return payments, nil
}
