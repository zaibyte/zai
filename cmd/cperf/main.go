package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"g.tesamc.com/IT/zaipkg/typeutil"
	"g.tesamc.com/IT/zaipkg/xredis"

	"g.tesamc.com/IT/zaipkg/xstrconv"
	memc "github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
)

var (
	target  = flag.String("target", "redis", "testing target")
	addr    = flag.String("addr", "", "target address")
	objSize = flag.String("size", "128mb", "object value size")
	threads = flag.Int("threads", 16, "concurrency")
)

func main() {

	flag.Parse()

	t := *target
	if t != "redis" && t != "memcached" {
	}

	var c client
	if t == "reids" {
		c = newRClient(*addr)
	} else if t == "memecached" {
		c = newMClient(*addr)
	} else {
		log.Fatal("cannot support target: ", t)
	}

	size := new(typeutil.ByteSize)
	err := size.UnmarshalText([]byte(*objSize))
	if err != nil {
		log.Fatal(err)
	}
	obj := make([]byte, int(*size))

	err = c.Put([]byte("1"), obj)
	if err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)
	wg.Add(*threads)

	for i := 0; i < *threads; i++ {
		go func(i int) {
			defer wg.Done()

			start := time.Now()
			_, err := c.Get([]byte("1"))
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d thread cost: %s", time.Now().Sub(start).String())
		}(i)
	}

	wg.Wait()

}

type client interface {
	Put(k, v []byte) error
	Get(k []byte) (v []byte, err error)
	Close()
}

type rClient struct {
	rc *redis.Client
}

func newRClient(addr string) *rClient {
	rc, err := xredis.NewClient(addr)
	if err != nil {
		log.Fatal(err)
	}
	return &rClient{rc}
}

func (r *rClient) Put(k, v []byte) error {
	panic("implement me")
}

func (r *rClient) Get(k []byte) (v []byte, err error) {
	panic("implement me")
}

func (r *rClient) Close() {
	panic("implement me")
}

type mClient struct {
	mc *memc.Client
}

func (m *mClient) Put(k, v []byte) error {
	return m.mc.Set(&memc.Item{Key: xstrconv.ToString(k), Value: v})
}

func (m *mClient) Get(k []byte) (v []byte, err error) {
	it, err := m.mc.Get(xstrconv.ToString(k))
	if err != nil {
		return nil, err
	}
	return it.Value, nil
}

func (m *mClient) Close() {
	return
}

func newMClient(addr string) *mClient {
	mc := memc.New(addr)
	return &mClient{mc: mc}
}
