package main

import (
	"fmt"
	"log"

	"github.com/macduyhai/golang-design-pattern/Creational-pattern/builder"
	"github.com/macduyhai/golang-design-pattern/Creational-pattern/factorymethod"
	"github.com/macduyhai/golang-design-pattern/Creational-pattern/singleton"
)

func main() {
	// TestSingleTon()
	// TestBuilder()
	TestFactoryMethod()
}
func TestFactoryMethod() {
	log.Println("Test factorymethod pattern")
	payment := factorymethod.GetPaymentMethod(factorymethod.DebitCard)
	log.Println(payment.Pay(23.5))
	payment = factorymethod.GetPaymentMethod(factorymethod.Cash)
	log.Println(payment.Pay(11))

}
func TestBuilder() {
	log.Println("Test builder pattern")
	account := builder.NewBankAccountBuilder().
		WithOwnerName("Tuan").
		WithOwnerIdentity(123456789).
		AtBranch("Sai Gon").
		OpeningBalance(1000).Build()

	account.Deposit(1000)
	account.WithDraw(1500)
	log.Printf("Balance: %d\n", account.GetBalance())
}
func TestSingleTon() {
	log.Println("Test singleton pattern")
	s := singleton.NewInstance()
	fmt.Println(s.AddOne())
	fmt.Println(s.AddOne())
	s2 := singleton.NewInstance()
	fmt.Println(s2.AddOne())
	fmt.Printf("Add s: %p\n", s)
	fmt.Printf("Add s2:%p\n", s2)
}
