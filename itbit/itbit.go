package itbit

import (
	"encoding/json"
	"strconv"
)

const ItBitEndpoint = "https://api.itbit.com/v1/"

// ItBitFloat is a float64 alias that implements UnmarshalJSON which
// the underlying bytes to be string representation of numbers.
//
// The itBit API returns strings to represent all numbers in order
// to avoid precision errors. ItBitFloat will parse the strings
// and store the float64 representation.
type ItBitFloat float64

func (f *ItBitFloat) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*f = ItBitFloat(float)
	return nil
}
