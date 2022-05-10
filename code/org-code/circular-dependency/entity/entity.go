package entity

import (
	"context"
)

type (
	// Box contains a list of items that are shipped to the customer.
	Box struct {
		// various fields
	}

	Item struct {
	}

	Packer interface {
		Pack(*Item, *Box) error
	}
)

// AcquirePacker returns a Packer instance.
var AcquirePacker func(context.Context) Packer

// Pack qty items of type i into the box.
func (b *Box) Pack(i *Item, qty int) error {
	_ = AcquirePacker(context.Background())
	return nil

}
