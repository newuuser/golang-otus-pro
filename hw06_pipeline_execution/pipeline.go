package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

func drain(in In) {
	go func() {
		for {
			select {
			case _, ok := <-in:
				if !ok {
					return
				}
			}
		}
	}()
}

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	ch := make(Bi)
	var res In = ch

	for _, stage := range stages {
		res = stage(res)
	}

	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				ch <- v
			}
		}
	}()
	ret := make(Bi)

	go func() {
		defer close(ret)
		for {
			select {
			case <-done:
				drain(res)
				return
			case v, ok := <-res:
				if !ok {
					return
				}
				ret <- v
			}
		}
	}()

	return ret
}
