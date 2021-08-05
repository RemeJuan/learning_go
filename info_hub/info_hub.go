package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	CategoryPortfolio     string `json:"categoryPortfolio"`
	CategoryId            string `json:"categoryId"`
	BusinessUnit          string `json:"businessUnit"`
	Division              string `json:"division"`
	DivisionName          string `json:"divisionName"`
	CategoryName          string `json:"categoryName"`
	BusinessUnitName      string `json:"businessUnitName"`
	CategoryPortfolioName string `json:"categoryPortfolioName"`
	Ttl                   int    `json:"ttl"`
}

type ProductSelection struct {
	Id            string `json:"id"`
	QueryEngineId string `json:"queryEngineId"`
	Name          string `json:"name"`
	ParentId      string `json:"parentId"`
}

func generateMasterDataItem() MasterDataItem {
	return MasterDataItem{
		CategoryId:            faker.Number().Number(18),
		BusinessUnit:          faker.Number().Number(18),
		Division:              faker.Number().Number(18),
		CategoryPortfolio:     faker.Number().Number(18),
		DivisionName:          faker.Company().Name(),
		CategoryName:          faker.Company().Name(),
		BusinessUnitName:      faker.Company().Name(),
		CategoryPortfolioName: faker.Company().Name(),
		Ttl:                   faker.Number().NumberInt(10),
	}
}

func buildProductId(item MasterDataItem, level ProductLevel) string {
	switch level {
	case Division:
		return item.Division
	case BusinessUnit:
		return item.Division + "-" + item.BusinessUnit
	case CategoryPortfolio:
		return item.Division + "-" + item.BusinessUnit + "-" + item.CategoryPortfolio
	case Category:
		return item.Division + "-" + item.BusinessUnit + "-" + item.CategoryPortfolio + "-" + item.CategoryId
	default:
		return ""
	}
}

func buildDivision(item MasterDataItem, id string) ProductSelection {
	return ProductSelection{
		Id:            id,
		QueryEngineId: id,
		Name:          item.DivisionName,
	}
}

func buildBusinessUnit(item MasterDataItem, id string) ProductSelection {
	return ProductSelection{
		Id:            id,
		QueryEngineId: item.BusinessUnit,
		Name:          item.BusinessUnitName,
	}
}

func buildCategoryPortfolio(item MasterDataItem, id string) ProductSelection {
	return ProductSelection{
		Id:            id,
		QueryEngineId: item.BusinessUnit,
		Name:          item.CategoryPortfolioName,
		ParentId:      buildProductId(item, BusinessUnit),
	}
}

func buildCategory(item MasterDataItem, id string) ProductSelection {
	return ProductSelection{
		Id:            id,
		QueryEngineId: item.CategoryPortfolio,
		Name:          item.CategoryName,
		ParentId:      buildProductId(item, CategoryPortfolio),
	}
}

func writeJSON(response map[string]map[string]ProductSelection, items []MasterDataItem) {
	modeled, _ := json.MarshalIndent(response, "", " ")
	md, _ := json.MarshalIndent(items, "", " ")

	_ = ioutil.WriteFile("modeled_data.json", modeled, 0644)
	_ = ioutil.WriteFile("master_data.json", md, 0644)
}

func writeCSV(items []MasterDataItem) {
	file, err := os.Create("./People.csv")

	if err != nil {
		fmt.Println(err)
	}

	writer := csv.NewWriter(file)

	csvData := [][]string{
		{"CategoryPortfolio", "BusinessUnit", "Division", "Category"},
	}

	for _, v := range items {
		tmp := []string{v.CategoryPortfolio, v.BusinessUnit, v.Division, v.CategoryId}
		csvData = append(csvData, tmp)
	}

	e := writer.WriteAll(csvData)
	if e != nil {
		fmt.Println(e)
	}
}

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

	writeJSON(response, items)
	writeCSV(items)
}
