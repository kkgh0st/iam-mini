package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"iam-mini/generic/opt"
	"net/http"
)

/*
设计原则是Config -> CompletedConfig -> GenericAPIServer ，
*/
type Config struct {
	Insecure *opt.InSecureServingInfo
}

func (c *Config) Complete() *CompletedConfig {
	return &CompletedConfig{c}
}

func NewConfig() *Config {
	return &Config{
		Insecure: &opt.InSecureServingInfo{},
	}
}

type CompletedConfig struct {
	*Config
}

func (c *CompletedConfig) New() (*GenericAPIServer, error) {

	s := &GenericAPIServer{
		InsecureServingInfo: c.Insecure,
		Engine:              gin.New(),
	}

	initGenericApiServer(s)

	return s, nil
}

func initGenericApiServer(s *GenericAPIServer) {
	// 来初始化各个通用的中间组件，无需多言
}

type GenericAPIServer struct {
	middleware []string

	// InsecureServingInfo ...
	InsecureServingInfo *opt.InSecureServingInfo
	// SecureServingInfo ..
	*gin.Engine

	insecureServer, secureServer *http.Server
}

func (s *GenericAPIServer) InstallAPIs() {
	// 这来负责安装全部API
}

func (s *GenericAPIServer) Run() error {
	// server 来绑定 gin.Engine，这个
	s.insecureServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.InsecureServingInfo.BindAddress, s.InsecureServingInfo.BindPort),
		Handler: s, // 其调用 gin.Engine对应的Handler函数
	}
	var eg errgroup.Group
	eg.Go(func() error {
		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	err := eg.Wait()
	if err != nil {
		return err
	}
	return nil
}
