package nacos

import (
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "testing"
)

func TestClient_Register(t *testing.T) {
    cc := constant.ClientConfig{
        //NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", //namespace id
        TimeoutMs:           5000,
        NotLoadCacheAtStart: true,
        //LogDir:              "/tmp/nacos/log",
        //CacheDir:            "/tmp/nacos/cache",
        RotateTime: "1h",
        MaxAge:     3,
        LogLevel:   "debug",
    }
    _ = cc
    //client := NewClient([]string{"127.0.0.1:8848"}, &Config{ClientConfig: cc}, log.NewNopLogger())
    //err := client.Register("127.0.0.1:11341", "user", nil)
    //if err != nil {
    //    logs.Error(err)
    //}
}

func TestClient_Deregister(t *testing.T) {
    cc := constant.ClientConfig{
        //NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", //namespace id
        TimeoutMs:           5000,
        NotLoadCacheAtStart: true,
        //LogDir:              "/tmp/nacos/log",
        //CacheDir:            "/tmp/nacos/cache",
        RotateTime: "1h",
        MaxAge:     3,
        LogLevel:   "debug",
    }
    _ = cc
    //client := NewClient([]string{"127.0.0.1:8848"}, &Config{ClientConfig: cc}, log.NewNopLogger())
    //err := client.Register("127.0.0.1:8848", "se.v1.Id", nil)
    //if err != nil {
    //    logs.Error(err)
    //}
    //err = client.Deregister()
    //if err != nil {
    //    logs.Error(err)
    //}
}

func TestNacosClient_Instancer(t *testing.T) {
    cc := constant.ClientConfig{
        //NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", //namespace id
        TimeoutMs:           5000,
        NotLoadCacheAtStart: true,
        //LogDir:              "/tmp/nacos/log",
        //CacheDir:            "/tmp/nacos/cache",
        RotateTime: "1h",
        MaxAge:     3,
        LogLevel:   "debug",
    }
    _ = cc
    //client := NewClient([]string{"127.0.0.1:8848"}, &Config{ClientConfig: cc}, log.NewNopLogger())
    //instances, err := client.client.SelectAllInstances(vo.SelectAllInstancesParam{ServiceName: "se.v1.Id"})
    //if err != nil {
    //    logs.Error(err)
    //} else {
    //    logs.Info(instances[0].ServiceName)
    //}
}
