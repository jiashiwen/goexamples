package etcdtest

import (
	"examples/globlezap"
	"go.etcd.io/etcd/embed"
	"net/url"

	//"github.com/coreos/etcd/embed"
	"io/ioutil"
	"log"
	"os"
	"time"

	//"go.etcd.io/etcd/clientv3"
)

var Logger = globlezap.GetLogger()

func StartEtcd() {

	lurl, _ := url.Parse("http://localhost:2379")
	urls := []url.URL{}
	urls = append(urls, *lurl)
	tdir, err := ioutil.TempDir(os.TempDir(), "token-test")
	cfg := embed.NewConfig()
	cfg.Dir = tdir
	//cfg.LCUrls = urls
	//cfg.APUrls = urls

	e, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()
	Logger.Info("server start!!")
	select {
	case <-e.Server.ReadyNotify():
		Logger.Info("Server is ready!")
		//log.Printf("Server is ready!")
	case <-time.After(60 * time.Second):
		e.Server.Stop() // trigger a shutdown
		os.Exit(1)
		log.Printf("Server took too long to start!")
	}

	for _, v := range e.Peers {
		Logger.Info(v.Addr().String())
	}
	log.Fatal(<-e.Err())
	//Logger.Fatal(e.Err())

	//client, err = clientv3.New(clientv3.Config{
	//	Endpoints:        parseClusterClients(db.Peers),
	//	AutoSyncInterval: time.Second * 5,
	//	DialTimeout:      5 * time.Second,
	//})
}
