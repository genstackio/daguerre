package services

import "github.com/genstackio/daguerre/commons"

func linkLayer(l *commons.LayerConfig, m *commons.Model, c *commons.Config) error {
	lt, err := getLayerTypeConfig(l, c)

	if nil != err {
		return err
	}

	populateModelLinks(lt, l, m)

	return nil
}
