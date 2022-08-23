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
	calldataSize = 32 + 32 + 32 + 32 // digest + dynamic array (position + length + data)
)

var (
	uintOne = uint256.NewInt(1)
)

func main() {
	var err error

	abi, _ := abi.JSON(strings.NewReader(IERC1271ABI))
	inputs := abi.Methods["isValidSignature"].Inputs

	prefix, _ := hex.DecodeString("1626ba7e")
	var digest [32]byte
	hex.Decode(digest[:], []byte("19bb34e293bba96bf0caeea54cdd3d2dad7fdf44cbea855173fa84534fcfb528"))

	signature := uint256.NewInt(0)

	data := make([]byte, len(prefix)+calldataSize)
	copy(data, prefix)

	start := time.Now()
	cnt := 0
	sha256 := sha256.New()
	calldata := make([]byte, calldataSize)

	for !bytes.HasPrefix(sha256.Sum(nil), prefix) {
		calldata, err = inputs.Pack(digest, signature.Bytes())
		if err != nil {
			log.Fatal(fmt.Errorf("failed to pack calldata: %w", err))
		}

		copy(data[len(prefix):], calldata)

		sha256.Reset()
		sha256.Write(data)

		cnt++
		signature.Add(signature, uintOne)

		if cnt%10_000_000 == 0 {
			fmt.Printf("%s: %d\n", time.Since(start), cnt)
		}
	}

	fmt.Printf("Done in %s: %x\n", time.Since(start), signature.Bytes())
}
