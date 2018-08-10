package heartbeat

// Application has a trigger function which can be
// used to trigger a heart beat for the application
// that has registered.
type Application interface {
	Trigger()
}
