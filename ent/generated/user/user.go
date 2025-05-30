// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldIsDemoUser holds the string denoting the isdemouser field in the database.
	FieldIsDemoUser = "is_demo_user"
	// FieldFirebaseUID holds the string denoting the firebaseuid field in the database.
	FieldFirebaseUID = "firebase_uid"
	// FieldFcmToken holds the string denoting the fcmtoken field in the database.
	FieldFcmToken = "fcm_token"
	// FieldExpoPushToken holds the string denoting the expopushtoken field in the database.
	FieldExpoPushToken = "expo_push_token"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldPhotoURL holds the string denoting the photourl field in the database.
	FieldPhotoURL = "photo_url"
	// FieldDepartment holds the string denoting the department field in the database.
	FieldDepartment = "department"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldBirthdate holds the string denoting the birthdate field in the database.
	FieldBirthdate = "birthdate"
	// FieldLastLogin holds the string denoting the lastlogin field in the database.
	FieldLastLogin = "last_login"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// EdgeAccountingEntries holds the string denoting the accountingentries edge name in mutations.
	EdgeAccountingEntries = "accountingEntries"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeAssignedRoles holds the string denoting the assignedroles edge name in mutations.
	EdgeAssignedRoles = "assignedRoles"
	// EdgeSubordinates holds the string denoting the subordinates edge name in mutations.
	EdgeSubordinates = "subordinates"
	// EdgeLeader holds the string denoting the leader edge name in mutations.
	EdgeLeader = "leader"
	// EdgeCreatedMemberSignupTokens holds the string denoting the createdmembersignuptokens edge name in mutations.
	EdgeCreatedMemberSignupTokens = "createdMemberSignupTokens"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgeIssuedInvoices holds the string denoting the issuedinvoices edge name in mutations.
	EdgeIssuedInvoices = "issuedInvoices"
	// EdgeCreatedProjects holds the string denoting the createdprojects edge name in mutations.
	EdgeCreatedProjects = "createdProjects"
	// EdgeLeaderedProjects holds the string denoting the leaderedprojects edge name in mutations.
	EdgeLeaderedProjects = "leaderedProjects"
	// EdgeAssignedProjectTasks holds the string denoting the assignedprojecttasks edge name in mutations.
	EdgeAssignedProjectTasks = "assignedProjectTasks"
	// EdgeParticipatedProjectTasks holds the string denoting the participatedprojecttasks edge name in mutations.
	EdgeParticipatedProjectTasks = "participatedProjectTasks"
	// EdgeCreatedTasks holds the string denoting the createdtasks edge name in mutations.
	EdgeCreatedTasks = "createdTasks"
	// EdgeTokens holds the string denoting the tokens edge name in mutations.
	EdgeTokens = "tokens"
	// EdgeApprovedWorkShifts holds the string denoting the approvedworkshifts edge name in mutations.
	EdgeApprovedWorkShifts = "approvedWorkShifts"
	// EdgeWorkShifts holds the string denoting the workshifts edge name in mutations.
	EdgeWorkShifts = "workShifts"
	// EdgeUploadedDocuments holds the string denoting the uploadeddocuments edge name in mutations.
	EdgeUploadedDocuments = "uploadedDocuments"
	// EdgeApprovedDocuments holds the string denoting the approveddocuments edge name in mutations.
	EdgeApprovedDocuments = "approvedDocuments"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AccountingEntriesTable is the table that holds the accountingEntries relation/edge.
	AccountingEntriesTable = "accounting_entries"
	// AccountingEntriesInverseTable is the table name for the AccountingEntry entity.
	// It exists in this package in order to avoid circular dependency with the "accountingentry" package.
	AccountingEntriesInverseTable = "accounting_entries"
	// AccountingEntriesColumn is the table column denoting the accountingEntries relation/edge.
	AccountingEntriesColumn = "user_accounting_entries"
	// CompanyTable is the table that holds the company relation/edge. The primary key declared below.
	CompanyTable = "company_users"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// AssignedRolesTable is the table that holds the assignedRoles relation/edge.
	AssignedRolesTable = "user_roles"
	// AssignedRolesInverseTable is the table name for the UserRole entity.
	// It exists in this package in order to avoid circular dependency with the "userrole" package.
	AssignedRolesInverseTable = "user_roles"
	// AssignedRolesColumn is the table column denoting the assignedRoles relation/edge.
	AssignedRolesColumn = "user_assigned_roles"
	// SubordinatesTable is the table that holds the subordinates relation/edge.
	SubordinatesTable = "users"
	// SubordinatesColumn is the table column denoting the subordinates relation/edge.
	SubordinatesColumn = "user_subordinates"
	// LeaderTable is the table that holds the leader relation/edge.
	LeaderTable = "users"
	// LeaderColumn is the table column denoting the leader relation/edge.
	LeaderColumn = "user_subordinates"
	// CreatedMemberSignupTokensTable is the table that holds the createdMemberSignupTokens relation/edge.
	CreatedMemberSignupTokensTable = "member_signup_tokens"
	// CreatedMemberSignupTokensInverseTable is the table name for the MemberSignupToken entity.
	// It exists in this package in order to avoid circular dependency with the "membersignuptoken" package.
	CreatedMemberSignupTokensInverseTable = "member_signup_tokens"
	// CreatedMemberSignupTokensColumn is the table column denoting the createdMemberSignupTokens relation/edge.
	CreatedMemberSignupTokensColumn = "user_created_member_signup_tokens"
	// EmployeeTable is the table that holds the employee relation/edge.
	EmployeeTable = "employees"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "user_employee"
	// IssuedInvoicesTable is the table that holds the issuedInvoices relation/edge.
	IssuedInvoicesTable = "invoices"
	// IssuedInvoicesInverseTable is the table name for the Invoice entity.
	// It exists in this package in order to avoid circular dependency with the "invoice" package.
	IssuedInvoicesInverseTable = "invoices"
	// IssuedInvoicesColumn is the table column denoting the issuedInvoices relation/edge.
	IssuedInvoicesColumn = "user_issued_invoices"
	// CreatedProjectsTable is the table that holds the createdProjects relation/edge.
	CreatedProjectsTable = "projects"
	// CreatedProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	CreatedProjectsInverseTable = "projects"
	// CreatedProjectsColumn is the table column denoting the createdProjects relation/edge.
	CreatedProjectsColumn = "user_created_projects"
	// LeaderedProjectsTable is the table that holds the leaderedProjects relation/edge.
	LeaderedProjectsTable = "projects"
	// LeaderedProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	LeaderedProjectsInverseTable = "projects"
	// LeaderedProjectsColumn is the table column denoting the leaderedProjects relation/edge.
	LeaderedProjectsColumn = "user_leadered_projects"
	// AssignedProjectTasksTable is the table that holds the assignedProjectTasks relation/edge.
	AssignedProjectTasksTable = "project_tasks"
	// AssignedProjectTasksInverseTable is the table name for the ProjectTask entity.
	// It exists in this package in order to avoid circular dependency with the "projecttask" package.
	AssignedProjectTasksInverseTable = "project_tasks"
	// AssignedProjectTasksColumn is the table column denoting the assignedProjectTasks relation/edge.
	AssignedProjectTasksColumn = "user_assigned_project_tasks"
	// ParticipatedProjectTasksTable is the table that holds the participatedProjectTasks relation/edge. The primary key declared below.
	ParticipatedProjectTasksTable = "user_participatedProjectTasks"
	// ParticipatedProjectTasksInverseTable is the table name for the ProjectTask entity.
	// It exists in this package in order to avoid circular dependency with the "projecttask" package.
	ParticipatedProjectTasksInverseTable = "project_tasks"
	// CreatedTasksTable is the table that holds the createdTasks relation/edge.
	CreatedTasksTable = "project_tasks"
	// CreatedTasksInverseTable is the table name for the ProjectTask entity.
	// It exists in this package in order to avoid circular dependency with the "projecttask" package.
	CreatedTasksInverseTable = "project_tasks"
	// CreatedTasksColumn is the table column denoting the createdTasks relation/edge.
	CreatedTasksColumn = "user_created_tasks"
	// TokensTable is the table that holds the tokens relation/edge.
	TokensTable = "tokens"
	// TokensInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokensInverseTable = "tokens"
	// TokensColumn is the table column denoting the tokens relation/edge.
	TokensColumn = "user_tokens"
	// ApprovedWorkShiftsTable is the table that holds the approvedWorkShifts relation/edge.
	ApprovedWorkShiftsTable = "workshifts"
	// ApprovedWorkShiftsInverseTable is the table name for the Workshift entity.
	// It exists in this package in order to avoid circular dependency with the "workshift" package.
	ApprovedWorkShiftsInverseTable = "workshifts"
	// ApprovedWorkShiftsColumn is the table column denoting the approvedWorkShifts relation/edge.
	ApprovedWorkShiftsColumn = "user_approved_work_shifts"
	// WorkShiftsTable is the table that holds the workShifts relation/edge.
	WorkShiftsTable = "workshifts"
	// WorkShiftsInverseTable is the table name for the Workshift entity.
	// It exists in this package in order to avoid circular dependency with the "workshift" package.
	WorkShiftsInverseTable = "workshifts"
	// WorkShiftsColumn is the table column denoting the workShifts relation/edge.
	WorkShiftsColumn = "user_work_shifts"
	// UploadedDocumentsTable is the table that holds the uploadedDocuments relation/edge.
	UploadedDocumentsTable = "company_documents"
	// UploadedDocumentsInverseTable is the table name for the CompanyDocument entity.
	// It exists in this package in order to avoid circular dependency with the "companydocument" package.
	UploadedDocumentsInverseTable = "company_documents"
	// UploadedDocumentsColumn is the table column denoting the uploadedDocuments relation/edge.
	UploadedDocumentsColumn = "user_uploaded_documents"
	// ApprovedDocumentsTable is the table that holds the approvedDocuments relation/edge.
	ApprovedDocumentsTable = "company_documents"
	// ApprovedDocumentsInverseTable is the table name for the CompanyDocument entity.
	// It exists in this package in order to avoid circular dependency with the "companydocument" package.
	ApprovedDocumentsInverseTable = "company_documents"
	// ApprovedDocumentsColumn is the table column denoting the approvedDocuments relation/edge.
	ApprovedDocumentsColumn = "user_approved_documents"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldIsDemoUser,
	FieldFirebaseUID,
	FieldFcmToken,
	FieldExpoPushToken,
	FieldEmail,
	FieldName,
	FieldAddress,
	FieldAvatar,
	FieldPhotoURL,
	FieldDepartment,
	FieldPhone,
	FieldBirthdate,
	FieldLastLogin,
	FieldGender,
	FieldActive,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_subordinates",
}

