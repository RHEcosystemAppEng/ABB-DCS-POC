package workflow

const (
	M_TEMP_MIN  = 70
	M_TEMP_MAX  = 85
	M_TEMP_UNIT = 1

	M_SPEED_MIN  = 5000
	M_SPEED_MAX  = 6500
	M_SPEED_UNIT = 100

	M_NOISE_MIN  = 90
	M_NOISE_MAX  = 97
	M_NOISE_UNIT = 0.5

	M_POWER_CONSUMPTION_MIN  = 14
	M_POWER_CONSUMPTION_MAX  = 21.5
	M_POWER_CONSUMPTION_UNIT = 0.5

	INCREMENT = "increment"
	DECREMENT = "decrement"
)

type Metrics struct {
	MotorTemp             Metric `json:"motor_temperature"`
	MotorSpeed            Metric `json:"motor_speed"`
	MotorNoise            Metric `json:"motor_noise"`
	MotorPowerConsumption Metric `json:"motor_power_consumption"`
}

type Metric struct {
	CurrentValue float64 `json:"value"`
	RangeMin     float64 `json:"-"`
	RangeMax     float64 `json:"-"`
	FluctUnit    float64 `json:"-"`
	Strategy     string  `json:"-"`
}

func InitMetrics() *Metrics {

	// init metrics with default values
	metrics := Metrics{
		MotorTemp:             InitMetric(M_TEMP_MIN, M_TEMP_MAX, M_TEMP_UNIT),
		MotorSpeed:            InitMetric(M_SPEED_MIN, M_SPEED_MAX, M_SPEED_UNIT),
		MotorNoise:            InitMetric(M_NOISE_MIN, M_NOISE_MAX, M_NOISE_UNIT),
		MotorPowerConsumption: InitMetric(M_POWER_CONSUMPTION_MIN, M_POWER_CONSUMPTION_MAX, M_POWER_CONSUMPTION_UNIT),
	}

	return &metrics
}

func InitMetric(min float64, max float64, unit float64) Metric {

	// init metric with default values
	metric := Metric{
		CurrentValue: min,
		RangeMin:     min,
		RangeMax:     max,
		FluctUnit:    unit,
		Strategy:     INCREMENT,
	}

	return metric
}

func (m *Metric) DetermineMetricStrategy() {

	if m.CurrentValue == m.RangeMax {
		// if metric has reached maximum range limit, change strategy to decrement
		m.Strategy = DECREMENT
	} else if m.CurrentValue == m.RangeMin {
		// if metric has reached minimum range limit, change strategy to increment
		m.Strategy = INCREMENT
	}
}

func (m *Metric) AdvanceMetricValue() {
	if m.Strategy == INCREMENT {
		// if metric strategy is increment, increase current metric value by one unit
		m.CurrentValue += m.FluctUnit
	} else {
		// if metric strategy is decrement, decrease current metric value by one unit
		m.CurrentValue -= m.FluctUnit
	}
}
