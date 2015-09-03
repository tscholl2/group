package group

import (
	"crypto/aes"
	"crypto/cipher"
	"math/big"
)

//Element represents an element of the group
type Element []byte

var n = new(big.Int)
var block cipher.Block

//G is the generator of the group
var G Element

func init() {
	n.Lsh(big.NewInt(1), 255)
	block, _ = aes.NewCipher([]byte("a joke should really fit in here"))
	G = encode(big.NewInt(1))
}

//Scale does repeated operation,
//Scale(c,k) = kc in additive or
//c^k in multiplicative notation.
func Scale(c Element, k *big.Int) Element {
	a := decode(c)
	return encode(new(big.Int).Mod(new(big.Int).Mul(a, k), n))
}

func encode(a *big.Int) Element {
	b := pad(a.Bytes())
	block.Encrypt(b, b)
	return b
}

func decode(c Element) *big.Int {
	b := pad(c)
	block.Decrypt(b, b)
	return new(big.Int).SetBytes(b)
}

func pad(b []byte) []byte {
	if len(b) <= block.BlockSize() {
		return append(make([]byte, block.BlockSize()-len(b)), b...)
	}
	return append(make([]byte, block.BlockSize()-len(b)%block.BlockSize()), b...)
}
