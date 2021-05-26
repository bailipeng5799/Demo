package common

import "log"

//使用令牌桶的方法限流
// 流控请求数量较大server连接数不够，导致系统crush
type ConnLimiter struct{
	ConcurrentConn int // 并发数
	Bucket chan int //利用channel中的阻塞机制来做到
}

func NewConnLimiter(cc int) *ConnLimiter{
	return &ConnLimiter{
		ConcurrentConn: cc,
		Bucket:         make(chan int,cc),
	}
}
func (c1 *ConnLimiter) GetConn() bool{
	if len(c1.Bucket) >= c1.ConcurrentConn {
		log.Printf("Reach the maximum rate limit.")
		return false
	}
	c1.Bucket <- 1
	return true
}
func (c1 *ConnLimiter) ReleaseConn() {
	c := <- c1.Bucket
	log.Printf("New conntection coming :%d. Number of current connections :%d",c,len(c1.Bucket))
}