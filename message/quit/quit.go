// 进程退出消息，用于服务优雅退出，通知一些模块做清理工作
// Closer接口的Close方法应该是同步清理的逻辑
//
//   func (s *SomeService) worker() {
//      for {
//			select {
//			//...other case
//
//			// quit
//			case <-s.quit
//			// ....
//			// 清理逻辑，例如buffer清空写入
//			// ACK，通知主线程清理已完成，可以退出
//			s.quit <- true
//			return
//			}
//		}
//   }
//
//   func (s *SomeService) Close() {
//		// 发送退出消息给worker线程
//		s.quit <- true
//		// 等待清理工作完成
//		<-s.quit
//	}
//
package quit

import (
	"sync"
)

type Closer interface {
	Close()
}

type CloserFunc func()

func (f CloserFunc) Close() {
	f()
}

var (
	closers map[Closer]bool
	mu      sync.Mutex
)

func init() {
	closers = make(map[Closer]bool)
}

// 注册清理对象
func Register(closer Closer) {
	mu.Lock()
	defer mu.Unlock()
	closers[closer] = true
}

// 取消注册
func Unregister(closer Closer) {
	mu.Lock()
	defer mu.Unlock()
	delete(closers, closer)
}

// 主线程退出时，调用
func Broadcast() {
	mu.Lock()
	defer mu.Unlock()
	for closer := range closers {
		closer.Close()
	}
}
