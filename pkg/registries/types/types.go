package types

import (
	"context"
	"net/http"

	"github.com/stackrox/rox/generated/storage"
)

const (
	// ArtifactRegistryType represents the Google Artifact registry image integration.
	ArtifactRegistryType = "artifactregistry"

	// ArtifactoryType represents the Artifactory image integration.
	ArtifactoryType = "artifactory"

	// AzureType represents the Azure container registry image integration.
	AzureType = "azure"

	// DockerType represents the docker image integration.
	DockerType = "docker"

	// ECRType represents the AWS ECR image integration.
	ECRType = "ecr"

	// GHCRType represents the GitHub container registry image integration.
	GHCRType = "ghcr"

	// GoogleType represents the Google container registry image integration.
	GoogleType = "google"

	// IBMType represents the IBM image integration.
	IBMType = "ibm"

	// NexusType represents the Nexus image integration.
	NexusType = "nexus"

	// QuayType represents the Quay image integration.
	QuayType = "quay"

	// RedHatType represents the RHEL image integration.
	RedHatType = "rhel"
)

// Config is the config of the registry, which can be utilized by 3rd party scanners
type Config struct {
	Username         string
	Password         string
	Insecure         bool
	URL              string
	RegistryHostname string
	Autogenerated    bool
}

// GetInsecure returns Insecure if Config is non-nil, false otherwise.
func (c *Config) GetInsecure() bool {
	if c != nil {
		return c.Insecure
	}
	return false
}

// GetRegistryHostname returns RegistryHostname if Config is non-nil, empty string otherwise.
func (c *Config) GetRegistryHostname() string {
	if c != nil {
		return c.RegistryHostname
	}

	return ""
}

// Registry is the interface that all image registries must implement
type Registry interface {
	Match(image *storage.ImageName) bool
	Metadata(image *storage.Image) (*storage.ImageMetadata, error)
	Test() error
	Config(ctx context.Context) *Config
	Name() string
	HTTPClient() *http.Client
}

// ImageRegistry adds a DataSource function to Registry that describes which
// integration formed the interface
//
//go:generate mockgen-wrapper
type ImageRegistry interface {
	Registry
	DataSource() *storage.DataSource
	Source() *storage.ImageIntegration
}

// DockerfileInstructionSet are the set of acceptable keywords in a Dockerfile
var DockerfileInstructionSet = map[string]struct{}{
	"ADD":         {},
	"ARG":         {},
	"CMD":         {},
	"COPY":        {},
	"ENTRYPOINT":  {},
	"ENV":         {},
	"EXPOSE":      {},
	"FROM":        {},
	"HEALTHCHECK": {},
	"LABEL":       {},
	"MAINTAINER":  {},
	"ONBUILD":     {},
	"RUN":         {},
	"SHELL":       {},
	"STOPSIGNAL":  {},
	"USER":        {},
	"VOLUME":      {},
	"WORKDIR":     {},
}

// Creator is the func stub that defines how to instantiate an image registry.
type Creator func(integration *storage.ImageIntegration, options ...CreatorOption) (Registry, error)

// CreatorWrapper is a wrapper around a Creator which also returns the registry's name.
type CreatorWrapper func() (string, Creator)
