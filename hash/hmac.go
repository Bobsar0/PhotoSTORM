// Package hash hashes strings/remember tokens using HMAC hashing algorithm
// Rather than storing raw remember tokens in our database we are going to store the hash of each remember token
package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// HMAC is a wrapper around the crypto/hmac package making
// it a little easier to use in our code.
type HMAC struct {
	hmac hash.Hash
}

// NewHMAC creates and returns a new HMAC object
func NewHMAC(key string) HMAC {
	//create an HMAC hash that is backed by the secret key stored in the key variable, and we intend to use the SHA256 hashing function once the data is digitally signed. h is of type hash.Hash
	h := hmac.New(sha256.New, []byte(key)) 
	return HMAC{
		hmac: h,
	}
}

// Hash will hash the provided input string using HMAC with
// the secret key provided when the HMAC object was created
func (h HMAC) Hash(input string) string {
	h.hmac.Reset() //Ensures that the hash has been flushed of any previously written data
	h.hmac.Write([]byte(input))// Write our own data (input) to the hashing function (HMAC)
	hashedData := h.hmac.Sum(nil) //Request that HMAC calculate a hash value for us as []bytes
	return base64.URLEncoding.EncodeToString(hashedData) //base 64 encode our hashedData byte slices into strings.
}
