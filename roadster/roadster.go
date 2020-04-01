package roadster

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jfoster/eunos/internal/date"
)

func ParseVIN(vin string) (v *VIN, err error) {
	v = new(VIN)

	m := regexp.MustCompile(`([A-Z]{2})([0-9])([A-Z]{1,2})-*(\d{6})`).FindStringSubmatch(vin)

	if m != nil && len(m) > 0 {
		v.Model = m[1]

		v.Engine, err = strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}

		v.Edition = m[3]

		v.Sequence, err = strconv.Atoi(m[4])
		if err != nil {
			return nil, err
		}

		return v, nil
	}
	return nil, fmt.Errorf("Error parsing %s", vin)
}

type VinDates []struct {
	Date date.YYYYMM `yaml:"Date"`

	VIN `yaml:"VIN"`
}

type VIN struct {
	Model    string `yaml:"Model"`
	Engine   int    `yaml:"Engine"`
	Edition  string `yaml:"Edition"`
	Sequence int    `yaml:"Sequence"`
}

func (v VIN) String() string {
	return v.Model + strconv.Itoa(v.Engine) + v.Edition + "-" + strconv.Itoa(v.Sequence)
}
