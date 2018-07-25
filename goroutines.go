package main

import "fmt"
//synchronization
import "time"

func f(from string)  {
	for i := 0;i < 3;i++ {
		fmt.Println(from, ":", i)
	}
}

//channel-synchronization
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

//channel-directions
func ping(pings chan<- string, msg2 string) {
	pings <- msg2
}

func pong(pings <-chan string, pongs chan<- string) {
	msg2 := <-pings
	pongs <- msg2
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

	fmt.Println("--------")
	//channel
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()

	msgs := <-messages
	fmt.Println(msgs)

	fmt.Println("-------")
	//channel-buffering
	msg1 := make(chan string, 2)

	msg1 <- "buffered"
	msg1 <- "channel"

	fmt.Println(<-msg1)
	fmt.Println(<-msg1)

	fmt.Println("------")
	//channel-synchronization
	done := make(chan bool, 1)
	go worker(done)
	<-done

	fmt.Println("-------")
	//channel-direction
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	fmt.Println("-------")
	//channel-select
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0;i < 2;i++ {
		select {
		case msg01 := <-c1:
			fmt.Println("received", msg01)
		case msg02 := <-c2:
			fmt.Println("received", msg02)
		}
	}

	fmt.Println("-------")
	//timeouts
	c3 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c3 <- "result 1"
	}()
	select {
		case res := <- c3:
			fmt.Println(res)
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1")
	}

	c4 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c4 <- "result 2"
	}()
	select {
	case res := <-c4:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	fmt.Println("-------")
	//non-blocking-channel-operations
	messages1 := make(chan string)
	signals := make(chan bool)

	select {
		case msg001 := <-messages1:
			fmt.Println("received message", msg001)
		default:
			fmt.Println("no message received")
	}

	msg000 := "hi"
	select {
	case messages1 <- msg000:
		fmt.Println("sent message", msg000)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages1:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	fmt.Println("--------")
	//closing-channels
	jobs := make(chan int, 5)
	done1 := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done1 <- true
				return
			}
		}
	}()
	for j := 1;j <= 3;j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done1

	fmt.Println("------")
	//range-over-channels
	queue := make(chan string, 2)
	queue <- "One"
	queue <- "Two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
