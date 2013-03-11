package gomounts

// Represents a mounted volume on the host system
type Volume struct {
	Path string
	Type string
}

// Gets a slice of all volumes that are currently
// mounted on the host system.
func GetMountedVolumes() ([]Volume, error) {
	return getMountedVolumes()
}