package goorderbook

import (
	"github.com/google/btree"
)

// Order represents an order on a book level
type Order struct {

	// Id is the id of the order on this book level
	Id string

	// Amount is the amount that this order is willing to buy/sell
	Amount uint64

	// Timesetamp is the time that this order was placed on the book in
	// nanoseconds.
	Timestamp uint64
}

// OrderBookLevel contains a list of orders that make up the entire level
type OrderBookLevel struct {
	// The orders in this level
	Orders []Order

	// Level is the bid or ask price that lies on the book
	Level uint64
}

// Less implements the btree.Item interface for our OrderBookLevel. Returns
// true if LevelA is less than LevelB
func (r OrderBookLevel) Less(b btree.Item) bool {
	return r.Level < b.(OrderBookLevel).Level
}

// IsNil returns true if the order book level contains no information.
func (r OrderBookLevel) IsNil() bool {
	return r.Orders == nil && r.Level == 0
}

// Total gets the total number amount of an orderbook level
func (r OrderBookLevel) Total() uint64 {
	var sum uint64 = 0
	for _, v := range r.Orders {
		sum += v.Amount
	}
	return sum
}
