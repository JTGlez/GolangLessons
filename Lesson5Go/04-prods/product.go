package main

/*
Ejercicio 2 - Productos

Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.

La empresa tiene 3 tipos de productos: Small, Medium y Large. (Se espera que sean muchos más)


Y los costos adicionales son:

Small: solo tiene el costo del producto
Medium: el precio del producto + un 3% de mantenerlo en la tienda
Large: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.

El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.


Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Product que tenga el método Price.


Se debe poder ejecutar el método Price y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga

*/

const (
	smallType  = "Small"
	mediumType = "Medium"
	largeType  = "Large"
)

type Product interface {
	Price() float64
}

type Small struct {
	price float64
}

func (s Small) Price() float64 {
	return s.price
}

type Medium struct {
	price float64
	fee   float64
}

func (m Medium) Price() float64 {
	return m.price + m.fee*m.price
}

type Large struct {
	price    float64
	fee      float64
	shipment float64
}

func (l Large) Price() float64 {
	return l.price + l.fee*l.price + l.shipment
}

func newProduct(prodType string, price float64) Product {

	switch prodType {
	case smallType:
		return Small{
			price: price,
		}
	case mediumType:
		return Medium{
			price: price,
			fee:   0.03,
		}
	case largeType:
		return Large{
			price:    price,
			fee:      0.06,
			shipment: 2500,
		}
	}

	return nil
}
