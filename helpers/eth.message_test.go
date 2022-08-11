package helpers

import (
	"fmt"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func ExampleSign() {
	types := []string{"address", "uint256"}
	values :=
		[]interface{}{
			"0x45f7967926e95fd161e56ed66b663c9114c5226f",
			"4685",
		}

	signature, _ := Sign("0x16870e97136d178735994d9ed537742276db9f30d80b2b35a05a798895679049", types, values)
	fmt.Println(signature)

	// Output:
	// 0x5e18c4ca19a2fcc6f740e5a7d834c10ebb883bba9891d5151c1d9e6e08a1ab216ce9a71c1673e2c5fdedb1e198fd6777175fe6d7534416cbbf4fbd11ef3614161b
}

func ExampleVerify() {
	types := []string{"address", "address", "uint256"}
	values :=
		[]interface{}{
			"0x45f7967926e95fd161e56ed66b663c9114c5226f",
			"0xa0F0546Eb5E3eE7e8cfC5DA12e5949F3AE622675",
			"4685",
		}

	signature, _ := Sign("0x16870e97136d178735994d9ed537742276db9f30d80b2b35a05a798895679049", types, values)
	fmt.Println(signature)

	hash := solsha3.SoliditySHA3(
		types,
		values,
	)
	fmt.Println(Verify("0x1111", signature, hash, false))
	fmt.Println(Verify("0xfE91b1E07b93fdfae1E02A985266fA3414aB844A", signature, hash, false))

	// Output:
	// 0x68a4ab4c439585f54aac6d16a51a2e9c83b5c0c5697b692aaac82c165326a2501600bfeb1690bb7f15d8e4821eb7fded5f2870903d7a126980132a63c50b90fd1b
	// false
	// true
}
