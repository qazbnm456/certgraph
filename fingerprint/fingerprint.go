// Package fingerprint defines types to define a certificate fingerprint for certgraph
package fingerprint

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Fingerprint sha256 of certificate bytes
type Fingerprint [sha256.Size]byte

// HexString print Fingerprint as uppercase hex
func (fp *Fingerprint) HexString() string {
	return fmt.Sprintf("%X", *fp)
}

// FromHashBytes returns a Fingerprint generated by the first len(Fingerprint) bytes
func FromHashBytes(data []byte) Fingerprint {
	var fp Fingerprint
	// if len(data) != len(fp) {
	// 	// TODO this should error....
	// }
	for i := 0; i < len(data) && i < len(fp); i++ {
		fp[i] = data[i]
	}
	return fp
}

// FromRawCertBytes returns a Fingerprint generated by the provided bytes
func FromRawCertBytes(data []byte) Fingerprint {
	fp := sha256.Sum256(data)
	return fp
}

// FromB64Hash returns a Fingerprint from a base64 encoded hash string
func FromB64Hash(hash string) Fingerprint {
	data, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		panic(err)
	}
	return FromHashBytes(data)
}

// FromHexHash returns a Fingerprint from a hex encoded hash string
func FromHexHash(hash string) Fingerprint {
	decoded, err := hex.DecodeString(hash)
	if err != nil {
		panic(err)
	}
	return FromHashBytes(decoded)
}

// B64Encode returns the b64 string of a Fingerprint
func (fp *Fingerprint) B64Encode() string {
	return base64.StdEncoding.EncodeToString(fp[:])
}
