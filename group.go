package group

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"math/big"
)

//Element is a generic group element
type Element [32]byte

var n = new(big.Int)
var block cipher.Block

//G is the generator of the group
var G Element

func init() {
	n.Lsh(big.NewInt(1), 255)
	block, _ = aes.NewCipher([]byte("a joke should really fit in here"))
	G = encode(big.NewInt(1))
}

func decode(h Element) *big.Int {
	block.Decrypt(h[:], h[:])
	block.Decrypt(h[16:], h[16:])
	return new(big.Int).SetBytes(h[:])
}

func encode(a *big.Int) Element {
	b := pad(new(big.Int).Mod(a, n).Bytes())
	block.Encrypt(b[:], b[:])
	block.Encrypt(b[16:], b[16:])
	return b
}

func (h Element) Print() string {
	return hex.EncodeToString(h[:])
}

func Load(s string) (h Element, err error) {
	if len(s) > 32 {
		err = errors.New("invalid string")
	}
	for len(s) < 32 {
		s = "0" + s
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return
	}
	return pad(b), nil
}

//Scale does repeated operation,
//Scale(h,k) = kh in additive or
//h^k in multiplicative notation.
func (h Element) Scale(k *big.Int) Element {
	return encode(new(big.Int).Mod(new(big.Int).Mul(decode(h), k), n))
}

func pad(b []byte) [32]byte {
	if len(b) > 32 {
		panic("too many bytes!")
	}
	if len(b) < 32 {
		b = append(make([]byte, 32-len(b)), b...)
	}
	var h [32]byte
	for i := 0; i < 32; i++ {
		h[i] = b[i]
	}
	return h
}
