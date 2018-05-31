package main
import "log"

func calculate(round int64, nElements int64, c chan<-float64) {
	start := round * nElements
	end := (round + 1) *nElements -1
	result := float64(0.0)
    for i := start; i <= end; i++ {
		flag := float64(1 - i%2*2)
		//log.Println(flag)
		result += flag / float64(2 * i +1)
	}
	c <- result
}

func main() {
	log.Println("starts!")
	round := int64(100000)
	nElements := int64(100000)
	queue := make(chan float64)
	for i:= int64(0); i< round; i++ {
		go calculate(i, nElements, queue)
	}
	result := float64(0)
	for m:=int64(0); m<round; m++ {
		message := <-queue
		//log.Println(message)
		result += float64(4) * message
	}
    log.Println(result)
}