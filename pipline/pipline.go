package pipline

import (
	"reflect"
	"sync"
	"mdocker/config"
)

func RunPipline( ob interface{}, method string, args []string, done chan struct{}) []string {
	p := Pipline{Ob:ob, Action:method, Args:args, Wn:config.Conf.Wn, Done:done}
	return p.Run()
}

type Pipline struct {
	Ob interface{}
	Action string
	Args []string
	Done chan struct{}
	Wn int
}

func NewPipline( ob interface{}, method string, args []string, done chan struct{}) *Pipline {
	return &Pipline{Ob:ob, Action:method, Args:args, Wn:4, Done:done}
}

func (p *Pipline) gen() <-chan string {
	out := make(chan string)
	go func() {
		for _, n := range p.Args {
			select {
			case out <- n:
			case <-p.Done:
				return
			}
		}
		close(out)
	}()
	return out
}

func (p *Pipline) work(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for arg := range in {
			select {
			case out <- p.runMethod(arg):
			case <-p.Done:
				return
			}
		}
		close(out)
	}()
	return out
}

func (p *Pipline) merge(cs []<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	output := func(c <-chan string) {
		for n := range c {
			select {
			case out <- n:
			case <-p.Done:
				return
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func (p *Pipline) runMethod(arg string) string {
	obValue := reflect.ValueOf(p.Ob)
	method := obValue.MethodByName(p.Action)
	if(method == reflect.Value{}) {
		return "no method "+ p.Action
	}
	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(arg)
	res := method.Call(args)
	return res[0].String()
}

func (p *Pipline) Run() []string {
	in := p.gen()
	cs := make([]<-chan string, p.Wn)
	for i := 0; i < p.Wn; i++ {
		cs[i] = p.work(in)
	}

	out := p.merge(cs)
	var res []string
	for n := range out {
		res = append(res, n)
	}
	return res
}
