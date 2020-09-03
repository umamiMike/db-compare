package storage

type Datasource struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Hostname string `json:"host"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}
