// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"mazza/ent/accountingentry"
	"mazza/ent/cashmovement"
	"mazza/ent/company"
	"mazza/ent/customer"
	"mazza/ent/employee"
	"mazza/ent/file"
	"mazza/ent/payable"
	"mazza/ent/product"
	"mazza/ent/productmovement"
	"mazza/ent/receivable"
	"mazza/ent/supplier"
	"mazza/ent/token"
	"mazza/ent/treasury"
	"mazza/ent/user"
	"mazza/ent/userrole"
	"mazza/ent/workshift"
	"mazza/ent/worktag"
	"mazza/ent/worktask"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

// IsNode implements the Node interface check for GQLGen.
func (n *AccountingEntry) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *CashMovement) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Company) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Customer) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Employee) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *File) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Payable) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Product) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *ProductMovement) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Receivable) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Supplier) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Token) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Treasury) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *User) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *UserRole) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Workshift) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Worktag) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Worktask) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case accountingentry.Table:
		query := c.AccountingEntry.Query().
			Where(accountingentry.ID(id))
		query, err := query.CollectFields(ctx, "AccountingEntry")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case cashmovement.Table:
		query := c.CashMovement.Query().
			Where(cashmovement.ID(id))
		query, err := query.CollectFields(ctx, "CashMovement")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case company.Table:
		query := c.Company.Query().
			Where(company.ID(id))
		query, err := query.CollectFields(ctx, "Company")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case customer.Table:
		query := c.Customer.Query().
			Where(customer.ID(id))
		query, err := query.CollectFields(ctx, "Customer")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case employee.Table:
		query := c.Employee.Query().
			Where(employee.ID(id))
		query, err := query.CollectFields(ctx, "Employee")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case file.Table:
		query := c.File.Query().
			Where(file.ID(id))
		query, err := query.CollectFields(ctx, "File")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case payable.Table:
		query := c.Payable.Query().
			Where(payable.ID(id))
		query, err := query.CollectFields(ctx, "Payable")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case product.Table:
		query := c.Product.Query().
			Where(product.ID(id))
		query, err := query.CollectFields(ctx, "Product")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case productmovement.Table:
		query := c.ProductMovement.Query().
			Where(productmovement.ID(id))
		query, err := query.CollectFields(ctx, "ProductMovement")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case receivable.Table:
		query := c.Receivable.Query().
			Where(receivable.ID(id))
		query, err := query.CollectFields(ctx, "Receivable")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case supplier.Table:
		query := c.Supplier.Query().
			Where(supplier.ID(id))
		query, err := query.CollectFields(ctx, "Supplier")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case token.Table:
		query := c.Token.Query().
			Where(token.ID(id))
		query, err := query.CollectFields(ctx, "Token")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case treasury.Table:
		query := c.Treasury.Query().
			Where(treasury.ID(id))
		query, err := query.CollectFields(ctx, "Treasury")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case userrole.Table:
		query := c.UserRole.Query().
			Where(userrole.ID(id))
		query, err := query.CollectFields(ctx, "UserRole")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case workshift.Table:
		query := c.Workshift.Query().
			Where(workshift.ID(id))
		query, err := query.CollectFields(ctx, "Workshift")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case worktag.Table:
		query := c.Worktag.Query().
			Where(worktag.ID(id))
		query, err := query.CollectFields(ctx, "Worktag")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case worktask.Table:
		query := c.Worktask.Query().
			Where(worktask.ID(id))
		query, err := query.CollectFields(ctx, "Worktask")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case accountingentry.Table:
		query := c.AccountingEntry.Query().
			Where(accountingentry.IDIn(ids...))
		query, err := query.CollectFields(ctx, "AccountingEntry")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case cashmovement.Table:
		query := c.CashMovement.Query().
			Where(cashmovement.IDIn(ids...))
		query, err := query.CollectFields(ctx, "CashMovement")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case company.Table:
		query := c.Company.Query().
			Where(company.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Company")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case customer.Table:
		query := c.Customer.Query().
			Where(customer.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Customer")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case employee.Table:
		query := c.Employee.Query().
			Where(employee.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Employee")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case file.Table:
		query := c.File.Query().
			Where(file.IDIn(ids...))
		query, err := query.CollectFields(ctx, "File")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case payable.Table:
		query := c.Payable.Query().
			Where(payable.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Payable")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case product.Table:
		query := c.Product.Query().
			Where(product.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Product")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case productmovement.Table:
		query := c.ProductMovement.Query().
			Where(productmovement.IDIn(ids...))
		query, err := query.CollectFields(ctx, "ProductMovement")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case receivable.Table:
		query := c.Receivable.Query().
			Where(receivable.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Receivable")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case supplier.Table:
		query := c.Supplier.Query().
			Where(supplier.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Supplier")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case token.Table:
		query := c.Token.Query().
			Where(token.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Token")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case treasury.Table:
		query := c.Treasury.Query().
			Where(treasury.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Treasury")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case userrole.Table:
		query := c.UserRole.Query().
			Where(userrole.IDIn(ids...))
		query, err := query.CollectFields(ctx, "UserRole")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workshift.Table:
		query := c.Workshift.Query().
			Where(workshift.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Workshift")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case worktag.Table:
		query := c.Worktag.Query().
			Where(worktag.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Worktag")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case worktask.Table:
		query := c.Worktask.Query().
			Where(worktask.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Worktask")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
