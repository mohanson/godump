// Package jany provides a set of functions to parse and interact with JSON data.
package jany

import (
	"encoding/json"
	"io"
	"iter"
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
	return j.Dict()[k]
}

// Idx returns the Jany[i].
func (j *Jany) Idx(k int) *Jany {
	return j.List()[k]
}

// Int returns the int representation of the current Jany instance.
func (j *Jany) Int() int {
	return int(doa.Try(strconv.ParseInt(j.j.(json.Number).String(), 0, 64)))
}

// Int32 returns the int32 representation of the current Jany instance.
func (j *Jany) Int32() int32 {
	return int32(doa.Try(strconv.ParseInt(j.j.(json.Number).String(), 0, 64)))
}

// Int64 returns the int64 representation of the current Jany instance.
func (j *Jany) Int64() int64 {
	return int64(doa.Try(strconv.ParseInt(j.j.(json.Number).String(), 0, 64)))
}

// Iter returns a sequence from the current Jany instance's list.
func (j *Jany) Iter() iter.Seq2[int, *Jany] {
	a := j.j.([]any)
	return func(yield func(int, *Jany) bool) {
		for i, e := range a {
			if !yield(i, &Jany{j: e}) {
				return
			}
		}
	}
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

// Uint32 returns the uint32 representation of the current Jany instance.
func (j *Jany) Uint32() uint32 {
	return uint32(doa.Try(strconv.ParseUint(j.j.(json.Number).String(), 0, 32)))
}

// Uint64 returns the uint64 representation of the current Jany instance.
func (j *Jany) Uint64() uint64 {
	return uint64(doa.Try(strconv.ParseUint(j.j.(json.Number).String(), 0, 64)))
}
