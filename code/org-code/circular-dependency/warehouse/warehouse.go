package warehouse

import (
	"context"

	"github.com/duckhue01/se/code/org-code/circular-dependency/entity"
)

type (
	// Robot navigates the warehouse floor and fetches times for packing
	Robot struct {
		//
	}
)

// AcquireRobot blocks until a Robot becomes available or until the
// context expires.
func AccquireRobot(ctx context.Context) *Robot { return nil }

// Pack instructs the robot to pick up an item from its shelf and place
// it into a box that will be shipped to the customer.
func (r *Robot) Pack(item *entity.Item, to *entity.Box) error { return nil }
