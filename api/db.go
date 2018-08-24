package api

type DB interface {
	Query(startTime int64, endTime int64) []interface{}
}
