package badger

type Datasource struct {
	ID       string `json:"string"`
	Username string `json:"username"`
	Hostname string `json:"host"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}
