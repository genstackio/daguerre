package services

func convertAliases(s string) string {
	if "bus" == s {
		return "buses/default"
	}
	if "tables" == s {
		return "tables/all"
	}

	return s
}