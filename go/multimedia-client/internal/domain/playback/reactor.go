package playback

type reactor struct {
	eventCh chan StreamEvent
}

func NewReactor() Reactor {
	return &reactor{
		eventCh: make(chan StreamEvent),
	}
}

func (r *reactor) Subscribe() <-chan StreamEvent {
	return r.eventCh
}

func (r *reactor) Publish(event StreamEvent) {
	r.eventCh <- event
}

func (r *reactor) Close() {
	close(r.eventCh)
}
