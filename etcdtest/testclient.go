package etcdtest

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func GetEtcdClient() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"114.67.67.7:2379", "114.67.83.163:2379", "114.67.112.67:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err == context.DeadlineExceeded {
		// handle errors
		Logger.Info("etcd connect time out!")
	}
	return cli
}

// 获取集群member信息
func GetClusterInfo() {
	cli := GetEtcdClient()
	cli.Get()
	for {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		select {
		case <-time.After(2 * time.Second):
			resp, err := cli.MemberList(ctx)
			if err != nil {
				Logger.Fatal(err.Error())
				continue
			}
			fmt.Println("time: ", time.Now(), "members:", len(resp.Members), "header: ", resp.Header)
			continue
		case <-ctx.Done():
			//fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			continue
		}

		//resp, err := cli.MemberList(ctx)
		//if err != nil {
		//	Logger.Fatal(err.Error())
		//	continue
		//}

		//fmt.Println("time: ", time.Now(), "members:", len(resp.Members), "header: ", resp.Header)

		//time.Sleep(2 * time.Second)
	}

	//for _, ep := range cli.Endpoints() {
	//	statusresponse, err := cli.Status(context.Background(), ep)
	//	if err != nil {
	//		Logger.Error(err.Error())
	//		continue
	//	}
	//	if statusresponse.Leader == statusresponse.Header.MemberId {
	//		fmt.Println("leader member is:", statusresponse.Leader)
	//	}
	//
	//}

}
