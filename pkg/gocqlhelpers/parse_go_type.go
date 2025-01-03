package gocqlhelpers

import (
	"fmt"
	"github.com/gocql/gocql"
	"gopkg.in/inf.v0"
	"math/big"
	"reflect"
	"time"
)

type GoType struct {
	Name       string
	ImportPath string
	reflect.Type
}

func ParseGoType(t gocql.TypeInfo) (*GoType, error) {
	switch t.Type() {
	case gocql.TypeVarchar, gocql.TypeAscii, gocql.TypeInet, gocql.TypeText:
		return reflectGoType(*new(string)), nil
	case gocql.TypeBigInt, gocql.TypeCounter:
		return reflectGoType(*new(int64)), nil
	case gocql.TypeTime:
		return reflectGoType(*new(time.Duration)), nil
	case gocql.TypeTimestamp:
		return reflectGoType(*new(time.Time)), nil
	case gocql.TypeBlob:
		return reflectGoType(*new([]byte)), nil
	case gocql.TypeBoolean:
		return reflectGoType(*new(bool)), nil
	case gocql.TypeFloat:
		return reflectGoType(*new(float32)), nil
	case gocql.TypeDouble:
		return reflectGoType(*new(float64)), nil
	case gocql.TypeInt:
		return reflectGoType(*new(int)), nil
	case gocql.TypeSmallInt:
		return reflectGoType(*new(int16)), nil
	case gocql.TypeTinyInt:
		return reflectGoType(*new(int8)), nil
	case gocql.TypeDecimal:
		typ := reflectGoType(*new(*inf.Dec))
		// For some reason PkgPath() is empty on the reflected type, therefore manually overriding.
		typ.ImportPath = "gopkg.in/inf.v0"
		return typ, nil
	case gocql.TypeUUID, gocql.TypeTimeUUID:
		return reflectGoType(*new(gocql.UUID)), nil
	case gocql.TypeList, gocql.TypeSet:
		elemType, err := ParseGoType(t.(gocql.CollectionType).Elem)
		if err != nil {
			return nil, err
		}
		return reflectSliceGoType(elemType), nil
	case gocql.TypeMap:
		keyType, err := ParseGoType(t.(gocql.CollectionType).Key)
		if err != nil {
			return nil, err
		}
		valueType, err := ParseGoType(t.(gocql.CollectionType).Elem)
		if err != nil {
			return nil, err
		}
		return reflectMapGoType(keyType, valueType), nil
	case gocql.TypeVarint:
		return reflectGoType(*new(*big.Int)), nil
	case gocql.TypeTuple:
		// what can we do here? all there is to do is to make a list of interface{}
		tuple := t.(gocql.TupleTypeInfo)
		return reflectGoType(make([]interface{}, len(tuple.Elems))), nil
	case gocql.TypeUDT:
		return reflectGoType(make(map[string]interface{})), nil
	case gocql.TypeDate:
		return reflectGoType(*new(time.Time)), nil
	case gocql.TypeDuration:
		return reflectGoType(*new(gocql.Duration)), nil
	default:
		return nil, fmt.Errorf("cannot create Go type for unknown CQL type %s", t)
	}
}

func reflectGoType(t any) *GoType {
	rt := reflect.TypeOf(t)
	return &GoType{
		Type:       rt,
		Name:       rt.String(),
		ImportPath: rt.PkgPath(),
	}
}

func reflectSliceGoType(t *GoType) *GoType {
	rt := reflect.SliceOf(t.Type)
	return &GoType{
		Type:       rt,
		Name:       rt.String(),
		ImportPath: rt.PkgPath(),
	}
}

func reflectMapGoType(k, v *GoType) *GoType {
	rt := reflect.MapOf(k.Type, v.Type)
	return &GoType{
		Type:       rt,
		Name:       rt.String(),
		ImportPath: rt.PkgPath(),
	}
}
