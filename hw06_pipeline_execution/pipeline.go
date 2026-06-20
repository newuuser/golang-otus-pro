package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

func drain(in In) {
	go func() {
		for range in {
		}
	}()
}

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	inChannel := make(Bi)
	var res In = inChannel

	for _, stage := range stages {
		res = stage(res)
	}

	go func() {
		defer close(inChannel)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				inChannel <- v
			}
		}
	}()
	outChannel := make(Bi)

	go func() {
		defer close(outChannel)
		for {
			select {
			case <-done:
				drain(res)
				return
			case v, ok := <-res:
				if !ok {
					return
				}
				outChannel <- v
			}
		}
	}()

	return outChannel
}
