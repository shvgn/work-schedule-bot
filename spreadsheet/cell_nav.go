package spreadsheet

import (
	"fmt"
)

func colAddrByIndex(i int) (string, error) {
	if i < 26 {
		return letterByIndex(i)
	}

	firstLetter, err := colAddrByIndex(i/26 - 1)
	if err != nil {
		return "", fmt.Errorf("cannot calc col first letter: %v", err)
	}

	secondLetter, err := letterByIndex(i % 26)
	if err != nil {
		return "", fmt.Errorf("cannot calc col second letter: %v", err)
	}
	return firstLetter + secondLetter, nil
}

func letterByIndex(i int) (string, error) {
	if i < 0 || i > 25 {
		return "", fmt.Errorf("letter index must be from 1 to 26, got %d", i)
	}
	return string(rune(65 + i)), nil
}
