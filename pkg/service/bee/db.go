package bee

type DB interface {
	Write(interface{})
}
