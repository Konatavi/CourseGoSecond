package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/Konatavi/CourseGoSecond/WebServerAndDB/internal/app/api"
	"github.com/joho/godotenv"
)

var (
	configPath     string
	typeConfigFile string
)

/*
Добавить в код необходимые блоки, для того, чтобы можно было запускать приложение следующими командами:
* Должна быть возможность запускать проект с конфигами в ```.toml```
```
api -format .env -path configs/.env
```
* Должна быть возможность запускать проект с конфигами в ```.env```
```
api -format .toml -path configs/api.toml
```
* Должна быть возможность запускать проект с дефолтными параметрами (дефолтным будем считать ```api.toml```,
если его нет, то запускаем с значениями из структуры ```Config```)
*/

func init() {
	//Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&typeConfigFile, "format", ".toml", "format config file in .toml format or .env")
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")

}

func main() {
	//В этот момент происходит инициализация переменной configPath значением
	flag.Parse()
	log.Println("typeConfigFile (after Parse2): ", typeConfigFile, "configPath (after Parse2): ", configPath)
	//server instance initialization
	config := api.NewConfig()
	switch typeConfigFile {
	case ".toml":
		_, err := toml.DecodeFile(configPath, config) // Десериалзиуете содержимое .toml файла
		if err != nil {
			log.Println("can not find configs file. using default values:", err)
		}
	case ".env":
		err := godotenv.Load(configPath)
		if err != nil {
			log.Fatal("could not find .env file:", err)
		}
		config.BindAddr = os.Getenv("bind_addr")
		config.LoggerLevel = os.Getenv("logger_level")
	default:
		{
		}

	}

	log.Println("config", config)
	server := api.New(config)

	//api server start
	log.Fatal(server.Start())
}
