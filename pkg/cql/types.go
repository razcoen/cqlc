package cql

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Reference: https://cassandra.apache.org/doc/stable/cassandra/cql/types.html

type DataType struct {
	NativeType     *NativeType
	CollectionType *CollectionType
}

var (
	ErrInvalidNativeType         = errors.New("invalid native type")
	ErrInvalidCollectionTypeSet  = errors.New("invalid collection type set")
	ErrInvalidCollectionTypeList = errors.New("invalid collection type list")
	ErrInvalidCollectionTypeMap  = errors.New("invalid collection type map")
	ErrInvalidCollectionType     = errors.New("invalid collection type")
)

func ParseDataType(str string) (*DataType, error) {
	str = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1 // Returning -1 removes the character
		}
		return r
	}, str)
	var parseErr error
	nt, err := ParseNativeType(str)
	if err == nil {
		return &DataType{NativeType: &nt}, nil
	}
	parseErr = errors.Join(parseErr, err)
	ct, err := ParseCollectionType(str)
	if err == nil {
		return &DataType{CollectionType: ct}, nil
	}
	parseErr = errors.Join(parseErr, err)
	return nil, parseErr
}

func NewNativeTypeAscii() *DataType     { return &DataType{NativeType: ntptr(NativeTypeAscii)} }
func NewNativeTypeBigint() *DataType    { return &DataType{NativeType: ntptr(NativeTypeBigint)} }
func NewNativeTypeBlob() *DataType      { return &DataType{NativeType: ntptr(NativeTypeBlob)} }
func NewNativeTypeBoolean() *DataType   { return &DataType{NativeType: ntptr(NativeTypeBoolean)} }
func NewNativeTypeCounter() *DataType   { return &DataType{NativeType: ntptr(NativeTypeCounter)} }
func NewNativeTypeDate() *DataType      { return &DataType{NativeType: ntptr(NativeTypeDate)} }
func NewNativeTypeDecimal() *DataType   { return &DataType{NativeType: ntptr(NativeTypeDecimal)} }
func NewNativeTypeDouble() *DataType    { return &DataType{NativeType: ntptr(NativeTypeDouble)} }
func NewNativeTypeDuration() *DataType  { return &DataType{NativeType: ntptr(NativeTypeDuration)} }
func NewNativeTypeFloat() *DataType     { return &DataType{NativeType: ntptr(NativeTypeFloat)} }
func NewNativeTypeInet() *DataType      { return &DataType{NativeType: ntptr(NativeTypeInet)} }
func NewNativeTypeInt() *DataType       { return &DataType{NativeType: ntptr(NativeTypeInt)} }
func NewNativeTypeSmallint() *DataType  { return &DataType{NativeType: ntptr(NativeTypeSmallint)} }
func NewNativeTypeText() *DataType      { return &DataType{NativeType: ntptr(NativeTypeText)} }
func NewNativeTypeTime() *DataType      { return &DataType{NativeType: ntptr(NativeTypeTime)} }
func NewNativeTypeTimestamp() *DataType { return &DataType{NativeType: ntptr(NativeTypeTimestamp)} }
func NewNativeTypeTimeuuid() *DataType  { return &DataType{NativeType: ntptr(NativeTypeTimeuuid)} }
func NewNativeTypeTinyint() *DataType   { return &DataType{NativeType: ntptr(NativeTypeTinyint)} }
func NewNativeTypeUUID() *DataType      { return &DataType{NativeType: ntptr(NativeTypeUUID)} }
func NewNativeTypeVarchar() *DataType   { return &DataType{NativeType: ntptr(NativeTypeVarchar)} }
func NewNativeTypeVarint() *DataType    { return &DataType{NativeType: ntptr(NativeTypeVarint)} }
func NewCollectionTypeSet(t NativeType) *DataType {
	return &DataType{CollectionType: &CollectionType{Set: &CollectionTypeSet{T: t}}}
}
func NewCollectionTypeList(t NativeType) *DataType {
	return &DataType{CollectionType: &CollectionType{List: &CollectionTypeList{T: t}}}
}
func NewCollectionTypeMap(k, v NativeType) *DataType {
	return &DataType{CollectionType: &CollectionType{Map: &CollectionTypeMap{K: k, V: v}}}
}

func ntptr(nt NativeType) *NativeType { return &nt }

type NativeType string

const (
	NativeTypeAscii     NativeType = "ascii"     // ASCII character string
	NativeTypeBigint    NativeType = "bigint"    // 64-bit signed long
	NativeTypeBlob      NativeType = "blob"      // Arbitrary bytes (no validation)
	NativeTypeBoolean   NativeType = "boolean"   // Either true or false
	NativeTypeCounter   NativeType = "counter"   // Counter column (64-bit signed value). See counters for details.
	NativeTypeDate      NativeType = "date"      // A date (with no corresponding time value). See dates below for details.
	NativeTypeDecimal   NativeType = "decimal"   // Variable-precision decimal
	NativeTypeDouble    NativeType = "double"    // 64-bit IEEE-754 floating point
	NativeTypeDuration  NativeType = "duration"  // A duration with nanosecond precision. See durations below for details.
	NativeTypeFloat     NativeType = "float"     // 32-bit IEEE-754 floating point
	NativeTypeInet      NativeType = "inet"      // An IP address, either IPv4 (4 bytes long) or IPv6 (16 bytes long). Note that there is no inet constant, IP address should be input as strings.
	NativeTypeInt       NativeType = "int"       // 32-bit signed int
	NativeTypeSmallint  NativeType = "smallint"  // 16-bit signed int
	NativeTypeText      NativeType = "text"      // UTF8 encoded string
	NativeTypeTime      NativeType = "time"      // A time (with no corresponding date value) with nanosecond precision. See times below for details.
	NativeTypeTimestamp NativeType = "timestamp" // A timestamp (date and time) with millisecond precision. See timestamps below for details.
	NativeTypeTimeuuid  NativeType = "timeuuid"  // Version 1 UUID, generally used as a “conflict-free” timestamp. Also see timeuuid-functions.
	NativeTypeTinyint   NativeType = "tinyint"   // 8-bit signed int
	NativeTypeUUID      NativeType = "uuid"      // A UUID (of any version)
	NativeTypeVarchar   NativeType = "varchar"   // UTF8 encoded string
	NativeTypeVarint    NativeType = "varint"    // Arbitrary-precision integer
)

