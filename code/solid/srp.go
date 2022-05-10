package main

import "image"

// in any well-designed system, objects should only have a single responsibility

type (
	Drone struct {
	}
	Vec3 struct {
	}
	Target struct {
	}

	MobileNet struct {
	}
)

// NavigateTo applies any required changes to the drone's speed
// vector so that its eventual position matches dst.
func (d *Drone) NavigateTo(dst Vec3) error {
	return nil
}

// Position returns the current drone position vector.
func (d *Drone) Position() Vec3 {
	return Vec3{}
}

// Position returns the current drone speed vector.
func (d *Drone) Speed() Vec3 {
	return Vec3{}
}

// CaptureImage records and returns an image of the drone's field of
// view using the on-board drone camera.
func (d *Drone) CaptureImage() (*image.RGBA, error) {
	return nil, nil
}

// DetectTargets captures an image of the drone's field of view and feeds
// it to a neural network to detect and classify interesting nearby
// targets.
func (mn *MobileNet) DetectTargets(d *Drone) ([]*Target, error) {
	return nil, nil
}
