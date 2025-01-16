package entities

type Votable struct {
	VotesCount int `json:"votes_count"`
}

func NewVotable() *Votable {
	return &Votable{
		VotesCount: 0,
	}
}

func (v *Votable) IncrementVotesCount() {
	v.VotesCount++
}

func (v *Votable) DecrementVotesCount() {
	if v.VotesCount > 0 {
		v.VotesCount--
	}
}
