package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"

	uuid "github.com/satori/go.uuid"
)

func main() {
	for i := 0; i < 4000; i++ {
		uuid, _ := uuid.NewV4()
		uuidStr := uuid.String()

		sum := sha256.Sum256([]byte(uuidStr))
		s := fmt.Sprintf("%x", sum)

		n, _ := new(big.Int).SetString(s, 16)
		var z big.Int
		mod := z.Mod(n, big.NewInt(4000))
		log.Printf("uuid: %s, hash: %s, mod: %d\n", uuidStr, s, mod)
	}
}
