package main

import (
	"encoding/json"
	"io/ioutil"
)

type ProductLevel string

const (
	Division = ProductLevel(rune(iota))
	BusinessUnit
	CategoryPortfolio
	Category
)

func main() {
	response := make(map[string]map[string]ProductSelection)

	divisions := make(map[string]ProductSelection)
	businessUnits := make(map[string]ProductSelection)
	categoryPortfolios := make(map[string]ProductSelection)
	categories := make(map[string]ProductSelection)
	items := make([]MasterDataItem, 0)

	for i := 0; i < 639; i++ {
		item := generateMasterDataItem()
		divisionId := buildProductId(item, Division)
		buId := buildProductId(item, BusinessUnit)
		cpId := buildProductId(item, CategoryPortfolio)
		catId := buildProductId(item, Category)

		divisions[divisionId] = buildDivision(item, divisionId)
		businessUnits[buId] = buildBusinessUnit(item, buId)
		categoryPortfolios[cpId] = buildCategoryPortfolio(item, cpId)
		categories[catId] = buildCategory(item, catId)
		items = append(items, item)
	}

	response["divisions"] = divisions
	response["businessUnits"] = businessUnits
	response["categoryPortfolios"] = categoryPortfolios
	response["categories"] = categories

	modeled, _ := json.MarshalIndent(response, "", " ")
	md, _ := json.MarshalIndent(items, "", " ")

	_ = ioutil.WriteFile("modeled_data.json", modeled, 0644)
	_ = ioutil.WriteFile("master_data.json", md, 0644)
}
