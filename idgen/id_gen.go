package idgen

import (
	"context"
	"errors"
	"github.com/dapr-platform/common"
	daprc "github.com/dapr/go-sdk/client"
	"log"
	"net"
	"strconv"
	"time"
)

var node *Node

func GetNode() *Node {
	if node == nil {
		/**
		query := "alarm-service:*"
		ctx := context.Background()
		ip := GetLocalIP()
		if ip == "" {
			log.Fatal("get local ip failed")
		}

		resp, err := common.GetDaprClient().QueryStateAlpha1(ctx, common.DAPR_STATESTORE_NAME, query, nil)
		if err != nil {
			panic(err)
		}

		setSelfState(ctx)
		go func() {
			time.Sleep(time.Second * 60)
			setSelfState(context.Background())
		}()
		for _, v := range resp.Results {
			log.Println("key:", v.Key)
			log.Println("value:", v.Value)

		}

		*/
		ip := GetLocalIP()
		if ip == "" {
			log.Fatal("get local ip failed")
		}
		id, err := IPString2Long(ip)
		if err != nil {
			log.Fatal("ip to long failed")
		}

		for {
			id = id % 1024
			key := "alarm-service:" + strconv.FormatInt(int64(id), 10)
			ctx := context.Background()
			ret, err := common.GetDaprClient().GetState(ctx, common.DAPR_STATESTORE_NAME, key, nil)
			if err != nil {
				log.Println("get state failed, key:", key, "err:", err)
				id = id + 1
				continue
			}
			log.Println("get state success, key:", key, "value:", string(ret.Value))
			if string(ret.Value) != "" {
				log.Println("id is exist, id:", string(ret.Value))

				id = id + 1
				continue
			}
			node, err = NewNode(int64(id))
			log.Println("node:", id)
			setSelfState(ctx, int64(id))
			go func() {
				time.Sleep(time.Second * 60)
				setSelfState(ctx, int64(id))
			}()
			break
		}

	}
	return node
}

func setSelfState(ctx context.Context, id int64) {

	key := "alarm-service:" + strconv.FormatInt(id, 10)
	value := []byte(strconv.FormatInt(id, 10))
	ttlstr := "300"
	item := &daprc.SetStateItem{
		Key: key,
		Etag: &daprc.ETag{
			Value: "1",
		},
		Metadata: map[string]string{
			"ttlInSeconds": ttlstr,
		},
		Value: value,
		Options: &daprc.StateOptions{
			Concurrency: daprc.StateConcurrencyLastWrite,
			Consistency: daprc.StateConsistencyStrong,
		},
	}

	common.GetDaprClient().SaveBulkState(ctx, common.DAPR_STATESTORE_NAME, item)
}

//get local ip
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func IPString2Long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}
func GenId() int64 {
	return int64(GetNode().Generate())
}
