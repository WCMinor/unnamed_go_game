package main

import (
	"math"
)

type circle struct {
	center vector
	radius float64
}

// Detects whether the distance between 2 circles is less or equal than the sum the 2 circles' radius
func collides(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c2.center.x - c1.center.x, 2) + math.Pow(c2.center.y - c1.center.y, 2))
	return dist <= c1.radius + c2.radius
}

func checkCollisions() error {
	for i := 0 ; i < len(gameElements) -1; i++ {
		for j := i+1; j < len(gameElements); j++ {
			for _, c1 := range gameElements[i].collisionPoints {
				for _, c2 := range gameElements[j].collisionPoints {
					if collides(c1, c2) && gameElements[i].active && gameElements[j].active {
						err := gameElements[i].collision(gameElements[j])
						if err != nil {
							return err
						}
						err = gameElements[j].collision(gameElements[i])
						if err != nil {
							return err
						}
					}
				} 
			}
		}
	}
	return nil
}