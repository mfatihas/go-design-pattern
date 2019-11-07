package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal("A Payment method of type 'Cash' must be exist")
	}

	msg := payment.Pay(10.30)

	if !strings.Contains(msg, "paid using cash") {
		t.Error("Payment method incorrect")
	}

	t.Log("Log: ", msg)
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(Debit)

	if err != nil {
		t.Fatal("A Payment method of type 'Debit' must be exist")
	}

	msg := payment.Pay(10.30)

	if !strings.Contains(msg, "paid using debit") {
		t.Error("Payment method incorrect")
	}

	t.Log("Log: ", msg)
}

func TestCreatePaymentMethodUnexisted(t *testing.T) {
	_, err := GetPaymentMethod(20)

	if err == nil {
		t.Fatal("A Payment method with ID = 20 must be error")
	}

	t.Log("Log: ", err)
}
