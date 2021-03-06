// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-message/ent/message"
	"go-message/ent/predicate"
	"go-message/ent/room"
	"go-message/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMessage = "Message"
	TypeRoom    = "Room"
	TypeUser    = "User"
)

// MessageMutation represents an operation that mutates the Message nodes in the graph.
type MessageMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	message       *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	user          *uuid.UUID
	cleareduser   bool
	room          *uuid.UUID
	clearedroom   bool
	done          bool
	oldValue      func(context.Context) (*Message, error)
	predicates    []predicate.Message
}

var _ ent.Mutation = (*MessageMutation)(nil)

// messageOption allows management of the mutation configuration using functional options.
type messageOption func(*MessageMutation)

// newMessageMutation creates new mutation for the Message entity.
func newMessageMutation(c config, op Op, opts ...messageOption) *MessageMutation {
	m := &MessageMutation{
		config:        c,
		op:            op,
		typ:           TypeMessage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMessageID sets the ID field of the mutation.
func withMessageID(id uuid.UUID) messageOption {
	return func(m *MessageMutation) {
		var (
			err   error
			once  sync.Once
			value *Message
		)
		m.oldValue = func(ctx context.Context) (*Message, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Message.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMessage sets the old Message of the mutation.
func withMessage(node *Message) messageOption {
	return func(m *MessageMutation) {
		m.oldValue = func(context.Context) (*Message, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MessageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MessageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Message entities.
func (m *MessageMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *MessageMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetMessage sets the "message" field.
func (m *MessageMutation) SetMessage(s string) {
	m.message = &s
}

// Message returns the value of the "message" field in the mutation.
func (m *MessageMutation) Message() (r string, exists bool) {
	v := m.message
	if v == nil {
		return
	}
	return *v, true
}

// OldMessage returns the old "message" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessage: %w", err)
	}
	return oldValue.Message, nil
}

// ResetMessage resets all changes to the "message" field.
func (m *MessageMutation) ResetMessage() {
	m.message = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *MessageMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *MessageMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *MessageMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUserID sets the "user" edge to the User entity by id.
func (m *MessageMutation) SetUserID(id uuid.UUID) {
	m.user = &id
}

// ClearUser clears the "user" edge to the User entity.
func (m *MessageMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared returns if the "user" edge to the User entity was cleared.
func (m *MessageMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the "user" edge ID in the mutation.
func (m *MessageMutation) UserID() (id uuid.UUID, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *MessageMutation) UserIDs() (ids []uuid.UUID) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *MessageMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// SetRoomID sets the "room" edge to the Room entity by id.
func (m *MessageMutation) SetRoomID(id uuid.UUID) {
	m.room = &id
}

// ClearRoom clears the "room" edge to the Room entity.
func (m *MessageMutation) ClearRoom() {
	m.clearedroom = true
}

// RoomCleared returns if the "room" edge to the Room entity was cleared.
func (m *MessageMutation) RoomCleared() bool {
	return m.clearedroom
}

// RoomID returns the "room" edge ID in the mutation.
func (m *MessageMutation) RoomID() (id uuid.UUID, exists bool) {
	if m.room != nil {
		return *m.room, true
	}
	return
}

// RoomIDs returns the "room" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// RoomID instead. It exists only for internal usage by the builders.
func (m *MessageMutation) RoomIDs() (ids []uuid.UUID) {
	if id := m.room; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetRoom resets all changes to the "room" edge.
func (m *MessageMutation) ResetRoom() {
	m.room = nil
	m.clearedroom = false
}

// Op returns the operation name.
func (m *MessageMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Message).
func (m *MessageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MessageMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.message != nil {
		fields = append(fields, message.FieldMessage)
	}
	if m.created_at != nil {
		fields = append(fields, message.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MessageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case message.FieldMessage:
		return m.Message()
	case message.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MessageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case message.FieldMessage:
		return m.OldMessage(ctx)
	case message.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Message field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case message.FieldMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessage(v)
		return nil
	case message.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MessageMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MessageMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Message numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MessageMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MessageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MessageMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Message nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MessageMutation) ResetField(name string) error {
	switch name {
	case message.FieldMessage:
		m.ResetMessage()
		return nil
	case message.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MessageMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.user != nil {
		edges = append(edges, message.EdgeUser)
	}
	if m.room != nil {
		edges = append(edges, message.EdgeRoom)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MessageMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case message.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	case message.EdgeRoom:
		if id := m.room; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MessageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MessageMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MessageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.cleareduser {
		edges = append(edges, message.EdgeUser)
	}
	if m.clearedroom {
		edges = append(edges, message.EdgeRoom)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MessageMutation) EdgeCleared(name string) bool {
	switch name {
	case message.EdgeUser:
		return m.cleareduser
	case message.EdgeRoom:
		return m.clearedroom
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MessageMutation) ClearEdge(name string) error {
	switch name {
	case message.EdgeUser:
		m.ClearUser()
		return nil
	case message.EdgeRoom:
		m.ClearRoom()
		return nil
	}
	return fmt.Errorf("unknown Message unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MessageMutation) ResetEdge(name string) error {
	switch name {
	case message.EdgeUser:
		m.ResetUser()
		return nil
	case message.EdgeRoom:
		m.ResetRoom()
		return nil
	}
	return fmt.Errorf("unknown Message edge %s", name)
}

// RoomMutation represents an operation that mutates the Room nodes in the graph.
type RoomMutation struct {
	config
	op              Op
	typ             string
	id              *uuid.UUID
	clearedFields   map[string]struct{}
	messages        map[uuid.UUID]struct{}
	removedmessages map[uuid.UUID]struct{}
	clearedmessages bool
	done            bool
	oldValue        func(context.Context) (*Room, error)
	predicates      []predicate.Room
}

var _ ent.Mutation = (*RoomMutation)(nil)

// roomOption allows management of the mutation configuration using functional options.
type roomOption func(*RoomMutation)

// newRoomMutation creates new mutation for the Room entity.
func newRoomMutation(c config, op Op, opts ...roomOption) *RoomMutation {
	m := &RoomMutation{
		config:        c,
		op:            op,
		typ:           TypeRoom,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRoomID sets the ID field of the mutation.
func withRoomID(id uuid.UUID) roomOption {
	return func(m *RoomMutation) {
		var (
			err   error
			once  sync.Once
			value *Room
		)
		m.oldValue = func(ctx context.Context) (*Room, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Room.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRoom sets the old Room of the mutation.
func withRoom(node *Room) roomOption {
	return func(m *RoomMutation) {
		m.oldValue = func(context.Context) (*Room, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RoomMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RoomMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Room entities.
func (m *RoomMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *RoomMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// AddMessageIDs adds the "messages" edge to the Message entity by ids.
func (m *RoomMutation) AddMessageIDs(ids ...uuid.UUID) {
	if m.messages == nil {
		m.messages = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.messages[ids[i]] = struct{}{}
	}
}

// ClearMessages clears the "messages" edge to the Message entity.
func (m *RoomMutation) ClearMessages() {
	m.clearedmessages = true
}

// MessagesCleared returns if the "messages" edge to the Message entity was cleared.
func (m *RoomMutation) MessagesCleared() bool {
	return m.clearedmessages
}

// RemoveMessageIDs removes the "messages" edge to the Message entity by IDs.
func (m *RoomMutation) RemoveMessageIDs(ids ...uuid.UUID) {
	if m.removedmessages == nil {
		m.removedmessages = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedmessages[ids[i]] = struct{}{}
	}
}

// RemovedMessages returns the removed IDs of the "messages" edge to the Message entity.
func (m *RoomMutation) RemovedMessagesIDs() (ids []uuid.UUID) {
	for id := range m.removedmessages {
		ids = append(ids, id)
	}
	return
}

// MessagesIDs returns the "messages" edge IDs in the mutation.
func (m *RoomMutation) MessagesIDs() (ids []uuid.UUID) {
	for id := range m.messages {
		ids = append(ids, id)
	}
	return
}

// ResetMessages resets all changes to the "messages" edge.
func (m *RoomMutation) ResetMessages() {
	m.messages = nil
	m.clearedmessages = false
	m.removedmessages = nil
}

// Op returns the operation name.
func (m *RoomMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Room).
func (m *RoomMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RoomMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RoomMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RoomMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Room field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoomMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Room field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RoomMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RoomMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoomMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Room numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RoomMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RoomMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RoomMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Room nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RoomMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Room field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RoomMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.messages != nil {
		edges = append(edges, room.EdgeMessages)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RoomMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case room.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.messages))
		for id := range m.messages {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RoomMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedmessages != nil {
		edges = append(edges, room.EdgeMessages)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RoomMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case room.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.removedmessages))
		for id := range m.removedmessages {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RoomMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedmessages {
		edges = append(edges, room.EdgeMessages)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RoomMutation) EdgeCleared(name string) bool {
	switch name {
	case room.EdgeMessages:
		return m.clearedmessages
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RoomMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Room unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RoomMutation) ResetEdge(name string) error {
	switch name {
	case room.EdgeMessages:
		m.ResetMessages()
		return nil
	}
	return fmt.Errorf("unknown Room edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op              Op
	typ             string
	id              *uuid.UUID
	username        *string
	password        *string
	profilePicture  *string
	clearedFields   map[string]struct{}
	messages        map[uuid.UUID]struct{}
	removedmessages map[uuid.UUID]struct{}
	clearedmessages bool
	done            bool
	oldValue        func(context.Context) (*User, error)
	predicates      []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetProfilePicture sets the "profilePicture" field.
func (m *UserMutation) SetProfilePicture(s string) {
	m.profilePicture = &s
}

// ProfilePicture returns the value of the "profilePicture" field in the mutation.
func (m *UserMutation) ProfilePicture() (r string, exists bool) {
	v := m.profilePicture
	if v == nil {
		return
	}
	return *v, true
}

// OldProfilePicture returns the old "profilePicture" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldProfilePicture(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldProfilePicture is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldProfilePicture requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProfilePicture: %w", err)
	}
	return oldValue.ProfilePicture, nil
}

// ResetProfilePicture resets all changes to the "profilePicture" field.
func (m *UserMutation) ResetProfilePicture() {
	m.profilePicture = nil
}

// AddMessageIDs adds the "messages" edge to the Message entity by ids.
func (m *UserMutation) AddMessageIDs(ids ...uuid.UUID) {
	if m.messages == nil {
		m.messages = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.messages[ids[i]] = struct{}{}
	}
}

// ClearMessages clears the "messages" edge to the Message entity.
func (m *UserMutation) ClearMessages() {
	m.clearedmessages = true
}

// MessagesCleared returns if the "messages" edge to the Message entity was cleared.
func (m *UserMutation) MessagesCleared() bool {
	return m.clearedmessages
}

// RemoveMessageIDs removes the "messages" edge to the Message entity by IDs.
func (m *UserMutation) RemoveMessageIDs(ids ...uuid.UUID) {
	if m.removedmessages == nil {
		m.removedmessages = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedmessages[ids[i]] = struct{}{}
	}
}

// RemovedMessages returns the removed IDs of the "messages" edge to the Message entity.
func (m *UserMutation) RemovedMessagesIDs() (ids []uuid.UUID) {
	for id := range m.removedmessages {
		ids = append(ids, id)
	}
	return
}

// MessagesIDs returns the "messages" edge IDs in the mutation.
func (m *UserMutation) MessagesIDs() (ids []uuid.UUID) {
	for id := range m.messages {
		ids = append(ids, id)
	}
	return
}

// ResetMessages resets all changes to the "messages" edge.
func (m *UserMutation) ResetMessages() {
	m.messages = nil
	m.clearedmessages = false
	m.removedmessages = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.profilePicture != nil {
		fields = append(fields, user.FieldProfilePicture)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUsername:
		return m.Username()
	case user.FieldPassword:
		return m.Password()
	case user.FieldProfilePicture:
		return m.ProfilePicture()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldProfilePicture:
		return m.OldProfilePicture(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldProfilePicture:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProfilePicture(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldProfilePicture:
		m.ResetProfilePicture()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.messages != nil {
		edges = append(edges, user.EdgeMessages)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.messages))
		for id := range m.messages {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedmessages != nil {
		edges = append(edges, user.EdgeMessages)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.removedmessages))
		for id := range m.removedmessages {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedmessages {
		edges = append(edges, user.EdgeMessages)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeMessages:
		return m.clearedmessages
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeMessages:
		m.ResetMessages()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
