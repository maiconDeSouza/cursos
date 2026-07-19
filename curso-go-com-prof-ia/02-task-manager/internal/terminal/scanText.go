package terminal

import (
	"errors"
)

func ScanText() (string, error) {
	if globalScanner.Scan() {
		return globalScanner.Text(), nil
	}

	if err := globalScanner.Err(); err != nil {
		return "", errors.New("Erro da leitura do terminal")
	}

	return "", nil
}
