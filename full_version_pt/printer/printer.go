package printer

type Printer struct {
	limit int
}

func NewPrinter(limit int) *Printer {
	return &Printer{limit: limit}
}

func (p *Printer) PrintNumbers(number <-chan int, resultChanel chan<- []int) {

	resultSlice := make([]int, 0)
	tempMap := map[int]struct{}{}

	for len(resultSlice) != p.limit {
		randomNumber := <-number
		if randomNumber == 0 {
			continue
		}
		tempMap[randomNumber] = struct{}{}
		resultSlice = append(resultSlice, randomNumber)

	}
	resultChanel <- resultSlice
}
