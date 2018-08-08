package go_rpc_thrift

import (
	"time"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"os"
	"batu/demo"
	"strconv"
	"context"
)

const (
	HOST = "127.0.0.1"
	PORT = "9090"
)



func ThriftClient()  {
	startTime := currentTimeMillis()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport, _ := transportFactory.GetTransport(transport)
	client := demo.NewBatuThriftClientFactory(useTransport, protocolFactory)

	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}

	defer transport.Close()

	ctx := context.Background()

	for i := 0; i < 10; i++ {
		paramMap := make(map[string]string)
		paramMap["a"] = "batu.demo"
		paramMap["b"] = "test" + strconv.Itoa(i+1)
		r1, _ := client.CallBack(ctx, time.Now().Unix(), "go client", paramMap)
		fmt.Println("GOCLient call ->", r1)
	}

	model := demo.Article{1, "WTF", "content", "chenjian"}
	client.Put(ctx, &model)
	endTime := currentTimeMillis()
	fmt.Printf("本次调用用时:%d-%d=%d毫秒\n", endTime, startTime, (endTime - startTime))
}
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}