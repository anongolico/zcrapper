package config

import (
	"net/http"

	"github.com/spf13/viper"
)

var (
	// Domain string ``

	// Sites Multiple sites can be added. Each time the program runs,
	// it reads existing sites, stores them in this map, and asks the
	// user to choose the desired site to download from
	// Sites = map[string]Site{}

	MainUrl  string
	MediaUrl string

	IdentityCookie = http.Cookie{
		Secure:   true,
		HttpOnly: true,
	}

	BufferSize int

	MaxParallelDownloads int
)

// type Site struct {
// 	Url      string
// 	MediaUrl string
// }

func ReadUrlParameters() error {
	vp := viper.New()

	// Tell viper where the configuration file is
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(".")
	vp.AddConfigPath("../.")

	err := vp.ReadInConfig()

	if err != nil {
		return err
	}

	// Reads the url settings from config file
	MainUrl = vp.GetString("MainUrl")
	MediaUrl = vp.GetString("MediaUrl")

	// Completes the cookie's values
	IdentityCookie.Name = vp.GetString("CookieName")
	IdentityCookie.Value = vp.GetString("CookieValue")

	// The buffer size is used to read the response size
	BufferSize = vp.GetInt("BufferSize")

	MaxParallelDownloads = vp.GetInt("MaxParallelDownloads")

	return nil
}
