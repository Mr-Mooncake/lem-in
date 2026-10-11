package structs

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
