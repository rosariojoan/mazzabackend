// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/membersignuptoken"
	"mazza/ent/generated/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MemberSignupToken is the model entity for the MemberSignupToken schema.
type MemberSignupToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deletedAt" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email *string `json:"email,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Must be an emoji
	Avatar *string `json:"avatar,omitempty"`
	// Role holds the value of the "role" field.
	Role membersignuptoken.Role `json:"role,omitempty"`
	// A description about the user and its role of this user
	Note string `json:"note,omitempty"`
	// NumberAccessed holds the value of the "numberAccessed" field.
	NumberAccessed int `json:"numberAccessed,omitempty"`
	// ExpiresAt holds the value of the "expiresAt" field.
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
	// AlreadyUsed holds the value of the "alreadyUsed" field.
	AlreadyUsed bool `json:"alreadyUsed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberSignupTokenQuery when eager-loading is set.
	Edges                             MemberSignupTokenEdges `json:"edges"`
	company_member_signup_tokens      *int
	user_created_member_signup_tokens *int
	selectValues                      sql.SelectValues
}

// MemberSignupTokenEdges holds the relations/edges for other nodes in the graph.
type MemberSignupTokenEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// CreatedBy holds the value of the createdBy edge.
	CreatedBy *User `json:"createdBy,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberSignupTokenEdges) CompanyOrErr() (*Company, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: company.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// CreatedByOrErr returns the CreatedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberSignupTokenEdges) CreatedByOrErr() (*User, error) {
	if e.CreatedBy != nil {
		return e.CreatedBy, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "createdBy"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MemberSignupToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case membersignuptoken.FieldAlreadyUsed:
			values[i] = new(sql.NullBool)
		case membersignuptoken.FieldID, membersignuptoken.FieldNumberAccessed:
			values[i] = new(sql.NullInt64)
		case membersignuptoken.FieldName, membersignuptoken.FieldEmail, membersignuptoken.FieldToken, membersignuptoken.FieldAvatar, membersignuptoken.FieldRole, membersignuptoken.FieldNote:
			values[i] = new(sql.NullString)
		case membersignuptoken.FieldCreatedAt, membersignuptoken.FieldUpdatedAt, membersignuptoken.FieldDeletedAt, membersignuptoken.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		case membersignuptoken.ForeignKeys[0]: // company_member_signup_tokens
			values[i] = new(sql.NullInt64)
		case membersignuptoken.ForeignKeys[1]: // user_created_member_signup_tokens
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MemberSignupToken fields.
func (mst *MemberSignupToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case membersignuptoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mst.ID = int(value.Int64)
		case membersignuptoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				mst.CreatedAt = value.Time
			}
		case membersignuptoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				mst.UpdatedAt = value.Time
			}
		case membersignuptoken.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				mst.DeletedAt = new(time.Time)
				*mst.DeletedAt = value.Time
			}
		case membersignuptoken.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mst.Name = value.String
			}
		case membersignuptoken.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				mst.Email = new(string)
				*mst.Email = value.String
			}
		case membersignuptoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				mst.Token = value.String
			}
		case membersignuptoken.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				mst.Avatar = new(string)
				*mst.Avatar = value.String
			}
		case membersignuptoken.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				mst.Role = membersignuptoken.Role(value.String)
			}
		case membersignuptoken.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				mst.Note = value.String
			}
		case membersignuptoken.FieldNumberAccessed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field numberAccessed", values[i])
			} else if value.Valid {
				mst.NumberAccessed = int(value.Int64)
			}
		case membersignuptoken.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiresAt", values[i])
			} else if value.Valid {
				mst.ExpiresAt = value.Time
			}
		case membersignuptoken.FieldAlreadyUsed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field alreadyUsed", values[i])
			} else if value.Valid {
				mst.AlreadyUsed = value.Bool
			}
		case membersignuptoken.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_member_signup_tokens", value)
			} else if value.Valid {
				mst.company_member_signup_tokens = new(int)
				*mst.company_member_signup_tokens = int(value.Int64)
			}
		case membersignuptoken.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_created_member_signup_tokens", value)
			} else if value.Valid {
				mst.user_created_member_signup_tokens = new(int)
				*mst.user_created_member_signup_tokens = int(value.Int64)
			}
		default:
			mst.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MemberSignupToken.
// This includes values selected through modifiers, order, etc.
func (mst *MemberSignupToken) Value(name string) (ent.Value, error) {
	return mst.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the MemberSignupToken entity.
func (mst *MemberSignupToken) QueryCompany() *CompanyQuery {
	return NewMemberSignupTokenClient(mst.config).QueryCompany(mst)
}

// QueryCreatedBy queries the "createdBy" edge of the MemberSignupToken entity.
func (mst *MemberSignupToken) QueryCreatedBy() *UserQuery {
	return NewMemberSignupTokenClient(mst.config).QueryCreatedBy(mst)
}

// Update returns a builder for updating this MemberSignupToken.
// Note that you need to call MemberSignupToken.Unwrap() before calling this method if this MemberSignupToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (mst *MemberSignupToken) Update() *MemberSignupTokenUpdateOne {
	return NewMemberSignupTokenClient(mst.config).UpdateOne(mst)
}

// Unwrap unwraps the MemberSignupToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mst *MemberSignupToken) Unwrap() *MemberSignupToken {
	_tx, ok := mst.config.driver.(*txDriver)
	if !ok {
		panic("generated: MemberSignupToken is not a transactional entity")
	}
	mst.config.driver = _tx.drv
	return mst
}

// String implements the fmt.Stringer.
func (mst *MemberSignupToken) String() string {
	var builder strings.Builder
	builder.WriteString("MemberSignupToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mst.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(mst.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(mst.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := mst.DeletedAt; v != nil {
		builder.WriteString("deletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(mst.Name)
	builder.WriteString(", ")
	if v := mst.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(mst.Token)
	builder.WriteString(", ")
	if v := mst.Avatar; v != nil {
		builder.WriteString("avatar=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", mst.Role))
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(mst.Note)
	builder.WriteString(", ")
	builder.WriteString("numberAccessed=")
	builder.WriteString(fmt.Sprintf("%v", mst.NumberAccessed))
	builder.WriteString(", ")
	builder.WriteString("expiresAt=")
	builder.WriteString(mst.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("alreadyUsed=")
	builder.WriteString(fmt.Sprintf("%v", mst.AlreadyUsed))
	builder.WriteByte(')')
	return builder.String()
}

// MemberSignupTokens is a parsable slice of MemberSignupToken.
type MemberSignupTokens []*MemberSignupToken
