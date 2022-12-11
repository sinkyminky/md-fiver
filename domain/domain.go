package domain

type Duty struct {
	parallelCount int
	urls          []string
}

func NewDuty(pc int, urls []string) *Duty {
	return &Duty{
		parallelCount: pc,
		urls:          urls,
	}
}

func (d Duty) ParallelCount() int {
	return d.parallelCount
}

func (d Duty) URLs() []string {
	return d.urls
}
