package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	bigOne = big.NewInt(1)
)

func main() {
	abi, _ := abi.JSON(strings.NewReader(IERC1271ABI))

	inputs := abi.Methods["isValidSignature"].Inputs

	prefix, _ := hex.DecodeString("1626ba7e")

	var digest [32]byte
	_, _ = hex.Decode(digest[:], []byte("19bb34e293bba96bf0caeea54cdd3d2dad7fdf44cbea855173fa84534fcfb528"))

	signature := big.NewInt(0)

	start := time.Now()
	cnt := 0
	sha256 := sha256.New()

	for {
		calldata, err := inputs.Pack(digest, signature.Bytes())
		if err != nil {
			log.Fatal(fmt.Errorf("failed to pack calldata: %w", err))
		}

		sha256.Reset()
		sha256.Write(append(prefix, calldata...))

		if bytes.HasPrefix(sha256.Sum(nil), prefix) {
			fmt.Printf("Done in %s: %x\n", time.Since(start), signature.Bytes())
			break
		}

		cnt++
		signature.Add(signature, bigOne)

		if cnt%10_000_000 == 0 {
			fmt.Printf("%s: %d\n", time.Since(start), cnt)
		}
	}

}
