package analytics

import (
	"context"
	"fmt"
	accountingentry "mazza/app/accountingEntry"
	"mazza/ent/generated"
	ac "mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
	"mazza/ent/generated/invoice"
	"mazza/ent/generated/project"
	"mazza/ent/generated/projecttask"
	"mazza/ent/utils"
	"mazza/mazza/generated/model"
	"time"
)

// var monthSales struc {
// 	date time.Time
// }

// Summary analytics for the initial app screen
// Output:
//   - monthly total sales since the begining of the year
//   - No. of pending invoices
//   - No. overdue tasks
//   - No. active projects
//   - Last 4 transactions
func HomepageAnalytics(ctx context.Context, client *generated.Client) (*model.HomepageAnalytics, error) {
	_, activeCompany := utils.GetSession(&ctx)

	// If the current month is January, get daily sales; otherwise, get monthly sales
	currentDate := time.Now()
	var frequency model.TimeRange
	if currentDate.Month() == 1 {
		frequency = model.TimeRangeMonth
	} else {
		frequency = model.TimeRangeYear
	}

	revenueAggregation, err := accountingentry.AggregateRevenue(ctx, *client.AccountingEntry, activeCompany.ID, frequency)
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	// Get number of pending invoice
	pendingInvoices, err := client.Invoice.Query().Where(
		invoice.HasCompanyWith(company.ID(activeCompany.ID)),
		invoice.StatusEQ(invoice.StatusPENDING),
	).Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	// Get the number of overdue tasks
	overdueTasks, err := client.ProjectTask.Query().Where(
		projecttask.HasProjectWith(project.HasCompanyWith(company.ID(activeCompany.ID))),
		projecttask.StatusNEQ(projecttask.StatusCompleted),
		projecttask.DueDateLT(time.Now()),
	).Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	// Get active projects
	activeProjects, err := client.Project.Query().Where(
		project.HasCompanyWith(company.ID(activeCompany.ID)),
		project.StatusEQ(project.StatusCompleted),
	).Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	// Last 4 transactions
	recentTransactions, err := client.AccountingEntry.Query().
		Where(ac.HasCompanyWith(company.ID(activeCompany.ID))).
		Order(generated.Desc(ac.FieldDate)).
		Limit(4).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("an error occurred")
	}

	result := model.HomepageAnalytics{
		RevenueFrequency: frequency,
		RevenueAggregation: revenueAggregation,
		PendingInvoices: pendingInvoices,
		OverdueTasks: overdueTasks,
		ActiveProjects: activeProjects,
		RecentTransactions: recentTransactions,
	}

	return &result, nil
}
