package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	// order, err := NewOrder("", 100, 10)
	// if err == nil {
	// 	t.Error("Expected an error, but received nil")
	// }
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptNewOrder_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	_, err := NewOrder("", 0, 0)
	assert.Error(t, err)
}

func TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{
		ID:    "123",
		Price: 10.0,
		Tax:   2.0,
	}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenAPriceAndTax_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
