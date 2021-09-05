package goorderbook

import (
  "testing"
)

// TestNewSide verifies the create a new btree to be used as an order book.
func TestNewSide(t *testing.T) {
  book := NewSide()

  if book == nil {
    t.Fatalf("%s", "unable to create new orderbookside");
  }
}

// TestPut tests the put order functionatlity of the orderbook.
func TestPut(t *testing.T) {

  book := NewSide()

  order := Order{
    "A",
    100,
    1,
  }

  book.Put(order, 1)
  ordern := Order{
    "B",
    100,
    2,
  }
  book.Put(ordern, 1);
  if book.Get(1).Orders[0] != order {
    t.Fatalf("%s", "put %+v but did not get it back");
  }
  if book.Get(1).Orders[1] != ordern {
    t.Fatalf("%s", "put %+v but did not get it back");
  }

}

// TestGet calls test put since TestPut and TestGet work hand in hand.
func TestGet(t *testing.T) {
  TestPut(t)

  book := NewSide()
  err := book.Get(1)

  if !err.IsNil() {
    t.Fatalf("%s", "expected empty orderbook level for empty orderbook");
  }
}

// TestRemove tests the removable of an added order to the book
func TestRemove(t *testing.T) {
  book := NewSide()

  order := Order{
    "A",
    100,
    1,
  }

  orderb := Order{
    "B",
    100,
    2,
  }

  book.Put(order, 1)
  book.Put(orderb, 1)

  err := book.Remove("ImNotInHere")
  if err == nil{
    t.Fatalf("expected error with order that doesn't exist");
  }

  err = book.Remove("A")
  if err == nil && len(book.Get(1).Orders) != 1 {
    t.Fatalf("%s", "order removal failed");
  }

  err = book.Remove("B")
  if err == nil && len(book.Get(1).Orders) != 0 {
    t.Fatalf("%s", "order removal failed");
  }
}

// TestMax verified the indenteded behavior of the Max function.
func TestMax(t *testing.T) {
  book := NewSide()

  var obl = book.Max()
  if !obl.IsNil() {
    t.Fatalf("%s", "expected nil because there are no orders");
  }

  book.Put(Order{"A", 100, 1}, 10)
  book.Put(Order{"B", 100, 2}, 20)

  if book.Max().Orders[0].Id != "B" {
    t.Fatalf("%s", "expected order B but got something else");
  }
}

// TestMin verified the indenteded behavior of the Max function.
func TestMin(t *testing.T) {
  book := NewSide()

  var obl = book.Min()
  if !obl.IsNil() {
    t.Fatalf("%s", "expected nil because there are no orders");
  }

  book.Put(Order{"A", 100, 1}, 10)
  book.Put(Order{"B", 100, 2}, 20)

  if book.Min().Orders[0].Id != "A" {
    t.Fatalf("%s", "expected order A but got something else");
  }
}

// TestModify verifies the functionality of order book order modifications
func TestModify(t *testing.T) {
  book := NewSide()

  order := Order{
    "A",
    100,
    1,
  }
  err := book.Modify(order, 1)

  if err == nil {
    t.Fatalf("%s", "expected error on modification of order that doesn't exist");
  }

  book.Put(order, 1);

  order.Amount = 200
  err = book.Modify(order, 1)

  if err != nil {
    t.Fatalf("%+v", err);
  }

  if (book.Get(1).Total() != 200) {
    t.Fatalf("%s", "modification failed expected 200 units");
  }

}
