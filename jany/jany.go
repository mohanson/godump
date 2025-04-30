// Package jany provides a set of functions to parse and interact with JSON data.
package jany

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/mohanson/godump/doa"
)

// Jany is a struct that holds any type of data, allowing for flexible parsing and manipulation.
type Jany struct {
	j any
}

// Reader returns a new Jany instance from the given reader.
func Reader(r io.Reader) (*Jany, error) {
	j := new(Jany)
	dec := json.NewDecoder(r)
	dec.UseNumber()
	err := dec.Decode(&j.j)
	return j, err
}

// Bool returns the bool representation of the current Jany instance.
func (j *Jany) Bool() bool {
	return j.j.(bool)
}

// Dict returns the dict representation of the current Jany instance.
func (j *Jany) Dict() map[string]*Jany {
	a := j.j.(map[string]any)
	r := map[string]*Jany{}
	for k, v := range a {
		r[k] = &Jany{j: v}
	}
	return r
}

// Float32 returns the float32 representation of the current Jany instance.
func (j *Jany) Float32() float32 {
	return float32(doa.Try(strconv.ParseFloat(j.j.(json.Number).String(), 32)))
}

// Float64 returns the float64 representation of the current Jany instance.
func (j *Jany) Float64() float64 {
	return float64(doa.Try(strconv.ParseFloat(j.j.(json.Number).String(), 64)))
}

// Get returns the Jany[k].
func (j *Jany) Get(k string) *Jany {
	a := j.j.(map[string]any)
	return &Jany{j: a[k]}
}

// Idx returns the Jany[i].
func (j *Jany) Idx(k int) *Jany {
	a := j.j.([]any)
	return &Jany{j: a[k]}
}

// Int8 returns the int8 representation of the current Jany instance.
func (j *Jany) Int8() int8 {
	return int8(doa.Try(strconv.ParseInt(j.j.(json.Number).String(), 0, 8)))
}

// Int16 returns the int16 representation of the current Jany instance.
func (j *Jany) Int16() int16 {
	return int16(doa.Try(strconv.ParseInt(j.j.(json.Number).String(), 0, 16)))
}

// Int32 returns the int32 representation of the current Jany instance.
func (j *Jany) Int32() int32 {
	return int32(doa.Try(strconv.ParseInt(j.j.(json.Number).String(), 0, 32)))
}

// Int64 returns the int64 representation of the current Jany instance.
func (j *Jany) Int64() int64 {
	return int64(doa.Try(strconv.ParseInt(j.j.(json.Number).String(), 0, 64)))
}

// Int returns the int representation of the current Jany instance.
func (j *Jany) Int() int {
	return int(doa.Try(strconv.ParseInt(j.j.(json.Number).String(), 0, 64)))
}

// List returns the list representation of the current Jany instance.
func (j *Jany) List() []*Jany {
	a := j.j.([]any)
	r := make([]*Jany, len(a))
	for i, e := range a {
		r[i] = &Jany{j: e}
	}
	return r
}

// String returns the string representation of the current Jany instance.
func (j *Jany) String() string {
	return j.j.(string)
}

// Uint8 returns the uint8 representation of the current Jany instance.
func (j *Jany) Uint8() uint8 {
	return uint8(doa.Try(strconv.ParseUint(j.j.(json.Number).String(), 0, 8)))
}

// Uint16 returns the uint16 representation of the current Jany instance.
func (j *Jany) Uint16() uint16 {
	return uint16(doa.Try(strconv.ParseUint(j.j.(json.Number).String(), 0, 16)))
}

// Uint32 returns the uint32 representation of the current Jany instance.
func (j *Jany) Uint32() uint32 {
	return uint32(doa.Try(strconv.ParseUint(j.j.(json.Number).String(), 0, 32)))
}

// Uint64 returns the uint64 representation of the current Jany instance.
func (j *Jany) Uint64() uint64 {
	return uint64(doa.Try(strconv.ParseUint(j.j.(json.Number).String(), 0, 64)))
}

// Uint returns the uint representation of the current Jany instance.
func (j *Jany) Uint() uint {
	return uint(doa.Try(strconv.ParseUint(j.j.(json.Number).String(), 0, 64)))
}
