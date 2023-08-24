package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type API struct {
	Port       string `json:"port"`
	URL        string `json:"url"`
	BCryptCost int64  `json:"bcrypt_cost"`
	JWTSecret  string `json:"jwt_secret"`
	WorkDir    string `json:"work_dir"`
}

func newAPI() *API {
	bcryptCost, err := strconv.ParseInt(goEnvVar("BCRYPT_COST"), 10, 36)
	if err != nil {
		bcryptCost = 10
	}
	wd, _ := os.Getwd()
	api := API{
		Port:       goEnvVar("API_PORT"),
		URL:        goEnvVar("API_URL"),
		BCryptCost: bcryptCost,
		JWTSecret:  goEnvVar("JWT_SECRET"),
		WorkDir:    fmt.Sprintf("%s/cmd/todolist/docs", wd),
	}
	if api.Port == "" {
		api.Port = "4000"
	}
	return &api
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func newDatabase() *Database {
	db := Database{
		Host:     goEnvVar("DB_HOST"),
		Port:     goEnvVar("DB_PORT"),
		User:     goEnvVar("DB_USER"),
		Password: goEnvVar("DB_PASSWORD"),
		Name:     goEnvVar("DB_NAME"),
	}

	if db.Host == "" ||
		db.Port == "" ||
		db.User == "" ||
		db.Password == "" ||
		db.Name == "" {
		db.Host = "localhost"
		db.Port = "5432"
		db.User = "postgres"
		db.Password = "admin"
		db.Name = "todo_list"
	}
	return &db
}

func (d *Database) ConnStr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=America/Sao_Paulo",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	)
}

type AppConf struct {
	API      `json:"api"`
	Database `json:"database"`
}

func NewAppConf() *AppConf {
	return &AppConf{
		API:      *newAPI(),
		Database: *newDatabase(),
	}
}

func goEnvVar(k string) string {
	godotenv.Load()
	return os.Getenv(k)
}
