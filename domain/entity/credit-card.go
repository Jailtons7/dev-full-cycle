package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number           string
	name             string
	expirationMonth  int
	expirationYear   int
	verificationCode int
}

func NewCreditCard(number string, name string, expirationMonth int, expirationYear int, verificationCode int) (*CreditCard, error) {
	cc := &CreditCard{
		number:           number,
		name:             name,
		expirationMonth:  expirationMonth,
		expirationYear:   expirationYear,
		verificationCode: verificationCode,
	}

	err := cc.IsValid()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *CreditCard) IsValid() error {
	err := c.ValidateNumber()
	if err != nil {
		return err
	}
	err = c.ValidateMonth()
	if err != nil {
		return err
	}
	err = c.ValidateYear()
	if err != nil {
		return err
	}
	return nil
}

func (c *CreditCard) ValidateNumber() error {
	re := regexp.MustCompile(`^4011(78|79)|^43(1274|8935)|^45(1416|7393|763(1|2))|^50(4175|6699|67[0-6][0-9]|677[0-8]|9[0-8][0-9]{2}|99[0-8][0-9]|999[0-9])|^627780|^63(6297|6368|6369)|^65(0(0(3([1-3]|[5-9])|4([0-9])|5[0-1])|4(0[5-9]|[1-3][0-9]|8[5-9]|9[0-9])|5([0-2][0-9]|3[0-8]|4[1-9]|[5-8][0-9]|9[0-8])|7(0[0-9]|1[0-8]|2[0-7])|9(0[1-9]|[1-6][0-9]|7[0-8]))|16(5[2-9]|[6-7][0-9])|50(0[0-9]|1[0-9]|2[1-9]|[3-4][0-9]|5[0-8]))`)
	if !re.MatchString(c.number) {
		return errors.New("Invalid credit card number")
	}
	return nil
}

func (c *CreditCard) ValidateMonth() error {
	if c.expirationMonth < 1 || c.expirationMonth > 12 {
		return errors.New("Invalid expiration month")
	}
	return nil
}

func (c *CreditCard) ValidateYear() error {
	actualYear := time.Now().Year()
	if c.expirationYear < actualYear {
		return errors.New("Invalid expiration year")
	}
	return nil
}
