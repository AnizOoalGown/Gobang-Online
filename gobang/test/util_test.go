package test

import (
	"fmt"
	"gobang/entity"
	"gobang/util"
	"testing"
)

func TestInteger(t *testing.T) {
	var a uint8 = 1
	var b uint8 = 2
	fmt.Println(a - b)
}

func TestCheckFiveOfLastStep(t *testing.T) {
	steps := make([]entity.Chess, 0)
	steps = append(steps, entity.Chess{
		I: 7,
		J: 7,
	}, entity.Chess{
		I: 1,
		J: 7,
	}, entity.Chess{
		I: 8,
		J: 8,
	}, entity.Chess{
		I: 2,
		J: 7,
	}, entity.Chess{
		I: 9,
		J: 9,
	}, entity.Chess{
		I: 3,
		J: 7,
	}, entity.Chess{
		I: 6,
		J: 6,
	}, entity.Chess{
		I: 4,
		J: 7,
	}, entity.Chess{
		I: 5,
		J: 5,
	})

	fmt.Println(util.CheckFiveOfLastStep(&steps))
}
