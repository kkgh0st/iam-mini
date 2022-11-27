package apiserver

import (
	"iam-mini/apiserver/server"
	"iam-mini/apiserver/store"
	"iam-mini/apiserver/store/mysql"
	genericoptions "iam-mini/generic/options"
)

type apiServer struct {
	genericAPIServer *server.GenericAPIServer
	gRPCAPIServer    *server.GrpcAPIServer
	// 这个只是在这里设置配置，之后会调用
	redisOptions *genericoptions.RedisOpt
}

type ExtraConfig struct {
	Addr string
	//	MaxMsgSize   int
	//	ServerCert   genericoptions.GeneratableKeyCert
	mysqlOptions *genericoptions.MySQLOptions
	// etcdOptions      *genericoptions.EtcdOptions
}

type completedExtraConfig struct {
	*ExtraConfig
}

func (c *ExtraConfig) complete() *completedExtraConfig {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}

	return &completedExtraConfig{c}
}

func buildExtraConfig(cfg *Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		//Addr:         fmt.Sprintf("%s:%d", cfg.GRPCOptions.BindAddress, cfg.GRPCOptions.BindPort),
		//MaxMsgSize:   cfg.GRPCOptions.MaxMsgSize,
		//ServerCert:   cfg.SecureServing.ServerCert,
		mysqlOptions: cfg.MySQLOpt,
		// etcdOptions:      cfg.EtcdOptions,
	}, nil
}

func (c *completedExtraConfig) New() (*server.GrpcAPIServer, error) {
	// 我们这里先来初始化其数据库相关细节，

	//creds, err := credentials.NewServerTLSFromFile(c.ServerCert.CertKey.CertFile, c.ServerCert.CertKey.KeyFile)
	//if err != nil {
	//	log.Fatalf("Failed to generate credentials %s", err.Error())
	//}
	//opts := []grpc.ServerOption{grpc.MaxRecvMsgSize(c.MaxMsgSize), grpc.Creds(creds)}
	//grpcServer := grpc.NewServer(opts...)

	storeIns, _ := mysql.GetMySQLFactoryOr(c.mysqlOptions)
	//storeIns, _ := etcd.GetEtcdFactoryOr(c.etcdOptions, nil)
	store.SetClient(storeIns)
	//cacheIns, err := cachev1.GetCacheInsOr(storeIns)
	//if err != nil {
	//	log.Fatalf("Failed to get cache instance: %s", err.Error())
	//}
	//
	//pb.RegisterCacheServer(grpcServer, cacheIns)
	//
	//reflection.Register(grpcServer)

	return nil, nil
	//return &grpcAPIServer{grpcServer, c.Addr}, nil
}

func buildGenericConfig(cfg *Config) (genericConfig *server.Config, lastErr error) {
	genericConfig = server.NewConfig()
	// 这里自己没有任何返回值
	cfg.InsecureOpt.ApplyTo(genericConfig.Insecure)
	//if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}
	//
	//if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}
	//
	//if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}
	//
	//if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}

	return
}

func CreateAPIServer(cfg *Config) (*apiServer, error) {
	/*
		这里各个服务器的创建逻辑，根据 config.Config，创建各种Config
		然后Config.Complete().New()，来创建对应的服务器
		服务器统一放入apiServer中，统一调度与运行!
	*/
	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}

	extraServer, err := extraConfig.complete().New()
	if err != nil {
		return nil, err
	}
	// 返回创建好的服务器
	return &apiServer{
		genericAPIServer: genericServer,
		gRPCAPIServer:    extraServer,
		redisOptions:     cfg.RedisOpt,
	}, nil
}

type preparedAPIServer struct {
	*apiServer
}

func (s *preparedAPIServer) Run() error {
	return s.genericAPIServer.Run()
}

func (s *apiServer) initRedisStore() {

}

func (s *apiServer) PrepareRun() *preparedAPIServer {
	server.InitRouter(s.genericAPIServer.Engine)

	s.initRedisStore()
	return &preparedAPIServer{s}
}
