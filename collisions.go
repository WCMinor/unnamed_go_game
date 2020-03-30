package main

import (
	"math"
)

type circle struct {
	center vector
	radius float64
}

type rect struct {
	center vector	
	width, height float64
}

// Detects whether the distance between 2 circles is less or equal than the sum the 2 circles' radius
func circleToCircleCollides(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c2.center.x - c1.center.x, 2) + math.Pow(c2.center.y - c1.center.y, 2))
	return dist <= c1.radius + c2.radius
}

func rectToRectCollides(r1, r2 rect) bool {
	return (
		//Is the RIGHT edge of r1 to the RIGHT of the LEFT edge of r2?
		(r1.center.x + (r1.width / 2.0)) > (r2.center.x - (r2.width / 2.0)) &&
		//Is the LEFT edge of r1 to the LEFT of the RIGHT edge of r2?
		(r1.center.x - (r1.width / 2.0)) < (r2.center.x + (r2.width / 2.0)) &&
		//Is the BOTTOM edge of r1 BELOW the TOP edge of r2?
		(r1.center.y + (r1.height / 2.0)) > (r2.center.y - (r2.height / 2.0)) &&
		//Is the TOP edge of r1 ABOVE the BOTTOM edge of r2?
		(r1.center.y - (r1.height / 2.0)) < (r2.center.y + (r2.height / 2.0)))
		// collision detected!
}

func checkCollisions() error {
	for i := 0 ; i < len(gameElements) -1; i++ {
		for j := i+1; j < len(gameElements); j++ {
			for _, c1 := range gameElements[i].collisionPoints {
				for _, c2 := range gameElements[j].collisionPoints {
					if circleToCircleCollides(c1, c2) && gameElements[i].active && gameElements[j].active {
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
	for i := 0 ; i < len(gameElements) -1; i++ {
		for j := i+1; j < len(gameElements); j++ {
			for _, r1 := range gameElements[i].collisionRects {
				for _, r2 := range gameElements[j].collisionRects {
					if rectToRectCollides(r1, r2) && gameElements[i].active && gameElements[j].active {
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