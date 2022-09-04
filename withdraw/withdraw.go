package withdraw

import "errors"

func Withdraw(balance, amount int) (newBalance int, err error) {
	newBalance = balance - amount
	if newBalance < 0 {
		return 0, errors.New("overdrawn")
	}
	return newBalance, nil
}
