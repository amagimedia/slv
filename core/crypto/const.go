package crypto

import (
	"errors"

	"dev.shib.me/xipher"
)

const (
	publicKeyAbbrev    = "PK" // PK = Public Key
	secretKeyAbbrev    = "SK" // SK = Secret Key
	wrappedKeyAbbrev   = "WK" // WK = Wrapped Key
	sealedSecretAbbrev = "SS" // SS = Sealed Secret

	publicKeyLength     = xipher.PublicKeyLength + 3
	secretKeyLength     = xipher.PrivateKeyLength + 3
	cipherTextMinLength = publicKeyLength + 2

	hashMaxLength = 4
)

var (
	errUnsupportedCryptoVersion = errors.New("unsupported cryptography version")
	errGeneratingSecretKey      = errors.New("error generating a new secret key")
	errDerivingPublicKey        = errors.New("error deriving public key from the secret key")
	errInvalidPublicKeyFormat   = errors.New("invalid public key format")
	errInvalidSecretKeyFormat   = errors.New("invalid secret key format")
	errEncryptionFailed         = errors.New("encryption failed")
	errDecryptionFailed         = errors.New("decryption failed")
	errSecretKeyMismatch        = errors.New("given secret key cannot decrypt the data")
	errInvalidCiphertextFormat  = errors.New("invalid ciphertext format")
)
