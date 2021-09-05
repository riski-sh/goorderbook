package goorderbook

// OrderBook builds upon the OrderBookSide int64erface. Orderbook simply holds
// two order book sides, one for the bid side and one for the ask side. The
// functions PutBid/PutAsk/Get*** etc... are pass through functions to the
// correct order book side.
type OrderBook interface {

	// PutBid calls Put on the bid side
	PutBid(order Order, level uint64)

	// PutAsk calls Put on the ask side
	PutAsk(order Order, level uint64)

	// GetBid calls Get on the bid side
	GetBid(level uint64) OrderBookLevel

	// GetAsk calls Get on the ask side
	GetAsk(level uint64) OrderBookLevel

	// RemoveBid calls Remove of the bid side
	RemoveBid(orderId string) error

	// RemoveAsk calls Remove on the ask side
	RemoveAsk(orderId string) error

	// ModifyBid modifies an order the bid side
	ModifyBid(order Order, level uint64) error

	// ModifyAsk modifies an order on the ask side
	ModifyAsk(order Order, level uint64) error

	// TopBid returns the top of the bid book
	TopBid() OrderBookLevel

	// TopAsk returns the top of the ask book
	TopAsk() OrderBookLevel

	// Sequence returns the sequence number of the last message
	GetSequence() uint64

	// SetSequence sets the sequence number for this orderbook
	SetSequence(sequence uint64)
}

// ob is the order book structure that holds both sides of the book
type ob struct {

	// Bid represents the bid side of the order book
	Bid OrderBookSide

	// Ask represents the ask side of the order book
	Ask OrderBookSide

	// Sequence is the sequence number of the last message to be processed
	// in the book. The caller maintining the book is responsible for calling
	// SetSequence to keep the recorded sequence number up to date.
	Sequence uint64
}

// New create a new order book
func New() OrderBook {
	return &ob{
		NewSide(),
		NewSide(),
		0,
	}
}

func (r *ob) ModifyBid(order Order, level uint64) error {
	return r.Bid.Modify(order, level)
}

func (r *ob) ModifyAsk(order Order, level uint64) error {
	return r.Ask.Modify(order, level)
}

func (r *ob) GetSequence() uint64 {
	return r.Sequence
}

func (r *ob) SetSequence(sequence uint64) {
	r.Sequence = sequence
}

func (r *ob) PutBid(order Order, level uint64) {
	r.Bid.Put(order, level)
}

func (r *ob) PutAsk(order Order, level uint64) {
	r.Ask.Put(order, level)
}

func (r *ob) GetBid(level uint64) OrderBookLevel {
	return r.Bid.Get(level)
}

func (r *ob) GetAsk(level uint64) OrderBookLevel {
	return r.Ask.Get(level)
}

func (r *ob) RemoveBid(orderId string) error {
	return r.Bid.Remove(orderId)
}

func (r *ob) RemoveAsk(orderId string) error {
	return r.Ask.Remove(orderId)
}

func (r *ob) TopBid() OrderBookLevel {
	return r.Bid.Max()
}

func (r *ob) TopAsk() OrderBookLevel {
	return r.Ask.Min()
}
