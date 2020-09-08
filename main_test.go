package gql_payload

import (
	"bytes"
	"reflect"
	"testing"
)


// TODO: Test cases missing: Add test cases
func TestGraphQL_GenerateRequestBody(t *testing.T) {
	type fields struct {
		q    string
		vars map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    *bytes.Buffer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphQL{
				q:    tt.fields.q,
				vars: tt.fields.vars,
			}
			got, err := g.GenerateRequestBody()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRequestBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRequestBody() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphQL_Var(t *testing.T) {
	type fields struct {
		q    string
		vars map[string]interface{}
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GraphQL
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GraphQL{
				q:    tt.fields.q,
				vars: tt.fields.vars,
			}
			if got := g.Var(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Var() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewField(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name string
		args args
		want *GraphQL
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewField(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewField() = %v, want %v", got, tt.want)
			}
		})
	}
}
