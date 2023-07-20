package main

import (
	"fmt"

	"github.com/tokoinofficial/go-sign-verify-message/helpers"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func main() {
	pk := "0x16870e97136d178735994d9ed537742276db9f30d80b2b35a05a798895679049"
	// sign message
	types := []string{"address", "uint256"}
	values :=
		[]interface{}{
			"0x45f7967926e95fd161e56ed66b663c9114c5226f",
			"4685",
		}
	signature, _ := helpers.Sign(pk, types, values)

	// verify message
	hash := solsha3.SoliditySHA3(
		types,
		values,
	)
	fmt.Println(helpers.Verify("0xfE91b1E07b93fdfae1E02A985266fA3414aB844A", signature, hash, false))
}
