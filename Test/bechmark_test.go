package Test

import (
	"testing"
	bus "Golang-CQRS/Bus"
	"strconv"
)


func Benchmark_EventBus_Publish(b *testing.B)  {
	b.Logf("b.N is %d\n", b.N)

	eventBus := bus.New()

	for i := 0; i < b.N; i++ {

		handler1 := &FakeHandler1{
			_event: "Event" + strconv.Itoa(int( i / 10)) ,
			_isDisableMessage: true,
			}

		handler2 := &FakeHandler2{}

		eventBus.Subscribe(handler1)
		eventBus.Subscribe(handler2)
	}

	for i := 0; i < b.N; i += 10 {
		//fmt.Printf("Run Event: i:%d,  Event%s", i, "Event" + strconv.Itoa(int( i / 10)))
		//fmt.Println()

		eventBus.Publish("Event" + strconv.Itoa(int( i / 10)))
	}
}
