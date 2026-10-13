package handlers

type RoomDetails struct {
	Name string
	X    int
	Y    int
}

type Rooms struct {
	Start RoomDetails
	End   RoomDetails
	All   []RoomDetails
	Links []Link
}

type Link struct {
	From string
	To   string
}
type Edge struct{ To, Cap, Cost int }

type FlowNet struct {
	Edges []Edge
	Adj   [][]int
}
