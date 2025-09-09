package res

import "embed"

//go:embed vendors
var vendorsFS embed.FS

func GetVendorBytes(path string) ([]byte, error) {
	return vendorsFS.ReadFile(path)
}
