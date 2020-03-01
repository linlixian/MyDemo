// 读取数据和写入数据各是1个队列，为了并发安全，使得处理数据串行化
// （比如多个线程同时处理读取的数据，如果不加队列就会竞争读取的数据，造成数据被重复处理）
// 当html点击close后，wsConn.wsSocket.ReadMessage()读取数据出错，然后可以执行关闭连接以及通道
// 心跳存在的意义是：关闭当前网页后，后端通过心跳定时向前端传信息以确定前端是否还存活，如果不存活则信息发送失败，关闭连接和通道（即为了关闭连接），socket的异常断开也是这样处理，轮询确定是否活跃
package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端读写消息
type wsMessage struct {
	messageType int
	data        []byte
}

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列

	mutex     sync.Mutex // 避免重复关闭管道
	isClosed  bool       // 避免重复关闭管道
	closeChan chan byte  // 关闭通知
}

func (wsConn *wsConnection) wsReadLoop() {
	for {
		// 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		fmt.Println("111111111", msgType)
		fmt.Println("222222222", data)

		if err != nil {
			goto error
		}
		req := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan: // 防止阻塞
			goto closed
		}
	}
error:
	fmt.Println("111111111")
	wsConn.wsClose()
closed:
}

func (wsConn *wsConnection) wsWriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				goto error
			}
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	fmt.Println("2222222")
	wsConn.wsClose()
closed:
}

func (wsConn *wsConnection) procLoop() {
	// 启动一个gouroutine发送心跳
	go func() {
		fmt.Println("zzzzzzzzzz")

		for {
			time.Sleep(2 * time.Second)
			if err := wsConn.wsWrite(websocket.TextMessage, []byte("heartbeat from server")); err != nil {
				fmt.Println("heartbeat fail")
				wsConn.wsClose()
				break
			}
		}
	}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			fmt.Println("read fail")
			wsConn.wsClose()
			break
		}
		fmt.Println(string(msg.data))
		err = wsConn.wsWrite(msg.messageType, msg.data)
		if err != nil {
			fmt.Println("write fail")
			wsConn.wsClose()
			break
		}
	}
}

func wsHandler(resp http.ResponseWriter, req *http.Request) {

	// 应答客户端告知升级连接为websocket
	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}

	// 处理器
	go wsConn.procLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}

func (wsConn *wsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
	}
	return nil, errors.New("websocket closed")
}

func (wsConn *wsConnection) wsClose() {
	wsConn.wsSocket.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}

func webws(w http.ResponseWriter, r *http.Request) {

	//打印请求的方法

	fmt.Println("method", r.Method)

	if r.Method == "GET" { //如果请求方法为get显示login.html,并相应给前端

		t, _ := template.ParseFiles("test/socket/websocket/websocket1.html")

		t.Execute(w, nil)

	} else {

		//否则走打印输出post接受的参数username和password

		fmt.Println(r.PostFormValue("username"))

		fmt.Println(r.PostFormValue("password"))

	}

}

func main() {
	http.HandleFunc("/ws", wsHandler)
	//html页面
	http.HandleFunc("/webws", webws)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
