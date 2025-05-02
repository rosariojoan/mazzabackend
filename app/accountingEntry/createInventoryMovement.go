package accountingentry

import (
	"context"
	"fmt"
	"math"
	"mazza/ent/generated"
	"mazza/ent/generated/company"
	"mazza/ent/generated/inventory"
	"mazza/ent/generated/inventorymovement"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
)

/* Pass in a database transaction. The caller function should commit the transaction if this operation returns no error */
func CreateInventoryMovement(ctx context.Context, tx *generated.Tx, input generated.CreateInventoryMovementInput, accountingEntry *model.BaseEntryRegistrationInput) (*generated.InventoryMovement, error) {
	_, activeCompany := utils.GetSession(&ctx)
	var quantity float64
	var value float64

	if input.Category == inventorymovement.CategoryOUT {
		// Check if there is enough inventory for this movement
		inventoryItem, err := tx.Inventory.Query().Where(
			inventory.ID(input.InventoryID),
			inventory.HasCompanyWith(company.ID(activeCompany.ID)),
			inventory.QuantityGTE(input.Quantity),
		).First(ctx)
		if err != nil {
			fmt.Println("CreateInventoryMovement inventory err:", err)
			return nil, fmt.Errorf("not enough inventory")
		}

		quantity = -input.Quantity
		if inventoryItem.Quantity > 0 {
			unitValue := inventoryItem.CurrentValue / inventoryItem.Quantity
			value = math.Round(100*unitValue*quantity) / 100 // round to 2 decimals
		} else {
			value = 0
		}
	} else {
		quantity = input.Quantity
		value = input.Value
	}

	movement, err := tx.InventoryMovement.Create().
		SetInput(input).
		SetValue(math.Abs(value)).
		SetCompanyID(activeCompany.ID).
		Save(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	// Update inventory
	_, err = tx.Inventory.UpdateOneID(input.InventoryID).
		Where(inventory.HasCompanyWith(company.ID(activeCompany.ID))).
		AddQuantity(quantity).
		AddCurrentValue(value).
		Save(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	if accountingEntry != nil {
		// Create accounting entry
		_, err = RegisterAccountingOperations(ctx, tx, *accountingEntry, nil)
		if err != nil {
			fmt.Println("err:", err)
			return nil, fmt.Errorf("an error occurred")
		}
	}

	return movement, nil
}
