package utils

import (
	"context"
	ent "mazza/ent/generated"
)

/* Prepare default items to be used during company creation */
func CreateDefaultItems(ctx context.Context, companyID int) error {
	client := ent.FromContext(ctx)
	companyClient := client.Company.UpdateOneID(companyID)
	descr := "Gerado automaticamente"
	isDefault := true
	stock := 0
	
	defaultProduct, err := client.Product.Create().SetInput(ent.CreateProductInput{
		Stock: &stock,
	}).Save(ctx)
	if err == nil {
		companyClient.AddProducts(defaultProduct)
	}

	defaultCustomer, err := client.Customer.Create().SetInput(ent.CreateCustomerInput{
		Description: &descr,
		Name:        "Clientes diversos",
		TaxId:       "----",
	}).SetIsDefault(true).Save(ctx)
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
		Balance: 0,
	}).Save(ctx)
	if err == nil {
		companyClient.AddTreasuries(defaultTreasury)
	}

	_, err = companyClient.Save(ctx)
	return err
}
