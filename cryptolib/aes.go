package cryptolib

import (
	"crypto/aes"
	"crypto/rand"
)

func AESECBDecrypt(ciphertext, key []byte) []byte {
	out := make([]byte, 0)
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		out = append(out, aESDecrypt(ciphertext[i:i+aes.BlockSize], key)...)
	}
	return out
}

func AESECBEncrypt(plaintext, key []byte) []byte {
	out := make([]byte, 0)
	for i := 0; i < len(plaintext); i += aes.BlockSize {
		out = append(out, aESEncrypt(plaintext[i:i+aes.BlockSize], key)...)
	}
	return out
}

func aESDecrypt(block, key []byte) []byte {
	ret := make([]byte, aes.BlockSize)
	crypter, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	crypter.Decrypt(ret, block)
	return ret
}

func aESEncrypt(block, key []byte) []byte {
	ret := make([]byte, aes.BlockSize)
	crypter, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	crypter.Encrypt(ret, block)
	return ret
}

//In CBC mode, each ciphertext block is added to the next plaintext block
//before the next call to the cipher core.

//The first plaintext block, which has no associated previous ciphertext block,
//is added to a "fake 0th ciphertext block" called the initialization vector, or IV.

//Implement CBC mode by hand by taking the ECB function you wrote earlier,
//making it encrypt instead of decrypt (verify this by decrypting whatever you encrypt to test),
//and using your XOR function from the previous exercise to combine them.

//replaceBytes replaces the bytes in b1 at index with b2
func replaceBytes(b1, b2 []byte, index int) []byte {
	ret := b1[:index]
	ret = append(ret, b2...)
	ret = append(ret, b1[index+len(b2):]...)
	return ret
}

func AESCBCEncrypt(plaintext, key, iv []byte) []byte {
	out := make([]byte, 0)
	//do IV
	//plaintext = append(iv, plaintext...)

	for i, block := range chunker(plaintext, aes.BlockSize) {
		if i == 0 {
			//do IV
			//xor
			x := XorBytes(iv, block)
			//encrypt
			out = append(out, aESEncrypt(x, key)...)
		} else {
			//do previous block
		}
		//fmt.Println(block)
	}

	return out
}

func AESCBCDecrypt(ciphertext, key, iv []byte) []byte {
	out := make([]byte, 0)
	//work backwards (need last block)
	blocks := chunker(ciphertext, aes.BlockSize)
	for i := len(blocks) - 1; i >= 0; i-- {
		if i == 0 {
			//fmt.Println("asdf")
			//decrypt block
			b := aESDecrypt(blocks[i], key)
			//xor against previous block
			x := XorBytes(b, iv)
			out = append(x, out...)
		} else {
			//fmt.Println("asdf")
			//decrypt block
			b := aESDecrypt(blocks[i], key)
			//xor against previous block
			x := XorBytes(b, blocks[i-1])
			out = append(x, out...)
		}
	}

	return out
}

func RandomKey() []byte {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
