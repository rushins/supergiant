package azure

import (
	"github.com/supergiant/supergiant/pkg/core"
	"github.com/supergiant/supergiant/pkg/kubernetes"
	"github.com/supergiant/supergiant/pkg/model"
)

// ValidateAccount validates that the AWS credentials entered work.
func (p *Provider) ValidateAccount(m *model.CloudAccount) error {
	return nil
}

// DeleteNode deletes a Kubernetes minion.
func (p *Provider) DeleteNode(m *model.Node, action *core.Action) error {
	return nil
}

// CreateVolume creates a kubernetes Volume.
func (p *Provider) CreateVolume(m *model.Volume, action *core.Action) error {
	return nil
}

// KubernetesVolumeDefinition defines object layout of a AWS volume.
func (p *Provider) KubernetesVolumeDefinition(m *model.Volume) *kubernetes.Volume {
	return &kubernetes.Volume{
		Name: m.Name,
		AwsElasticBlockStore: &kubernetes.AwsElasticBlockStore{
			VolumeID: m.ProviderID,
			FSType:   "ext4",
		},
	}
}

// ResizeVolume resizes a AWS volume.
func (p *Provider) ResizeVolume(m *model.Volume, action *core.Action) error {
	return nil
}

// WaitForVolumeAvailable waits for AWS volume to be available.
func (p *Provider) WaitForVolumeAvailable(m *model.Volume, action *core.Action) error {
	return nil
}

// DeleteVolume deletes a aws volume.
func (p *Provider) DeleteVolume(m *model.Volume, action *core.Action) error {
	return nil
}

// CreateEntrypoint creates a AWS LoadBalancer
func (p *Provider) CreateEntrypoint(m *model.Entrypoint, action *core.Action) error {
	return nil
}

// DeleteEntrypoint deletes a aws loadbalancer.
func (p *Provider) DeleteEntrypoint(m *model.Entrypoint, action *core.Action) error {
	return nil
}

// CreateEntrypointListener creates a listener for a aws loadbalancer.
func (p *Provider) CreateEntrypointListener(m *model.EntrypointListener, action *core.Action) error {
	return nil
}

// DeleteEntrypointListener deletes a listener form an aws loadbalancer.
func (p *Provider) DeleteEntrypointListener(m *model.EntrypointListener, action *core.Action) error {
	return nil
}
