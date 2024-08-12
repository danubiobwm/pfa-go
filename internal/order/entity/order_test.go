package entity_test

import (
	"testing"

	"github.com/danubiobwm/pfa-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGiveAnEmptyId_WhenCreateANewOrder_ThenSouldReceiveAndError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), " invalid id")
}

func TestGiveAnEmptyPrice_WhenCreateANewOrder_ThenSouldReceiveAndError(t *testing.T) {
	order := entity.Order{ID: "123"}
	assert.Error(t, order.IsValid(), " invalid price")
}

func TestGiveAnEmptyTax_WhenCreateANewOrder_ThenSouldReceiveAndError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 12.0}
	assert.Error(t, order.IsValid(), " invalid tax")
}

func TestGiveAValidParams_WhenCallNewOrder_ThenShould_ReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 12.0, 0.5)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 12.0, order.Price)
	assert.Equal(t, 0.5, order.Tax)
}

func TestGiveAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)
	err = order.CalculteFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)
}
