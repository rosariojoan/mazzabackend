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
	// FcmToken holds the value of the "fcmToken" field.
	FcmToken *string `json:"fcmToken,omitempty"`
	// Email holds the value of the "email" field.
	Email *string `json:"email,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled *bool `json:"disabled,omitempty"`
	// NotVerified holds the value of the "notVerified" field.
	NotVerified *bool `json:"notVerified,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// AccountingEntries holds the value of the accountingEntries edge.
	AccountingEntries []*AccountingEntry `json:"accountingEntries,omitempty"`
	// Company holds the value of the company edge.
	Company []*Company `json:"company,omitempty"`
	// a user should be assigned to at least one role in the company
	AssignedRoles []*UserRole `json:"assignedRoles,omitempty"`
	// Tasks created by this user
	CreatedTasks []*Worktask `json:"createdTasks,omitempty"`
	// Employee holds the value of the employee edge.
	Employee *Employee `json:"employee,omitempty"`
	// Represent the projects created by the user
	CreatedProjects []*Project `json:"createdProjects,omitempty"`
	// Represent the projects leadered or supervised by the user
	LeaderedProjects []*Project `json:"leaderedProjects,omitempty"`
	// These are the project tasks assigned to the user and he is responsible for them
	AssignedProjectTasks []*ProjectTask `json:"assignedProjectTasks,omitempty"`
	// These are the project tasks in which the user is a member. E.g. a meeting
	ParticipatedProjectTasks []*ProjectTask `json:"participatedProjectTasks,omitempty"`
	// Tokens holds the value of the tokens edge.
	Tokens []*Token `json:"tokens,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [10]bool
	// totalCount holds the count of the edges above.
	totalCount [10]map[string]int

	namedAccountingEntries        map[string][]*AccountingEntry
	namedCompany                  map[string][]*Company
	namedAssignedRoles            map[string][]*UserRole
	namedCreatedTasks             map[string][]*Worktask
	namedCreatedProjects          map[string][]*Project
	namedLeaderedProjects         map[string][]*Project
	namedAssignedProjectTasks     map[string][]*ProjectTask
	namedParticipatedProjectTasks map[string][]*ProjectTask
	namedTokens                   map[string][]*Token
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

// CreatedTasksOrErr returns the CreatedTasks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedTasksOrErr() ([]*Worktask, error) {
	if e.loadedTypes[3] {
		return e.CreatedTasks, nil
	}
	return nil, &NotLoadedError{edge: "createdTasks"}
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) EmployeeOrErr() (*Employee, error) {
	if e.Employee != nil {
		return e.Employee, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: employee.Label}
	}
	return nil, &NotLoadedError{edge: "employee"}
}

// CreatedProjectsOrErr returns the CreatedProjects value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[5] {
		return e.CreatedProjects, nil
	}
	return nil, &NotLoadedError{edge: "createdProjects"}
}

// LeaderedProjectsOrErr returns the LeaderedProjects value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LeaderedProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[6] {
		return e.LeaderedProjects, nil
	}
	return nil, &NotLoadedError{edge: "leaderedProjects"}
}

// AssignedProjectTasksOrErr returns the AssignedProjectTasks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AssignedProjectTasksOrErr() ([]*ProjectTask, error) {
	if e.loadedTypes[7] {
		return e.AssignedProjectTasks, nil
	}
	return nil, &NotLoadedError{edge: "assignedProjectTasks"}
}

// ParticipatedProjectTasksOrErr returns the ParticipatedProjectTasks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ParticipatedProjectTasksOrErr() ([]*ProjectTask, error) {
	if e.loadedTypes[8] {
		return e.ParticipatedProjectTasks, nil
	}
	return nil, &NotLoadedError{edge: "participatedProjectTasks"}
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TokensOrErr() ([]*Token, error) {
	if e.loadedTypes[9] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldDisabled, user.FieldNotVerified:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldFcmToken, user.FieldEmail, user.FieldName, user.FieldPassword, user.FieldUsername:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
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
				u.Email = new(string)
				*u.Email = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				u.Disabled = new(bool)
				*u.Disabled = value.Bool
			}
		case user.FieldNotVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field notVerified", values[i])
			} else if value.Valid {
				u.NotVerified = new(bool)
				*u.NotVerified = value.Bool
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

// QueryCreatedTasks queries the "createdTasks" edge of the User entity.
func (u *User) QueryCreatedTasks() *WorktaskQuery {
	return NewUserClient(u.config).QueryCreatedTasks(u)
}

// QueryEmployee queries the "employee" edge of the User entity.
func (u *User) QueryEmployee() *EmployeeQuery {
	return NewUserClient(u.config).QueryEmployee(u)
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

// QueryTokens queries the "tokens" edge of the User entity.
func (u *User) QueryTokens() *TokenQuery {
	return NewUserClient(u.config).QueryTokens(u)
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
	if v := u.FcmToken; v != nil {
		builder.WriteString("fcmToken=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	if v := u.Disabled; v != nil {
		builder.WriteString("disabled=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.NotVerified; v != nil {
		builder.WriteString("notVerified=")
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

// NamedCreatedTasks returns the CreatedTasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCreatedTasks(name string) ([]*Worktask, error) {
	if u.Edges.namedCreatedTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCreatedTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCreatedTasks(name string, edges ...*Worktask) {
	if u.Edges.namedCreatedTasks == nil {
		u.Edges.namedCreatedTasks = make(map[string][]*Worktask)
	}
	if len(edges) == 0 {
		u.Edges.namedCreatedTasks[name] = []*Worktask{}
	} else {
		u.Edges.namedCreatedTasks[name] = append(u.Edges.namedCreatedTasks[name], edges...)
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

// Users is a parsable slice of User.
type Users []*User
