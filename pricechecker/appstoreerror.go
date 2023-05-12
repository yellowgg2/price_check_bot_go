package pricechecker

import "fmt"

type AppStoreError struct {
	ErrMsg string
	ID     string
}

func (ase AppStoreError) Error() string {
	return fmt.Sprintf("ID: %v, Error Message : %v", ase.ID, ase.ErrMsg)
}
