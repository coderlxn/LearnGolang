package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

/*
测试Redis的事务功能，启动多个routine读写同一个变量，测试是否出现错误的值
如果只是用Incr这一类的原子操作，则不会产生值问题
 */

func redisReadWrite(index int)  {
	client := redis.NewClient(&redis.Options{Addr:"localhost:6379", Password:"", DB:0})
	for {
		pipe := client.TxPipeline()
		before := pipe.Incr("pipeline_val")

		mid := pipe.Incr("pipeline_val")

		after := pipe.Incr("pipeline_val")

		_, err := pipe.Exec()
		if err != nil {
			panic(err.Error())
		}

		if before.Val() + 1 != mid.Val() || mid.Val() + 1 != after.Val() {
			panic(fmt.Sprintf("value not match %d %d %d", before.Val(), mid.Val(), after.Val()))
		} else {
			println("value match")
		}
	}


	//incr := pipe.Incr("tx_pipeline_counter")
	//pipe.Expire("tx_pipeline_counter", time.Hour)
	//_, err := pipe.Exec()
	//println("index : ", index, incr.Val(), err)
}

func main()  {
	forever := make(chan bool)

	for i := 0; i < 1000; i++ {
		println("Start func ", i)
		go redisReadWrite(i)
	}

	<- forever
}