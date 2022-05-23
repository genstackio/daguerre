package services

import "github.com/genstackio/daguerre/commons"

func alignCloudfrontDistributionLinkEndpoint(ll commons.LinkEndpointModel, m *commons.Model) commons.LinkEndpointModel {
	tt := ll.Type
	x := m.Lists[tt][ll.Name]

	if len(x.LayerType.Cloudfuncs) > 0 {
		return commons.LinkEndpointModel{Name: "redirect", Type: "cloudfuncs"}
	}
	return ll
}

func alignLinkEndpoint(ll commons.LinkEndpointModel, m *commons.Model) commons.LinkEndpointModel {
	if "cdns" == ll.Type {
		return alignCloudfrontDistributionLinkEndpoint(ll, m)
	}
	return ll
}
