package server

type apiServer struct {
	genericAPIServer *GenericAPIServer
}

func createAPIServer(cfg *Config) (*apiServer, error) {

	complete := cfg.Complete()
	genericServer, err := cfg.Complete().New()

}
