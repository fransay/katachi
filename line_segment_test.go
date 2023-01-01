package shapelib

import (
	"math"
	"shapelib/types"
	"testing"
)

// distance test
func TestLineSegmentDistance(t *testing.T) {
	// line segment one
	var lineSegmentOne types.LineSegment = types.LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 5, Y: 10}, PointB: struct {
		X float64
		Y float64
	}{X: 10, Y: 15}}
	// line segment two
	var lineSegmentTwo types.LineSegment = types.LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 234.543, Y: 24043.354}, PointB: struct {
		X float64
		Y float64
	}{X: 544.654, Y: 354.544}}

	// line segment one result
	var distanceResultLineSegmentOne = lineSegmentOne.Distance()
	var distanceExpectedLineSegmentOne = 7.017
	if math.Trunc(distanceResultLineSegmentOne) != math.Trunc(distanceExpectedLineSegmentOne) {
		t.Errorf("Expected %f, got %f", distanceExpectedLineSegmentOne, distanceResultLineSegmentOne)
	}

	// line segment two results
	var distanceResultLineSegmentTwo = lineSegmentTwo.Distance()
	var distanceExpectedLineSegmentTwo = 23690.84
	if math.Trunc(distanceResultLineSegmentTwo) != math.Trunc(distanceExpectedLineSegmentTwo) {
		t.Errorf("Expected %f, got %f", distanceExpectedLineSegmentTwo, distanceResultLineSegmentTwo)
	}

}

// midpoint test
func TestLineSegmentMidPoint(t *testing.T) {
	var lineSegmentOne types.LineSegment = types.LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 10, Y: 14}, PointB: struct {
		X float64
		Y float64
	}{X: 20, Y: 22}}

	// midpoint test one
	var midPointLineResultLineSegmentOne types.Point2D = lineSegmentOne.MidPoint()
	var midPointLineExpectedLineSegmentOne types.Point2D = types.Point2D{X: 15, Y: 18}
	if midPointLineExpectedLineSegmentOne != midPointLineResultLineSegmentOne {
		t.Errorf("Expected %f, got %f", midPointLineExpectedLineSegmentOne, midPointLineResultLineSegmentOne)
	}
	// midpoint test two
	var lineSegmentTwo types.LineSegment = types.LineSegment{PointA: struct {
		X float64
		Y float64
	}{X: 10, Y: 14}, PointB: struct {
		X float64
		Y float64
	}{X: 20, Y: 22}}

	// midpoint test one
	var midPointLineResultLineSegmentTwo types.Point2D = lineSegmentTwo.MidPoint()
	var midPointLineExpectedLineSegmentTwo types.Point2D = types.Point2D{X: 15, Y: 18}
	if midPointLineExpectedLineSegmentTwo != midPointLineResultLineSegmentTwo {
		t.Errorf("Expected %f, got %f", midPointLineExpectedLineSegmentTwo, midPointLineResultLineSegmentTwo)
	}

}

// bearing test
