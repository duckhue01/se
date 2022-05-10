# Best practices for writing clean and maintainable Go code

## the SOLID principles of OOP design
contrary to other, traditional OOP languages, Go has no built in support for classes. However, its does supports the concepts of interfaces and structs 

### Single responsibility
"in any well-designed system, objects should only have a single responsibility"
1. the problem:

```go
package main

// in any well-designed system, objects should only have a single responsibility
type (
	Drone struct {
	}
	Vec3 struct {
	}
	Target struct {
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

// DetectTargets captures an image of the drone's field of view (FoV) using
// the on-board camera and feeds it to a pre-trained SSD MobileNet V1 neural
// network to detect and classify interesting nearby targets. For more info
// on this model see:
func (d *Drone) DetectTargets() ([]*Target, error) {
	return nil, nil
}
```

- the code violates the SRP by conflating two separate responsibility: navigating the drone, detecting targets in close proximity to the drone
- the extra coupling that is introduced makes the implementation harder to maintain and extend for instance, what if we want to evaluate different neural network models? what if we want to use the same object recognition code for different Drone type?
  
2. the solution
- we can expose a method on the Drone object to capture and return an image using the camera. at this point, you may be thinking: wait, isn't image capturing a different responsibility than navigation? The answer is: it's all a matter of perspective! Describing and assigning responsibilities to objects is an art in itself, and quite a subjective one. Conversely, we could counter-argue that navigation needs access to various sources of sensor data, and the camera is one of them.
- in a second refactoring step, we can extract the target detection code into a separate, standalone object that would allow us to move on with the object-recognition model evaluation without having to modify any of the code in the drone type

```go
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
```

3. the benefit
   
- we reduced the coupling between our code.
- easier to maintain and extend.
- solve the problem when we want to evaluate different neural network models for object-recognition purposes
- need to cooperate with to solve the problem when we want to use the same object recognition code for different `Drone` types



### Open/closed principle
"a software module should be open for extension but closed for modification"
- (on external package scope) almost all go program import and use types from a host of other packages, some being part of the go standard library whereas other packages are provided by third parties. Any software engineer that imports a package into their code base should always safely assume that all the types of imported package is immutable. valid closed principle
- Go does not support inheritance, that leaves composition as the only viable approach for extending existing code.

1. the problem
- we want to create an object that easy to extend but close close to modify

```go
package main

import "fmt"

// a software module should be open for extension but closed for modification

type (
	Sword struct {
		name string // Important tip for RPG players: always name your swords!
	}

	EnchantedSword struct {
		Sword
	}
)

// Damage returns the damage dealt by this sword.
func (Sword) Damage() int {
	return 2
}

// String implements fmt.Stringer for the Sword type.
func (s Sword) String() string {
	return fmt.Sprint(s.Damage())
}

// Damage returns the damage dealt by the enchanted sword.
func (EnchantedSword) Damage() int {
	return 42
}
```
- since our enchanted sword is merely a generic sword that deals a different amount of damage, we will apply the `open principle` and use composition to create a new type that embeds the sword type and overrides the implementation of the `Damage` method
- (on internal scope) when we call String method of `EnchantedSword` instance it doesn't return the damage of `EnchantedSword`. because the method it nothing more than a function with a argument is `struct` called `receiver` so when we call that method it receive `Sword` instance as the argument. In order to solve that we need to implement the `String` method for `EnchantedSword` instance. the close principle is applied, because we `EnchantedSword` can't modify the implementation of the methods that have been defined on the embedded `Sword` instance.
  
- *inheritance will automatically re implement the method - as i remember*

1. the benefits
### Liskov substitution

### interface segregation

### dependency inversion