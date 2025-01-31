package gocqlhelpers

import (
	"reflect"
	"testing"

	"github.com/gocql/gocql"
)

func TestParseCassandraType_Set(t *testing.T) {
	typ := ParseCassandraType("set<text>", nil)
	set, ok := typ.(gocql.CollectionType)
	if !ok {
		t.Fatalf("expected CollectionType got %T", typ)
	} else if set.Type() != gocql.TypeSet {
		t.Fatalf("expected type %v got %v", gocql.TypeSet, set.Type())
	}

	inner, ok := set.Elem.(gocql.NativeType)
	if !ok {
		t.Fatalf("expected to get NativeType got %T", set.Elem)
	} else if inner.Type() != gocql.TypeText {
		t.Fatalf("expected to get %v got %v for set value", gocql.TypeText, set.Type())
	}
}

func TestParseCassandraType(t *testing.T) {
	tests := []struct {
		input string
		exp   gocql.TypeInfo
	}{
		{
			"set<text>", gocql.CollectionType{
				NativeType: newNativeType(gocql.TypeSet),
				Elem:       newNativeType(gocql.TypeText),
			},
		},
		{
			"map<text, varchar>", gocql.CollectionType{
				NativeType: newNativeType(gocql.TypeMap),

				Key:  newNativeType(gocql.TypeText),
				Elem: newNativeType(gocql.TypeVarchar),
			},
		},
		{
			"list<int>", gocql.CollectionType{
				NativeType: newNativeType(gocql.TypeList),
				Elem:       newNativeType(gocql.TypeInt),
			},
		},
		{
			"tuple<int, int, text>", gocql.TupleTypeInfo{
				NativeType: newNativeType(gocql.TypeTuple),

				Elems: []gocql.TypeInfo{
					newNativeType(gocql.TypeInt),
					newNativeType(gocql.TypeInt),
					newNativeType(gocql.TypeText),
				},
			},
		},
		{
			"frozen<map<text, frozen<list<frozen<tuple<int, int>>>>>>", gocql.CollectionType{
				NativeType: newNativeType(gocql.TypeMap),

				Key: newNativeType(gocql.TypeText),
				Elem: gocql.CollectionType{
					NativeType: newNativeType(gocql.TypeList),
					Elem: gocql.TupleTypeInfo{
						NativeType: newNativeType(gocql.TypeTuple),

						Elems: []gocql.TypeInfo{
							newNativeType(gocql.TypeInt),
							newNativeType(gocql.TypeInt),
						},
					},
				},
			},
		},
		{
			"frozen<tuple<frozen<tuple<text, frozen<list<frozen<tuple<int, int>>>>>>, frozen<tuple<text, frozen<list<frozen<tuple<int, int>>>>>>,  frozen<map<text, frozen<list<frozen<tuple<int, int>>>>>>>>",
			gocql.TupleTypeInfo{
				NativeType: newNativeType(gocql.TypeTuple),
				Elems: []gocql.TypeInfo{
					gocql.TupleTypeInfo{
						NativeType: newNativeType(gocql.TypeTuple),
						Elems: []gocql.TypeInfo{
							newNativeType(gocql.TypeText),
							gocql.CollectionType{
								NativeType: newNativeType(gocql.TypeList),
								Elem: gocql.TupleTypeInfo{
									NativeType: newNativeType(gocql.TypeTuple),
									Elems: []gocql.TypeInfo{
										newNativeType(gocql.TypeInt),
										newNativeType(gocql.TypeInt),
									},
								},
							},
						},
					},
					gocql.TupleTypeInfo{
						NativeType: newNativeType(gocql.TypeTuple),
						Elems: []gocql.TypeInfo{
							newNativeType(gocql.TypeText),
							gocql.CollectionType{
								NativeType: newNativeType(gocql.TypeList),
								Elem: gocql.TupleTypeInfo{
									NativeType: newNativeType(gocql.TypeTuple),
									Elems: []gocql.TypeInfo{
										newNativeType(gocql.TypeInt),
										newNativeType(gocql.TypeInt),
									},
								},
							},
						},
					},
					gocql.CollectionType{
						NativeType: newNativeType(gocql.TypeMap),
						Key:        newNativeType(gocql.TypeText),
						Elem: gocql.CollectionType{
							NativeType: newNativeType(gocql.TypeList),
							Elem: gocql.TupleTypeInfo{
								NativeType: newNativeType(gocql.TypeTuple),
								Elems: []gocql.TypeInfo{
									newNativeType(gocql.TypeInt),
									newNativeType(gocql.TypeInt),
								},
							},
						},
					},
				},
			},
		},
		{
			"frozen<tuple<frozen<tuple<int, int>>, int, frozen<tuple<int, int>>>>", gocql.TupleTypeInfo{
				NativeType: newNativeType(gocql.TypeTuple),

				Elems: []gocql.TypeInfo{
					gocql.TupleTypeInfo{
						NativeType: newNativeType(gocql.TypeTuple),

						Elems: []gocql.TypeInfo{
							newNativeType(gocql.TypeInt),
							newNativeType(gocql.TypeInt),
						},
					},
					newNativeType(gocql.TypeInt),
					gocql.TupleTypeInfo{
						NativeType: newNativeType(gocql.TypeTuple),

						Elems: []gocql.TypeInfo{
							newNativeType(gocql.TypeInt),
							newNativeType(gocql.TypeInt),
						},
					},
				},
			},
		},
		{
			"frozen<map<frozen<tuple<int, int>>, int>>", gocql.CollectionType{
				NativeType: newNativeType(gocql.TypeMap),

				Key: gocql.TupleTypeInfo{
					NativeType: newNativeType(gocql.TypeTuple),

					Elems: []gocql.TypeInfo{
						newNativeType(gocql.TypeInt),
						newNativeType(gocql.TypeInt),
					},
				},
				Elem: newNativeType(gocql.TypeInt),
			},
		},
		{
			"set<smallint>", gocql.CollectionType{
				NativeType: newNativeType(gocql.TypeSet),
				Elem:       newNativeType(gocql.TypeSmallInt),
			},
		},
		{
			"list<tinyint>", gocql.CollectionType{
				NativeType: newNativeType(gocql.TypeList),
				Elem:       newNativeType(gocql.TypeTinyInt),
			},
		},
		{"smallint", newNativeType(gocql.TypeSmallInt)},
		{"tinyint", newNativeType(gocql.TypeTinyInt)},
		{"duration", newNativeType(gocql.TypeDuration)},
		{"date", newNativeType(gocql.TypeDate)},
		{
			"list<date>", gocql.CollectionType{
				NativeType: newNativeType(gocql.TypeList),
				Elem:       newNativeType(gocql.TypeDate),
			},
		},
		{
			"set<duration>", gocql.CollectionType{
				NativeType: newNativeType(gocql.TypeSet),
				Elem:       newNativeType(gocql.TypeDuration),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := ParseCassandraType(test.input, nil)
			if !reflect.DeepEqual(got, test.exp) {
				t.Fatalf("expected %v got %v", test.exp, got)
			}
		})
	}
}
