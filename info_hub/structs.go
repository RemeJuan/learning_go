package main

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
