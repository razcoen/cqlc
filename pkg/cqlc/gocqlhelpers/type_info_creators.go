package gocqlhelpers

import (
	"github.com/gocql/gocql"
)

func NewTypeAscii() gocql.TypeInfo     { return newNativeType(gocql.TypeAscii) }
func NewTypeBigInt() gocql.TypeInfo    { return newNativeType(gocql.TypeBigInt) }
func NewTypeBlob() gocql.TypeInfo      { return newNativeType(gocql.TypeBlob) }
func NewTypeBoolean() gocql.TypeInfo   { return newNativeType(gocql.TypeBoolean) }
func NewTypeCounter() gocql.TypeInfo   { return newNativeType(gocql.TypeCounter) }
func NewTypeDecimal() gocql.TypeInfo   { return newNativeType(gocql.TypeDecimal) }
func NewTypeDouble() gocql.TypeInfo    { return newNativeType(gocql.TypeDouble) }
func NewTypeFloat() gocql.TypeInfo     { return newNativeType(gocql.TypeFloat) }
func NewTypeInt() gocql.TypeInfo       { return newNativeType(gocql.TypeInt) }
func NewTypeText() gocql.TypeInfo      { return newNativeType(gocql.TypeText) }
func NewTypeTimestamp() gocql.TypeInfo { return newNativeType(gocql.TypeTimestamp) }
func NewTypeUUID() gocql.TypeInfo      { return newNativeType(gocql.TypeUUID) }
func NewTypeVarchar() gocql.TypeInfo   { return newNativeType(gocql.TypeVarchar) }
func NewTypeVarint() gocql.TypeInfo    { return newNativeType(gocql.TypeVarint) }
func NewTypeTimeUUID() gocql.TypeInfo  { return newNativeType(gocql.TypeTimeUUID) }
func NewTypeInet() gocql.TypeInfo      { return newNativeType(gocql.TypeInet) }
func NewTypeDate() gocql.TypeInfo      { return newNativeType(gocql.TypeDate) }
func NewTypeTime() gocql.TypeInfo      { return newNativeType(gocql.TypeTime) }
func NewTypeSmallInt() gocql.TypeInfo  { return newNativeType(gocql.TypeSmallInt) }
func NewTypeTinyInt() gocql.TypeInfo   { return newNativeType(gocql.TypeTinyInt) }
func NewTypeDuration() gocql.TypeInfo  { return newNativeType(gocql.TypeDuration) }
func NewTypeList(elemType gocql.TypeInfo) gocql.TypeInfo {
	return gocql.CollectionType{NativeType: newNativeType(gocql.TypeList), Elem: elemType}
}
func NewTypeMap(keyType, valueType gocql.TypeInfo) gocql.TypeInfo {
	return gocql.CollectionType{NativeType: newNativeType(gocql.TypeMap), Key: keyType, Elem: valueType}
}
func NewTypeSet(elemType gocql.TypeInfo) gocql.TypeInfo {
	return gocql.CollectionType{NativeType: newNativeType(gocql.TypeSet), Elem: elemType}
}
func NewTypeCustom(name string) gocql.TypeInfo {
	return gocql.NewNativeType(byte(0), gocql.TypeCustom, name)
}

// TODO: Tuple
func newNativeType(typ gocql.Type) gocql.NativeType { return gocql.NewNativeType(byte(0), typ, "") }
