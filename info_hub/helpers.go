package main

import "syreclabs.com/go/faker"

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
