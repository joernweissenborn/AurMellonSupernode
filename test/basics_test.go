package test

import (
	"github.com/joernweissenborn/aursir4go/aurarath/implementation/zaurarath"
	"github.com/joernweissenborn/aursir4go/aurarath"
	"github.com/joernweissenborn/aursir4go/aurmellon"
	"time"
	"testing"
	"AurMellonSupernode/boot"
	"sync"
	"github.com/joernweissenborn/stream2go"
	"github.com/joernweissenborn/future2go"
)

func TetSupernodeGreeting(t *testing.T) {
	boot.Boot()
	clientnode := aurarath.NewNode()
	clientnode.RegisterImplementation(new(zaurarath.Implementation))
	client := aurmellon.NewInterface(clientnode)

	clientnode.Run()

	time.Sleep(2*time.Second)

	if !client.IsSupernodeConnected() {
		t.Error("supernode not registered")

	}

}

func Test1000Apps(t *testing.T) {
	boot.Boot()
	wg := new(sync.WaitGroup)
	for i:=0;i<200;i++{
		wg.Add(1)
		time.Sleep(100* time.Millisecond)
		go testapp(wg,t, i)
	}
	time.Sleep(100* time.Millisecond)
	wg.Wait()


}

func testapp(wg *sync.WaitGroup, t *testing.T, i int){
	go func(){
		clientnode := aurarath.NewNode()
		clientnode.RegisterImplementation(new(zaurarath.Implementation))
		client := aurmellon.NewInterface(clientnode)
		c := make(chan interface {})
		client.SupernodeConnected.Then(testcompleter(c))
		clientnode.Run()
		select {
		case <-time.After(50*time.Second):
			t.Error("supernode not Found",i,clientnode.Self.Id)
		case <-c:

		}


		wg.Done()
		clientnode.Stop()

	}()
}


func testlistener(c chan interface{}) stream2go.Suscriber {
	return func(d interface{}) {
		c <- d
	}
}


func testcompleter(c chan interface{}) future2go.CompletionFunc {
	return func(d interface{}) interface {}{
		c <- d
		return nil
	}
}
func testcompletererr(c chan interface{}) future2go.ErrFunc {
	return func(d error) (interface {},error){
		c <- d
		return nil,nil
	}
}
