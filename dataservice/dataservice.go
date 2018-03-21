package dataservice

// DumbDictionary describes what this dumb dictionary should do.
type DumbDictionary interface {
	Init(filePath string) error
	Get(name string) (int32, error)
	Delete(name string) (string, error)
}

// DumbEntry is a very dumb entry
type DumbEntry struct {
	Name   string
	Length int32
}

// DumbList it's a list of dumb entries
type DumbList struct {
	List []DumbEntry
}
