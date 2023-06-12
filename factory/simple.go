package factory

type StudentConfig struct {
	DB   string
	Url  string
	Port int
}

// 將 struct 需要的資料，用參數傳入，並返回
func GetStudentDao(db, url string, port int) *StudentConfig {
	return &StudentConfig{
		DB:   db,
		Url:  url,
		Port: port,
	}
}
