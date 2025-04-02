// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"mazza/ent/generated/employee"
	"mazza/ent/generated/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deletedAt" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// FirebaseUID holds the value of the "firebaseUID" field.
	FirebaseUID string `json:"-"`
	// FcmToken holds the value of the "fcmToken" field.
	FcmToken *string `json:"-"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Address holds the value of the "address" field.
	Address *string `json:"address,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar *string `json:"avatar,omitempty"`
	// PhotoURL holds the value of the "photoURL" field.
	PhotoURL *string `json:"photoURL,omitempty"`
	// Department holds the value of the "department" field.
	Department *string `json:"department,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone *string `json:"phone,omitempty"`
	// Birthdate holds the value of the "birthdate" field.
	Birthdate *time.Time `json:"birthdate,omitempty"`
	// It can be the last time the user opened the app and synced with the backend.
	LastLogin *time.Time `json:"lastLogin,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender user.Gender `json:"gender,omitempty"`
	// Active holds the value of the "active" field.
	Active *bool `json:"active,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges             UserEdges `json:"edges"`
	user_subordinates *int
	selectValues      sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// AccountingEntries holds the value of the accountingEntries edge.
	AccountingEntries []*AccountingEntry `json:"accountingEntries,omitempty"`
	// Company holds the value of the company edge.
	Company []*Company `json:"company,omitempty"`
	// a user should be assigned to only one role in the company
	AssignedRoles []*UserRole `json:"assignedRoles,omitempty"`
	// Subordinates holds the value of the subordinates edge.
	Subordinates []*User `json:"subordinates,omitempty"`
	// Leader holds the value of the leader edge.
	Leader *User `json:"leader,omitempty"`
	// CreatedMemberSignupTokens holds the value of the createdMemberSignupTokens edge.
	CreatedMemberSignupTokens []*MemberSignupToken `json:"createdMemberSignupTokens,omitempty"`
	// Employee holds the value of the employee edge.
	Employee *Employee `json:"employee,omitempty"`
	// IssuedInvoices holds the value of the issuedInvoices edge.
	IssuedInvoices []*Invoice `json:"issuedInvoices,omitempty"`
	// Represents the projects created by the user
	CreatedProjects []*Project `json:"createdProjects,omitempty"`
	// Represents the projects leadered or supervised by the user
	LeaderedProjects []*Project `json:"leaderedProjects,omitempty"`
	// These are the project tasks assigned to the user and he is responsible for them
	AssignedProjectTasks []*ProjectTask `json:"assignedProjectTasks,omitempty"`
	// These are the project tasks in which the user is a member. E.g. a meeting
	ParticipatedProjectTasks []*ProjectTask `json:"participatedProjectTasks,omitempty"`
	// Represents the tasks created by a user
	CreatedTasks []*ProjectTask `json:"createdTasks,omitempty"`
	// Tokens holds the value of the tokens edge.
	Tokens []*Token `json:"tokens,omitempty"`
	// ApprovedWorkShifts holds the value of the approvedWorkShifts edge.
	ApprovedWorkShifts []*Workshift `json:"approvedWorkShifts,omitempty"`
	// WorkShifts holds the value of the workShifts edge.
	WorkShifts []*Workshift `json:"workShifts,omitempty"`
	// UploadedDocuments holds the value of the uploadedDocuments edge.
	UploadedDocuments []*CompanyDocument `json:"uploadedDocuments,omitempty"`
	// ApprovedDocuments holds the value of the approvedDocuments edge.
	ApprovedDocuments []*CompanyDocument `json:"approvedDocuments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [18]bool
	// totalCount holds the count of the edges above.
	totalCount [18]map[string]int

	namedAccountingEntries         map[string][]*AccountingEntry
	namedCompany                   map[string][]*Company
	namedAssignedRoles             map[string][]*UserRole
	namedSubordinates              map[string][]*User
	namedCreatedMemberSignupTokens map[string][]*MemberSignupToken
	namedIssuedInvoices            map[string][]*Invoice
	namedCreatedProjects           map[string][]*Project
	namedLeaderedProjects          map[string][]*Project
	namedAssignedProjectTasks      map[string][]*ProjectTask
	namedParticipatedProjectTasks  map[string][]*ProjectTask
	namedCreatedTasks              map[string][]*ProjectTask
	namedTokens                    map[string][]*Token
	namedApprovedWorkShifts        map[string][]*Workshift
	namedWorkShifts                map[string][]*Workshift
	namedUploadedDocuments         map[string][]*CompanyDocument
	namedApprovedDocuments         map[string][]*CompanyDocument
}

