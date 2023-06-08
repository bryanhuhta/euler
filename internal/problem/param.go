package problem

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

type Params map[string]string

func (p Params) assertKeysExist(keys ...string) error {
	missingKeys := []string{}
	for _, key := range keys {
		_, ok := p[key]
		if !ok {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) != 0 {
		return fmt.Errorf("missing parameters: %v", missingKeys)
	}
	return nil
}

func (p Params) getInt(key string) (int, error) {
	val, ok := p[key]
	if !ok {
		return 0, fmt.Errorf("no such key: %s", key)
	}

	n, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("key %s is not an integer: %v", key, err)
	}

	return n, nil
}

func (p Params) getFileText(key string) (string, error) {
	const basePath = "data"

	val, ok := p[key]
	if !ok {
		return "", fmt.Errorf("no such key: %s", key)
	}

	filepath := path.Join(basePath, val)
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
