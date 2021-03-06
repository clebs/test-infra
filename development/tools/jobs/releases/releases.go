// Code generated by rendertemplates. DO NOT EDIT.

package releases

// List of currently supported releases
var (
	Release124 = mustParse("1.24")
	Release123 = mustParse("1.23")
	Release122 = mustParse("1.22")
	Release121 = mustParse("1.21")
)

// GetAllKymaReleases returns all supported kyma release branches
func GetAllKymaReleases() []*SupportedRelease {
	return []*SupportedRelease{
		Release123,
		Release122,
		Release121,
	}
}

// GetNextKymaRelease returns the version of kyma currently under development
func GetNextKymaRelease() *SupportedRelease {
	return Release124
}
