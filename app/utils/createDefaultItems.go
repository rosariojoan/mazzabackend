package utils

import (
	"context"
	"mazza/ent"
	"mazza/ent/product"
	"mazza/ent/treasury"
)

/* Prepare default items to be used during company creation */
func CreateDefaultItems(ctx context.Context, companyID int) error {
	client := ent.FromContext(ctx)
	companyClient := client.Company.UpdateOneID(companyID)
	descr := "Gerado automaticamente"
	isDefault := true
	defaultProduct, err := client.Product.Create().SetInput(ent.CreateProductInput{
		Category:    product.CategoryOther,
		Description: descr,
		IsDefault:   &isDefault,
		Name:        "Produtos/serviços diversos",
		Sku:         "P0001",
		UnitCost:    0,
	}).Save(ctx)
	if err == nil {
		companyClient.AddProducts(defaultProduct)
	}

	defaultCustomer, err := client.Customer.Create().SetInput(ent.CreateCustomerInput{
		Description: descr,
		IsDefault:   &isDefault,
		Name:        "Clientes diversos",
		TaxId:       "----",
	}).Save(ctx)
	if err == nil {
		companyClient.AddCustomers(defaultCustomer)
	}

	defaultSupplier, err := client.Supplier.Create().SetInput(ent.CreateSupplierInput{
		Description: descr,
		IsDefault:   &isDefault,
		Name:        "Fornecedores diversos",
		TaxId:       "----",
	}).Save(ctx)
	if err == nil {
		companyClient.AddSuppliers(defaultSupplier)
	}

	defaultTreasury, err := client.Treasury.Create().SetInput(ent.CreateTreasuryInput{
		Category:      treasury.CategoryCash,
		Currency:      treasury.CurrencyMzn,
		Description:   &descr,
		IsDefault:     &isDefault,
		IsMainAccount: &isDefault,
		Name:          "Caixa e bancos - diversos",
	}).Save(ctx)
	if err == nil {
		companyClient.AddTreasuries(defaultTreasury)
	}

	_, err = companyClient.Save(ctx)
	return err
}
