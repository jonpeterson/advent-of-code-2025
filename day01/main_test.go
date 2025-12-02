package main

import "testing"

func TestDial_Turn_0_1(t *testing.T) {
	dial := Dial{}
	dial.Turn(1)
	dial.asset(t, 1, 0, 0)
}

func TestDial_Turn_0_100(t *testing.T) {
	dial := Dial{}
	dial.Turn(100)
	dial.asset(t, 0, 1, 1)
}

func TestDial_Turn_0_101(t *testing.T) {
	dial := Dial{}
	dial.Turn(101)
	dial.asset(t, 1, 1, 0)
}

func TestDial_Turn_0_500(t *testing.T) {
	dial := Dial{}
	dial.Turn(500)
	dial.asset(t, 0, 5, 1)
}

func TestDial_Turn_0_599(t *testing.T) {
	dial := Dial{}
	dial.Turn(599)
	dial.asset(t, 99, 5, 0)
}

func TestDial_Turn_50_1(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(1)
	dial.asset(t, 51, 0, 0)
}

func TestDial_Turn_50_50(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(50)
	dial.asset(t, 0, 1, 1)
}

func TestDial_Turn_50_151(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(151)
	dial.asset(t, 1, 2, 0)
}

func TestDial_Turn_50_110(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(110)
	dial.asset(t, 60, 1, 0)
}

func TestDial_Turn_50_1000(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(1000)
	dial.asset(t, 50, 10, 0)
}

func TestDial_Turn_0m1(t *testing.T) {
	dial := Dial{}
	dial.Turn(-1)
	dial.asset(t, 99, 0, 0)
}

func TestDial_Turn_0m100(t *testing.T) {
	dial := Dial{}
	dial.Turn(-100)
	dial.asset(t, 0, 1, 1)
}

func TestDial_Turn_0m101(t *testing.T) {
	dial := Dial{}
	dial.Turn(-101)
	dial.asset(t, 99, 1, 0)
}

func TestDial_Turn_0m500(t *testing.T) {
	dial := Dial{}
	dial.Turn(-500)
	dial.asset(t, 0, 5, 1)
}

func TestDial_Turn_0m599(t *testing.T) {
	dial := Dial{}
	dial.Turn(-599)
	dial.asset(t, 1, 5, 0)
}

func TestDial_Turn_50m1(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(-1)
	dial.asset(t, 49, 0, 0)
}

func TestDial_Turn_50m50(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(-50)
	dial.asset(t, 0, 1, 1)
}

func TestDial_Turn_50m151(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(-151)
	dial.asset(t, 99, 2, 0)
}

func TestDial_Turn_50m110(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(-110)
	dial.asset(t, 40, 1, 0)
}

func TestDial_Turn_50m1000(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(-1000)
	dial.asset(t, 50, 10, 0)
}

func TestDial_Turn_50m150p200(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(-150)
	dial.Turn(200)
	dial.asset(t, 0, 4, 2)
}

func TestDial_Turn_50m50p100(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(-50)
	dial.Turn(100)
	dial.asset(t, 0, 2, 2)
}

func TestDial_Turn_50p150m200(t *testing.T) {
	dial := Dial{Current: 50}
	dial.Turn(150)
	dial.Turn(-200)
	dial.asset(t, 0, 4, 2)
}

func TestDial_Turn_10p180(t *testing.T) {
	dial := Dial{Current: 10}
	dial.Turn(180)
	dial.asset(t, 90, 1, 0)
}

func TestDial_Turn_80p110(t *testing.T) {
	dial := Dial{Current: 80}
	dial.Turn(110)
	dial.asset(t, 90, 1, 0)
}

func TestDial_Turn_90m180(t *testing.T) {
	dial := Dial{Current: 90}
	dial.Turn(-180)
	dial.asset(t, 10, 1, 0)
}

func TestDial_Turn_20m110(t *testing.T) {
	dial := Dial{Current: 20}
	dial.Turn(-110)
	dial.asset(t, 10, 1, 0)
}

func (dial *Dial) asset(t *testing.T, current int, passedZero int, endedOnZero int) {
	if dial.Current != current {
		t.Errorf("dial.Current = %d, want %d", dial.Current, current)
	}
	if dial.PassedZero != passedZero {
		t.Errorf("dial.PassedZero = %d, want %d", dial.PassedZero, passedZero)
	}
	if dial.EndedOnZero != endedOnZero {
		t.Errorf("dial.EndedOnZero = %d, want %d", dial.EndedOnZero, endedOnZero)
	}
}
