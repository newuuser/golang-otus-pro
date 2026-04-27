package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

// без : n stage первым обрабатывает done, n-1 пробует передать -> deadlock
func drain(in Out) {
	for range in {
	}
}

// wraps in chan + done chan to in chan that closes at done signal
func gen(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				drain(in)
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- v:
				case <-done:
					drain(in)
					return
				}
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = gen(in, done)
		in = stage(in)
	}
	return in
}
