package hcl2template

import (
	"github.com/hashicorp/hcl/v2"
	packerregistry "github.com/hashicorp/packer/internal/packer_registry"
)

// ConfiguredArtifactMetadataPublisher returns a configured image bucket that can be used for publishing
// build image artifacts to a configured Packer Registry destination.
func (cfg *PackerConfig) ConfiguredArtifactMetadataPublisher() (*packerregistry.Bucket, hcl.Diagnostics) {
	if cfg.Bucket == nil {
		return nil, hcl.Diagnostics{
			&hcl.Diagnostic{
				Summary: "Publishing build artifacts to HCP Packer Registry not enabled",
				Detail: "No Packer Registry configuration detected; skipping all publishing steps " +
					"See publishing to a Packer registry for Packer configuration details",
				Severity: hcl.DiagWarning,
			},
		}
	}

	err := cfg.Bucket.Validate()
	if err != nil {
		return nil, hcl.Diagnostics{
			&hcl.Diagnostic{
				Summary:  "Invalid HCP Packer Registry configuration",
				Detail:   err.Error(),
				Severity: hcl.DiagError,
			},
		}
	}

	return cfg.Bucket, nil
}