func (nt NativeType) String() string {
	return string(nt)
}

func ParseNativeType(nt string) (NativeType, error) {
	nt = strings.ToLower(nt)
	nativeTypeSet := map[NativeType]bool{
		NativeTypeAscii:     true,
		NativeTypeBigint:    true,
		NativeTypeBlob:      true,
		NativeTypeBoolean:   true,
		NativeTypeCounter:   true,
		NativeTypeDate:      true,
		NativeTypeDecimal:   true,
		NativeTypeDouble:    true,
		NativeTypeDuration:  true,
		NativeTypeFloat:     true,
		NativeTypeInet:      true,
		NativeTypeInt:       true,
		NativeTypeSmallint:  true,
		NativeTypeText:      true,
		NativeTypeTime:      true,
		NativeTypeTimestamp: true,
		NativeTypeTimeuuid:  true,
		NativeTypeTinyint:   true,
		NativeTypeUUID:      true,
		NativeTypeVarchar:   true,
		NativeTypeVarint:    true,
	}
	parsed := NativeType(nt)
	if ok := nativeTypeSet[parsed]; !ok {
		return NativeType(""), ErrInvalidNativeType
	}
	return parsed, nil
}

type CollectionType struct {
	Set  *CollectionTypeSet
	List *CollectionTypeList
	Map  *CollectionTypeMap
}

func (ct *CollectionType) String() string {
	switch {
	case ct.Set != nil:
		return ct.Set.String()
	case ct.List != nil:
		return ct.List.String()
	case ct.Map != nil:
		return ct.Map.String()
	default:
		// TODO: Handle such error
		return ""
	}
}

func ParseCollectionType(collection string) (*CollectionType, error) {
	switch collection[0] {
	case 'S':
		set, err := ParseCollectionTypeSet(collection)
		if err != nil {
			return nil, fmt.Errorf("parse collection type set: %w", err)
		}
		return &CollectionType{Set: set}, nil
	case 'L':
		list, err := ParseCollectionTypeList(collection)
		if err != nil {
			return nil, fmt.Errorf("parse collection type list: %w", err)
		}
		return &CollectionType{List: list}, nil
	case 'M':
		mp, err := ParseCollectionTypeMap(collection)
		if err != nil {
			return nil, fmt.Errorf("parse collection type map: %w", err)
		}
		return &CollectionType{Map: mp}, nil
	}
	return nil, ErrInvalidCollectionType
}

type CollectionTypeSet struct{ T NativeType }

func (s CollectionTypeSet) String() string { return fmt.Sprintf("SET<%s>", s.T.String()) }

func ParseCollectionTypeSet(set string) (*CollectionTypeSet, error) {
	if set[0:4] != "SET<" || set[len(set)-1:] != ">" {
		return nil, ErrInvalidCollectionTypeSet
	}
	nt, err := ParseNativeType(set[4 : len(set)-1])
	if err != nil {
		return nil, fmt.Errorf("parse set native type: %w", err)
	}
	return &CollectionTypeSet{T: nt}, nil
}

type CollectionTypeList struct{ T NativeType }

func (l CollectionTypeList) String() string { return fmt.Sprintf("LIST<%s>", l.T.String()) }

func ParseCollectionTypeList(list string) (*CollectionTypeList, error) {
	if list[0:5] != "LIST<" || list[len(list)-1:] != ">" {
		return nil, ErrInvalidCollectionTypeList
	}
	nt, err := ParseNativeType(list[5 : len(list)-1])
	if err != nil {
		return nil, fmt.Errorf("parse list native type: %w", err)
	}
	return &CollectionTypeList{T: nt}, nil
}

type CollectionTypeMap struct{ K, V NativeType }

func (m CollectionTypeMap) String() string {
	return fmt.Sprintf("MAP<%s,%s>", m.K.String(), m.V.String())
}

func ParseCollectionTypeMap(mp string) (*CollectionTypeMap, error) {
	comma := strings.Index(mp, ",")
	if mp[0:4] != "MAP<" || mp[len(mp)-1:] != ">" || comma == -1 {
		return nil, ErrInvalidCollectionTypeSet
	}
	k, err := ParseNativeType(mp[4:comma])
	if err != nil {
		return nil, fmt.Errorf("parse map key native type: %w", err)
	}
	v, err := ParseNativeType(mp[comma+1 : len(mp)-1])
	if err != nil {
		return nil, fmt.Errorf("parse map value native type: %w", err)
	}
	return &CollectionTypeMap{K: k, V: v}, nil
}
