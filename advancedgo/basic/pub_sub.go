package main

import (
	"sync"
	"time"
	"strings"
	"fmt"
)

// 发布订阅模型
// publish and subscribe

type (
	subscriber chan interface{}        // 订阅者是一个管道
	topicFunc func(v interface{}) bool // 主体是一个过滤器
)

// Publisher 发布者对象
type Publisher struct {
	m           sync.RWMutex             // 读写锁
	buffer      int                      // 订阅队列的缓存
	timeout     time.Duration            // 发布超时时间
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// NewPublisher 构建一个发布者对象
func NewPublisher(timeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个新的订阅者 订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// 添加一个新的订阅者 订阅过滤器筛选过的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()
	wg := sync.WaitGroup{}
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 发送消息
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		// 过滤掉不在订阅内的消息
		return
	}
	select {
	// 要么发送成功 要么超时
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

// 关闭发布者对象，同时关闭所有的订阅者管道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()
	all := p.Subscribe()
	golang := p.SubscribeTopic(
		func(v interface{}) bool {
			if s, ok := v.(string); ok {
				return strings.Contains(s, "golang")
			}
			return false
		})

	p.Publish("hello,world")
	p.Publish("hello,golang!")
	go func() {
		for msg:= range all {
			fmt.Println("all:",msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:",msg)
		}
	}()
	time.Sleep(3*time.Second)
}
