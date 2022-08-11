# sign-verify-message

This repos shows the way to sign/verify an ethereum message. It's used for backend of third party where interacts with Tokoin Payment System.

### Tokoin Payment System [WIP]

### How to use
- Sign message
```golang
func Sign(pk string, types []string, values []interface{}) (string, error) {
	privateKey, err := crypto.HexToECDSA(pk[2:])
	if err != nil {
		return "", err
	}

	hash := solsha3.SoliditySHA3(
		types,
		values,
	)

	signatureBytes, _ := crypto.Sign(SignHash(hash), privateKey)
	sig := hexutil.Encode(signatureBytes)
	signature := ""
	if sig[len(sig)-2:] == "00" {
		signature = sig[0:len(sig)-2] + "1b"
	} else {
		signature = sig[0:len(sig)-2] + "1c"
	}

	return signature, nil
}
```

- Verify message
```golang
func Verify(from, sigHex string, msg []byte, ignoreVersion bool) bool {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)

	if !ignoreVersion {
		// If the version(v) is correct return the signer address
		if sig[64] != 27 && sig[64] != 28 {
			return false
		}

		// Version(v) of signature should be 27 or 28, but 0 and 1 are also possible versions
		sig[64] -= 27
	}

	pubKey, err := crypto.SigToPub(SignHash(msg), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return fromAddr == recoveredAddr
}
```

### How to run sample
- go mod tidy
- go run main.go

### Demo
- [x] golang
- [x] python (https://github.com/tokoinofficial/python-sign-verify-message)

### Created & Maintained By

[Trong Dinh](https://github.com/trongdth)
