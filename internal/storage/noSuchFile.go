package storage

type NoSuchFile struct{}

func (n *NoSuchFile) Error() string {
	return "Error: no such file exists"
}
