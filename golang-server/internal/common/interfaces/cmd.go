package interfaces

type CMD struct {
	Server       Server
	ConsumeEvent <-chan Event
	Util         *Util
}
