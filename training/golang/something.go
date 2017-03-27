package main

type Db struct {
	Shards []string
	Fd map[string] int
}

func main()  {
	db := Db{}
	db.Shards = []string{"ab"}
	db.Fd = map[string]int{}
	db.Fd = nil
}
