package gce

import (
	"fmt"
	"strconv"
	"time"

	"github.com/supergiant/supergiant/pkg/core"
	"github.com/supergiant/supergiant/pkg/model"
	compute "google.golang.org/api/compute/v1"
)

// CreateKube creates a new GCE kubernetes cluster.
func (p *Provider) CreateKube(m *model.Kube, action *core.Action) error {

	// setup provider steps.
	procedure := &core.Procedure{
		Core:   p.Core,
		Name:   "Create Kube",
		Model:  m,
		Action: action,
	}

	// fetch client.
	client, err := p.Client(m)
	if err != nil {
		return err
	}

	// find the core os image.
	image, err := client.Images.GetFromFamily("coreos-cloud", "coreos-stable").Do()
	if err != nil {
		return err
	}

	// get master machine type.
	instType, err := client.MachineTypes.Get(m.CloudAccount.Credentials["project_id"], m.GCEConfig.Zone, m.MasterNodeSize).Do()
	if err != nil {
		return err
	}

	prefix := "https://www.googleapis.com/compute/v1/projects/" + m.CloudAccount.Credentials["project_id"]

	// Create Master Instance group.
	procedure.AddStep("Creating Kubernetes Master Instance Group...", func() error {
		instanceGroup := &compute.InstanceGroup{
			Name:        m.Name + "-kubernetes-masters",
			Description: "Kubernetes master group for cluster:" + m.Name,
		}
		group, serr := client.InstanceGroups.Insert(m.CloudAccount.Credentials["project_id"], m.GCEConfig.Zone, instanceGroup).Do()
		if serr != nil {
			return serr
		}

		m.GCEConfig.MasterInstanceGroup = group.SelfLink
		serr = p.Core.DB.Save(m)
		if serr != nil {
			return serr
		}
		return nil
	})

	// Create Minion Instance group
	procedure.AddStep("Creating Kubernetes Minion Instance Group...", func() error {
		instanceGroup := &compute.InstanceGroup{
			Name:        m.Name + "-kubernetes-minions",
			Description: "Kubernetes minion group for cluster:" + m.Name,
		}
		group, serr := client.InstanceGroups.Insert(m.CloudAccount.Credentials["project_id"], m.GCEConfig.Zone, instanceGroup).Do()
		if serr != nil {
			return serr
		}

		m.GCEConfig.MinionInstanceGroup = group.SelfLink
		serr = p.Core.DB.Save(m)
		if serr != nil {
			return serr
		}
		return nil
	})

	if m.GCEConfig.KubeMasterCount == 0 {
		m.GCEConfig.KubeMasterCount = 1
	}

	for i := 1; i <= m.GCEConfig.KubeMasterCount; i++ {
		// Create master(s)
		count := strconv.Itoa(i)
		procedure.AddStep("Creating Kubernetes Master Node "+count+"...", func() error {
			// launch master.
			role := "master"
			instance := &compute.Instance{
				Name:        m.Name + "-master" + count,
				Description: "Kubernetes master node for cluster:" + m.Name,
				MachineType: instType.SelfLink,
				Metadata: &compute.Metadata{
					Items: []*compute.MetadataItems{
						&compute.MetadataItems{
							Key:   "KubernetesCluster",
							Value: &m.Name,
						},
						&compute.MetadataItems{
							Key:   "Role",
							Value: &role,
						},
					},
				},
				Disks: []*compute.AttachedDisk{
					{
						AutoDelete: true,
						Boot:       true,
						Type:       "PERSISTENT",
						InitializeParams: &compute.AttachedDiskInitializeParams{
							DiskName:    m.Name + "-master" + count + "-root-pd",
							SourceImage: image.SelfLink,
						},
					},
				},
				NetworkInterfaces: []*compute.NetworkInterface{
					&compute.NetworkInterface{
						AccessConfigs: []*compute.AccessConfig{
							&compute.AccessConfig{
								Type: "ONE_TO_ONE_NAT",
								Name: "External NAT",
							},
						},
						Network: prefix + "/global/networks/default",
					},
				},
			}

			// create the instance.
			_, serr := client.Instances.Insert(m.CloudAccount.Credentials["project_id"], m.GCEConfig.Zone, instance).Do()
			if serr != nil {
				return serr
			}

			return action.CancellableWaitFor("Kubernetes master launch", 5*time.Minute, 3*time.Second, func() (bool, error) {
				resp, serr := client.Instances.Get(m.CloudAccount.Credentials["project_id"], m.GCEConfig.Zone, instance.Name).Do()
				if serr != nil {
					return false, serr
				}

				// Save Master info when ready
				if resp.Status == "RUNNING" {
					m.GCEConfig.MasterNodes = append(m.GCEConfig.MasterNodes, resp.SelfLink)
					if serr := p.Core.DB.Save(m); serr != nil {
						return false, serr
					}
				}
				return resp.Status == "RUNNING", nil
			})
		})
	}
	procedure.AddStep("Adding Kubernetes master(s) to instance group...", func() error {
		for _, masterSelfLink := range m.GCEConfig.MasterNodes {
			fmt.Println("Adding self link:", masterSelfLink)
			_, err = client.InstanceGroups.AddInstances(
				m.CloudAccount.Credentials["project_id"],
				m.GCEConfig.Zone,
				m.Name+"-kubernetes-masters",
				&compute.InstanceGroupsAddInstancesRequest{
					Instances: []*compute.InstanceReference{
						&compute.InstanceReference{
							Instance: masterSelfLink,
						},
					},
				},
			).Do()

			if err != nil {
				return err
			}

		}
		return nil
	})

	return procedure.Run()
}
