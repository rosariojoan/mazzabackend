package utils

import (
	"context"
	ent "mazza/ent/generated"
	"mazza/ent/generated/inventory"
)

/* Prepare default items to be used during company creation */
func CreateDefaultItems(ctx context.Context, companyID int) error {
	client := ent.FromContext(ctx)
	companyClient := client.Company.UpdateOneID(companyID)
	descr := "Gerado automaticamente"
	stock := 0.0

	defaultProduct, err := client.Inventory.Create().SetInput(ent.CreateInventoryInput{
		Category:     inventory.CategoryMerchandise,
		CurrentValue: 0,
		MinimumLevel: 0,
		Name:         "Miscellaneous",
		Notes:        "",
		Quantity:     stock,
		Unit:         "unit",
	}).Save(ctx)
	if err == nil {
		companyClient.AddInventory(defaultProduct)
	}

	emptryString := ""
	defaultCustomer, err := client.Customer.Create().SetInput(ent.CreateCustomerInput{
		Description: &descr,
		Name:        "Clientes diversos",
		TaxID:       &emptryString,
	}).SetIsDefault(true).Save(ctx)
	if err == nil {
		companyClient.AddCustomers(defaultCustomer)
	}

	defaultSupplier, err := client.Supplier.Create().SetInput(ent.CreateSupplierInput{
		Description: &descr,
		Name:        "Fornecedores diversos",
		TaxID:       &emptryString,
	}).SetIsDefault(true).Save(ctx)
	if err == nil {
		companyClient.AddSuppliers(defaultSupplier)
	}

	defaultTreasury, err := client.Treasury.Create().SetInput(ent.CreateTreasuryInput{
		Balance: 0,
	}).Save(ctx)
	if err == nil {
		companyClient.AddTreasuries(defaultTreasury)
	}

	_, err = companyClient.Save(ctx)
	return err
}
