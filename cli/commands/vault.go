package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/shibme/slv/core/crypto"
	"github.com/shibme/slv/core/vaults"
	"github.com/spf13/cobra"
)

func vaultCommand() *cobra.Command {
	if vaultCmd != nil {
		return vaultCmd
	}
	vaultCmd = &cobra.Command{
		Use:   "vault",
		Short: "Vault operations",
		Long:  `Vault operations in SLV`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	vaultCmd.AddCommand(vaultNewCommand())
	vaultCmd.AddCommand(vaultShareCommand())
	return vaultCmd
}

func vaultNewCommand() *cobra.Command {
	if vaultNewCmd != nil {
		return vaultNewCmd
	}
	vaultNewCmd = &cobra.Command{
		Use:   "new",
		Short: "Creates a new vault",
		Run: func(cmd *cobra.Command, args []string) {
			vaultFile := cmd.Flag("file").Value.String()
			publicKeyStrings, err := cmd.Flags().GetStringSlice("public-keys")
			if err != nil {
				PrintErrorAndExit(err)
			}
			var publicKeys []crypto.PublicKey
			for _, publicKeyString := range publicKeyStrings {
				publicKey, err := crypto.PublicKeyFromString(publicKeyString)
				if err != nil {
					PrintErrorAndExit(err)
				}
				publicKeys = append(publicKeys, *publicKey)
			}
			_, err = vaults.New(vaultFile, publicKeys...)
			if err != nil {
				PrintErrorAndExit(err)
			}
			fmt.Println("Created vault:", color.GreenString(vaultFile))
			os.Exit(0)
		},
	}
	vaultNewCmd.Flags().StringP("file", "f", "", "Name of the environment")
	vaultNewCmd.Flags().StringSliceP("public-keys", "k", []string{}, "Public keys of environments or groups that can access the vault")
	vaultNewCmd.MarkFlagRequired("file")
	vaultNewCmd.MarkFlagRequired("public-keys")
	return vaultNewCmd
}

func vaultShareCommand() *cobra.Command {
	if vaultShareCmd != nil {
		return vaultShareCmd
	}
	vaultShareCmd = &cobra.Command{
		Use:   "share",
		Short: "Shares a vault with another environment or group",
		Run: func(cmd *cobra.Command, args []string) {
			envPrivateKeyString := getEnvSecretKey()
			envPrivateKey, err := crypto.PrivateKeyFromString(envPrivateKeyString)
			if err != nil {
				PrintErrorAndExit(err)
			}
			vaultFile := cmd.Flag("file").Value.String()
			publicKeyStrings, err := cmd.Flags().GetStringSlice("public-keys")
			if err != nil {
				PrintErrorAndExit(err)
			}
			var publicKeys []crypto.PublicKey
			for _, publicKeyString := range publicKeyStrings {
				publicKey, err := crypto.PublicKeyFromString(publicKeyString)
				if err != nil {
					PrintErrorAndExit(err)
				}
				publicKeys = append(publicKeys, *publicKey)
			}
			vault, err := vaults.Get(vaultFile)
			if err == nil {
				err = vault.Unlock(*envPrivateKey)
				if err == nil {
					for _, pubKey := range publicKeys {
						if err = vault.ShareAccessToKey(pubKey); err != nil {
							break
						}
					}
					if err == nil {
						fmt.Println("Shared vault:", color.GreenString(vaultFile))
						os.Exit(0)
					}
				}
			}
			PrintErrorAndExit(err)
		},
	}
	vaultNewCmd.Flags().StringP("file", "f", "", "Name of the environment")
	vaultNewCmd.Flags().StringSliceP("public-keys", "k", []string{}, "Public keys of environments or groups that can access the vault")
	vaultNewCmd.MarkFlagRequired("file")
	vaultNewCmd.MarkFlagRequired("public-keys")
	return vaultShareCmd
}
