package azure

import (
	"github.com/supergiant/supergiant/pkg/core"
	"github.com/supergiant/supergiant/pkg/model"
)

// DeleteKube deletes a Kubernetes cluster.
func (p *Provider) DeleteKube(m *model.Kube, action *core.Action) error {
	return nil
}
