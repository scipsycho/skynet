package root

type Crypto interface {
	Salt(s string) (string, error)
	Compare(hash string, s string) (bool, error)
	GenerateRandomASCIIString(int) (string, error)
}
