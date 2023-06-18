package  main

import (
	"fmt"
	"github.com/jefersonf/taxons/pkg/taxon"
)

func main() {
	for i:=0; i<8; i++ {
		name, _ := taxon.GetTaxonNameByRank(taxon.TaxonRank(i))
		fmt.Println(name)
	}
	
}
