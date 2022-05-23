package services

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/nodes"
)

func populateModelNodes(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	if nil == m.Clusters {
		m.Clusters = map[string]commons.ClusterModel{}
	}
	m.Clusters[l.Name] = commons.ClusterModel{
		Nodes: &[]commons.Node{},
	}

	nodes.Lambdas(lt, l, m)
	nodes.Opensearches(lt, l, m)
	nodes.Cdns(lt, l, m)
	nodes.Apigws(lt, l, m)
	nodes.Buckets(lt, l, m)
	nodes.Zones(lt, l, m)
	nodes.Records(lt, l, m)
	nodes.Ses(lt, l, m)
	nodes.Tables(lt, l, m)
	nodes.Topics(lt, l, m)
	nodes.Edges(lt, l, m)
	nodes.Cloudfuncs(lt, l, m)
	nodes.Buses(lt, l, m)
	nodes.Personae(lt, l, m)
	nodes.Partners(lt, l, m)
}
