package stack

import (
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/filters"
	"docker.io/go-docker/api/types/swarm"
	"github.com/appcelerator/amp/docker/cli/cli/compose/convert"
	"github.com/appcelerator/amp/docker/cli/opts"
	"golang.org/x/net/context"
)

func getStackFilter(namespace string) filters.Args {
	filter := filters.NewArgs()
	filter.Add("label", convert.LabelNamespace+"="+namespace)
	return filter
}

func getServiceFilter(namespace string) filters.Args {
	return getStackFilter(namespace)
}

func getStackFilterFromOpt(namespace string, opt opts.FilterOpt) filters.Args {
	filter := opt.Value()
	filter.Add("label", convert.LabelNamespace+"="+namespace)
	return filter
}

func getAllStacksFilter() filters.Args {
	filter := filters.NewArgs()
	filter.Add("label", convert.LabelNamespace)
	return filter
}

func getServices(
	ctx context.Context,
	apiclient docker.APIClient,
	namespace string,
) ([]swarm.Service, error) {
	return apiclient.ServiceList(
		ctx,
		types.ServiceListOptions{Filters: getServiceFilter(namespace)})
}

func getStackNetworks(
	ctx context.Context,
	apiclient docker.APIClient,
	namespace string,
) ([]types.NetworkResource, error) {
	return apiclient.NetworkList(
		ctx,
		types.NetworkListOptions{Filters: getStackFilter(namespace)})
}

func getStackSecrets(
	ctx context.Context,
	apiclient docker.APIClient,
	namespace string,
) ([]swarm.Secret, error) {
	return apiclient.SecretList(
		ctx,
		types.SecretListOptions{Filters: getStackFilter(namespace)})
}

func getStackConfigs(
	ctx context.Context,
	apiclient docker.APIClient,
	namespace string,
) ([]swarm.Config, error) {
	return apiclient.ConfigList(
		ctx,
		types.ConfigListOptions{Filters: getStackFilter(namespace)})
}
