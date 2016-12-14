package ui

import (
	"net/http"

	"github.com/supergiant/supergiant/pkg/client"
)

func ListApps(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	fields := []map[string]interface{}{
		{
			"title": "Kube name",
			"type":  "field_value",
			"field": "kube_name",
		},
		{
			"title": "Kind",
			"type":  "field_value",
			"field": "kind",
		},
		{
			"title": "Namespace",
			"type":  "field_value",
			"field": "namespace",
		},
		{
			"title": "Name",
			"type":  "field_value",
			"field": "name",
		},
	}
	return renderTemplate(sg, w, "apps", map[string]interface{}{
		"title":         "Apps",
		"uiBasePath":    "/ui/apps",
		"apiBasePath":   "/api/v0/kube_resources",
		"fields":        fields,
		"showNewLink":   true,
		"showStatusCol": true,
		"newOptions": map[string]string{
			"pod":     "Pod",
			"service": "Service",
			"other":   "Other",
		},
		"actionPaths": map[string]string{
			"Edit": "/edit",
		},
		"batchActionPaths": map[string]map[string]string{
			"Delete": map[string]string{
				"method":       "DELETE",
				"relativePath": "",
			},
			"Start": map[string]string{
				"method":       "POST",
				"relativePath": "/start",
			},
			"Stop": map[string]string{
				"method":       "POST",
				"relativePath": "/stop",
			},
		},
	})
}
