package k8s

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/internal/core"
	k8s "github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type k8sKubeconfigGetRequest struct {
	ClusterID string
	Region    scw.Region
}

func k8sKubeconfigGetCommand() *core.Command {
	return &core.Command{
		Short:     `Retrieve a kubeconfig`,
		Long:      `Retrieve the kubeconfig for a specified cluster.`,
		Namespace: "k8s",
		Verb:      "get",
		Resource:  "kubeconfig",
		ArgsType:  reflect.TypeOf(k8sKubeconfigGetRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      "Cluster ID from which to retrieve the kubeconfig",
				Required:   true,
				Positional: true,
			},
			core.RegionArgSpec(),
		},
		Run: k8sKubeconfigGetRun,
	}
}

func k8sKubeconfigGetRun(ctx context.Context, argsI interface{}) (i interface{}, e error) {
	request := argsI.(*k8sKubeconfigGetRequest)

	kubeconfigRequest := &k8s.GetClusterKubeConfigRequest{
		Region:    request.Region,
		ClusterID: request.ClusterID,
	}

	client := core.ExtractClient(ctx)
	apiK8s := k8s.NewAPI(client)

	kubeconfig, err := apiK8s.GetClusterKubeConfig(kubeconfigRequest)
	if err != nil {
		return nil, err
	}

	return string(kubeconfig.GetRaw()), nil
}
