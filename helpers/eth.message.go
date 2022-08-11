package helpers

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// https://ethereum.stackexchange.com/questions/76810/sign-message-with-web3-and-verify-with-openzeppelin-solidity-ecdsa-sol
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

func SignHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

// https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
// https://developer.cargox.digital/examples/signing_with_go.html
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
