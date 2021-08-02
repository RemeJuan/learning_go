package main

import (
	"fmt"
	"syreclabs.com/go/faker"
)

type ProductLevel string

const (
	Division = ProductLevel(rune(iota))
	BusinessUnit
	CategoryPortfolio
	Category
)

// MasterDataItem
/* {
  "categoryId": "000000000000040621",
  "businessUnit": "000000000000020096",
  "division": "000000000000000006",
  "divisionName": "Clothing",
  "categoryName": "Curtains Bombs",
  "businessUnitName": "Womens, F&A",
  "categoryPortfolio": "000000000000030290",
  "categoryPortfolioName": "Bombs",
  "ttl": 1617598838
},
*/
type MasterDataItem struct {
	categoryPortfolio     string
	categoryId            string
	businessUnit          string
	division              string
	divisionName          string
	categoryName          string
	businessUnitName      string
	categoryPortfolioName string
	ttl                   int
}

type ProductSelection struct {
	Id            string `json:"id"`
	QueryEngineId string `json:"queryEngineId"`
	Name          string `json:"name"`
	ParentId      string `json:"parentId"`
}

func generateMasterDataItem() MasterDataItem {
	return MasterDataItem{
		categoryId:            faker.Number().Number(18),
		businessUnit:          faker.Number().Number(18),
		division:              faker.Number().Number(18),
		categoryPortfolio:     faker.Number().Number(18),
		divisionName:          faker.Company().Name(),
		categoryName:          faker.Company().Name(),
		businessUnitName:      faker.Company().Name(),
		categoryPortfolioName: faker.Company().Name(),
		ttl:                   faker.Number().NumberInt(10),
	}
}

func buildProductId(item MasterDataItem, level ProductLevel) string {
	switch level {
	case Division:
		return item.division
	case BusinessUnit:
		return item.division + "-" + item.businessUnit
	case CategoryPortfolio:
		return item.division + "-" + item.businessUnit + "-" + item.categoryPortfolio
	case Category:
		return item.division + "-" + item.businessUnit + "-" + item.categoryPortfolio + "-" + item.categoryId
	default:
		return ""
	}
}

func buildDivision(item MasterDataItem, id string) ProductSelection {
	return ProductSelection{
		Id:            id,
		QueryEngineId: id,
		Name:          item.divisionName,
	}
}

func buildBusinessUnit(item MasterDataItem, id string) ProductSelection {
	return ProductSelection{
		Id:            id,
		QueryEngineId: item.businessUnit,
		Name:          item.businessUnitName,
		ParentId:      buildProductId(item, Division),
	}
}

func buildCategoryPortfolio(item MasterDataItem, id string) ProductSelection {
	return ProductSelection{
		Id:            id,
		QueryEngineId: item.businessUnit,
		Name:          item.categoryPortfolioName,
		ParentId:      buildProductId(item, BusinessUnit),
	}
}

func buildCategory(item MasterDataItem, id string) ProductSelection {
	return ProductSelection{
		Id:            id,
		QueryEngineId: item.categoryPortfolio,
		Name:          item.categoryName,
		ParentId:      buildProductId(item, CategoryPortfolio),
	}
}

func main() {
	response := make(map[string][]ProductSelection)

	divisions := make([]ProductSelection, 0)
	businessUnits := make([]ProductSelection, 0)
	categoryPortfolios := make([]ProductSelection, 0)
	categories := make([]ProductSelection, 0)

	for i := 0; i < 1; i++ {
		item := generateMasterDataItem()
		divisionId := buildProductId(item, Division)
		buId := buildProductId(item, BusinessUnit)
		cpId := buildProductId(item, CategoryPortfolio)
		catId := buildProductId(item, Category)

		response["divisions"] = append(divisions, buildDivision(item, divisionId))
		response["businessUnits"] = append(businessUnits, buildBusinessUnit(item, buId))
		response["categoryPortfolios"] = append(categoryPortfolios, buildCategoryPortfolio(item, cpId))
		response["categories"] = append(categories, buildCategory(item, catId))
	}

	fmt.Println(response)
}
