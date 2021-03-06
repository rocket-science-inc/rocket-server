package types

import (
	"time"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
	"errors"
)

type Event struct {
	ID      		string   	`json:"id"`
	Title    		string 		`json:"title"`
	Info    		string 		`json:"info"`
	CreatedAt 		time.Time	`json:"created"`
	UpdatedAt 		time.Time	`json:"updated"`
	DeletedAt   	*time.Time 	`json:"deleted"`
}

type NewEvent struct {
	Title			string		`json:"title"`
	Info			string		`json:"info"`
}

// MarshalTimestamp marshal custom Timestamp scalar for GraphQL schema
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix()
	if timestamp < 0 {
		timestamp = 0
	}
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}
// UnmarshalTimestamp unmarshal custom Timestamp scalar for GraphQL schema
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}
