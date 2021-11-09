package fileutils_test

import (
	"testing"

	fileutils "github.com/maybeineedhelp/utils/pkg/file"
	jsonutils "github.com/maybeineedhelp/utils/pkg/json"

	"github.com/go-playground/assert"
)

type Config struct {
	Database  *CongfigDatabase `json:"database"`
	AppID     string           `json:"app_id"`
	AppSecret string           `json:"app_secret"`
}

type CongfigDatabase struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	SSLMode  string `json:"sslmode"`
}

func TestParse(t *testing.T) {
	config := &Config{}
	err := fileutils.Parse("../not_found", config)
	assert.NotEqual(t, err, nil)

	config = &Config{}
	err = fileutils.Parse("../fakedata/file/config.json", config)
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonutils.MarshalJSONOrDie(config), `{"database":{"host":"localhost","port":5432,"dbname":"postgres","user":"postgres","password":"postgres","sslmode":"disable"},"app_id":"xxx","app_secret":"xxxxx"}`)
}
