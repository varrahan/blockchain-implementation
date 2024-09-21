package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	utils "blockchain-emulator/src/utils"
)

type Transaction struct {
	ID        string
	From      string
	To        string
	Amount    float64
	Timestamp time.Time
}

type TransactionPool struct {
	Transactions []*Transaction
}

// Create a new transaction
func NewTransaction(from string, to string, amount float64) *Transaction {
	transaction := &Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	transaction.ID = transaction.calculateHash()
	return transaction
}

// Hash for transaction ID
func (t *Transaction) calculateHash() string {
	data := t.From + t.To + utils.Int64ToString(t.Timestamp.Unix()) + utils.IntToString(int(t.Amount))
	h := sha256.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

// Converts transaction into a slice of bytes
func Convert(ts []*Transaction) []byte {
	bytes, err := json.Marshal(ts)
	if err != nil {
		return []byte{}
	}
	return bytes
}

// Create a new transaction pool; it's purpose is to store unprocessed transactions
func NewTransactionPool() *TransactionPool {
	return &TransactionPool{
		Transactions: make([]*Transaction, 0),
	}
}

// Add a new transaction to pool
func (tp *TransactionPool) AddTransaction(transaction *Transaction) {
	tp.Transactions = append(tp.Transactions, transaction)
}

// Clear transaction pool
func (tp *TransactionPool) Clear() {
	tp.Transactions = make([]*Transaction, 0)
}

func (ts *Transaction) IsValid() bool {
	return ts != nil
}
