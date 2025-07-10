package faker

import (
	"fmt"
	"l0/internal/models"
	"math/rand"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func GenerateOrder() *models.Order {
	orderUID := uuid.New()
	trackNumber := fmt.Sprintf("WBIL%dTRACK", rand.Intn(900000)+100000)

	delivery := models.Delivery{
		OrderUID: orderUID,
		Name:     gofakeit.Name(),
		Phone:    gofakeit.Phone(),
		Zip:      gofakeit.Zip(),
		City:     gofakeit.City(),
		Address:  gofakeit.Street(),
		Region:   gofakeit.State(),
		Email:    gofakeit.Email(),
	}

	goodsTotal := rand.Intn(901) + 100
	deliveryCost := rand.Intn(451) + 50
	amount := goodsTotal + deliveryCost
	payment := models.Payment{
		Transaction:  orderUID,
		RequestID:    "",
		Currency:     gofakeit.RandomString([]string{"USD", "EUR", "RUB"}),
		Provider:     gofakeit.RandomString([]string{"wbpay", "stripe", "paypal"}),
		Amount:       amount,
		PaymentDt:    time.Now().Add(-time.Duration(rand.Intn(365)) * 24 * time.Hour).Unix(),
		Bank:         gofakeit.RandomString([]string{"alpha", "sber", "tinkoff"}),
		DeliveryCost: deliveryCost,
		GoodsTotal:   goodsTotal,
		CustomFee:    rand.Intn(101),
	}

	items := []models.Item{}
	for i := 0; i < rand.Intn(3)+1; i++ {
		price := rand.Intn(401) + 100
		sale := rand.Intn(51)
		totalPrice := int(float64(price) * (1 - float64(sale)/100))
		status := gofakeit.RandomString([]string{"101", "202", "303"})
		intStatus, _ := strconv.Atoi(status)

		items = append(items, models.Item{
			OrderUID:    orderUID,
			ChrtID:      rand.Intn(9000000) + 1000000,
			TrackNumber: trackNumber,
			Price:       price,
			Rid:         uuid.New().String(),
			Name:        gofakeit.RandomString([]string{"Mascaras", "Lipstick", "Shampoo", "Perfume"}),
			Sale:        sale,
			Size:        gofakeit.RandomString([]string{"0", "S", "M", "L"}),
			TotalPrice:  totalPrice,
			NmID:        rand.Intn(9000000) + 1000000,
			Brand:       gofakeit.RandomString([]string{"Vivienne Sabo", "L'Oréal", "Nivea"}),
			Status:      intStatus,
		})
	}

	return &models.Order{
		OrderUID:          orderUID,
		TrackNumber:       trackNumber,
		Entry:             "WBIL",
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            gofakeit.RandomString([]string{"en", "ru", "fr"}),
		InternalSignature: "",
		CustomerID:        fmt.Sprintf("cust_%d", rand.Intn(9000)+1000),
		DeliveryService:   gofakeit.RandomString([]string{"meest", "dhl", "fedex"}),
		ShardKey:          fmt.Sprintf("%d", rand.Intn(9)+1),
		SmID:              rand.Intn(100) + 1,
		DateCreated:       time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		OofShard:          fmt.Sprintf("%d", rand.Intn(5)+1),
	}
}
