package accountingentry

import (
	"context"
	"fmt"
	ent "mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/product"
	"mazza/mazza/generated/model"
)

// Check if products are registered in the inventory
// Only check for tangible or generic product
// Stock restriction: whether to check if product input should not exceed the current stock
func CheckAndUpdateProductInput(ctx context.Context, client *ent.Client, companyID int, input []*model.EntryItem, retrievedProducts []*ent.Product, stockRestriction bool) error {
	var err error
	productIDs := []int{}
	for _, p := range input {
		if p.ProductID != nil {
			productIDs = append(productIDs, *p.ProductID)
		}
	}

	retrievedProducts, err = client.Product.Query().
		Where(product.HasCompanyWith(company.IDEQ(companyID)), product.IDIn(productIDs...)).
		All(ctx)
	if err != nil {
		return err
	}
	if len(productIDs) != len(retrievedProducts) {
		return fmt.Errorf("invalid product input")
	}

	// Update stock
	for _, prod := range retrievedProducts {
		if prod.Category != product.CategoryMERCHANDISE {
			continue
		}
		for _, p := range input {
			inputQuantity := 0
			if p.Quantity != nil {
				inputQuantity = *p.Quantity
			}
			if stockRestriction && inputQuantity > int(prod.Stock) {
				return fmt.Errorf("not enough stock")
			}
			if p.ProductID == &prod.ID {
				if p.IsDebit {
					// Product inflow: increase current stock and update unit cost
					prod.UnitCost += (prod.Stock * prod.UnitCost + p.Amount) / (prod.Stock + float64(inputQuantity))
				} else {
					// Product outflow: decrease current stock (if it isn't a default product) and keep the unit cost
					inputQuantity *= -1
				}
				
				prod.Stock += float64(inputQuantity)
				continue
			}
		}
	}

	return nil
}
