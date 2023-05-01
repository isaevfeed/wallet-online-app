package wallet

import "log"

type Wallet struct {
	Amount float32 `json:"money"`
}

func NewWallet() *Wallet {
	return &Wallet{
		Amount: 10000,
	}
}

func (w *Wallet) Take(wallet *Wallet, amount float32) {
	if w.Amount < amount {
		log.Println("Недостаточно средств")
	}

	w.Amount -= amount
	wallet.Add(amount)
}

func (w *Wallet) Add(amount float32) {
	w.Amount += amount
}
