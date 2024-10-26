// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package realmv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type RealmTable interface {
	Insert(ctx context.Context, realm *Realm) error
	Update(ctx context.Context, realm *Realm) error
	Save(ctx context.Context, realm *Realm) error
	Delete(ctx context.Context, realm *Realm) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Realm, error)
	HasByName(ctx context.Context, name string) (found bool, err error)
	// GetByName returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByName(ctx context.Context, name string) (*Realm, error)
	List(ctx context.Context, prefixKey RealmIndexKey, opts ...ormlist.Option) (RealmIterator, error)
	ListRange(ctx context.Context, from, to RealmIndexKey, opts ...ormlist.Option) (RealmIterator, error)
	DeleteBy(ctx context.Context, prefixKey RealmIndexKey) error
	DeleteRange(ctx context.Context, from, to RealmIndexKey) error

	doNotImplement()
}

type RealmIterator struct {
	ormtable.Iterator
}

func (i RealmIterator) Value() (*Realm, error) {
	var realm Realm
	err := i.UnmarshalMessage(&realm)
	return &realm, err
}

type RealmIndexKey interface {
	id() uint32
	values() []interface{}
	realmIndexKey()
}

// primary key starting index..
type RealmPrimaryKey = RealmIdIndexKey

type RealmIdIndexKey struct {
	vs []interface{}
}

func (x RealmIdIndexKey) id() uint32            { return 0 }
func (x RealmIdIndexKey) values() []interface{} { return x.vs }
func (x RealmIdIndexKey) realmIndexKey()        {}

func (this RealmIdIndexKey) WithId(id uint64) RealmIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type RealmNameIndexKey struct {
	vs []interface{}
}

func (x RealmNameIndexKey) id() uint32            { return 1 }
func (x RealmNameIndexKey) values() []interface{} { return x.vs }
func (x RealmNameIndexKey) realmIndexKey()        {}

func (this RealmNameIndexKey) WithName(name string) RealmNameIndexKey {
	this.vs = []interface{}{name}
	return this
}

type realmTable struct {
	table ormtable.Table
}

func (this realmTable) Insert(ctx context.Context, realm *Realm) error {
	return this.table.Insert(ctx, realm)
}

func (this realmTable) Update(ctx context.Context, realm *Realm) error {
	return this.table.Update(ctx, realm)
}

func (this realmTable) Save(ctx context.Context, realm *Realm) error {
	return this.table.Save(ctx, realm)
}

func (this realmTable) Delete(ctx context.Context, realm *Realm) error {
	return this.table.Delete(ctx, realm)
}

func (this realmTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this realmTable) Get(ctx context.Context, id uint64) (*Realm, error) {
	var realm Realm
	found, err := this.table.PrimaryKey().Get(ctx, &realm, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &realm, nil
}

func (this realmTable) HasByName(ctx context.Context, name string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		name,
	)
}

func (this realmTable) GetByName(ctx context.Context, name string) (*Realm, error) {
	var realm Realm
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &realm,
		name,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &realm, nil
}

func (this realmTable) List(ctx context.Context, prefixKey RealmIndexKey, opts ...ormlist.Option) (RealmIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return RealmIterator{it}, err
}

func (this realmTable) ListRange(ctx context.Context, from, to RealmIndexKey, opts ...ormlist.Option) (RealmIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return RealmIterator{it}, err
}

func (this realmTable) DeleteBy(ctx context.Context, prefixKey RealmIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this realmTable) DeleteRange(ctx context.Context, from, to RealmIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this realmTable) doNotImplement() {}

var _ RealmTable = realmTable{}

func NewRealmTable(db ormtable.Schema) (RealmTable, error) {
	table := db.GetTable(&Realm{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Realm{}).ProtoReflect().Descriptor().FullName()))
	}
	return realmTable{table}, nil
}

type StateStore interface {
	RealmTable() RealmTable

	doNotImplement()
}

type stateStore struct {
	realm RealmTable
}

func (x stateStore) RealmTable() RealmTable {
	return x.realm
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	realmTable, err := NewRealmTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		realmTable,
	}, nil
}