package apiserver

import (
	"github.com/ATWp/Go_pet_projects/REST_API/internal/app/store"
)

// APIServer
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

// New APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
}

// Start APIServer
func (s *APIServer) Start() error {
	if err:= s.configureLogger(); err != nil{
		return err
	}

	s.configureRouter()

	if err:= s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("strarting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParsLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err:= st.Open(); err != nil{
		return err 
	}

	s.store = st

	return nil
}

func (s *APIServer) HandleHello() http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "hello")
	}
}