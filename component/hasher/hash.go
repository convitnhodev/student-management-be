package hasher

type HasherInfo interface {
	HashMd5(data string) string
}
