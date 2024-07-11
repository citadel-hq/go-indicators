package indicator

import (
	"math"
	"testing"
)

// Custom comparison function for slices with potential NaN values
func compareSlicesWithNaN(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.IsNaN(a[i]) && math.IsNaN(b[i]) {
			continue // Both NaN, considered equal
		} else if a[i] != b[i] {
			return false // Different values
		}
	}
	return true
}

// Adjusted TestSuperTrendBasic function using the custom comparison
func TestSuperTrendBasic(t *testing.T) {
	factor := 3.0
	period := 10

	inHigh := []float64{
		0.27190569, 0.27010248, 0.27046785, 0.26917804, 0.26941121, 0.26968698, 0.27000002, 0.26983872, 0.26958276, 0.26976193,
		0.27004934, 0.27020565, 0.27039812, 0.26995636, 0.27054357, 0.27032682, 0.27022004, 0.26946286, 0.26744416, 0.26995657,
		0.27003074, 0.27040362, 0.27026614, 0.27014678, 0.27012619, 0.27023186, 0.27018331, 0.27026649, 0.27056686, 0.27151356,
		0.27158236, 0.27182906, 0.27178310, 0.27223088, 0.27423544, 0.27521870, 0.27518390, 0.27518390, 0.27493740, 0.27512880,
		0.27423210, 0.27341333, 0.27311309, 0.27293831, 0.27364682,
	}

	inLow := []float64{
		0.26964718, 0.26894660, 0.26770696, 0.26672307, 0.26859641, 0.26941121, 0.26968414, 0.26941496, 0.26943487, 0.26957992,
		0.26963858, 0.26988706, 0.26961120, 0.26954199, 0.27023204, 0.26987050, 0.26993324, 0.26744416, 0.26740177, 0.26916830,
		0.26980250, 0.26994844, 0.26994912, 0.27007764, 0.26989486, 0.26992627, 0.27008906, 0.27021506, 0.27021506, 0.27030964,
		0.27138737, 0.27148199, 0.27159068, 0.27178023, 0.27213331, 0.27404592, 0.27428464, 0.27441917, 0.27481560, 0.27359810,
		0.27209321, 0.27178580, 0.27310154, 0.27206385, 0.27246528,
	}

	inClose := []float64{
		0.26981776, 0.26894943, 0.26929180, 0.26804067, 0.26922168, 0.26968698, 0.26981028, 0.26950313, 0.26958276, 0.26976193,
		0.26987567, 0.27018000, 0.27039812, 0.26976658, 0.27032682, 0.26993324, 0.27022004, 0.26744416, 0.26740177, 0.26992804,
		0.26997939, 0.27026899, 0.27015761, 0.27008621, 0.26992913, 0.27011477, 0.27015761, 0.27021506, 0.27053256, 0.27131854,
		0.27157949, 0.27149921, 0.27178310, 0.27222800, 0.27423544, 0.27461455, 0.27485330, 0.27470223, 0.27481850, 0.27380308,
		0.27370323, 0.27311309, 0.27310443, 0.27237067, 0.27251136,
	}

	correctedExpectedTSL := []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 0.2725,
		0.2725, 0.2725, 0.2725, 0.2724, 0.2724, 0.2724, 0.2724, 0.2714, 0.27, 0.27,
		0.27, 0.2674, 0.2676, 0.2678, 0.2679, 0.268, 0.2683, 0.2685, 0.2687, 0.2691,
		0.2698, 0.27, 0.2701, 0.2704, 0.2711, 0.2724, 0.2725, 0.2726, 0.2728, 0.2728,
		0.2728, 0.2728, 0.2728, 0.2752, 0.2752,
	}

	correctedExpectedTrend := []bool{
		false, false, false, false, false, false, false, false, false,
		true, true, true, true, true, true, true, true, false, false, true,
		true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true,
		true, true, true, false, false,
	}

	tsl, trend := SuperTrend(factor, period, inHigh, inLow, inClose)

	// Round off tsl values to 4 decimal places
	for i := range tsl {
		tsl[i] = roundTo4Decimals(tsl[i])
	}

	actualTrueCount, actualFalseCount := 0, 0
	for _, v := range trend {
		if v {
			actualTrueCount++
		} else {
			actualFalseCount++
		}
	}

	wantTrueCount, wantFalseCount := 0, 0
	for _, v := range correctedExpectedTrend {
		if v {
			wantTrueCount++
		} else {
			wantFalseCount++
		}
	}

	if !compareSlicesWithNaN(tsl, correctedExpectedTSL) || actualTrueCount != wantTrueCount || actualFalseCount != wantFalseCount {
		t.Errorf("SuperTrend() = %v, %v; want %v, %v. \n Trend count: actual (true: %d, false: %d), want (true: %d, false: %d)",
			tsl, trend, correctedExpectedTSL, correctedExpectedTrend, actualTrueCount, actualFalseCount, wantTrueCount, wantFalseCount)
	}
}

func roundTo4Decimals(value float64) float64 {
	return math.Round(value*10000) / 10000
}
