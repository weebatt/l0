package validation

import (
	"errors"
	"fmt"
	"l0/internal/models"
	"regexp"
)

var (
	emailRegexp = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	phoneRegexp = regexp.MustCompile(`^\+?\d{7,15}$`)
	trackRegexp = regexp.MustCompile(`^WBIL\d{6}TRACK$`)
)

/*
Я честно не совсем понимаю, что можно было бы отвалидировать кроме полей email,
phone и track_number, поэтому встал вопрос о том что валидировать в таблице
payment. Посему решил просто оставить валидацию на currency и все)

Возможно стоило бы проверять поля на дефолтные значения nil/0/"", но как будто бы
это вышло бы достаточно громоздко, так что решил не делать!
*/

func ValidateOrder(order *models.Order) error {
	if order.OrderUID == [16]byte{} {
		return errors.New("order_uid is empty")
	}
	if !trackRegexp.MatchString(order.TrackNumber) {
		return errors.New("track_number must match WBILddddddTRACK")
	}

	return nil
}

func ValidateDelivery(delivery *models.Delivery) error {
	if delivery.Name == "" {
		return errors.New("delivery name is empty")
	}
	if !emailRegexp.MatchString(delivery.Email) {
		return errors.New("invalid email format")
	}
	if !phoneRegexp.MatchString(delivery.Phone) {
		return errors.New("invalid phone format")
	}

	return nil
}

func ValidatePayment(payment *models.Payment) error {
	if payment.Currency == "" {
		return errors.New("payment currency is empty")
	}
	return nil
}

func ValidateItems(items []models.Item) error {
	if len(items) == 0 {
		return errors.New("items list is empty")
	}

	for i, item := range items {
		if !trackRegexp.MatchString(item.TrackNumber) {
			return errors.New(fmt.Sprintf("item track_number invalid at index %d", i))
		}
	}
	return nil
}
