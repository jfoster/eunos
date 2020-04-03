package roadster

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/jfoster/eunos/internal/date"
	"gopkg.in/yaml.v2"
)

const VINRegex = `([A-Z]{2})([0-9])([A-Z]{1,2})-*(\d{6})`

func ParseVIN(vin string, tablepath string) (v *VIN, err error) {
	v = new(VIN)

	m := regexp.MustCompile(VINRegex).FindStringSubmatch(vin)
	if m == nil || len(m) == 0 {
		return nil, fmt.Errorf("Error parsing %s", vin)
	}

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

	date, err := v.GetDate(tablepath)
	if err != nil {
		return nil, err
	}
	v.Date = date

	return v, nil
}

type VIN struct {
	Model    string `yaml:"Model"`
	Engine   int    `yaml:"Engine"`
	Edition  string `yaml:"Edition"`
	Sequence int    `yaml:"Sequence"`

	Date *date.YYYYMM `yaml:"Date"`
}

func (v VIN) GetDate(tablepath string) (*date.YYYYMM, error) {
	table, err := ParseTable(tablepath)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(table)-2; i++ {
		x := table[i]
		y := table[i+1]

		if x.Sequence <= v.Sequence && v.Sequence < y.Sequence && x.Model == v.Model && x.Engine == v.Engine {
			return x.Date, nil
		}
	}

	return nil, fmt.Errorf("Could not find manufacture date for vin %s", v)
}

func (v VIN) String() string {
	return v.Model + strconv.Itoa(v.Engine) + v.Edition + "-" + strconv.Itoa(v.Sequence)
}

type VINs []VIN

func ParseTable(tablepath string) (VINs, error) {
	var table VINs

	file, err := os.Open(tablepath)
	if err != nil {
		log.Fatal(err)
	}

	if filepath.Ext(tablepath) == ".yml" {
		err = yaml.NewDecoder(file).Decode(&table)
		if err != nil {
			return nil, err
		}
	}

	return table, nil
}
