package service

import (
	"github.com/abdolrhman/simple-go-lang-rest-api/model"
	"github.com/abdolrhman/simple-go-lang-rest-api/pkg/logger"
	"gorm.io/gorm"
	"regexp"
)

type Countries struct {
	Cameroon   string
	Ethiopia   string
	Morocco    string
	Mozambique string
	Uganda     string
}

func GetCustomers(db *gorm.DB, args model.Args) (map[string][]model.Customer, int64, int64, error) {
	var customers []model.Customer
	var filteredData, totalData int64

	table := "customer"
	query := db.Select(table + ".*")
	query = query.Offset(Offset(args.Offset))
	query = query.Limit(Limit(args.Limit))
	query = query.Order(SortOrder(table, args.Sort, args.Order))
	query = query.Scopes(Search(args.Search))

	if err := query.Find(&customers).Error; err != nil {
		logger.Errorf("GetCustomers error: %v", err)
		return nil, filteredData, totalData, err
	}
	// // Count filtered table
	// // We are resetting offset to 0 to return total number.
	// // This is a fix for Gorm offset issue
	query = query.Offset(0)
	query.Table(table).Count(&filteredData)

	groupedCustomersByCountry := groupCustomersByCountry(customers)

	// // Count total table
	db.Table(table).Count(&totalData)
	return groupedCustomersByCountry, filteredData, totalData, nil
}

func groupCustomersByCountry(customers []model.Customer) map[string][]model.Customer {
	var cameronCustomers []model.Customer
	var ethiopiaCustomers []model.Customer
	var moroccoCustomers []model.Customer
	var mozambiqueCustomers []model.Customer
	var ugandaCustomers []model.Customer
	for _, customer := range customers {
		isCameron, _ := phonePatternCheck("\\(237\\)\\ ?[2368]\\d{7,8}$", customer.Phone)
		isEthiopia, _ := phonePatternCheck("\\(251\\)\\ ?[1-59]\\d{8}$", customer.Phone)
		isMorocco, _ := phonePatternCheck("\\(212\\)\\ ?[5-9]\\d{8}$", customer.Phone)
		isMozambique, _ := phonePatternCheck("\\(258\\)\\ ?[28]\\d{7,8}$", customer.Phone)
		isUganda, _ := phonePatternCheck("\\(256\\)\\ ?\\d{9}$", customer.Phone)
		if isCameron {
			cameronCustomers = append(cameronCustomers, customer)
		}
		if isEthiopia {
			ethiopiaCustomers = append(ethiopiaCustomers, customer)
		}
		if isMorocco {
			moroccoCustomers = append(moroccoCustomers, customer)
		}
		if isMozambique {
			mozambiqueCustomers = append(mozambiqueCustomers, customer)
		}
		if isUganda {
			ugandaCustomers = append(ugandaCustomers, customer)
		}

	}
	groupedCustomersByCountry := make(map[string][]model.Customer)
	groupedCustomersByCountry["cameronCustomers"] = cameronCustomers
	groupedCustomersByCountry["ethiopiaCustomers"] = ethiopiaCustomers
	groupedCustomersByCountry["moroccoCustomers"] = moroccoCustomers
	groupedCustomersByCountry["mozambiqueCustomers"] = mozambiqueCustomers
	groupedCustomersByCountry["ugandaCustomers"] = ugandaCustomers
	return groupedCustomersByCountry
}

func phonePatternCheck(pattern string, phone string) (bool, error) {
	return regexp.MatchString(pattern, phone)
}
