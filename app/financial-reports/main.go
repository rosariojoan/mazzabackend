package financialreports

import (
	balancesheet "mazza/app/financial-reports/balance-sheet"
	incomestatement "mazza/app/financial-reports/income-statement"
	"mazza/app/financial-reports/ledger"
	trialbalance "mazza/app/financial-reports/trial-balance"
)

var DownloadLedger = ledger.BuildReport
var DownloadTrialBalance = trialbalance.BuildReport
var DownloadIncomeStatement = incomestatement.BuildReport
var DownloadBalanceSheet = balancesheet.BuildReport

// var DownloadCashFlow = cashflow.BuildReport

var TrialBalance = trialbalance.GetTrialBalance
var IncomeStatement = incomestatement.GetIncomeStatement
var BalanceSheet = balancesheet.GetBalanceSheet

// var CashFlow = cashflow.GetCashFlow
