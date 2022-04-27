package typedefine

type Price int

const (
	ApplePrice  Price = 1
	OrangePrice Price = 2
)

func (p Price) String() string {
	switch p {
	case ApplePrice:
		return "リンゴ"
	case OrangePrice:
		return "オレンジ"
	default:
		return "その他"
	}
}
