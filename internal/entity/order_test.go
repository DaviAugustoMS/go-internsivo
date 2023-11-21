package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlanck(t *testing.T) {
	order := Order{}

	assert.Error(t, order.Validate(), "ID is requeres")

	// if order.Validate() == nil {
	// 	t.Error("ID is requered")
	// }
}

func Test_If_It_Gets_An_Error_Id_Price_Is_Black(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func Test_If_It_Gets_An_Error_Id_Tax_Is_Black(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFilnalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