var (
	// CompanyPrimaryKey and CompanyColumn2 are the table columns denoting the
	// primary key for the company relation (M2M).
	CompanyPrimaryKey = []string{"company_id", "user_id"}
	// ParticipatedProjectTasksPrimaryKey and ParticipatedProjectTasksColumn2 are the table columns denoting the
	// primary key for the participatedProjectTasks relation (M2M).
	ParticipatedProjectTasksPrimaryKey = []string{"user_id", "project_task_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsDemoUser holds the default value on creation for the "isDemoUser" field.
	DefaultIsDemoUser bool
	// FirebaseUIDValidator is a validator for the "firebaseUID" field. It is called by the builders before save.
	FirebaseUIDValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
)

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMale, GenderFemale:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for gender field: %q", ge)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deletedAt field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByIsDemoUser orders the results by the isDemoUser field.
func ByIsDemoUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDemoUser, opts...).ToFunc()
}

// ByFirebaseUID orders the results by the firebaseUID field.
func ByFirebaseUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirebaseUID, opts...).ToFunc()
}

// ByFcmToken orders the results by the fcmToken field.
func ByFcmToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFcmToken, opts...).ToFunc()
}

// ByExpoPushToken orders the results by the expoPushToken field.
func ByExpoPushToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpoPushToken, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByPhotoURL orders the results by the photoURL field.
func ByPhotoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL, opts...).ToFunc()
}

