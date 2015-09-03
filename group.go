package group

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"math/big"
)

type Element struct {
	a    *big.Int //up to 32 bytes
	size uint8    //1 byte, number of bytes a has
}

func NewElement(b []byte) (e Element, err error) {
	if len(b)%block.BlockSize() != 0 {
		err = errors.New("wrong size")
		return
	}

	//#TODO fix

	return
}

var n = new(big.Int)
var block cipher.Block
var size = 32 //number of bytes per element

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
	if len(b) <= size {
		return append(make([]byte, size-len(b)), b...)
	}
	return append(make([]byte, size-len(b)%block.BlockSize()), b...)
}
