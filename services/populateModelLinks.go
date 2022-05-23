package services

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/links"
)

func populateModelLinks(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	populateLinks(lt.Links, m)
	populateLinks(l.Links, m)
	links.Lambdas(lt, l, m)
	links.Opensearches(lt, l, m)
	links.Cdns(lt, l, m)
	links.Apigws(lt, l, m)
	links.Buckets(lt, l, m)
	links.Zones(lt, l, m)
	links.Records(lt, l, m)
	links.Ses(lt, l, m)
	links.Tables(lt, l, m)
	links.Topics(lt, l, m)
	links.Edges(lt, l, m)
	links.Cloudfuncs(lt, l, m)
	links.Buses(lt, l, m)
	links.Personae(lt, l, m)
	links.Partners(lt, l, m)
}
