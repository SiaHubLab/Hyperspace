package compatibility

// siag.go checks that any changes made to the code retain compatibility with
// old versions of siag.

import (
	"errors"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/HyperspaceApp/Hyperspace/crypto"
	"github.com/HyperspaceApp/Hyperspace/encoding"
	"github.com/HyperspaceApp/Hyperspace/types"
)

// KeyPairSiag_1_0 matches the KeyPair struct of the siag 1.0 code.
type KeyPairSiag_1_0 struct {
	Header           string
	Version          string
	Index            int
	SecretKey        crypto.SecretKey
	UnlockConditions types.UnlockConditions
}

// TestVerifyKeysSiag_1_0 loads some keys generated by siag1.0.
// Verification must still work.
func TestVerifyKeysSiag_1_0(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	var kp KeyPairSiag_1_0

	// 1 of 1
	err := encoding.ReadFile("siag_1.0_1of1_Key0.siakey", &kp)
	if err != nil {
		t.Fatal(err)
	}
	err = verifyKeysSiag_1_0(kp.UnlockConditions, "", "siag_1.0_1of1")
	if err != nil {
		t.Fatal(err)
	}

	// 1 of 2
	err = encoding.ReadFile("siag_1.0_1of2_Key0.siakey", &kp)
	if err != nil {
		t.Fatal(err)
	}
	err = verifyKeysSiag_1_0(kp.UnlockConditions, "", "siag_1.0_1of2")
	if err != nil {
		t.Fatal(err)
	}

	// 2 of 3
	err = encoding.ReadFile("siag_1.0_2of3_Key0.siakey", &kp)
	if err != nil {
		t.Fatal(err)
	}
	err = verifyKeysSiag_1_0(kp.UnlockConditions, "", "siag_1.0_2of3")
	if err != nil {
		t.Fatal(err)
	}

	// 3 of 3
	err = encoding.ReadFile("siag_1.0_3of3_Key0.siakey", &kp)
	if err != nil {
		t.Fatal(err)
	}
	err = verifyKeysSiag_1_0(kp.UnlockConditions, "", "siag_1.0_3of3")
	if err != nil {
		t.Fatal(err)
	}

	// 4 of 9
	err = encoding.ReadFile("siag_1.0_4of9_Key0.siakey", &kp)
	if err != nil {
		t.Fatal(err)
	}
	err = verifyKeysSiag_1_0(kp.UnlockConditions, "", "siag_1.0_4of9")
	if err != nil {
		t.Fatal(err)
	}
}
