package advanced

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func Advanced1() {
	readRegistryValues()
}

/* Read windows registry key values */
func readRegistryValues() {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	instype, _, err := key.GetStringValue("InstallationType")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Windows installation type is %q\n", instype)
}
