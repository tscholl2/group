package main

import (
	"encoding/hex"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/tscholl2/group"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dir := strings.Split(r.URL.Path, "/")
	if len(dir) < 3 || dir[1] != "scale" || len(dir) > 4 {
		w.Write([]byte("hello"))
		return
	}
	switch len(dir) {
	case 3:
		b, err := hex.DecodeString(dir[2])
		if err != nil {
			w.Write([]byte("err"))
			log.Printf("Err: %s", err)
			return
		}
		a := new(big.Int).SetBytes(b)
		s := hex.EncodeToString(group.Scale(group.G, a))
		w.Write([]byte(s))
		return
	case 4:
		a, err := hex.DecodeString(dir[2])
		if err != nil {
			w.Write([]byte("err"))
			log.Printf("Err: %s", err)
			return
		}
		k, err := hex.DecodeString(dir[3])
		if err != nil {
			w.Write([]byte("err"))
			log.Printf("Err: %s", err)
			return
		}
		s := hex.EncodeToString(group.Scale(a, new(big.Int).SetBytes(k)))
		w.Write([]byte(s))
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServeTLS(":8889", "certificate.pem", "key.pem", mux))
}
