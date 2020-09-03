package adding

type Datasource struct {
	Username string `json:"username"`
	Hostname string `json:"host"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}
