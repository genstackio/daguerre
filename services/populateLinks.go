package services

import (
	"github.com/genstackio/daguerre/commons"
)

func populateLinks(links *[]string, m *commons.Model) {
	if nil == links {
		return
	}
	for _, v := range *links {
		for _, vv := range parseLink(v, m) {
			*m.Links = append(*m.Links, vv)
		}
	}
}
