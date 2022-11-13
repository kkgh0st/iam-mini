package server

type apiServer struct {
	genericAPIServer *GenericAPIServer
}

func CreateAPIServer(cfg *Config) (*apiServer, error) {

	genericServer, err := cfg.Complete().New()
	if err != nil {
		return nil, err
	}

	// 返回创建好的服务器
	return &apiServer{
		genericAPIServer: genericServer,
	}, nil
}

type preparedAPIServer struct {
	*apiServer
}

func (s *preparedAPIServer) Run() error {
	return s.genericAPIServer.Run()
}

func (s *apiServer) PrepareRun() *preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)
	return &preparedAPIServer{s}
}
