## Golang cryptography library 

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## Features/Functionality

- **Library that provides cryptographic functions and algorithms, 
including encryption, decryption, hashing, and digital signatures, generation**
  
- **Encryption**:
  
	- **EncryptAES(key []byte, plaintext []byte) ([]byte, error)**
   
	    - Encrypts the provided plaintext using the AES encryption algorithm.
	    - Utilizes AES in CBC mode with a random initialization vector (IV).
       
	- **EncryptDES(key []byte, plaintext []byte) ([]byte, error)**
   
		- Encrypts the provided plaintext using the DES encryption algorithm.
		- Supports DES encryption for data blocks of the correct size.
    
	- **Encrypt3DES(key []byte, plaintext []byte) ([]byte, error**)
   
	    - Encrypts the provided plaintext using the 3DES (Triple DES) encryption algorithm.
	    - Requires a 24-byte (192-bit) key.

- **Decryption**:
  
	- **DecryptAES(key []byte, ciphertext []byte) ([]byte, error)**
   
		- Decrypts AES-encrypted ciphertext using the provided key.
    
	- **DecryptDES(key []byte, ciphertext []byte) ([]byte, error)**
   
		- Decrypts DES-encrypted ciphertext using the provided key.
    
	- **Decrypt3DES(key []byte, ciphertext []byte) ([]byte, error)**
   
		- Decrypts 3DES (Triple DES)-encrypted ciphertext using the provided key.

- **Hashing**:
  
	- **CalculateSHA256(data []byte) []byte**
   
		- Calculates the SHA-256 hash of the input data.
    
	- **CalculateSHA224(data []byte) []byte**
   
		- Calculates the SHA-224 hash of the input data.
    
	- **CalculateSHA3_256(data []byte) []byte**
   
		- Calculates the SHA3-256 hash of the input data.
    
	- **CalculateSHA1(data []byte) []byte**
   
		- Calculates the SHA-1 hash of the input data.
    
	- **CalculateMD5(data []byte) []byte**
   
		- Valculates the MD5 hash of the input data.
    
	- **CalculateSHA512(data []byte) []byte**
   
		- Valculates the SHA-512 hash of the input data.
    
	- **CalculateSHA384(data []byte) []byte**
   
		- Calculates the SHA-384 hash of the input data.

- **Randomization**:
  
    - **GenerateRandomBytes(length int) ([]byte, error)**
      
		- Generates a specified number of random bytes.
		- Uses the crypto/rand package for secure random byte generation.

	- **GenerateRandomHex(length int) (string, error)**
   
		- Generates a random hexadecimal string of the specified length.
		- Utilizes GenerateRandomBytes internally and encodes the result in hexadecimal.

- **Digital signatures**:

	- **GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)****
 
	    - Generates an RSA key pair for digital signatures with the specified number of bits.
	    - Returns a private key and its corresponding public key.
       
	- **SignDataWithRSA(data []byte, privateKey *rsa.PrivateKey) ([]byte, error)***
   
		- Signs a byte array of data using an RSA private key.
		- Uses the PKCS1v15 signature scheme with SHA-256 hash.
    
	- **VerifySignatureWithRSA(data, signature []byte, publicKey *rsa.PublicKey) error***
   
	    - Verifies a digital signature using an RSA public key.
	    - Validates the signature against the data using the PKCS1v15 signature scheme and SHA-256 hash.
       
	- **ExportRSAPrivateKeyToPEM(privateKey *rsa.PrivateKey) (string, error)***
   
		- Exports an RSA private key to a PEM-encoded string.
    
	 - **ExportRSAPublicKeyToPEM(publicKey *rsa.PublicKey) (string, error)***
    
		- Exports an RSA public key to a PEM-encoded string.
    
	 - **ImportRSAPrivateKeyFromPEM(keyPEM string) (*rsa.PrivateKey, error)***
    
		- Imports an RSA private key from a PEM-encoded string.
    
	- **ImportRSAPublicKeyFromPEM(keyPEM string) (*rsa.PublicKey, error)***
   
		- Imports an RSA public key from a PEM-encoded string.
 
## Project structure:

```go
crypty
│
├── decrypt
│   └── decrypt.go
├── digital_signatures
│   └── digital_signatures.go
├── encrypt
│   └── encrypt.go
├── go.mod
├── go.sum
├── hashing
│   └── hashing.go
├── random
│   └── random.go
└── README.md
```

## Installation

```sh
go get github.com/kenjitheman/crypty
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

## License

- [MIT](./LICENSE)
