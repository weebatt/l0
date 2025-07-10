package cache

import (
	"l0/internal/models"
	"sync"

	"github.com/google/uuid"
)

type OrderCache struct {
	mu     sync.RWMutex
	orders map[uuid.UUID]*models.Order
}

func NewOrderCache() *OrderCache {
	return &OrderCache{
		orders: make(map[uuid.UUID]*models.Order),
	}
}

func (c *OrderCache) Set(order *models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.orders[order.OrderUID] = order
}

func (c *OrderCache) Get(id uuid.UUID) (*models.Order, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	order, ok := c.orders[id]
	return order, ok
}

func (c *OrderCache) Load(orders []*models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, order := range orders {
		c.orders[order.OrderUID] = order
	}
}
