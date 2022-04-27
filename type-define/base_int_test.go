package typedefine_test

import (
	"testing"

	typedefine "github.com/okatti-et14/go-examples/type-define"
)

func TestPrice(t *testing.T) {
	i := 1
	if typedefine.Price(i) != typedefine.ApplePrice {
		t.Errorf("Price is %d, want 1", i)
	}
}

func TestPriceString(t *testing.T) {
	if typedefine.ApplePrice.String() != "リンゴ" {
		t.Errorf("Price is %d, want 1", typedefine.ApplePrice)
	}

	if typedefine.OrangePrice.String() != "オレンジ" {
		t.Errorf("Price is %d, want 2", typedefine.OrangePrice)
	}

	if typedefine.Price(3).String() != "その他" {
		t.Errorf("Price is %d, want 3", 3)
	}
}
