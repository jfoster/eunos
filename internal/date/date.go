package date

import "time"

const fmtYYYYMM = `2006-01`

// YYYYMM provides a year and month type.
type YYYYMM time.Time

// MarshalJSON outputs JSON.
func (d YYYYMM) MarshalJSON() ([]byte, error) {
	return []byte(wrap(d.String())), nil
}

// UnmarshalJSON handles incoming JSON.
func (d *YYYYMM) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(wrap(fmtYYYYMM), string(b))
	if err != nil {
		return err
	}
	*d = YYYYMM(t)
	return nil
}

func (d YYYYMM) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

func (d *YYYYMM) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string

	if err := unmarshal(&str); err != nil {
		return err
	}

	t, err := time.Parse(fmtYYYYMM, str)
	if err != nil {
		return err
	}
	*d = YYYYMM(t)
	return nil
}

func (d YYYYMM) String() string {
	return time.Time(d).Format(fmtYYYYMM)
}

func wrap(s string) string {
	return "\"" + s + "\""
}
