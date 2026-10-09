package validation

type Room struct {
	Name string
	X    int
	Y    int
}

type Rooms struct {
	Start Room
	End   Room
	All   []Room
	Links []Link
}

type Link struct {
	From string
	To   string
}