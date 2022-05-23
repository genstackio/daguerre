package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func loadModel(d *diagram.Diagram, m *commons.Model, o *commons.Order) error {
	ctx := commons.Ctx{}

	initCtx(&ctx)

	expandeds := map[string]bool{}
	collapseds := map[string]bool{}
	requireds := map[string]bool{}
	requiredSpecials := map[string]map[string]bool{
		"personae": {},
		"partners": {},
	}

	for _, v := range o.Expand {
		expandeds[v] = true
	}
	for _, v := range o.Clusters {
		requireds[v] = true
	}
	for _, v := range o.Personae {
		requiredSpecials["personae"][v] = true
	}
	for _, v := range o.Partners {
		requiredSpecials["partners"][v] = true
	}
	for _, v := range o.Show {
		_, found1 := expandeds[v]
		_, found2 := collapseds["all"]
		if !found1 && !found2 {
			collapseds[v] = true
		}
	}

	return nil
}

/*

for l in m["links"]:
for ll in parse_link(l, ctx, m):
a, b, c, d = extract_cluster_from_links(ll, ctx, expandeds, m)
if a in expandeds or b in expandeds or "all" in expandeds:
if a in ["personae", "partners"]:
required_specials[a][c] = True
else:
requireds[a] = True
if b in ["personae", "partners"]:
required_specials[b][d] = True
else:
requireds[b] = True

with Cluster("platform"):
for kk in m["clusters"]:
if kk in ["personae", "partners"]:
continue
with Cluster(kk):
if (kk in expandeds or "all" in expandeds):
for k in m["clusters"][kk]["nodes"]:
if k["type"] in m:
dd = m[k["type"]][k["name"]]
if "hidden" not in dd or False == dd["hidden"]:
if k["type"] in node_types:
ctx[k["type"]][k["name"]] = create_node_by_type(k["type"], k["name"])
else:
if kk in collapseds or "all" in collapseds:
ctx["clusters"][kk] = create_cluster_node(kk, m)
else:
if kk in requireds or "all" in requireds:
ctx["clusters"][kk] = create_cluster_node(kk, m)


for i in m["personae"]:
if i in required_specials["personae"] or "all" in required_specials["personae"]:
if "multiple" in m["personae"][i] and True == m["personae"][i]["multiple"]:
ctx["personae"][i] = create_node_by_type("users", i)
else:
ctx["personae"][i] = create_node_by_type("user", i)

for i in m["partners"]:
if i in required_specials["partners"] or "all" in required_specials["partners"]:
if i in node_types:
ctx["partners"][i] = create_node_by_type(i, i)
else:
ctx["partners"][i] = create_node_by_type("unknown", i)

kept_links = {}

for l in m["links"]:
for ll in parse_link(l, ctx, m):
a, b, md, c, d, e = explode_link(ll, ctx, expandeds, collapseds, m)
kk = c + md + d
if kk in kept_links:
kept_links[kk]["count"] = kept_links[kk]["count"] + 1
if e is not None:
kept_links[kk]["labels"][e] = True
else:
kept_links[kk] = {"from": a, "to": b, "count": 1, "labels": {}, "mode": md}
if e is not None:
kept_links[kk]["labels"][e] = True

for lll in kept_links:
l = kept_links[lll]
a = l["from"]
b = l["to"]
cc = l["count"]
md = l["mode"]
if None != a and None != b:
if "-" == md:
if (cc > 1):
a - Edge(label="(" + str(cc) + ")") - b
else:
a - b
else:
if ">" == md:
if (cc > 1):
a >> Edge(label="(" + str(cc) + ")") >> b
else:
a >> b
else:
if "<" == md:
if (cc > 1):
a << Edge(label="(" + str(cc) + ")") << b
else:
a << b
else:
if (cc > 1):
a - Edge(label="(" + str(cc) + ")") - b
else:
a - b

return ctx

*/
