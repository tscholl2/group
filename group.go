package group

import (
	"crypto/aes"
	"crypto/cipher"
	"math/big"
)

type Element []byte

var n = new(big.Int)
var block cipher.Block
var G Element

func init() {
	n.Rsh(big.NewInt(1), 255)
	block, _ = aes.NewCipher([]byte("a joke should really fit in here"))
	G = encode(big.NewInt(1))
}

func ScaleGen(k *big.Int) Element {
	return Scale(G, k)
}

func Scale(c Element, k *big.Int) Element {
	a := decode(c)
	return encode(new(big.Int).Mul(a, k))
}

func encode(a *big.Int) Element {
	b := pad(a.Bytes())
	block.Encrypt(b, b)
	return b
}

func decode(c Element) *big.Int {
	block.Decrypt(c, c)
	return new(big.Int).SetBytes(c)
}

func pad(b []byte) []byte {
	if len(b) <= block.BlockSize() {
		return append(make([]byte, block.BlockSize()-len(b)), b...)
	}
	return append(make([]byte, block.BlockSize()-len(b)%block.BlockSize()), b...)
}
