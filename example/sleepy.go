package main

import (
	"time"
	"log"
	//"errors"
	//"sync/atomic"
	"github.com/liujianping/consumer"
)

type context struct{
	main chan bool
	count int32
}

func (c *context) Do(req interface{}) error {
	r, _ := req.(*MyProduct)
	return r.Do(c)
}

func (c *context) Encode(request interface{}) ([]byte, error) {
	return nil,nil
}
func (c *context) Decode(data []byte) (interface{}, error) {

	return nil,nil
}


func (p *MyProduct) Do(c *context) error {
	log.Printf("product No(%d) do", p.No)

	time.Sleep(time.Millisecond * 250)
	
	if p.No == 30 {
		c.main <- true
	}
	log.Printf("product No(%d) Done", p.No)
	return nil
}


type MyProduct struct{
	No int
}

func main() {
	core := &context{ main: make(chan bool, 0), count:0}

	consumer := consumer.NewMemoryConsumer("sleepy", 10)
	consumer.Resume(core, 4)
	
	for i:= 1; i <= 30; i++ {
		consumer.Put(&MyProduct{i})
	}
	log.Printf("consumer running %v", consumer.Running())
	consumer.Close()
	log.Printf("consumer running %v", consumer.Running())
}