package go_rpc_thrift

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"os"
	"batu/demo"
	"context"
	"time"
)

const (
	NetworkAddr = "127.0.0.1:9090"
)

type batuThrift struct {

}

func (this *batuThrift) CallBack(ctx context.Context, callTime int64, name string, paramMap map[string]string) (r []string, err error) {
	//panic("implement me")
	fmt.Println("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
	r = append(r, "key:"+paramMap["a"]+"    value:"+paramMap["b"])
	return
}

func (this *batuThrift) Put(ctx context.Context, s *demo.Article) (err error) {
	fmt.Printf("Article--->id: %d\tTitle:%s\tContent:%t\tAuthor:%d\n", s.ID, s.Title, s.Content, s.Author)
	return nil
}

func ThriftServer()  {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	handler := &batuThrift{}
	processor := demo.NewBatuThriftProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}