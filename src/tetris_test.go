package main

import "testing"

func testAt(p *Playfield, x, y int) (val int, panicked bool) {
	panicked = false
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	val = p.at(x,y)
	return val, panicked
}


func testSet(p *Playfield, x, y, val int) (panicked bool) {
	panicked = false
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	p.set(x,y,val)
	return panicked
}

func TestAtSetBounds(t *testing.T) {
	const width = 5
	const height = 7
	p := NewPlayfield(width, height)

	v, panicked := testAt(&p,0,0)
	if panicked { t.Error("Unwarranted panic") }
	if v != 0 { t.Error("Bad value") }
	panicked = testSet(&p,width-1,height-1,100)
	if panicked { t.Error("Unwarranted panic") }
	v, panicked = testAt(&p,width-1,height-1)
	if panicked { t.Error("Unwarranted panic") }
	if v != 100 { t.Error("Bad value") }

	_, panicked = testAt(&p,width,0)
	if !panicked { t.Error("At failed to trigger panic") }
	_, panicked = testAt(&p,0,height)
	if !panicked { t.Error("At failed to trigger panic") }
	panicked = testSet(&p,width,0,1)
	if !panicked { t.Error("Set failed to trigger panic") }
	panicked = testSet(&p,0,height,1)
	if !panicked { t.Error("Set failed to trigger panic") }
}
