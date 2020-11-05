package main

var (
	configPath string
)

func init() {
	flag.StringVar(&congifPath, "config-path", "configs/apiserver.toml", "path to config file")
}

/* make
./apiserver -help
*/
func main() {
	//Parsing variable
	flag.Parse()

	//Load Configue
	config := apiserver.NewConfig()

	//Read toml config
	_, err := toml.DecodeFile(configPath, config)
	//Check Error
	if err != nil {
		log.Fatal(err)
	}

	//Create APIServer
	server := apiserver.New()
	//Start APIServer
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
