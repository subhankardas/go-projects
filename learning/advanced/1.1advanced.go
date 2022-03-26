package advanced

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

const (
	WIN_VERSION_PATH  string = "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion"
	INSTALLATION_TYPE string = "InstallationType"
)

func Advanced1() {
	str, err := GetWindowsVersionRegistryValue(INSTALLATION_TYPE)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Windows installation type is %q\n", str)
}

/* Read windows registry key values */
func GetWindowsVersionRegistryValue(val string) (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, WIN_VERSION_PATH, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer key.Close()

	regval, _, err := key.GetStringValue(val)
	if err != nil {
		return "", err
	}

	return regval, nil
}
