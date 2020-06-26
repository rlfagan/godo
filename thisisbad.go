// Package fossademo is a demo for FOSSA that will not be merged.
package fossademo

import (
	"crypto/sha1"
	"fmt"
)

// NeverUseSHA1 indicates that SHA1 should never be used
func NeverUseSHA1() error {

	password := "pleasedontshare1234"

	var output []byte
	hash := sha1.New()
	hash.Sum([]byte(password))
	_, err := hash.Write(output)
	if err != nil {
		return fmt.Errorf("so sad %w", err)
	}

	fmt.Println("Output: " + string(output))

	return nil
}
