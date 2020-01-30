package chemical

type Reaction struct {
	Reactors []Chemical
	Product  Chemical
}

type Chemical struct {
	Units int
	Id    string
}
