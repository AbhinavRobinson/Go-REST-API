package newsfeed

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

// Add new Repo
func New() *Repo {
	return &Repo{}
}

// Add Item to repo
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// Get Data from repo
func (r *Repo) GetAll() []Item {
	return r.Items
}
