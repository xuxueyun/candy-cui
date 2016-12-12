package candy

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/zeazen/candy-cui/meta"
	"github.com/zeazen/candy-cui/util"
	"github.com/zeazen/candy-cui/util/log"
)

const (
	networkTimeout = time.Second * 5
)

// MessageHandler 接收服务器端推送来的消息
type MessageHandler interface {
	// OnRecv 这函数理论上是多线程调用，客户端需要注意下
	OnRecv(event int32, operate int32, ID int64, group int64, from int64, to int64, body string)

	// OnError 连接被服务器断开，或其它错误
	OnError(msg string)

	// OnHealth 连接正常
	OnHealth()

	// OnUnHealth 连接异常
	OnUnHealth(msg string)
}

// CandyClient 客户端提供和服务器交互的接口
type CandyClient struct {
	host    string
	broken  bool
	last    time.Time
	conn    *grpc.ClientConn
	gate    meta.GateClient
	handler MessageHandler
	stream  meta.Gate_StreamClient
	id      int64
	token   int64
	user    string
	pass    string
	device  string
	closer chan struct{}
	mu sync.RWMutex
}

// CuiHandler 原cmdclient 满世界用
type CuiHandler struct{}

// OnRecv 这函数理论上是多线程调用，客户端需要注意下
func (c *CuiHandler) OnRecv(event int32, operate int32, id int64, group int64, from int64, to int64, body string) {
	log.Debugf("recv msg id:%d event:%v, operate:%v, group:%d, from:%d, to:%d, body:%s\n", id, meta.Event(event), meta.Relation(operate), group, from, to, body)
}

// OnError 连接被服务器断开，或其它错误
func (c *CuiHandler) OnError(msg string) {
	log.Errorf("rpc error:%s\n", msg)
}

// OnHealth 连接恢复
func (c *CuiHandler) OnHealth() {
	log.Debugf("connection recovery\n")
}

// OnUnHealth 连接异常
func (c *CuiHandler) OnUnHealth(msg string) {
	log.Errorf("connection UnHealth, msg:%v\n", msg)
}

// CClient 声明一个 client 满世界用
var CandyCUIClient *CandyClient

// NewCandyClient - create an new CandyClient
func NewCandyClient(dev,host string, handler MessageHandler) *CandyClient {
	return &CandyClient{device:dev,host: host, handler: handler, broken: true,closer:make(chan struct{})}
}

// Start 连接服务端.
func (c *CandyClient) Start() error {
	var err error

	c.conn, err = grpc.Dial(c.host, grpc.WithInsecure(), grpc.WithTimeout(networkTimeout), grpc.WithBackoffMaxDelay(networkTimeout))
	if err != nil {
		log.Errorf("dial:%s error:%s", c.host, err.Error())
		return err
	}

	c.gate = meta.NewGateClient(c.conn)
	c.last = time.Now()
	c.broken = false

	go c.healthCheck()

	return nil
}

// service 调用服务器接口, 带上token
func (c *CandyClient) service(call func(context.Context, meta.GateClient) error) {
	ctx := util.ContextSet(context.Background(), "token", fmt.Sprintf("%d", c.token))
	ctx = util.ContextSet(ctx, "id", fmt.Sprintf("%d", c.id))
	if err := call(ctx, c.gate); err != nil {
		log.Infof("call:%s error:%s", c.host, err.Error())
		return
	}
	c.mu.Lock()
	c.last = time.Now()
	c.mu.Unlock()
}

func (c *CandyClient) openStream() (resp meta.Gate_StreamClient, err error) {
	req := &meta.GateStreamRequest{Token: c.token, ID: c.id}
	c.service(func(ctx context.Context, api meta.GateClient) error {
		if resp, err = api.Stream(ctx, req); err != nil {
			return err
		}
		return nil
	})
	return
}

// receiver 一直接收服务器返回消息, 直到出错.
func (c *CandyClient) receiver(stream meta.Gate_StreamClient) {
	for {
		pm, err := stream.Recv()
		if err != nil {
			log.Errorf("recv error:%s", err)
			c.onError(err.Error())
			break
		}
		c.handler.OnRecv(int32(pm.Event), int32(pm.Operate), pm.Msg.ID, pm.Msg.Group, pm.Msg.From, pm.Msg.To, pm.Msg.Body)
	}
}

func (c *CandyClient) onError(msg string) {
	c.mu.Lock()
	c.last = time.Now().Add(-time.Minute)
	if c.broken {
		c.mu.Unlock()
		return
	}
	c.broken = true
	c.mu.Unlock()

	if strings.Contains(msg, "invalid context") && c.user != "" && c.pass != "" {
		c.Login(c.user, c.pass)
	}

	c.handler.OnError(msg)
}


func (c *CandyClient) startReceiver() {
	if c.token != 0 && c.id != 0 {
		stream, err := c.openStream()
		if err != nil {
			c.onError(err.Error())
		}

		go c.receiver(stream)

	}
}

//onHealth 如果网络正常了，要尝试启动Push Stream
func (c *CandyClient) onHealth() {
	c.mu.Lock()
	c.last = time.Now()
	if !c.broken {
		c.mu.Unlock()
		return
	}
	c.broken = false
	c.mu.Unlock()

	c.handler.OnHealth()

	c.startReceiver()

	log.Debugf("heartbeat ok")
}


// OnNetStateChange 移动端如果网络状态发生变化要通知这边
func (c *CandyClient) OnNetStateChange() {
	//TODO 细分
	c.mu.Lock()
	c.last = time.Now().Add(-time.Minute)
	c.mu.Unlock()
}

// healthCheck 健康检查,60秒发一次, 目前服务器超过90秒会发探活
func (c *CandyClient) healthCheck() {
	t := time.NewTicker(networkTimeout)
	defer t.Stop()

	req := &meta.HeartbeatRequest{}
	var resp *meta.HeartbeatResponse
	var err error

	for {
		select {
		case <-c.closer:
			t.Stop()
			return
		case <-t.C:
		}
		c.mu.RLock()
		if c.token == 0 || time.Now().Sub(c.last) < time.Minute {
			c.mu.RUnlock()
			continue
		}
		c.mu.RUnlock()

		c.service(func(ctx context.Context, client meta.GateClient) error {
			resp, err = client.Heartbeat(ctx, req)
			return err
		})

		if err != nil {
			c.onError(err.Error())
			continue
		}

		if resp.Header.Error() != nil {
			log.Errorf("Heartbeat response error:%v", resp.Header.Error())
			if c.user != "" && c.pass != "" {
				c.Login(c.user, c.pass)
			}
			continue
		}

		c.onHealth()
	}
}
