package newsfeed

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

// Add new Repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// Add Item to repo
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// Get Data from repo
func (r *Repo) GetAll() []Item {
	return r.Items
}
