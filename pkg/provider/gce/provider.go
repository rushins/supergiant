package gce

import (
	"github.com/supergiant/supergiant/pkg/core"
	"github.com/supergiant/supergiant/pkg/kubernetes"
	"github.com/supergiant/supergiant/pkg/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	compute "google.golang.org/api/compute/v1"
)

// Provider Holds DO account info.
type Provider struct {
	Core   *core.Core
	Client func(*model.Kube) (*compute.Service, error)
}

// ValidateAccount Valitades Open Stack account info.
func (p *Provider) ValidateAccount(m *model.CloudAccount) error {
	client, err := p.Client(&model.Kube{CloudAccount: m})
	if err != nil {
		return err
	}

	_, err = client.Projects.Get(m.Credentials["project_id"]).Do()
	if err != nil {
		return err
	}
	return nil
}

// CreateKube creates a new DO kubernetes cluster.
func (p *Provider) CreateKube(m *model.Kube, action *core.Action) error {
	return nil
}

// DeleteKube deletes a DO kubernetes cluster.
func (p *Provider) DeleteKube(m *model.Kube, action *core.Action) error {
	// Initialize steps
	return nil
}

// CreateNode creates a new minion on DO kubernetes cluster.
func (p *Provider) CreateNode(m *model.Node, action *core.Action) error {
	return nil
}

// DeleteNode deletes a minsion on a DO kubernetes cluster.
func (p *Provider) DeleteNode(m *model.Node, action *core.Action) error {
	return nil
}

// CreateVolume createss a Volume on DO for Kubernetes
func (p *Provider) CreateVolume(m *model.Volume, action *core.Action) error {
	return nil
}

func (p *Provider) KubernetesVolumeDefinition(m *model.Volume) *kubernetes.Volume {
	return &kubernetes.Volume{
		Name: m.Name,
		Cinder: &kubernetes.Cinder{
			VolumeID: m.ProviderID,
			FSType:   m.Type,
		},
	}
}

// ResizeVolume re-sizes volume on DO kubernetes cluster.
func (p *Provider) ResizeVolume(m *model.Volume, action *core.Action) error {

	return nil
}

// WaitForVolumeAvailable waits for DO volume to become available.
func (p *Provider) WaitForVolumeAvailable(m *model.Volume, action *core.Action) error {
	return nil
}

// DeleteVolume deletes a DO volume.
func (p *Provider) DeleteVolume(m *model.Volume, action *core.Action) error {

	return nil
}

// CreateEntrypoint creates a new Load Balancer for Kubernetes in DO
func (p *Provider) CreateEntrypoint(m *model.Entrypoint, action *core.Action) error {
	return nil
}

// DeleteEntrypoint deletes load balancer from DO.
func (p *Provider) DeleteEntrypoint(m *model.Entrypoint, action *core.Action) error {
	return nil
}

func (p *Provider) CreateEntrypointListener(m *model.EntrypointListener, action *core.Action) error {
	return nil
}

func (p *Provider) DeleteEntrypointListener(m *model.EntrypointListener, action *core.Action) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Private methods                                                            //
////////////////////////////////////////////////////////////////////////////////

// Client creates the client for the provider.
func Client(kube *model.Kube) (*compute.Service, error) {

	clientScopes := []string{
		"https://www.googleapis.com/auth/compute",
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/ndev.clouddns.readwrite",
		"https://www.googleapis.com/auth/devstorage.full_control",
	}

	conf := jwt.Config{
		Email:      kube.CloudAccount.Credentials["client_email"],
		PrivateKey: []byte(kube.CloudAccount.Credentials["private_key"]),
		Scopes:     clientScopes,
		TokenURL:   kube.CloudAccount.Credentials["token_uri"],
	}

	client := conf.Client(oauth2.NoContext)

	computeService, err := compute.New(client)
	if err != nil {
		return nil, err
	}
	return computeService, nil
}
