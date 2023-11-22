package mempool

import (
	"context"

	tx "cosmossdk.io/core/transaction"
)

var _ Mempool[tx.Tx] = NoOpMempool[tx.Tx]{}

type NoOpMempool[T tx.Tx] struct {
	txValidator tx.Validator[T]
}

func NewNoopMempool[T tx.Tx](txv tx.Validator[T]) *NoOpMempool[T] {
	return &NoOpMempool[T]{txValidator: txv}
}

func (s *NoOpMempool[T]) Start() error {
	// NoOpMempool[T] does not require any initialization
	return nil
}

func (s *NoOpMempool[T]) Stop() error {
	// NoOpMempool[T] does not require any cleanup
	return nil
}

func (npm NoOpMempool[T]) Insert(ctx context.Context, tx T) map[[32]byte]error {
	_, err := npm.txValidator.Validate(ctx, []T{tx}, false)
	return err
}

func (NoOpMempool[T]) GetTxs(ctx context.Context, size uint32) (any, error) {
	return nil, nil
}

func (NoOpMempool[T]) CountTx() uint32  { return 0 }
func (NoOpMempool[T]) Remove([]T) error { return nil }
