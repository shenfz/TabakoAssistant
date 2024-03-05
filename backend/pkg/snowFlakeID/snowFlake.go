package snowFlakeID

/**
 * @Author shenfz
 * @Date 2024/3/4 12:21
 * @Email 1328919715@qq.com
 * @Description:
 **/

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
 1位最高位：符号位不使用（0表正数，1表负数）

41位时间戳：2^41-1个数字代表69年，所以设置发号起始时间最好为发号器首次运行时间

10位工作机器id：也会分为5位datacenterId和5位workerId

12位序列号：2^12-1个数字总共4095，同一毫秒同一机器节点可以并发产生4095个不同Id

*/

// global var
var sequence int = 0
var lastTime int = -1

// every segment bit
var workerIdBits = 5
var datacenterIdBits = 5
var sequenceBits = 12

// every segment max number
var maxWorkerId int = -1 ^ (-1 << workerIdBits)
var maxDatacenterId int = -1 ^ (-1 << datacenterIdBits)
var maxSequence int = -1 ^ (-1 << sequenceBits)

// bit operation shift
var workerIdShift = sequenceBits
var datacenterShift = workerIdBits + sequenceBits
var timestampShift = datacenterIdBits + workerIdBits + sequenceBits

type Snowflake struct {
	datacenterId int
	workerId     int
	epoch        int
	mt           *sync.Mutex
}

var (
	std *Snowflake
)

func init() {
	datacenterId := 0
	workerId := 0
	epoch := 1596850974657
	std, _ = NewSnowflake(datacenterId, workerId, epoch)
}

func GetSnowFlakeID() string {
	return strconv.Itoa(std.getUniqueId())
}

func NewSnowflake(datacenterId int, workerId int, epoch int) (*Snowflake, error) {
	if datacenterId > maxDatacenterId || datacenterId < 0 {
		return nil, errors.New(fmt.Sprintf("datacenterId cant be greater than %d or less than 0", maxDatacenterId))
	}
	if workerId > maxWorkerId || workerId < 0 {
		return nil, errors.New(fmt.Sprintf("workerId cant be greater than %d or less than 0", maxWorkerId))
	}
	if epoch > getCurrentTime() {
		return nil, errors.New(fmt.Sprintf("epoch time cant be after now"))
	}
	sf := Snowflake{datacenterId, workerId, epoch, new(sync.Mutex)}
	return &sf, nil
}

func (sf *Snowflake) getUniqueId() int {
	sf.mt.Lock()
	defer sf.mt.Unlock()
	//get current time
	currentTime := getCurrentTime()
	//compute sequence
	if currentTime < lastTime { //occur clock back
		//panic or wait,wait is not the best way.can be optimized.
		currentTime = waitUntilNextTime(lastTime)
		sequence = 0
	} else if currentTime == lastTime { //at the same time(micro-second)
		sequence = (sequence + 1) & maxSequence
		if sequence == 0 { //overflow max num,wait next time
			currentTime = waitUntilNextTime(lastTime)
		}
	} else if currentTime > lastTime { //next time
		sequence = 0
		lastTime = currentTime
	}
	//generate id
	return (currentTime-sf.epoch)<<timestampShift | sf.datacenterId<<datacenterShift |
		sf.workerId<<workerIdShift | sequence
}

func waitUntilNextTime(lasttime int) int {
	currentTime := getCurrentTime()
	for currentTime <= lasttime {
		time.Sleep(1 * time.Second / 1000) //sleep micro second
		currentTime = getCurrentTime()
	}
	return currentTime
}

func getCurrentTime() int {
	return int(time.Now().UnixNano() / 1e6) //micro second
}

/*//  生产100 0000 万个
// sort -nr y|uniq -c|head -2
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	datacenterId := 0
	workerId := 0
	epoch := 1596850974657
	s, err := NewSnowflake(datacenterId, workerId, epoch)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(s.getUniqueId())
			wg.Done()
		}()
	}
	wg.Wait()
}*/
