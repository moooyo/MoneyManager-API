package model

type App struct {
	LogType   string
	LogPath   string
	WallPaper string
	Debug     bool
	Port      uint
}

type Database struct {
	Type     string
	FilePath string
	Host     string
	Port     uint
	UserName string
	Password string
	DBName   string
}

type API struct {
	Type      string
	Host      string
	Secret    string
	UserAgent string
}

type Config struct {
	App      App
	Database Database
	API      API
}
