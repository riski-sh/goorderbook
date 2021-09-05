package goorderbook

import (
  "testing"
)

// TestTotal asserts that the aggregate total on the level of the book
// is calculated correctly.
func TestTotal(t *testing.T) {
  var obl OrderBookLevel

  obl.Level = 1
  obl.Orders = []Order{
    Order{
      "A",
      100,
      0,
    },
    Order{
      "B",
      100,
      1,
    },
  }

  if obl.Total() != 200 {
    t.Fatalf("%d != 200", obl.Total())
  }

}

// TestLess tests to make sure that orderbook levels have a correctly
// defined less than operator
func TestLess(t *testing.T) {
  var oblA OrderBookLevel
  oblA.Level = 1

  var oblB OrderBookLevel
  oblB.Level = 2

  if !oblA.Less(oblB) {
    t.Fatalf("%d >= %d", oblA.Level, oblB.Level);
  }
}

// TestIsNil verifies that the nil order book is well defined.
func TestIsNil(t *testing.T) {
  var oblNill OrderBookLevel
  oblNill.Level = 0
  oblNill.Orders = nil

  if !(oblNill.IsNil()) {
    t.Fatalf("OrderBookLevel%+v != nil", oblNill);
  }
}
