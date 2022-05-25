package utils

import (
	"github.com/genstackio/daguerre/commons"
)

func MergeOrderAndSchemaConfig(order *commons.Order, schema *commons.SchemaConfig) *commons.Order {
	o := commons.Order{
		Input:    order.Input,
		Schema:   order.Schema,
		Clusters: []string{},
		Personae: []string{},
		Partners: []string{},
		Output:   order.Output,
		Expand:   []string{},
		Collapse: []string{},
		Show:     []string{},
	}

	o.Clusters = mergeStringList(&order.Clusters, &schema.Clusters)
	o.Personae = mergeStringList(&order.Personae, &schema.Personae)
	o.Partners = mergeStringList(&order.Partners, &schema.Partners)
	o.Expand = mergeStringList(&order.Expand, &schema.Expand)
	o.Collapse = mergeStringList(&order.Collapse, &schema.Collapse)
	o.Show = mergeStringList(&order.Show, &schema.Show)

	return &o
}

func mergeStringList(a *[]string, b *[]string) []string {
	items := map[string]bool{}
	hasAll := false
	if nil != a && len(*a) > 0 {
		for _, v := range *a {
			items[v] = true
			if v == "all" {
				hasAll = true
			}
		}
	}
	if nil != b && len(*b) > 0 {
		for _, v := range *b {
			items[v] = true
			if v == "all" {
				hasAll = true
			}
		}
	}

	if hasAll {
		return []string{"all"}
	}

	list := []string{}

	for k := range items {
		list = append(list, k)
	}

	return list
}
