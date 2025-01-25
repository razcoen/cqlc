package gocqlhelpers

import (
	"github.com/gocql/gocql"
	"strings"
)

func ParseCassandraType(str string, logger gocql.StdLogger) gocql.TypeInfo {
	str = strings.TrimSpace(strings.ToLower(str))
	if strings.HasPrefix(str, "frozen<") {
		return ParseCassandraType(strings.TrimPrefix(str[:len(str)-1], "frozen<"), logger)
	} else if strings.HasPrefix(str, "set<") {
		return gocql.CollectionType{
			NativeType: gocql.NewNativeType(byte(0), gocql.TypeSet, ""),
			Elem:       ParseCassandraType(strings.TrimPrefix(str[:len(str)-1], "set<"), logger),
		}
	} else if strings.HasPrefix(str, "list<") {
		return gocql.CollectionType{
			NativeType: gocql.NewNativeType(byte(0), gocql.TypeList, ""),
			Elem:       ParseCassandraType(strings.TrimPrefix(str[:len(str)-1], "list<"), logger),
		}
	} else if strings.HasPrefix(str, "map<") {
		names := splitCompositeTypes(strings.TrimPrefix(str[:len(str)-1], "map<"))
		if len(names) != 2 {
			logger.Printf("Error parsing map type, it has %d subelements, expecting 2\n", len(names))
			return gocql.NewNativeType(byte(0), gocql.TypeCustom, "")
		}
		return gocql.CollectionType{
			NativeType: gocql.NewNativeType(byte(0), gocql.TypeMap, ""),
			Key:        ParseCassandraType(names[0], logger),
			Elem:       ParseCassandraType(names[1], logger),
		}
	} else if strings.HasPrefix(str, "tuple<") {
		names := splitCompositeTypes(strings.TrimPrefix(str[:len(str)-1], "tuple<"))
		types := make([]gocql.TypeInfo, len(names))

		for i, name := range names {
			types[i] = ParseCassandraType(name, logger)
		}

		return gocql.TupleTypeInfo{
			NativeType: gocql.NewNativeType(byte(0), gocql.TypeTuple, ""),
			Elems:      types,
		}
	} else {
		typ := getCassandraBaseType(str)
		if typ != gocql.TypeCustom {
			return gocql.NewNativeType(byte(0), typ, "")
		}
		return gocql.NewNativeType(byte(0), typ, str)
	}
}

func splitCompositeTypes(name string) []string {
	if !strings.Contains(name, "<") {
		return strings.Split(name, ",")
	}
	var parts []string
	lessCount := 0
	segment := ""
	for _, char := range name {
		if char == ',' && lessCount == 0 {
			if segment != "" {
				parts = append(parts, strings.TrimSpace(segment))
			}
			segment = ""
			continue
		}
		segment += string(char)
		if char == '<' {
			lessCount++
		} else if char == '>' {
			lessCount--
		}
	}
	if segment != "" {
		parts = append(parts, strings.TrimSpace(segment))
	}
	return parts
}

func getCassandraBaseType(name string) gocql.Type {
	switch name {
	case "ascii":
		return gocql.TypeAscii
	case "bigint":
		return gocql.TypeBigInt
	case "blob":
		return gocql.TypeBlob
	case "boolean":
		return gocql.TypeBoolean
	case "counter":
		return gocql.TypeCounter
	case "date":
		return gocql.TypeDate
	case "decimal":
		return gocql.TypeDecimal
	case "double":
		return gocql.TypeDouble
	case "duration":
		return gocql.TypeDuration
	case "float":
		return gocql.TypeFloat
	case "int":
		return gocql.TypeInt
	case "smallint":
		return gocql.TypeSmallInt
	case "tinyint":
		return gocql.TypeTinyInt
	case "time":
		return gocql.TypeTime
	case "timestamp":
		return gocql.TypeTimestamp
	case "uuid":
		return gocql.TypeUUID
	case "varchar":
		return gocql.TypeVarchar
	case "text":
		return gocql.TypeText
	case "varint":
		return gocql.TypeVarint
	case "timeuuid":
		return gocql.TypeTimeUUID
	case "inet":
		return gocql.TypeInet
	case "Mapgocql.Type":
		return gocql.TypeMap
	case "Listgocql.Type":
		return gocql.TypeList
	case "Setgocql.Type":
		return gocql.TypeSet
	case "Tuplegocql.Type":
		return gocql.TypeTuple
	default:
		return gocql.TypeCustom
	}
}
