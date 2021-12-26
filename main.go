package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Cannot load .env file")
	}
	fmt.Println("Membuka konfigurasi menggunakan file .env")
	fmt.Println("==========================================")
	fmt.Println(os.Getenv("APP-NAME"))
	fmt.Println(os.Getenv("APP-GITHUB-PROJECT"))
	fmt.Println(os.Getenv("OWNER"))

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic("Cannot load configuration file " + err.Error())
	}
	fmt.Println("Membuka konfigurasi menggunakan file json (tanpa menentukan tipe data dari konfigurasi)")
	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("app_github_project"))
	fmt.Println(viper.Get("owner"))
	fmt.Println(viper.Get("port"))

	fmt.Println("Membuka konfigurasi menggunakan file json (dengan menentukan tipe data dari konfigurasi)")
	fmt.Println(viper.GetString("app_name"))
	fmt.Println(viper.GetString("app_github_project"))
	fmt.Println(viper.GetString("owner"))
	fmt.Println(viper.GetInt("port"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic("Cannot load configuration file " + err.Error())
	}
	fmt.Println("Membuka konfigurasi menggunakan file yaml (tanpa menentukan tipe data dari konfigurasi)")
	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("app_github_project"))
	fmt.Println(viper.Get("owner"))
	fmt.Println(viper.Get("port"))

	fmt.Println("Membuka konfigurasi menggunakan file yaml (dengan menentukan tipe data dari konfigurasi)")
	fmt.Println(viper.GetString("app_name"))
	fmt.Println(viper.GetString("app_github_project"))
	fmt.Println(viper.GetString("owner"))
	fmt.Println(viper.GetInt("port"))
}
