package goorderbook

import (
	"fmt"
	"github.com/google/btree"
)

// OrderBookSide represents either the bid or ask side of the order book
type OrderBookSide interface {

	// Puts a level on the order book
	Put(order Order, level uint64)

	// Gets a level from the order book
	Get(level uint64) OrderBookLevel

	// Removes an order from the order book
	Remove(orderId string) error

	// Modifies an order already on the book, throws an error if the order
	// can not be found
	Modify(order Order, level uint64) error

	// Max returns the maximum value in the orderbook or a nil orderbooklevel.
	// The nil orderbook level is only returned if there are no elements in
	// the book.
	Max() OrderBookLevel

	// Min returns the minimum value in the orderbook or a nil orderbooklevel.
	// The nil orderbook level is only returned if there are no elements in
	// the book.
	Min() OrderBookLevel
}

// ob is the internal struct of an interface
type obs struct {

	// Tree is the order book stored as a B-Tree
	Tree *(btree.BTree)

	// LevelMapping is a mapping between any arbitrary order id and its
	// level in the orderbook
	LevelMapping map[string]uint64

	// HiddenOrderMapping maps orderId's to orders that have no price or amount
	// because they are hidden.
	HiddenOrderMapping map[string]Order
}

// New creates a new order book
func NewSide() OrderBookSide {
	return &obs{
		btree.New(btree.DefaultFreeListSize),
		make(map[string]uint64),
		make(map[string]Order),
	}
}

// Put puts the order on the book at the specified level
func (r *obs) Put(order Order, level uint64) {

	// treeItme represents the order book level as a btree item
	var treeItem btree.Item

	// Grab the order book level
	treeItem = r.Tree.Get(OrderBookLevel{nil, level})

	// obl is the order book level that we will manipulate to replace or
	// insert back into the book
	var obl OrderBookLevel

	// Test to see if the level exists. Cast if the level exists or create
	// and empty level
	if treeItem == nil {
		obl = OrderBookLevel{[]Order{}, level}
	} else {
		obl = treeItem.(OrderBookLevel)
	}

	// Add the order to the  level
	obl.Orders = append(obl.Orders, order)

	// Replace it or insert back into the tree
	r.Tree.ReplaceOrInsert(obl)

	// Place the order Id into the LevelMapping in order to understand what level
	// the orderId when asked "what level does this ID live in"
	r.LevelMapping[order.Id] = level
}

// Get returns an arbitrary level if it exists or the nil orderbooklevel.
func (r *obs) Get(level uint64) OrderBookLevel {
	var treeItem btree.Item
	treeItem = r.Tree.Get(OrderBookLevel{nil, level})

	if treeItem == nil {
		return OrderBookLevel{nil, 0}
	} else {
		return treeItem.(OrderBookLevel)
	}
}

// Modifies an order on the book
func (r *obs) Modify(order Order, level uint64) error {
	if bookLevel, ok := r.LevelMapping[order.Id]; ok {
		// Grab the order book level
		obl := r.Get(bookLevel)

		// Search the Orders slice for the orderId
		for idx, placedOrder := range obl.Orders {
			if placedOrder.Id == order.Id {
				obl.Orders[idx] = order
				r.Tree.ReplaceOrInsert(obl)
				return nil
			}
		}
	}
  return fmt.Errorf("unable to modify order on level %d, level does not exist", level)
}

// Remove removes the order from the book
func (r *obs) Remove(orderId string) error {
	if bookLevel, ok := r.LevelMapping[orderId]; ok {
		// Grab the order book level
		obl := r.Get(bookLevel)

		// Serch the Orders slice for the orderId which MUST exist
		for idx, order := range obl.Orders {
			if order.Id == orderId {
				obl.Orders = append(obl.Orders[:idx], obl.Orders[idx+1:]...)
				if len(obl.Orders) == 0 {
					// No more orders in the level delete the level
					r.Tree.Delete(obl)
				} else {
					// Orders still exist so update the level
					r.Tree.ReplaceOrInsert(obl)
				}

				// Remove the order from the mapping because it does not exist on the
				// book anymore
				delete(r.LevelMapping, orderId)
				return nil
			}
		}
  }
  return fmt.Errorf("OrderID %s does not exist", orderId)
}

// Max returns the maximum level in the orderbook or the nil orderbooklevel
// if no levels exist in the book
func (r *obs) Max() OrderBookLevel {
	var treeItem btree.Item
	treeItem = r.Tree.Max()

	// Return nil orderbook level if treeItem returns nil otherwise cast
	// to OrderBookLevel and return
	if treeItem == nil {
		return OrderBookLevel{nil, 0}
	} else {
		return treeItem.(OrderBookLevel)
	}
}

// Min returns the minimum level in the orderbook or the nil orderbooklevel
// if no levels exist in the book
func (r *obs) Min() OrderBookLevel {
	var treeItem btree.Item
	treeItem = r.Tree.Min()

	// Return nil orderbook level if treeItem returns nil otherwise cast
	// to OrderBookLevel and return
	if treeItem == nil {
		return OrderBookLevel{nil, 0}
	} else {
		return treeItem.(OrderBookLevel)
	}
}
