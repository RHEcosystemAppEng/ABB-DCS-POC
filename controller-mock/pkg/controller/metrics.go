package controller

const (
	INCREMENT = "increment"
	DECREMENT = "decrement"
)

type Metric struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	RangeMin  float64 `json:"range_min"`
	RangeMax  float64 `json:"range_max"`
	FluctUnit float64 `json:"fluct_unit"`
	Strategy  string  `json:"strategy"`
}

func (m *Metric) determineMetricStrategy() {

	if m.Value == m.RangeMax {
		// if metric has reached maximum range limit, change strategy to decrement
		m.Strategy = DECREMENT
	} else if m.Value == m.RangeMin {
		// if metric has reached minimum range limit, change strategy to increment
		m.Strategy = INCREMENT
	}
}

func (m *Metric) advanceMetricValue() {
	if m.Strategy == INCREMENT {
		// if metric strategy is increment, increase current metric value by one unit
		m.Value += m.FluctUnit
	} else {
		// if metric strategy is decrement, decrease current metric value by one unit
		m.Value -= m.FluctUnit
	}
}
