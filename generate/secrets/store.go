package secrets

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fkondej/gocli/generate"
)

func StoreEthereumWallet(
	ethereumPrivateKey string,
	walletPassphrase string,
	ethWalletPath string,
) error {
	var errMsg = "failed to store ethereum wallet, %w"
	// Create tmp directory inside directory with result
	rndSuffix, err := generate.GenerateRandomWord(6)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	tmpDir := path.Join(
		filepath.Dir(ethWalletPath),
		fmt.Sprintf("tmp-keystore-%s", rndSuffix),
	)
	defer os.RemoveAll(tmpDir)

	// Create new KeyStore
	ethPrivateKey, err := crypto.HexToECDSA(ethereumPrivateKey)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	ethKeystore := keystore.NewKeyStore(tmpDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// Import secret into KeyStore
	_, err = ethKeystore.ImportECDSA(ethPrivateKey, walletPassphrase)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	// Copy generated file into desired location
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	if len(files) != 1 {
		return fmt.Errorf("expected to have one file in directory %s\n", tmpDir)
	}
	ethWalletOrigPath := path.Join(tmpDir, files[0].Name())
	if err := os.Rename(ethWalletOrigPath, ethWalletPath); err != nil {
		return fmt.Errorf(errMsg, err)
	}

	return nil
}

func encodeToBase64(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	content, _ := io.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded, nil
}
