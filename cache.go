package changenow

type Pair struct {
	Name      string
	Source    string
	Target    string
	MinAmount float64
	IsActive  bool
}

var pairs = make(map[string]*Pair)

func NewPair(name string) *Pair {
	source, target := sourceTarget(name)
	return &Pair{
		Name:   name,
		Source: source,
		Target: target,
	}
}

func addPair(name string) *Pair {
	pair, ok := pairs[name]
	if !ok {
		pair = NewPair(name)
		pairs[name] = pair
	}
	return pair
}

func getPair(name string) *Pair {
	return addPair(name)
}
