package Set2

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net/url"

	"github.com/c-sto/gocryptopals/pkg/aes"
	"github.com/c-sto/gocryptopals/pkg/cryptobytes"
)

/*
ECB cut-and-paste
Write a k=v parsing routine, as if for a structured cookie. The routine should take:

foo=bar&baz=qux&zap=zazzle
... and produce:

{
  foo: 'bar',
  baz: 'qux',
  zap: 'zazzle'
}
(you know, the object; I don't care if you convert it to JSON).

Now write a function that encodes a user profile in that format, given an email address. You should have something like:

profile_for("foo@bar.com")
... and it should produce:

{
  email: 'foo@bar.com',
  uid: 10,
  role: 'user'
}
... encoded as:

email=foo@bar.com&uid=10&role=user
Your "profile_for" function should not allow encoding metacharacters (& and =). Eat them, quote them, whatever you want to do, but don't let people set their email address to "foo@bar.com&role=admin".

Now, two more easy functions. Generate a random AES key, then:

Encrypt the encoded user profile under the key; "provide" that to the "attacker".
Decrypt the encoded user profile and parse it.
Using only the user input to profile_for() (as an oracle to generate "valid" ciphertexts) and the ciphertexts themselves, make a role=admin profile.
*/
var key = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}

func Challenge13() {
	key = aes.RandomKey()
	//the goal of this challenge is to turn an encrypted profile blob into a 'valid' decryption that sets the user's password as 'admin'
	c := profile_for("te@test.net")
	fmt.Println(isAdmin(c))
	//set is up so the ECB encoding makes the cookie look like(0 in email is a null):
	//email=aaaaaaaaaaadmin&uid=10&role=admin
	//|--------------||--------------||--------------||--------------|
	//take the second encrypted block, and replace the last block in the 'te@test.net' email:
	//email=te%40test.net&uid=10&role=user
	//|--------------||--------------||--------------||--------------|
	//which makes it:
	//email=te%40test.net&uid=10&role=admin&uid=10&rol
	//|--------------||--------------||--------------||--------------|
	pad := profile_for("aaaaaaaaaaadmin")
	decodedPad, e := hex.DecodeString(pad)
	if e != nil {
		panic(e)
	}
	decodedC, e := hex.DecodeString(c)
	if e != nil {
		panic(e)
	}
	chunkPad := cryptobytes.Chunker(decodedPad, 16)
	chunkC := cryptobytes.Chunker(decodedC, 16)
	chunkC[2] = chunkPad[1]

	forgedString := hex.EncodeToString(bytes.Join(chunkC, nil))

	fmt.Println(cryptobytes.Chunker(decodedPad, 16))
	fmt.Println(isAdmin(forgedString))

}

//profile_for should take an email address, and return an encrypted profile blob
func profile_for(email string) string {
	s := "email=%s&uid=10&role=user"
	encoded := url.QueryEscape(email)
	r := fmt.Sprintf(s, encoded)
	v := aes.AESECBEncrypt([]byte(r), key)
	r = hex.EncodeToString(v)
	return r
}

//isAdmin takes an encrypted string (in hex) and determines if the user is an admin or note (simulates a web server cookie or somesuch)
func isAdmin(s string) bool {
	v, e := parseCookie(s)
	if e != nil {
		panic(e)
	}
	if v["role"] == nil {
		return false
	}
	if v["role"][0] == "admin" {
		return true
	}
	return false
}

func parseCookie(s string) (values url.Values, err error) {
	b, e := hex.DecodeString(s)
	if e != nil {
		return nil, e
	}
	p := aes.AESECBDecrypt(b, key)
	return url.ParseQuery(string(p))
}
