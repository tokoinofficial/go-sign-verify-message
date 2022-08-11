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

### Demo
- [x] golang
- [ ] python

### Created & Maintained By

[Trong Dinh](https://github.com/trongdth)
