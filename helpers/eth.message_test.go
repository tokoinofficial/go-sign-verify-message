package helpers

import (
	"fmt"
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
	// defer func() {
	// 	if r := recover(); r == nil {
	// 		t.Errorf("The code did not panic")
	// 	}
	// }()

	// Sign("1", types, values)
}
