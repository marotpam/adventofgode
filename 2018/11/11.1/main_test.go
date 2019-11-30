package main

import "testing"

func TestCalculatingThePowerOfFuelCells(t *testing.T) {
	type testCase struct {
		cell          cell
		expectedPower int
	}

	tcs := []testCase{
		{
			cell{xy{3, 5}, 8},
			4,
		},{
			cell{xy{122, 79}, 57},
			-5,
		},{
			cell{xy{217, 196}, 39},
			0,
		},{
			cell{xy{101, 153}, 71},
			4,
		},
	}

	for _, tc := range tcs {
		if pl := tc.cell.powerLevel(); pl != tc.expectedPower {
			t.Errorf("the power level for %d should be %d, got %d instead\n", tc.cell.serialNumber, tc.expectedPower, pl)
		}
	}
}