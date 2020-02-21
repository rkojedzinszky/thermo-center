package sensorvalue

// htu21DTemperature holds temperature from HTU21D sensor
type htu21DTemperature struct {
	value float64
}

func (t *htu21DTemperature) Metric() string {
	return "Temperature"
}

func (t *htu21DTemperature) Value() float64 {
	return t.value
}

// htu21DHumidity holds humidity from HTU21D sensor
type htu21DHumidity struct {
	value float64
}

func (h *htu21DHumidity) Metric() string {
	return "Humidity"
}

func (h *htu21DHumidity) Value() float64 {
	return h.value
}

// Compensate compensates humdity according to HTU21D
// datasheet recommendation
func (h *htu21DHumidity) Compensate(t *htu21DTemperature) {
	h.value -= 0.15 * (25.0 - t.value)
}

func htu21dTemperatureParser(r RawValue) (SensorValue, int32) {
	raw := int32(r.Unsigned)

	return &htu21DTemperature{
		value: -46.85 + 175.72*float64(raw)/16384.0,
	}, raw
}

func htu21dHumidityParser(r RawValue) (SensorValue, int32) {
	raw := int32(r.Unsigned)

	return &htu21DHumidity{
		value: -6 + 125*float64(r.Unsigned)/4096.0,
	}, raw
}

func init() {
	register(0, 0, htu21dTemperatureParser)
	register(1, 0, htu21dHumidityParser)
}