// AccountingEntriesOrErr returns the AccountingEntries value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AccountingEntriesOrErr() ([]*AccountingEntry, error) {
	if e.loadedTypes[0] {
		return e.AccountingEntries, nil
	}
	return nil, &NotLoadedError{edge: "accountingEntries"}
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CompanyOrErr() ([]*Company, error) {
	if e.loadedTypes[1] {
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// AssignedRolesOrErr returns the AssignedRoles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AssignedRolesOrErr() ([]*UserRole, error) {
	if e.loadedTypes[2] {
		return e.AssignedRoles, nil
	}
	return nil, &NotLoadedError{edge: "assignedRoles"}
}

// SubordinatesOrErr returns the Subordinates value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SubordinatesOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.Subordinates, nil
	}
	return nil, &NotLoadedError{edge: "subordinates"}
}

// LeaderOrErr returns the Leader value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) LeaderOrErr() (*User, error) {
	if e.Leader != nil {
		return e.Leader, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "leader"}
}

// CreatedMemberSignupTokensOrErr returns the CreatedMemberSignupTokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedMemberSignupTokensOrErr() ([]*MemberSignupToken, error) {
	if e.loadedTypes[5] {
		return e.CreatedMemberSignupTokens, nil
	}
	return nil, &NotLoadedError{edge: "createdMemberSignupTokens"}
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) EmployeeOrErr() (*Employee, error) {
	if e.Employee != nil {
		return e.Employee, nil
	} else if e.loadedTypes[6] {
		return nil, &NotFoundError{label: employee.Label}
	}
	return nil, &NotLoadedError{edge: "employee"}
}

// IssuedInvoicesOrErr returns the IssuedInvoices value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) IssuedInvoicesOrErr() ([]*Invoice, error) {
	if e.loadedTypes[7] {
		return e.IssuedInvoices, nil
	}
	return nil, &NotLoadedError{edge: "issuedInvoices"}
}

// CreatedProjectsOrErr returns the CreatedProjects value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[8] {
		return e.CreatedProjects, nil
	}
	return nil, &NotLoadedError{edge: "createdProjects"}
}

// LeaderedProjectsOrErr returns the LeaderedProjects value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LeaderedProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[9] {
		return e.LeaderedProjects, nil
	}
	return nil, &NotLoadedError{edge: "leaderedProjects"}
}

// AssignedProjectTasksOrErr returns the AssignedProjectTasks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AssignedProjectTasksOrErr() ([]*ProjectTask, error) {
	if e.loadedTypes[10] {
		return e.AssignedProjectTasks, nil
	}
	return nil, &NotLoadedError{edge: "assignedProjectTasks"}
}

// ParticipatedProjectTasksOrErr returns the ParticipatedProjectTasks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ParticipatedProjectTasksOrErr() ([]*ProjectTask, error) {
	if e.loadedTypes[11] {
		return e.ParticipatedProjectTasks, nil
	}
	return nil, &NotLoadedError{edge: "participatedProjectTasks"}
}

// CreatedTasksOrErr returns the CreatedTasks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedTasksOrErr() ([]*ProjectTask, error) {
	if e.loadedTypes[12] {
		return e.CreatedTasks, nil
	}
	return nil, &NotLoadedError{edge: "createdTasks"}
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TokensOrErr() ([]*Token, error) {
	if e.loadedTypes[13] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// ApprovedWorkShiftsOrErr returns the ApprovedWorkShifts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ApprovedWorkShiftsOrErr() ([]*Workshift, error) {
	if e.loadedTypes[14] {
		return e.ApprovedWorkShifts, nil
	}
	return nil, &NotLoadedError{edge: "approvedWorkShifts"}
}

// WorkShiftsOrErr returns the WorkShifts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WorkShiftsOrErr() ([]*Workshift, error) {
	if e.loadedTypes[15] {
		return e.WorkShifts, nil
	}
	return nil, &NotLoadedError{edge: "workShifts"}
}

