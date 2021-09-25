package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

const (
	testKey     string = "30770201010420af8236f6ba7fe69343da47fab5bc7086e1098c0b660fd75adc4945800fcdda26a00a06082a8648ce3d030107a14403420004dcedbaa357eb97f30e5beccf11e0187ce88de09795a7925519b175237f087cc9a3b65f1edda29962ba1a08889872e978c1844de5d98fe01e2cdf8dd59600a7d7"
	testPayload string = "0011b0cf0838c0c873965a2bfe6152833c8ee30ea21d7d0c4d4d75eb4d89b8a5"
	testSign    string = "6c2d9c3c1209b4dbc23534f9c29d597933fef6c30b616f03036b1709f7ac13a435b762d3eb508bbb68844b250bb753824f1449771ac3901e2069c87e698afc26"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = addrFromKey(key)
	return w
}

func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()
}

func (fakeLayer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}
	tests := []test{
		{testPayload, true},
		{"1058d1e8d78697b7577d96d80051d3268757fc688fcf058e7eb8c4e0718340cb", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSign, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("x")
	if err == nil {
		t.Error("restoreBigInts() should return error when payload is not hex")
	}
}

func TestWallet(t *testing.T) {
	t.Run("New Wallet is created", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return false },
		}
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
	t.Run("Wallet is restored", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return true },
		}
		w = nil
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
}
