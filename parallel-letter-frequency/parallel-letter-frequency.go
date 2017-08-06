package letter

const testVersion = 1

func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap)
	defer close(ch)
	for _, s := range texts {
		go func(s string, ch chan FreqMap) {
			ch <- Frequency(s)
		}(s, ch)
	}

	var res = make(FreqMap)
	for i := 0; i < len(texts); i++ {
		f := <-ch
		for r, c := range f {
			res[r] += c
		}
	}

	return res
}
