package collider

type Colliderer interface {
	OnCollision(otherCollider Collider)
}
