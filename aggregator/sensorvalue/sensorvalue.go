package sensorvalue

// RawValue holds raw value in signed and unsigned form
type RawValue struct {
	Unsigned uint16
	Signed   int16
}

// SensorValue represents a parsed metric
type SensorValue interface {
	// Metric returns name for SensorValue
	Metric() string
	// Value returns value for SensorValue
	Value() float64
}

// plainValue represents a plain parsed metric
type plainValue struct {
	metric string
	value  float64
}

func (p *plainValue) Metric() string {
	return p.metric
}

func (p *plainValue) Value() float64 {
	return p.value
}

// metricParser represents a metric type, supports parsing
type metricParser func(RawValue) (SensorValue, int32)

const maxT = 4

// Parse 2-byte encoded sensor value
func Parse(d0, d1 uint8) (SensorValue, int32) {
	var T uint8 = 0
	var b uint8 = 0x80

	for (d0 & b) != 0 {
		b >>= 1
		T++

		if T == maxT {
			return nil, 0
		}
	}

	t := (d0 >> (6 - 2*T)) & (uint8(0xff) >> (7 - T))
	d := uint16(d0&(uint8(0x3f)>>(2*T)))<<8 | uint16(d1)

	parser := parsers[T][t]
	if parser == nil {
		return nil, 0
	}

	rv := RawValue{
		Unsigned: d,
	}

	if d&(1<<(13-2*T)) != 0 {
		d |= uint16(0xffff) << (14 - 2*T)
	}

	rv.Signed = int16(d)

	return parser(rv)
}

// ParseBytes parses bytestream
func ParseBytes(data []byte) []SensorValue {
	var t *htu21DTemperature
	var h *htu21DHumidity
	values := make([]SensorValue, 0, 5)

	for len(data) >= 2 {
		if sv, _ := Parse(data[0], data[1]); sv != nil {
			values = append(values, sv)

			switch sv.(type) {
			case *htu21DTemperature:
				t = sv.(*htu21DTemperature)
			case *htu21DHumidity:
				h = sv.(*htu21DHumidity)
			}
		}

		data = data[2:]
	}

	if t != nil && h != nil {
		h.Compensate(t)
	}

	return values
}

var parsers = func() [][]metricParser {
	p := make([][]metricParser, maxT)
	for i := range p {
		p[i] = make([]metricParser, (2 << i))
	}
	return p
}()

func register(T, t uint8, m metricParser) {
	parsers[T][t] = m
}
