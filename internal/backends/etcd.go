package backends

import (
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func Init() {
	initEtcd()
	loadEtcdConfig()
}

var etcdClient *clientv3.Client

func initEtcd() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   viper.GetStringSlice("backends.etcd.endpoints"),
		DialTimeout: time.Duration(viper.GetInt("backends.etcd.dialTimeout")) * time.Second,
	})
	if err != nil {
		panic(err)
	}
	etcdClient = client
}

func loadEtcdConfig() {
	err := viper.AddRemoteProvider(
		"etcd",
		viper.GetStringSlice("backends.etcd.endpoints")[0],
		"/gGateway")
	if err != nil {
		panic(err)
	}
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
}

func GetEtcdClient() *clientv3.Client {
	return etcdClient
}
