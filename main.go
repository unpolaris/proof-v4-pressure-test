package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/unpolaris/proof-v4-go/swagger"
)

var (
	addr = flag.String("a", "", "the remote addr")
	num  = flag.String("n", "", "the number of simulated users")
)

var RemoteAddr = "http://183.134.99.131:46789"
var UserNumber = 200

func main() {
	flag.Parse()
	if *addr != "" {
		RemoteAddr = *addr
	}
	if *num != "" {
		n, err := strconv.Atoi(*num)
		if err != nil {
			panic("the number of simulated users is error value")
		}
		UserNumber = n
	}

	for i := 0; i < UserNumber; i++ {
		go Send(RemoteAddr)
	}

	time.Sleep(10 * time.Second)
}

func Send(remoteAddr string) {
	config := swagger.NewConfiguration()
	config.BasePath = remoteAddr

	client := swagger.NewAPIClient(config)

	data := []swagger.ModelReqAutoProof{
		{
			Id: "1",
			Ext: &swagger.ModelProofExtInfo{
				TemplateId: 25,
				Version:    "v4",
			},
			Detail: "\"{\\\"订单\\\":{\\\"订单id\\\":\\\"1\\\",\\\"下订单时间\\\":\\\"2009-01-01:12:00:00\\\",\\\"订单状态\\\":\\\"已完成\\\",\\\"交易编号\\\":\\\"001\\\",\\\"商品名称\\\":\\\"订单测试\\\",\\\"商品金额\\\":\\\"10\\\",\\\"商品数量\\\":\\\"1\\\",\\\"企业名称\\\":\\\"测试企业\\\",\\\"收货人姓名\\\":\\\"张三\\\",\\\"收货人地址\\\":\\\"浙江杭州\\\",\\\"收获人手机号\\\":\\\"13800000000\\\"}}\"",
		},
	}

	signOpt := &swagger.ProofApiApiAutoProofsPostOpts{}

	for {
		got, _, err := client.ProofApi.ApiAutoProofsPost(context.Background(), data, signOpt)

		if err != nil {
			fmt.Println("err=", err)
		} else {
			fmt.Println("success", "response", got, "data", got.Data)
		}
	}
}
