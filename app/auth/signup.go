package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/ent/generated/employee"
	"mazza/ent/generated/userrole"
	"mazza/inits"
)

// Create company and a new user. Associate the new company with the new user.
func Signup(ctx *gin.Context) {
	var body struct {
		Company ent.CreateCompanyInput
		User    ent.CreateUserInput
	}

	tx, err := inits.Client.Tx(ctx)
	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "houve um erro ao registar usuário"})
		return
	}

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "houve um erro ao registar usuário"})
		return
	}

	pwdHash, err := utils.HashPwd(&body.User.Password)
	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ocorreu um erro ao registar usuário"})
		return
	}

	body.User.Password = pwdHash
	body.Company.LastEntryDate = utils.StartOfYear(time.Now())
	trueValue := true
	customerDescription := "Este cliente foi gerado automaticamente"
	supplierDescription := "Este fornecedor foi gerado automaticamente"

	newCompany, err := tx.Company.Create().SetInput(body.Company).Save(ctx)
	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ocorreu um erro ao registar usuário"})
		return
	}

	body.User.CompanyIDs = append(body.User.CompanyIDs, newCompany.ID)
	stock := 0
	_, err = newCompany.Update().AddUsers(
		tx.User.Create().SetInput(body.User).AddAssignedRoles(
			tx.UserRole.Create().SetInput(ent.CreateUserRoleInput{
				Role:      userrole.RoleAdmin,
				CompanyID: &newCompany.ID,
			}).SaveX(ctx),
		).SetEmployee(
			tx.Employee.Create().SetInput(ent.CreateEmployeeInput{
				Name:      body.User.Name,
				Gender:    employee.GenderMale,
				Phone:     utils.GetValue(newCompany.Phone, ""),
				CompanyID: &newCompany.ID,
			}).SaveX(ctx),
		).SaveX(ctx),
	).AddProducts(
		tx.Product.Create().SetInput(ent.CreateProductInput{
			Stock: &stock,
		}).SaveX(ctx),
	).AddCustomers(
		tx.Customer.Create().SetInput(ent.CreateCustomerInput{
			Address:     "--",
			City:        "--",
			Country:     "--",
			Description: &customerDescription,
			IsDefault:   &trueValue,
			Name:        "Clientes Diversos",
			Phone:       "--",
			TaxId:       "--",
		}).SaveX(ctx),
	).AddSuppliers(
		tx.Supplier.Create().SetInput(ent.CreateSupplierInput{
			Address:     "--",
			City:        "--",
			Country:     "--",
			Description: supplierDescription,
			IsDefault:   &trueValue,
			Name:        "Clientes Diversos",
			Phone:       "--",
			TaxId:       "--",
		}).SaveX(ctx),
	).AddTreasuries(
		tx.Treasury.Create().SetInput(ent.CreateTreasuryInput{
			Balance: 0,
		}).SaveX(ctx),
	).Save(ctx)

	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ocorreu um erro ao registar usuário"})
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ocorreu um erro ao registar usuário"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": "usuário registado"})
}
