package taxon

import "errors"

type TaxonRank uint

const (
	Domain TaxonRank = iota
	Kingdom
	Phylum
	Class
	Order
	Family
	Genus
	Species
)

const TotalTaxons = 8

func GetTaxonNameByRank(rank TaxonRank) (string, error) {
	var name string
	switch rank {
	case Domain:
		name = "Domain"
	case Kingdom:
		name = "Kingdom"
	case Phylum:
		name = "Phylum"
	case Class:
		name = "Class"
	case Order:
		name = "Order"
	case Family:
		name = "Family"
	case Genus:
		name = "Genus"
	case Species:
		name = "Species"
	default:
		return name, errors.New("Invalid taxon rank")
	}
	return name, nil
}
