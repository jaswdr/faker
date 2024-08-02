package faker

import (
	"testing"
)

func TestBankAccountNumber(t *testing.T) {
	v := New().BankAccount().GenerateBankAccount("")
	NotExpect(t, "", v)
	ExpectIntWithSize(t, 15, v)

	v1 := New().BankAccount().GenerateBankAccount("uK")
	NotExpect(t, "", v1)
	ExpectIntWithSize(t, 14, v1)

	v2 := New().BankAccount().GenerateBankAccount("ru")
	NotExpect(t, "", v2)
	ExpectIntWithSize(t, 20, v2)


	v4 := New().BankAccount().GenerateBankAccount("SG")
	NotExpect(t, "", v4)
	ExpectIntWithSize(t, 10, v4)

	v5 := New().BankAccount().GenerateBankAccount("de")
	NotExpect(t, "", v1)
	ExpectIntWithSize(t, 15, v5)
}