// ByDepartment orders the results by the department field.
func ByDepartment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartment, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByBirthdate orders the results by the birthdate field.
func ByBirthdate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthdate, opts...).ToFunc()
}

// ByLastLogin orders the results by the lastLogin field.
func ByLastLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLogin, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByAccountingEntriesCount orders the results by accountingEntries count.
func ByAccountingEntriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAccountingEntriesStep(), opts...)
	}
}

// ByAccountingEntries orders the results by accountingEntries terms.
func ByAccountingEntries(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountingEntriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCompanyCount orders the results by company count.
func ByCompanyCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCompanyStep(), opts...)
	}
}

// ByCompany orders the results by company terms.
func ByCompany(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAssignedRolesCount orders the results by assignedRoles count.
func ByAssignedRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssignedRolesStep(), opts...)
	}
}

// ByAssignedRoles orders the results by assignedRoles terms.
func ByAssignedRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssignedRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubordinatesCount orders the results by subordinates count.
func BySubordinatesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubordinatesStep(), opts...)
	}
}

// BySubordinates orders the results by subordinates terms.
func BySubordinates(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubordinatesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLeaderField orders the results by leader field.
func ByLeaderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLeaderStep(), sql.OrderByField(field, opts...))
	}
}

// ByCreatedMemberSignupTokensCount orders the results by createdMemberSignupTokens count.
func ByCreatedMemberSignupTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedMemberSignupTokensStep(), opts...)
	}
}

// ByCreatedMemberSignupTokens orders the results by createdMemberSignupTokens terms.
func ByCreatedMemberSignupTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedMemberSignupTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEmployeeField orders the results by employee field.
func ByEmployeeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmployeeStep(), sql.OrderByField(field, opts...))
	}
}

// ByIssuedInvoicesCount orders the results by issuedInvoices count.
func ByIssuedInvoicesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIssuedInvoicesStep(), opts...)
	}
}

