package types

import (
	"fmt"
)

const (
	DefaultName = "Tirt Assistant"

	VersionMajor = 0
	VersionMinor = 1
	VersionPatch = 0

	PathToConfig = "configs/config.json"
)

func GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}
