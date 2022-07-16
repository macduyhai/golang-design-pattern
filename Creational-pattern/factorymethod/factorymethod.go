package factorymethod

import (
	"fmt"
)

// Tạo một đối tượng mà không cần thiết chỉ ra chính xác lớp nào sẽ được tạo, bằng các nhóm đối tượng
// liên quan đến với nhau và sử dụng một đối tượng trung gian(method cùng tên) để khởi tạo
// đối tượng cần tạo

type PaymentMenthod interface {
	Pay(amount float64) string
}

type PaymentType int
type CashPM struct {
}

type DebitCrashPM struct {
}

const (
	Cash PaymentType = iota
	DebitCard
)

func (c *CashPM) Pay(amount float64) string {
	money := fmt.Sprintf("%f", amount)
	return money + "$ :Pay by Cash success"
}
func (c *DebitCrashPM) Pay(amount float64) string {
	money := fmt.Sprintf("%f", amount)
	return money + "$ :Pay by DebitCard success"
}
func GetPaymentMethod(p PaymentType) PaymentMenthod {
	switch p {
	case Cash:
		return new(CashPM)
	case DebitCard:
		return new(DebitCrashPM)
	default:
		return nil
	}
}
