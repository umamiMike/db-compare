package main

type Datasource struct {
	//db key
	Name        string
	ID          string
	Credentials DatasourceCredentials
}
type DatasourceCredentials struct {
	Username string `json:"username"`
	Hostname string `json:"host"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

func (ds *Datasource) Get(id string) *Datasource {

	return ds
}
