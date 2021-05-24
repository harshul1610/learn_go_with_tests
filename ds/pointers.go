// Wallet system
package ds

type Wallet struct {
	bal int
}

func balance(w *Wallet) int {
	return (*w).bal
}

func deposit(w *Wallet, amount int) {
	(*w).bal += amount
}
