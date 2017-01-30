package azure

import (
<<<<<<< HEAD
	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/go-autorest/autorest/azure"
=======
>>>>>>> ac387429444c83d44e391c4d266d466c681f8fe7
	"github.com/supergiant/supergiant/pkg/core"
	"github.com/supergiant/supergiant/pkg/kubernetes"
	"github.com/supergiant/supergiant/pkg/model"
)

<<<<<<< HEAD
// Provider Holds DO account info.
type Provider struct {
	Core     *core.Core
	VMClient func(*model.Kube) *compute.VirtualMachinesClient
}

// ValidateAccount validates that the AWS credentials entered work.
func (p *Provider) ValidateAccount(m *model.CloudAccount) error {
	client := p.VMClient(&model.Kube{CloudAccount: m})

	_, err := client.ListAll()
	if err != nil {
		return err
	}
=======
// ValidateAccount validates that the AWS credentials entered work.
func (p *Provider) ValidateAccount(m *model.CloudAccount) error {
>>>>>>> ac387429444c83d44e391c4d266d466c681f8fe7
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
<<<<<<< HEAD

func VMClient(kube *model.Kube) *compute.VirtualMachinesClient {
	oauthConfig, err := azure.PublicCloud.OAuthConfigForTenant(kube.CloudAccount.Credentials["tenant_id"])
	if err != nil {
		return nil
	}
	token, _ := azure.NewServicePrincipalToken(
		*oauthConfig,
		kube.CloudAccount.Credentials["client_id"],
		kube.CloudAccount.Credentials["client_secret"],
		azure.PublicCloud.ResourceManagerEndpoint,
	)
	client := compute.NewVirtualMachinesClient(kube.CloudAccount.Credentials["subscription_id"])
	client.Authorizer = token
	return &client
}
=======
>>>>>>> ac387429444c83d44e391c4d266d466c681f8fe7
