package main

import (
	"context"

	"github.com/duckhue01/se/code/org-code/circular-dependency/entity"
	"github.com/duckhue01/se/code/org-code/circular-dependency/warehouse"
)

func main() {
	entity.AcquirePacker = func(ctx context.Context) entity.Packer {
		return warehouse.AccquireRobot(context.Background())

	}
}
