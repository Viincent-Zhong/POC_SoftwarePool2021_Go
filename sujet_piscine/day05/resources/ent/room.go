// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-message/ent/room"
	"strings"

	"github.com/google/uuid"
)

// Room is the model entity for the Room schema.
type Room struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges RoomEdges `json:"edges"`
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[0] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Room", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		}
	}
	return nil
}

// QueryMessages queries the "messages" edge of the Room entity.
func (r *Room) QueryMessages() *MessageQuery {
	return (&RoomClient{config: r.config}).QueryMessages(r)
}

// Update returns a builder for updating this Room.
// Note that you need to call Room.Unwrap() before calling this method if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return (&RoomClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Room entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Rooms is a parsable slice of Room.
type Rooms []*Room

func (r Rooms) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
