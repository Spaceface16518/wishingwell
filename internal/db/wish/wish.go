package wish

type Wish struct {
	Title   string   `bson:"title"`
	Body    string   `bson:"body"`
	Owner   string   `bson:"owner"`
	Link    string   `bson:"link,omitempty"`
	Claimed []string `bson:"claimed"`
}

func (wish Wish) IsClaimed() bool {
	return len(wish.Claimed) > 0
}
