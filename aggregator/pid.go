package aggregator

import "time"

type pidError struct {
	Ts    time.Time `json:"t"`
	Error float64   `json:"e"`
}

func newPidError(err float64) *pidError {
	return &pidError{
		Ts:    time.Now(),
		Error: err,
	}
}

type pidController struct {
	Integral  float64   `json:"i"`
	LastError *pidError `json:"l"`
	CurError  *pidError `json:"c"`
}

func (p *pidController) feed(err float64, intabsmax *float64) {
	p.LastError = p.CurError
	p.CurError = newPidError(err)

	if p.LastError != nil {
		newint := p.Integral + (p.LastError.Error+err)*(p.CurError.Ts.Sub(p.LastError.Ts).Seconds()/2.0)

		if intabsmax != nil {
			if newint < -*intabsmax {
				newint = -*intabsmax
			} else if newint > *intabsmax {
				newint = *intabsmax
			}
		}

		p.Integral = newint
	}
}

func (p *pidController) value(kp, ki, kd float64) float64 {
	pv := p.CurError.Error
	iv := p.Integral
	dv := float64(0)

	if p.LastError != nil {
		dv = (p.CurError.Error - p.LastError.Error) / p.CurError.Ts.Sub(p.LastError.Ts).Seconds()
	}

	return kp*pv + kd*dv + ki*iv
}
