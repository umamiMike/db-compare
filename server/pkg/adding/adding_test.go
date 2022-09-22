package adding

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDatasource(t *testing.T) {
	ds := Datasource{
		Username: "mike",
		Hostname: "localhost",
		Password: "postgres",
		DbName:   "test",
	}
	d2 := Datasource{
		Username: "mike",
		Hostname: "localfost",
		Password: "postgres",
		DbName:   "test",
	}

	mR := new(mockStorage)
	s := NewService(mR)
	err := s.AddDatasource(ds, d2)
	storedDatasources := mR.GetAllDatasources()

	assert.Len(t, storedDatasources, 2)
	require.NoError(t, err)
}

type mockStorage struct {
	datasources []Datasource
}

func (m *mockStorage) AddDatasource(ds Datasource) error {
	m.datasources = append(m.datasources, ds)
	return nil
}

func (m *mockStorage) GetAllDatasources() []Datasource {
	return m.datasources
}