// UploadedDocumentsOrErr returns the UploadedDocuments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UploadedDocumentsOrErr() ([]*CompanyDocument, error) {
	if e.loadedTypes[16] {
		return e.UploadedDocuments, nil
	}
	return nil, &NotLoadedError{edge: "uploadedDocuments"}
}

// ApprovedDocumentsOrErr returns the ApprovedDocuments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ApprovedDocumentsOrErr() ([]*CompanyDocument, error) {
	if e.loadedTypes[17] {
		return e.ApprovedDocuments, nil
	}
	return nil, &NotLoadedError{edge: "approvedDocuments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldActive:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldFirebaseUID, user.FieldFcmToken, user.FieldEmail, user.FieldName, user.FieldAddress, user.FieldAvatar, user.FieldPhotoURL, user.FieldDepartment, user.FieldPhone, user.FieldGender:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt, user.FieldBirthdate, user.FieldLastLogin:
			values[i] = new(sql.NullTime)
		case user.ForeignKeys[0]: // user_subordinates
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				u.DeletedAt = new(time.Time)
				*u.DeletedAt = value.Time
			}
		case user.FieldFirebaseUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firebaseUID", values[i])
			} else if value.Valid {
				u.FirebaseUID = value.String
			}
		case user.FieldFcmToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fcmToken", values[i])
			} else if value.Valid {
				u.FcmToken = new(string)
				*u.FcmToken = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				u.Address = new(string)
				*u.Address = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = new(string)
				*u.Avatar = value.String
			}
		case user.FieldPhotoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photoURL", values[i])
			} else if value.Valid {
				u.PhotoURL = new(string)
				*u.PhotoURL = value.String
			}
		case user.FieldDepartment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field department", values[i])
			} else if value.Valid {
				u.Department = new(string)
				*u.Department = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = new(string)
				*u.Phone = value.String
			}
		case user.FieldBirthdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birthdate", values[i])
			} else if value.Valid {
				u.Birthdate = new(time.Time)
				*u.Birthdate = value.Time
			}
		case user.FieldLastLogin:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastLogin", values[i])
			} else if value.Valid {
				u.LastLogin = new(time.Time)
				*u.LastLogin = value.Time
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = user.Gender(value.String)
			}
		case user.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				u.Active = new(bool)
				*u.Active = value.Bool
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_subordinates", value)
			} else if value.Valid {
				u.user_subordinates = new(int)
				*u.user_subordinates = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryAccountingEntries queries the "accountingEntries" edge of the User entity.
func (u *User) QueryAccountingEntries() *AccountingEntryQuery {
	return NewUserClient(u.config).QueryAccountingEntries(u)
}

// QueryCompany queries the "company" edge of the User entity.
func (u *User) QueryCompany() *CompanyQuery {
	return NewUserClient(u.config).QueryCompany(u)
}

// QueryAssignedRoles queries the "assignedRoles" edge of the User entity.
func (u *User) QueryAssignedRoles() *UserRoleQuery {
	return NewUserClient(u.config).QueryAssignedRoles(u)
}

// QuerySubordinates queries the "subordinates" edge of the User entity.
func (u *User) QuerySubordinates() *UserQuery {
	return NewUserClient(u.config).QuerySubordinates(u)
}

// QueryLeader queries the "leader" edge of the User entity.
func (u *User) QueryLeader() *UserQuery {
	return NewUserClient(u.config).QueryLeader(u)
}

// QueryCreatedMemberSignupTokens queries the "createdMemberSignupTokens" edge of the User entity.
func (u *User) QueryCreatedMemberSignupTokens() *MemberSignupTokenQuery {
	return NewUserClient(u.config).QueryCreatedMemberSignupTokens(u)
}

// QueryEmployee queries the "employee" edge of the User entity.
func (u *User) QueryEmployee() *EmployeeQuery {
	return NewUserClient(u.config).QueryEmployee(u)
}

// QueryIssuedInvoices queries the "issuedInvoices" edge of the User entity.
func (u *User) QueryIssuedInvoices() *InvoiceQuery {
	return NewUserClient(u.config).QueryIssuedInvoices(u)
}

// QueryCreatedProjects queries the "createdProjects" edge of the User entity.
func (u *User) QueryCreatedProjects() *ProjectQuery {
	return NewUserClient(u.config).QueryCreatedProjects(u)
}

// QueryLeaderedProjects queries the "leaderedProjects" edge of the User entity.
func (u *User) QueryLeaderedProjects() *ProjectQuery {
	return NewUserClient(u.config).QueryLeaderedProjects(u)
}

// QueryAssignedProjectTasks queries the "assignedProjectTasks" edge of the User entity.
func (u *User) QueryAssignedProjectTasks() *ProjectTaskQuery {
	return NewUserClient(u.config).QueryAssignedProjectTasks(u)
}

// QueryParticipatedProjectTasks queries the "participatedProjectTasks" edge of the User entity.
func (u *User) QueryParticipatedProjectTasks() *ProjectTaskQuery {
	return NewUserClient(u.config).QueryParticipatedProjectTasks(u)
}

// QueryCreatedTasks queries the "createdTasks" edge of the User entity.
func (u *User) QueryCreatedTasks() *ProjectTaskQuery {
	return NewUserClient(u.config).QueryCreatedTasks(u)
}

// QueryTokens queries the "tokens" edge of the User entity.
func (u *User) QueryTokens() *TokenQuery {
	return NewUserClient(u.config).QueryTokens(u)
}

// QueryApprovedWorkShifts queries the "approvedWorkShifts" edge of the User entity.
func (u *User) QueryApprovedWorkShifts() *WorkshiftQuery {
	return NewUserClient(u.config).QueryApprovedWorkShifts(u)
}

// QueryWorkShifts queries the "workShifts" edge of the User entity.
func (u *User) QueryWorkShifts() *WorkshiftQuery {
	return NewUserClient(u.config).QueryWorkShifts(u)
}

// QueryUploadedDocuments queries the "uploadedDocuments" edge of the User entity.
func (u *User) QueryUploadedDocuments() *CompanyDocumentQuery {
	return NewUserClient(u.config).QueryUploadedDocuments(u)
}

// QueryApprovedDocuments queries the "approvedDocuments" edge of the User entity.
func (u *User) QueryApprovedDocuments() *CompanyDocumentQuery {
	return NewUserClient(u.config).QueryApprovedDocuments(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("generated: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := u.DeletedAt; v != nil {
		builder.WriteString("deletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("firebaseUID=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("fcmToken=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	if v := u.Address; v != nil {
		builder.WriteString("address=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Avatar; v != nil {
		builder.WriteString("avatar=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.PhotoURL; v != nil {
		builder.WriteString("photoURL=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Department; v != nil {
		builder.WriteString("department=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Phone; v != nil {
		builder.WriteString("phone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Birthdate; v != nil {
		builder.WriteString("birthdate=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.LastLogin; v != nil {
		builder.WriteString("lastLogin=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", u.Gender))
	builder.WriteString(", ")
	if v := u.Active; v != nil {
		builder.WriteString("active=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedAccountingEntries returns the AccountingEntries named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedAccountingEntries(name string) ([]*AccountingEntry, error) {
	if u.Edges.namedAccountingEntries == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedAccountingEntries[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedAccountingEntries(name string, edges ...*AccountingEntry) {
	if u.Edges.namedAccountingEntries == nil {
		u.Edges.namedAccountingEntries = make(map[string][]*AccountingEntry)
	}
	if len(edges) == 0 {
		u.Edges.namedAccountingEntries[name] = []*AccountingEntry{}
	} else {
		u.Edges.namedAccountingEntries[name] = append(u.Edges.namedAccountingEntries[name], edges...)
	}
}

// NamedCompany returns the Company named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCompany(name string) ([]*Company, error) {
	if u.Edges.namedCompany == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCompany[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCompany(name string, edges ...*Company) {
	if u.Edges.namedCompany == nil {
		u.Edges.namedCompany = make(map[string][]*Company)
	}
	if len(edges) == 0 {
		u.Edges.namedCompany[name] = []*Company{}
	} else {
		u.Edges.namedCompany[name] = append(u.Edges.namedCompany[name], edges...)
	}
}

// NamedAssignedRoles returns the AssignedRoles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedAssignedRoles(name string) ([]*UserRole, error) {
	if u.Edges.namedAssignedRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedAssignedRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedAssignedRoles(name string, edges ...*UserRole) {
	if u.Edges.namedAssignedRoles == nil {
		u.Edges.namedAssignedRoles = make(map[string][]*UserRole)
	}
	if len(edges) == 0 {
		u.Edges.namedAssignedRoles[name] = []*UserRole{}
	} else {
		u.Edges.namedAssignedRoles[name] = append(u.Edges.namedAssignedRoles[name], edges...)
	}
}

// NamedSubordinates returns the Subordinates named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedSubordinates(name string) ([]*User, error) {
	if u.Edges.namedSubordinates == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedSubordinates[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedSubordinates(name string, edges ...*User) {
	if u.Edges.namedSubordinates == nil {
		u.Edges.namedSubordinates = make(map[string][]*User)
	}
	if len(edges) == 0 {
		u.Edges.namedSubordinates[name] = []*User{}
	} else {
		u.Edges.namedSubordinates[name] = append(u.Edges.namedSubordinates[name], edges...)
	}
}

// NamedCreatedMemberSignupTokens returns the CreatedMemberSignupTokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCreatedMemberSignupTokens(name string) ([]*MemberSignupToken, error) {
	if u.Edges.namedCreatedMemberSignupTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCreatedMemberSignupTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCreatedMemberSignupTokens(name string, edges ...*MemberSignupToken) {
	if u.Edges.namedCreatedMemberSignupTokens == nil {
		u.Edges.namedCreatedMemberSignupTokens = make(map[string][]*MemberSignupToken)
	}
	if len(edges) == 0 {
		u.Edges.namedCreatedMemberSignupTokens[name] = []*MemberSignupToken{}
	} else {
		u.Edges.namedCreatedMemberSignupTokens[name] = append(u.Edges.namedCreatedMemberSignupTokens[name], edges...)
	}
}

// NamedIssuedInvoices returns the IssuedInvoices named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedIssuedInvoices(name string) ([]*Invoice, error) {
	if u.Edges.namedIssuedInvoices == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedIssuedInvoices[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedIssuedInvoices(name string, edges ...*Invoice) {
	if u.Edges.namedIssuedInvoices == nil {
		u.Edges.namedIssuedInvoices = make(map[string][]*Invoice)
	}
	if len(edges) == 0 {
		u.Edges.namedIssuedInvoices[name] = []*Invoice{}
	} else {
		u.Edges.namedIssuedInvoices[name] = append(u.Edges.namedIssuedInvoices[name], edges...)
	}
}

// NamedCreatedProjects returns the CreatedProjects named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCreatedProjects(name string) ([]*Project, error) {
	if u.Edges.namedCreatedProjects == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCreatedProjects[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCreatedProjects(name string, edges ...*Project) {
	if u.Edges.namedCreatedProjects == nil {
		u.Edges.namedCreatedProjects = make(map[string][]*Project)
	}
	if len(edges) == 0 {
		u.Edges.namedCreatedProjects[name] = []*Project{}
	} else {
		u.Edges.namedCreatedProjects[name] = append(u.Edges.namedCreatedProjects[name], edges...)
	}
}

// NamedLeaderedProjects returns the LeaderedProjects named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedLeaderedProjects(name string) ([]*Project, error) {
	if u.Edges.namedLeaderedProjects == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedLeaderedProjects[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedLeaderedProjects(name string, edges ...*Project) {
	if u.Edges.namedLeaderedProjects == nil {
		u.Edges.namedLeaderedProjects = make(map[string][]*Project)
	}
	if len(edges) == 0 {
		u.Edges.namedLeaderedProjects[name] = []*Project{}
	} else {
		u.Edges.namedLeaderedProjects[name] = append(u.Edges.namedLeaderedProjects[name], edges...)
	}
}

// NamedAssignedProjectTasks returns the AssignedProjectTasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedAssignedProjectTasks(name string) ([]*ProjectTask, error) {
	if u.Edges.namedAssignedProjectTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedAssignedProjectTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedAssignedProjectTasks(name string, edges ...*ProjectTask) {
	if u.Edges.namedAssignedProjectTasks == nil {
		u.Edges.namedAssignedProjectTasks = make(map[string][]*ProjectTask)
	}
	if len(edges) == 0 {
		u.Edges.namedAssignedProjectTasks[name] = []*ProjectTask{}
	} else {
		u.Edges.namedAssignedProjectTasks[name] = append(u.Edges.namedAssignedProjectTasks[name], edges...)
	}
}

// NamedParticipatedProjectTasks returns the ParticipatedProjectTasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedParticipatedProjectTasks(name string) ([]*ProjectTask, error) {
	if u.Edges.namedParticipatedProjectTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedParticipatedProjectTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedParticipatedProjectTasks(name string, edges ...*ProjectTask) {
	if u.Edges.namedParticipatedProjectTasks == nil {
		u.Edges.namedParticipatedProjectTasks = make(map[string][]*ProjectTask)
	}
	if len(edges) == 0 {
		u.Edges.namedParticipatedProjectTasks[name] = []*ProjectTask{}
	} else {
		u.Edges.namedParticipatedProjectTasks[name] = append(u.Edges.namedParticipatedProjectTasks[name], edges...)
	}
}

// NamedCreatedTasks returns the CreatedTasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCreatedTasks(name string) ([]*ProjectTask, error) {
	if u.Edges.namedCreatedTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCreatedTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCreatedTasks(name string, edges ...*ProjectTask) {
	if u.Edges.namedCreatedTasks == nil {
		u.Edges.namedCreatedTasks = make(map[string][]*ProjectTask)
	}
	if len(edges) == 0 {
		u.Edges.namedCreatedTasks[name] = []*ProjectTask{}
	} else {
		u.Edges.namedCreatedTasks[name] = append(u.Edges.namedCreatedTasks[name], edges...)
	}
}

// NamedTokens returns the Tokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedTokens(name string) ([]*Token, error) {
	if u.Edges.namedTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedTokens(name string, edges ...*Token) {
	if u.Edges.namedTokens == nil {
		u.Edges.namedTokens = make(map[string][]*Token)
	}
	if len(edges) == 0 {
		u.Edges.namedTokens[name] = []*Token{}
	} else {
		u.Edges.namedTokens[name] = append(u.Edges.namedTokens[name], edges...)
	}
}

// NamedApprovedWorkShifts returns the ApprovedWorkShifts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedApprovedWorkShifts(name string) ([]*Workshift, error) {
	if u.Edges.namedApprovedWorkShifts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedApprovedWorkShifts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedApprovedWorkShifts(name string, edges ...*Workshift) {
	if u.Edges.namedApprovedWorkShifts == nil {
		u.Edges.namedApprovedWorkShifts = make(map[string][]*Workshift)
	}
	if len(edges) == 0 {
		u.Edges.namedApprovedWorkShifts[name] = []*Workshift{}
	} else {
		u.Edges.namedApprovedWorkShifts[name] = append(u.Edges.namedApprovedWorkShifts[name], edges...)
	}
}

// NamedWorkShifts returns the WorkShifts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedWorkShifts(name string) ([]*Workshift, error) {
	if u.Edges.namedWorkShifts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedWorkShifts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedWorkShifts(name string, edges ...*Workshift) {
	if u.Edges.namedWorkShifts == nil {
		u.Edges.namedWorkShifts = make(map[string][]*Workshift)
	}
	if len(edges) == 0 {
		u.Edges.namedWorkShifts[name] = []*Workshift{}
	} else {
		u.Edges.namedWorkShifts[name] = append(u.Edges.namedWorkShifts[name], edges...)
	}
}

// NamedUploadedDocuments returns the UploadedDocuments named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedUploadedDocuments(name string) ([]*CompanyDocument, error) {
	if u.Edges.namedUploadedDocuments == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedUploadedDocuments[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedUploadedDocuments(name string, edges ...*CompanyDocument) {
	if u.Edges.namedUploadedDocuments == nil {
		u.Edges.namedUploadedDocuments = make(map[string][]*CompanyDocument)
	}
	if len(edges) == 0 {
		u.Edges.namedUploadedDocuments[name] = []*CompanyDocument{}
	} else {
		u.Edges.namedUploadedDocuments[name] = append(u.Edges.namedUploadedDocuments[name], edges...)
	}
}

// NamedApprovedDocuments returns the ApprovedDocuments named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedApprovedDocuments(name string) ([]*CompanyDocument, error) {
	if u.Edges.namedApprovedDocuments == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedApprovedDocuments[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedApprovedDocuments(name string, edges ...*CompanyDocument) {
	if u.Edges.namedApprovedDocuments == nil {
		u.Edges.namedApprovedDocuments = make(map[string][]*CompanyDocument)
	}
	if len(edges) == 0 {
		u.Edges.namedApprovedDocuments[name] = []*CompanyDocument{}
	} else {
		u.Edges.namedApprovedDocuments[name] = append(u.Edges.namedApprovedDocuments[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
