package column

import (
	"fmt"
	"github.com/tobgu/qframe/config/rolling"

	"github.com/tobgu/qframe/internal/hash"
	"github.com/tobgu/qframe/internal/index"
	"github.com/tobgu/qframe/types"
)

type Column interface {
	fmt.Stringer
	Filter(index index.Int, comparator interface{}, comparatee interface{}, bIndex index.Bool) error
	Subset(index index.Int) Column
	Equals(index index.Int, other Column, otherIndex index.Int) bool
	Comparable(reverse, equalNull, nullLast bool) Comparable
	Aggregate(indices []index.Int, fn interface{}) (Column, error)
	StringAt(i uint32, naRep string) string
	AppendByteStringAt(buf []byte, i uint32) []byte
	ByteSize() int
	Len() int

	Apply1(fn interface{}, ix index.Int) (interface{}, error)
	Apply2(fn interface{}, s2 Column, ix index.Int) (Column, error)

	Rolling(fn interface{}, ix index.Int, config rolling.Config) (Column, error)

	FunctionType() types.FunctionType
	DataType() types.DataType
}

type CompareResult byte

const (
	LessThan CompareResult = iota
	GreaterThan
	Equal

	// Used when comparing null with null
	NotEqual
)

type Comparable interface {
	Compare(i, j uint32) CompareResult

	// Write bytes to be used for hashing into buf. Using specific type for hash.Reset
	// here rather than the general Writer interface to avoid that objects escape to heap
	// due to unknown target.
	HashBytes(i uint32, buf *hash.MemHash)
}
