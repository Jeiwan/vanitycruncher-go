package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/holiman/uint256"
)

const (
	methodName = "isValidSignature"
)

var (
	uintOne = uint256.NewInt(1)
)

func main() {
	var err error

	abi, err := abi.JSON(strings.NewReader(IERC1271ABI))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse ABI: %w", err))
	}

	prefix := abi.Methods[methodName].ID
	digest := mustDecodeHash("19bb34e293bba96bf0caeea54cdd3d2dad7fdf44cbea855173fa84534fcfb528")
	signature := uint256.NewInt(0)

	var calldata []byte

	cnt := 0
	sha256 := sha256.New()
	start := time.Now()

	for !bytes.HasPrefix(sha256.Sum(nil), prefix) {
		calldata, err = abi.Pack(methodName, digest, signature.Bytes())
		if err != nil {
			log.Fatal(fmt.Errorf("failed to pack calldata: %w", err))
		}

		sha256.Reset()
		sha256.Write(calldata)

		cnt++
		signature.Add(signature, uintOne)

		if cnt%10_000_000 == 0 {
			fmt.Printf("%s: %d\n", time.Since(start), cnt)
		}
	}

	fmt.Printf("Done in %s: %x\n", time.Since(start), signature.Bytes())
}

func mustDecodeHash(s string) (hash [32]byte) {
	_, err := hex.Decode(hash[:], []byte(s))
	if err != nil {
		log.Fatal(err)
	}
	return
}
