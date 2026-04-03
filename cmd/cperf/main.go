package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/zaibyte/zaipkg/typeutil"
	"github.com/zaibyte/zaipkg/xredis"

	memc "github.com/bradfitz/gomemcache/memcache"
	"github.com/elastic/go-hdrhistogram"
	"github.com/go-redis/redis/v8"
	"github.com/zaibyte/zaipkg/xstrconv"
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
	if t == "redis" {
		c = newRClient(*addr)
	} else if t == "memcached" {
		c = newMClient(*addr)
	} else {
		log.Fatal("cannot support target: ", t)
	}

	size := new(typeutil.ByteSize)
	err1 := size.UnmarshalText([]byte(*objSize))
	if err1 != nil {
		log.Fatal(err1)
	}
	obj := make([]byte, int(*size))

	cnt := *threads

	for i := 1; i <= cnt; i++ {
		err1 = c.Put([]byte(strconv.Itoa(i)), obj)
		if err1 != nil {
			log.Fatal("failed to put:", err1)
		}
	}

	lats := hdrhistogram.New(0, int64(time.Second), 3)

	eachThread := 4 * 1024 * 1024 / int(*size)
	if eachThread < 4 {
		eachThread = 4
	}

	wg := new(sync.WaitGroup)
	wg.Add(*threads)

	for i := 0; i < *threads; i++ {
		go func(i int) {
			defer wg.Done()

			for j := 0; j < eachThread; j++ {
				start := time.Now()
				_, err2 := c.Get([]byte(strconv.Itoa(i + 1)))
				if err2 != nil {
					log.Fatal("failed to get:", err2, i)
				}
				cost := time.Now().Sub(start)
				_ = lats.RecordValueAtomic(int64(cost))
			}

		}(i)
	}

	wg.Wait()

	printDelta(lats)
}

func printDelta(lats *hdrhistogram.Histogram) {

	name := *target
	fmt.Println(fmt.Sprintf("%s min: %d, avg: %.2f, max: %d",
		name, lats.Min(), lats.Mean(), lats.Max()))
	fmt.Println("percentiles (nsec):")
	fmt.Print(fmt.Sprintf(
		"|  1.00th=[%d],  5.00th=[%d], 10.00th=[%d], 20.00th=[%d],\n"+
			"| 30.00th=[%d], 40.00th=[%d], 50.00th=[%d], 60.00th=[%d],\n"+
			"| 70.00th=[%d], 80.00th=[%d], 90.00th=[%d], 95.00th=[%d],\n"+
			"| 99.00th=[%d], 99.50th=[%d], 99.90th=[%d], 99.95th=[%d],\n"+
			"| 99.99th=[%d]\n",
		lats.ValueAtQuantile(1), lats.ValueAtQuantile(5), lats.ValueAtQuantile(10), lats.ValueAtQuantile(20),
		lats.ValueAtQuantile(30), lats.ValueAtQuantile(40), lats.ValueAtQuantile(50), lats.ValueAtQuantile(60),
		lats.ValueAtQuantile(70), lats.ValueAtQuantile(80), lats.ValueAtQuantile(90), lats.ValueAtQuantile(95),
		lats.ValueAtQuantile(99), lats.ValueAtQuantile(99.5), lats.ValueAtQuantile(99.9), lats.ValueAtQuantile(99.95),
		lats.ValueAtQuantile(99.99)))
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

	cmd := r.rc.Set(context.Background(), xstrconv.ToString(k), v, 0)
	return cmd.Err()
}

func (r *rClient) Get(k []byte) (v []byte, err error) {

	cmd := r.rc.Get(context.Background(), xstrconv.ToString(k))

	return cmd.Bytes()
}

func (r *rClient) Close() {
	_ = r.rc.Close()
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
	mc.MaxIdleConns = 200
	mc.Timeout = 3 * time.Second
	return &mClient{mc: mc}
}
