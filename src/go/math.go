package liquidation

import "math/big"

func cloneBig(x *big.Int) *big.Int {
	if x == nil {
		return big.NewInt(0)
	}
	return new(big.Int).Set(x)
}

func bigFromUint64(v uint64) *big.Int {
	return new(big.Int).SetUint64(v)
}

func assetUnit(decimals uint8) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
}

func percentMul(value, percentage *big.Int) *big.Int {
	if value == nil || percentage == nil || value.Sign() == 0 || percentage.Sign() == 0 {
		return big.NewInt(0)
	}
	numer := new(big.Int).Mul(value, percentage)
	numer.Add(numer, halfPercentageFactor)
	return numer.Div(numer, percentageFactor)
}

func percentDiv(value, percentage *big.Int) *big.Int {
	if percentage == nil || percentage.Sign() == 0 {
		return big.NewInt(0)
	}
	numer := new(big.Int).Mul(value, percentageFactor)
	numer.Add(numer, new(big.Int).Div(percentage, big.NewInt(2)))
	return numer.Div(numer, percentage)
}

func wadDiv(a, b *big.Int) *big.Int {
	if b == nil || b.Sign() == 0 {
		return big.NewInt(0)
	}
	numer := new(big.Int).Mul(a, wad)
	numer.Add(numer, new(big.Int).Div(b, big.NewInt(2)))
	return numer.Div(numer, b)
}

func rayMul(a, b *big.Int) *big.Int {
	if a == nil || b == nil || a.Sign() == 0 || b.Sign() == 0 {
		return big.NewInt(0)
	}
	numer := new(big.Int).Mul(a, b)
	numer.Add(numer, halfRay)
	return numer.Div(numer, ray)
}

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(cloneBig(a), cloneBig(b))
}

func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(cloneBig(a), cloneBig(b))
}

func mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(cloneBig(a), cloneBig(b))
}

func div(a, b *big.Int) *big.Int {
	if b == nil || b.Sign() == 0 {
		return big.NewInt(0)
	}
	return new(big.Int).Div(cloneBig(a), cloneBig(b))
}

func minBig(a, b *big.Int) *big.Int {
	if cloneBig(a).Cmp(cloneBig(b)) <= 0 {
		return cloneBig(a)
	}
	return cloneBig(b)
}

func maxBig(a, b *big.Int) *big.Int {
	if cloneBig(a).Cmp(cloneBig(b)) >= 0 {
		return cloneBig(a)
	}
	return cloneBig(b)
}

func isZero(i *big.Int) bool {
	return i == nil || i.Sign() == 0
}
