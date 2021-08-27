package collider

import (
	"github.com/faiface/pixel"
)

type Collider struct {
	CatBits               int // Category bits, what collision type this is.
	MaskBits              int // What this collides with.
	Rect                  pixel.Rect
	Position              pixel.Vec
	Intersection          pixel.Vec
	IntersectionArea      float64
	IntersectingColliders []Collider
	ParentInterface       Colliderer
}

func NewCollider(rect pixel.Rect, parentInterface Colliderer, position pixel.Vec, thisFilter int, thisMask int) *Collider {
	collider := &Collider{
		Rect:                  rect,
		Position:              position,
		CatBits:               thisFilter,
		MaskBits:              thisMask,
		Intersection:          pixel.V(0, 0), // idk
		IntersectionArea:      0,             // idk
		IntersectingColliders: []Collider{},
		ParentInterface:       parentInterface,
	}

	return collider
}

func (c *Collider) OnCollision(otherCollider Collider) {
	c.ParentInterface.OnCollision(otherCollider)
}
