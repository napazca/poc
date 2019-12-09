package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

func main(){
	// See http://redis.io/topics/cluster-tutorial for instructions
	// how to setup Redis Cluster.
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(`redis-1:6379,redis-2:6379`, ","),
	})
	cmd := rdb.Ping()
	fmt.Println(cmd.Result())
	fmt.Println(cmd.Err())

	rdb.Set("a", "aaa", 1*time.Second)
	fmt.Println(rdb.Get("a"))
}

