package api

import (
	"net/http"

	"github.com/supergiant/supergiant/pkg/core"
	"github.com/supergiant/supergiant/pkg/model"
)

func ListApps(core *core.Core, user *model.User, r *http.Request) (*Response, error) {
	return handleList(core, r, new(model.KubeResource), new(model.KubeResourceList))
}
