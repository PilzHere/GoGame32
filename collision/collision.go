package collision

import (
	"github.com/faiface/pixel"
	"github.com/pilzhere/GoGame32/collider"
	"math"
)

const (
	DefaultBit     = 0
	TerrainBit     = 1
	PlayerBit      = 2
	PlayerBlockBit = 4
	EnemyBit       = 8

	TerrainMask     = PlayerBlockBit | EnemyBit
	PlayerMask      = TerrainBit | PlayerBlockBit | EnemyBit
	PlayerBlockMask = TerrainBit | PlayerBit | EnemyBit
	EnemyMask       = TerrainBit | PlayerBit | PlayerBlockBit | EnemyBit
)

func GetIntersectionDepth(thisCollider collider.Collider, otherCollider collider.Collider) pixel.Vec {
	distanceX := thisCollider.Rect.Center().X - otherCollider.Rect.Center().X
	distanceY := thisCollider.Rect.Center().Y - otherCollider.Rect.Center().Y
	minDistanceX := thisCollider.Rect.W()/2 + otherCollider.Rect.W()/2
	minDistanceY := thisCollider.Rect.H()/2 + otherCollider.Rect.H()/2

	if math.Abs(distanceX) >= minDistanceX || math.Abs(distanceY) >= minDistanceY {
		return pixel.ZV
	}

	var depthX float64
	if distanceX > 0 {
		depthX = minDistanceX - distanceX
	} else {
		depthX = -minDistanceX - distanceX
	}

	var depthY float64
	if distanceY > 0 {
		depthY = minDistanceY - distanceY
	} else {
		depthY = -minDistanceY - distanceY
	}

	return pixel.Vec{
		X: depthX,
		Y: depthY,
	}
}

func GetIntersectionArea(intersection pixel.Vec) float64 {
	return math.Abs(intersection.X * intersection.Y)
}
