package service

import entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"

func CalculateCartTotal(cart *entities.Cart) float64 {
	var total float64 = 0
	for _, item := range cart.Items {
		total += float64(item.Quantity) * item.Price
	}
	return total
}
