package group

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math/big"
)

//Element is a generic group element
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

func decode(h Element) *big.Int {
	block.Decrypt(h, h)
	return new(big.Int).SetBytes(h)
}

func encode(a *big.Int) Element {
	b := pad(new(big.Int).Mod(a, n).Bytes())
	block.Encrypt(b, b)
	return b
}

func (h Element) Print() string {
	fmt.Println(h)
	return hex.EncodeToString(h)
}

func Load(s string) (Element, error) {
	if len(s)%2 == 1 {
		s = "0" + s
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return encode(decode(b)), nil
}

//Scale does repeated operation,
//Scale(c,k) = kc in additive or
//c^k in multiplicative notation.
func (h Element) Scale(k *big.Int) Element {
	return encode(new(big.Int).Mod(new(big.Int).Mul(decode(h), k), n))
}

func pad(b []byte) []byte {
	if len(b) > 32 {
		return append(make([]byte, 32-len(b)%32), b...)
	}
	if len(b) < 32 {
		return append(make([]byte, 32-len(b)), b...)
	}
	return b
}
