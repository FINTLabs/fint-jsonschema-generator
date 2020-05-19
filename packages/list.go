package packages

import (
	"github.com/FINTLabs/fint-jsonschema-generator/common/types"
	"github.com/FINTLabs/fint-jsonschema-generator/common/utils"
)

func DistinctPackageList(classes []*types.Class) []string {

	var p []string
	for _, c := range classes {
		p = append(p, c.Package)
	}

	return utils.Distinct(p)
}
