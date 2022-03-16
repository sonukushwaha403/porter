package configadapter

import (
	"get.porter.sh/porter/pkg/cnab"
	"get.porter.sh/porter/pkg/context"
	"get.porter.sh/porter/pkg/manifest"
)

// ConvertToTestBundle is suitable for taking a test manifest (porter.yaml)
// and making a bundle.json for it. Does not make an accurate representation
// of the bundle, but is suitable for testing.
func ConvertToTestBundle(cxt *context.Context, manifest *manifest.Manifest) (cnab.ExtendedBundle, error) {
	converter := NewManifestConverter(cxt, manifest, nil, nil)
	return converter.ToBundle()
}