// ByIssuedInvoices orders the results by issuedInvoices terms.
func ByIssuedInvoices(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssuedInvoicesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedProjectsCount orders the results by createdProjects count.
func ByCreatedProjectsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedProjectsStep(), opts...)
	}
}

// ByCreatedProjects orders the results by createdProjects terms.
func ByCreatedProjects(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedProjectsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLeaderedProjectsCount orders the results by leaderedProjects count.
func ByLeaderedProjectsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLeaderedProjectsStep(), opts...)
	}
}

// ByLeaderedProjects orders the results by leaderedProjects terms.
func ByLeaderedProjects(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLeaderedProjectsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAssignedProjectTasksCount orders the results by assignedProjectTasks count.
func ByAssignedProjectTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssignedProjectTasksStep(), opts...)
	}
}

// ByAssignedProjectTasks orders the results by assignedProjectTasks terms.
func ByAssignedProjectTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssignedProjectTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParticipatedProjectTasksCount orders the results by participatedProjectTasks count.
func ByParticipatedProjectTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newParticipatedProjectTasksStep(), opts...)
	}
}

// ByParticipatedProjectTasks orders the results by participatedProjectTasks terms.
func ByParticipatedProjectTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParticipatedProjectTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedTasksCount orders the results by createdTasks count.
func ByCreatedTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedTasksStep(), opts...)
	}
}

// ByCreatedTasks orders the results by createdTasks terms.
func ByCreatedTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTokensCount orders the results by tokens count.
func ByTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTokensStep(), opts...)
	}
}

// ByTokens orders the results by tokens terms.
func ByTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByApprovedWorkShiftsCount orders the results by approvedWorkShifts count.
func ByApprovedWorkShiftsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newApprovedWorkShiftsStep(), opts...)
	}
}

// ByApprovedWorkShifts orders the results by approvedWorkShifts terms.
func ByApprovedWorkShifts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApprovedWorkShiftsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkShiftsCount orders the results by workShifts count.
func ByWorkShiftsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkShiftsStep(), opts...)
	}
}

// ByWorkShifts orders the results by workShifts terms.
func ByWorkShifts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkShiftsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUploadedDocumentsCount orders the results by uploadedDocuments count.
func ByUploadedDocumentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUploadedDocumentsStep(), opts...)
	}
}

// ByUploadedDocuments orders the results by uploadedDocuments terms.
func ByUploadedDocuments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUploadedDocumentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByApprovedDocumentsCount orders the results by approvedDocuments count.
func ByApprovedDocumentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newApprovedDocumentsStep(), opts...)
	}
}

// ByApprovedDocuments orders the results by approvedDocuments terms.
func ByApprovedDocuments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApprovedDocumentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAccountingEntriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountingEntriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AccountingEntriesTable, AccountingEntriesColumn),
	)
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CompanyTable, CompanyPrimaryKey...),
	)
}
func newAssignedRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssignedRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AssignedRolesTable, AssignedRolesColumn),
	)
}
func newSubordinatesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubordinatesTable, SubordinatesColumn),
	)
}
func newLeaderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, LeaderTable, LeaderColumn),
	)
}
func newCreatedMemberSignupTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedMemberSignupTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedMemberSignupTokensTable, CreatedMemberSignupTokensColumn),
	)
}
func newEmployeeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmployeeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, EmployeeTable, EmployeeColumn),
	)
}
func newIssuedInvoicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssuedInvoicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IssuedInvoicesTable, IssuedInvoicesColumn),
	)
}
func newCreatedProjectsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedProjectsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedProjectsTable, CreatedProjectsColumn),
	)
}
func newLeaderedProjectsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LeaderedProjectsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LeaderedProjectsTable, LeaderedProjectsColumn),
	)
}
func newAssignedProjectTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssignedProjectTasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AssignedProjectTasksTable, AssignedProjectTasksColumn),
	)
}
func newParticipatedProjectTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ParticipatedProjectTasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ParticipatedProjectTasksTable, ParticipatedProjectTasksPrimaryKey...),
	)
}
func newCreatedTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedTasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedTasksTable, CreatedTasksColumn),
	)
}
func newTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TokensTable, TokensColumn),
	)
}
func newApprovedWorkShiftsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApprovedWorkShiftsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ApprovedWorkShiftsTable, ApprovedWorkShiftsColumn),
	)
}
func newWorkShiftsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkShiftsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkShiftsTable, WorkShiftsColumn),
	)
}
func newUploadedDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UploadedDocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UploadedDocumentsTable, UploadedDocumentsColumn),
	)
}
func newApprovedDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApprovedDocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ApprovedDocumentsTable, ApprovedDocumentsColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Gender) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Gender) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Gender(str)
	if err := GenderValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}
