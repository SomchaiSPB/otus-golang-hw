package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return in
	}

	for _, stage := range stages {
		if stage == nil {
			continue
		}
		wrapped := wrap(in, done)
		in = stage(wrapped)
	}

	return in
}

func wrap(in In, done In) Out {
	out := make(Bi)

	go func() {
		defer func() {
			close(out)
			for range in {
				// skip
			}
		}()
		for {
			select {
			case <-done:
				return
			case data, ok := <-in:
				if !ok {
					return
				}
				out <- data
			}
		}
	}()

	return out
}
