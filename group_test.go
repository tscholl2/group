package group

import (
	"math/big"
	"testing"
)

func TestPad(t *testing.T) {
	b := block.BlockSize()
	a := big.NewInt(1)
	m := len(pad(a.Bytes()))
	if m%b != 0 {
		t.Errorf("Expected %d got %d", b, m)
	}
	a.Lsh(a, 255)
	m = len(pad(a.Bytes()))
	if m%b != 0 {
		t.Errorf("Expected %d got %d", b, m)
	}
	a.Lsh(a, 35)
	m = len(pad(a.Bytes()))
	if m%b != 0 {
		t.Errorf("Expected %d got %d", b, m)
	}
	c := new(big.Int).SetBytes(pad(a.Bytes()))
	if c.Cmp(a) != 0 {
		t.Errorf("padded number not the same as original")
	}
}

func TestEncodeDecode(t *testing.T) {
	tests := []int64{1, 7, 100, 14324}
	for _, i := range tests {
		a := big.NewInt(i)
		E := encode(a)
		b := decode(E)
		if a.Cmp(b) != 0 {
			t.Errorf("Expected %d got %d", a, b)
		}
	}
}

func TestEncode(t *testing.T) {
	s := "f70b69f4a26bac8c740912424b6dcd71"
	h, err := Load(s)
	if err != nil {
		t.Errorf("expected nil got %s", err)
	}
	if h.Print() != s {
		t.Errorf("expected %s, got %s", s, h.Print())
	}
	if decode(h).Cmp(big.NewInt(1)) != 0 {
		t.Errorf("decoded G and got %d instead of 1", decode(h))
	}
	if s != G.Print() {
		t.Errorf("encoded G and got %s instead of %s", G.Print(), s)
	}
}

func TestScale(t *testing.T) {
	tests := []struct {
		a *big.Int
		k *big.Int
		b *big.Int
	}{
		{big.NewInt(1), big.NewInt(2), big.NewInt(2)},
		{big.NewInt(1), big.NewInt(2), big.NewInt(2)},
		{big.NewInt(112), big.NewInt(2), big.NewInt(224)},
		{big.NewInt(7), big.NewInt(3), big.NewInt(21)},
	}
	for _, c := range tests {
		x := decode(encode(c.a).Scale(c.k))
		if x.Cmp(c.b) != 0 {
			t.Errorf("expected %d got %d", c.b, x)
		}
	}
}
