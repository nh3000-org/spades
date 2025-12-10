package config

type PS struct {
	Name         string
	Bid          int
	Bags         int
	Tricks       int
	Score        int
	ScoreHistory []string

}

func NewPlayer(name string) PS {
	p := PS{}
	p.Name = name
	p.Bid = 0
	p.Bags = 0
	p.Tricks = 0
	p.Score = 0
	p.ScoreHistory = []string{""}

	return p
}
