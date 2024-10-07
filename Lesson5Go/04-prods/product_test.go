package main

import (
	"testing"
)

func TestSmallProductPrice(t *testing.T) {
	product := newProduct(smallType, 100.0)
	expectedPrice := 100.0
	if product.Price() != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, product.Price())
	}
}

func TestMediumProductPrice(t *testing.T) {
	product := newProduct(mediumType, 100.0)
	expectedPrice := 100 + 100*.03
	if product.Price() != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, product.Price())
	}
}

func TestLargeProductPrice(t *testing.T) {
	product := newProduct(largeType, 100.0)
	expectedPrice := 100 + 100*.06 + 2500
	if product.Price() != expectedPrice {
		t.Errorf("Expected price %.2f, but got %.2f", expectedPrice, product.Price())
	}
}
