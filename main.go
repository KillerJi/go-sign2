package main

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// ...
func main() {
	// Replace this with the address of the user's wallet
	// 548364400416034343698204186575808495617

	// var prime1, _ = math.ParseBig256("548364400416034343698204186575808495617")
	// var prime1, boo = math.ParseBig256("54836440041603434369820418575808495611")
	// if !boo {
	// 	fmt.Print("parse big256 error")
	// }
	var prime1, boo = math.ParseBig256("125000000")
	if !boo {
		fmt.Print("parse big256 error")
	}
	var number = math.HexOrDecimal256(*prime1)
	var prime2, boo1 = math.ParseBig256("1")
	if !boo1 {
		fmt.Print("parse big256 error")
	}
	var nonce = math.HexOrDecimal256(*prime2)
	salt := "0xf78d4feD6467ca6AFCB579E409552433152F88E2"
	token := "0x63a03409E9cfC5983E30568F384000E3d4dC1357"
	account := "0x004ec07d2329997267Ec62b4166639513386F32E"
	privateKey, er2 := crypto.HexToECDSA("13962cc606545b8a706ee4fad4ccf6cfd21add41e24f4c9abd667ceeaa0a74aa")
	// privateKey, er2 := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if er2 != nil {
		fmt.Print(er2)
	}
	signerData := apitypes.TypedData{
		Types: apitypes.Types{
			"Claim": []apitypes.Type{
				{Name: "token", Type: "address"},
				{Name: "account", Type: "address"},
				{Name: "number", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "Claim",
		Domain: apitypes.TypedDataDomain{
			Name:              "XP721",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(123),
			VerifyingContract: salt,
		},
		Message: apitypes.TypedDataMessage{
			"token":   token,
			"account": account,
			"number":  &number,
			"nonce":   &nonce,
		},
	}

	domainSeparator, err1 := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err1 != nil {
		fmt.Print(err1)
	}
	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		fmt.Print("err", err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := crypto.Keccak256Hash(rawData)
	// fmt.Print(hash)
	signature, err3 := crypto.Sign(hash.Bytes(), privateKey)
	if err3 != nil {
		fmt.Print(err3)
	}

	a := hexutil.Encode(signature)
	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62
	r := fmt.Sprintf(a[2:66])
	fmt.Print("0x", r, "\n")
	s := fmt.Sprintf(a[66:130])
	fmt.Print("0x", s, "\n")
	v := fmt.Sprintf(a[130:132])
	aa, err := strconv.ParseInt(v, 10, 64)
	fmt.Print(aa+27, "\n")
	// fmt.Print(a[66:130], "\n")
	// fmt.Print(a[130:132], "\n")

	pubKeyRaw, err := crypto.Ecrecover(hash.Bytes(), signature)
	pubKey, err33 := crypto.UnmarshalPubkey(pubKeyRaw)
	if err33 != nil {
		fmt.Print(err33)
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Print("yanzheng ", recoveredAddr, "\n")
}
