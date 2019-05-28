package server

import (
	"log"
	"net/http"
	"os"

	root "skynet/pkg"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server is used to access the mux router
type Server struct {
	router *mux.Router
	config *root.ServerConfig
}

// NewServer creates an instance of a new server
func NewServer(config *root.Config) *Server {
	s := Server{
		router: mux.NewRouter(),
		config: config.Server}

	return &s
}

// Start initiates the server and listens for calls
func (s *Server) Start() {
	log.Println("Listening on port " + s.config.Port)

	if err := http.ListenAndServe(s.config.Port, handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) getSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}

// CreateUserRouter creates UserRouter for handling user related functions
func (s *Server) CreateUserRouter(u root.UserService) {
	NewUserRouter(u, s.getSubrouter("/user"))
}

// CreateRecordRouter creates RecordRouter for handling record related functions
func (s *Server) CreateRecordRouter(rec root.RecordService) {
	NewRecordRouter(rec, s.getSubrouter("/record"), s.config.Port)
}

// CreateClaimRouter creates RecordRouter for handling record related functions
func (s *Server) CreateClaimRouter(rec root.ClaimService) {
	NewClaimRouter(rec, s.getSubrouter("/claim"), s.config.Port)
}

// CreateRoutes registers the independent handler functions
func (s *Server) CreateRoutes() {

	s.router.HandleFunc("/", frontPage).Methods("GET")
	s.router.HandleFunc("/signup", displaySignUpHandler).Methods("GET")
	s.router.HandleFunc("/login", displayLoginHandler).Methods("GET")
	s.router.HandleFunc("/display", randomDisplay).Methods("GET")

}
