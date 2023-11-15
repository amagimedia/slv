package vaults

import (
	"github.com/shibme/slv/core/commons"
	"github.com/shibme/slv/core/crypto"
)

func (vlt *Vault) putSecretWithoutCommit(secretName string, secretValue []byte) (err error) {
	if !secretNameRegex.MatchString(secretName) {
		return ErrInvalidSecretName
	}
	var sealedSecret *crypto.SealedSecret
	sealedSecret, err = vlt.Config.PublicKey.EncryptSecret(secretValue, vlt.Config.HashLength)
	if err == nil {
		if vlt.vault.Secrets == nil {
			vlt.vault.Secrets = make(map[string]*string)
		}
		vlt.vault.Secrets[secretName] = commons.String(sealedSecret.String())
	}
	return
}

func (vlt *Vault) PutSecret(secretName string, secretValue []byte) (err error) {
	if err = vlt.putSecretWithoutCommit(secretName, secretValue); err == nil {
		err = vlt.commit()
	}
	return
}

func (vlt *Vault) SecretExists(secretName string) (exists bool) {
	if vlt.vault.Secrets != nil {
		_, exists = vlt.vault.Secrets[secretName]
	}
	return exists
}

func (vlt *Vault) ListSecrets() []string {
	names := make([]string, 0, len(vlt.vault.Secrets))
	for name := range vlt.vault.Secrets {
		names = append(names, name)
	}
	return names
}

func (vlt *Vault) GetSecret(secretName string) (decryptedSecret []byte, err error) {
	if vlt.IsLocked() {
		return decryptedSecret, ErrVaultLocked
	}
	sealedSecretData := vlt.vault.Secrets[secretName]
	if sealedSecretData == nil {
		return nil, ErrVaultSecretNotFound
	}
	if decryptedSecret = vlt.getSecretFromCache(secretName); decryptedSecret == nil {
		sealedSecret := &crypto.SealedSecret{}
		if err = sealedSecret.FromString(*sealedSecretData); err == nil {
			if decryptedSecret, err = vlt.secretKey.DecryptSecret(*sealedSecret); err == nil {
				vlt.putSecretToCache(secretName, decryptedSecret)
			}
		}
	}
	return
}

func (vlt *Vault) DeleteSecret(secretName string) error {
	delete(vlt.vault.Secrets, secretName)
	vlt.deleteSecretFromCache(secretName)
	return vlt.commit()
}
