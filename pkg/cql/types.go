package cql

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// References:
// Cassandra: https://cassandra.apache.org/doc/stable/cassandra/cql/types.html
// Datastax: https://docs.datastax.com/en/cql-oss/3.3/cql/cql_reference/collection_type_r.html

type DataType struct {
	NativeType      *NativeType
	CollectionType  *CollectionType
	FrozenType      *FrozenType
	UserDefinedType *UserDefinedType
}

func (dt DataType) String() string {
	switch {
	case dt.NativeType != nil:
		return dt.NativeType.String()
	case dt.CollectionType != nil:
		return dt.CollectionType.String()
	case dt.FrozenType != nil:
		return dt.FrozenType.String()
	case dt.UserDefinedType != nil:
		return ""
		// return dt.UserDefinedType.String()
	default:
		// TODO: Handle such error
		return ""
	}
}

var (
	ErrInvalidNativeType         = errors.New("invalid native type")
	ErrInvalidCollectionTypeSet  = errors.New("invalid collection type set")
	ErrInvalidFrozenType         = errors.New("invalid frozen type")
	ErrInvalidCollectionTypeList = errors.New("invalid collection type list")
	ErrInvalidCollectionTypeMap  = errors.New("invalid collection type map")
	ErrInvalidCollectionType     = errors.New("invalid collection type")
)

func ParseDataType(str string) (*DataType, error) {
	str = formatIntoLowercaseWithoutWhitespace(str)
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
	ft, err := ParseFrozenType(str)
	if err == nil {
		return &DataType{FrozenType: ft}, nil
	}
	parseErr = errors.Join(parseErr, err)
	return nil, parseErr
}

func formatIntoLowercaseWithoutWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1 // Returning -1 removes the character
		}
		return r
	}, strings.ToLower(str))
}

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

func (nt NativeType) IntoDataType() *DataType {
	return &DataType{
		NativeType: &nt,
	}
}

func (nt NativeType) IntoCollectableType() *CollectableType {
	return &CollectableType{
		NativeType: &nt,
	}
}

func ParseNativeType(nt string) (NativeType, error) {
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

func (ct CollectionType) IntoDataType() *DataType {
	return &DataType{CollectionType: &ct}
}

func (ct CollectionType) String() string {
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
	case 's':
		set, err := ParseCollectionTypeSet(collection)
		if err != nil {
			return nil, fmt.Errorf("parse collection type set: %w", err)
		}
		return &CollectionType{Set: set}, nil
	case 'l':
		list, err := ParseCollectionTypeList(collection)
		if err != nil {
			return nil, fmt.Errorf("parse collection type list: %w", err)
		}
		return &CollectionType{List: list}, nil
	case 'm':
		mp, err := ParseCollectionTypeMap(collection)
		if err != nil {
			return nil, fmt.Errorf("parse collection type map: %w", err)
		}
		return &CollectionType{Map: mp}, nil
	}
	return nil, ErrInvalidCollectionType
}

type CollectionTypeSet struct{ T *CollectableType }

func (s CollectionTypeSet) IntoDataType() *DataType {
	return &DataType{CollectionType: &CollectionType{Set: &s}}
}

func (s CollectionTypeSet) String() string { return fmt.Sprintf("set<%s>", s.T.String()) }

func ParseCollectionTypeSet(set string) (*CollectionTypeSet, error) {
	if set[0:4] != "set<" || set[len(set)-1:] != ">" {
		return nil, ErrInvalidCollectionTypeSet
	}
	nt, err := ParseCollectableType(set[4 : len(set)-1])
	if err != nil {
		return nil, fmt.Errorf("parse set native type: %w", err)
	}
	return &CollectionTypeSet{T: nt}, nil
}

type CollectionTypeList struct{ T *CollectableType }

func (l CollectionTypeList) IntoDataType() *DataType {
	return &DataType{CollectionType: &CollectionType{List: &l}}
}

func (l CollectionTypeList) String() string { return fmt.Sprintf("list<%s>", l.T.String()) }

func ParseCollectionTypeList(list string) (*CollectionTypeList, error) {
	if list[0:5] != "list<" || list[len(list)-1:] != ">" {
		return nil, ErrInvalidCollectionTypeList
	}
	nt, err := ParseCollectableType(list[5 : len(list)-1])
	if err != nil {
		return nil, fmt.Errorf("parse list native type: %w", err)
	}
	return &CollectionTypeList{T: nt}, nil
}

type CollectionTypeMap struct{ K, V *CollectableType }

func (m CollectionTypeMap) IntoDataType() *DataType {
	return &DataType{CollectionType: &CollectionType{Map: &m}}
}

func (m CollectionTypeMap) String() string {
	return fmt.Sprintf("map<%s,%s>", m.K.String(), m.V.String())
}

func ParseCollectionTypeMap(mp string) (*CollectionTypeMap, error) {
	comma := strings.Index(mp, ",")
	if mp[0:4] != "map<" || mp[len(mp)-1:] != ">" || comma == -1 {
		return nil, ErrInvalidCollectionTypeSet
	}
	k, err := ParseCollectableType(mp[4:comma])
	if err != nil {
		return nil, fmt.Errorf("parse map key native type: %w", err)
	}
	v, err := ParseCollectableType(mp[comma+1 : len(mp)-1])
	if err != nil {
		return nil, fmt.Errorf("parse map value native type: %w", err)
	}
	return &CollectionTypeMap{K: k, V: v}, nil
}

type FrozenType struct {
	DataType *DataType
}

func (ft FrozenType) IntoDataType() *DataType {
	return &DataType{FrozenType: &ft}
}

func (ft FrozenType) String() string {
	return fmt.Sprintf("frozen<%s>", ft.DataType.String())
}

func ParseFrozenType(str string) (*FrozenType, error) {
	if str[0:7] != "frozen<" || str[len(str)-1:] != ">" {
		return nil, ErrInvalidFrozenType
	}
	dt, err := ParseDataType(str[7 : len(str)-1])
	if err != nil {
		return nil, fmt.Errorf("parse frozen type: %w", err)
	}
	return &FrozenType{DataType: dt}, nil
}

type CollectableType struct {
	NativeType *NativeType
	FrozenType *FrozenType
}

func (t CollectableType) String() string {
	switch {
	case t.NativeType != nil:
		return t.NativeType.String()
	case t.FrozenType != nil:
		return t.FrozenType.String()
	default:
		// TODO: Handle such error
		return ""
	}
}

func ParseCollectableType(str string) (*CollectableType, error) {
	var parseErr error
	nt, err := ParseNativeType(str)
	if err == nil {
		return &CollectableType{NativeType: &nt}, nil
	}
	parseErr = errors.Join(parseErr, err)
	ft, err := ParseFrozenType(str)
	if err == nil {
		return &CollectableType{FrozenType: ft}, nil
	}
	parseErr = errors.Join(parseErr, err)
	return nil, parseErr
}

type UserDefinedType struct {
	Name   string
	Fields []*UserDefinedTypeField
}

func (t UserDefinedType) IntoDataType() *DataType {
	return &DataType{
		UserDefinedType: &t,
	}
}

type UserDefinedTypeField struct {
	Name     string
	DataType *DataType
}
