package excel

import (
	"encoding/json"
	"fmt"
	"reflect"

	"school_project/internal/pkg/util/convert"
)

type Temp struct {
	Foo string
}

type Standard struct {
	// use field name as default column name
	ID int
	// column means to map the column name
	Name string `xlsx:"column(NameOf)"`
	// you can map a column into more than one field
	NamePtr *string `xlsx:"column(NameOf)"`
	// omit `column` if only want to map to column name, it's equal to `column(AgeOf)`
	Age int `xlsx:"AgeOf"`
	// split means to split the string into slice by the `|`
	Slice []int `xlsx:"split(|)"`
	// *Temp implement the `encoding.BinaryUnmarshaler`
	Temp *Temp `xlsx:"column(UnmarshalString)"`
	// use '-' to ignore.
	WantIgnored string `xlsx:"-"`
}

func strPtr(s string) *string {
	return &s
}
func (this *Temp) UnmarshalBinary(d []byte) error {
	return json.Unmarshal(d, this)
}

var expectStandardList = []Standard{
	{
		ID:      1,
		Name:    "Andy",
		NamePtr: strPtr("Andy"),
		Age:     1,
		Slice:   []int{1, 2},
		Temp: &Temp{
			Foo: "Andy",
		},
	},
	{
		ID:      2,
		Name:    "Leo",
		NamePtr: strPtr("Leo"),
		Age:     2,
		Slice:   []int{2, 3, 4},
		Temp: &Temp{
			Foo: "Leo",
		},
	},
	{
		ID:      3,
		Name:    "Ben",
		NamePtr: strPtr("Ben"),
		Age:     3,
		Slice:   []int{3, 4, 5, 6},
		Temp: &Temp{
			Foo: "Ben",
		},
	},
	{
		ID:      4,
		Name:    "Ming",
		NamePtr: strPtr("Ming"),
		Age:     4,
		Slice:   []int{1},
		Temp: &Temp{
			Foo: "Ming",
		},
	},
}

func ExampleUnmarshalXLSX_struct() {
	var stdList []Standard
	err := UnmarshalXLSX("simple.xlsx", &stdList)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !reflect.DeepEqual(stdList, expectStandardList) {
		fmt.Printf("unexprect std list: %s", convert.MustJsonPrettyString(stdList))
	}

	fmt.Printf(convert.MustJsonString(stdList))

	// output:
	// [{"ID":1,"Name":"Andy","NamePtr":"Andy","Age":1,"Slice":[1,2],"Temp":{"Foo":"Andy"},"WantIgnored":""},{"ID":2,"Name":"Leo","NamePtr":"Leo","Age":2,"Slice":[2,3,4],"Temp":{"Foo":"Leo"},"WantIgnored":""},{"ID":3,"Name":"Ben","NamePtr":"Ben","Age":3,"Slice":[3,4,5,6],"Temp":{"Foo":"Ben"},"WantIgnored":""},{"ID":4,"Name":"Ming","NamePtr":"Ming","Age":4,"Slice":[1],"Temp":{"Foo":"Ming"},"WantIgnored":""}]
}

func ExampleUnmarshalXLSXCustom_struct() {
	var stdList []Standard
	err := UnmarshalXLSXCustom("simple.xlsx", "Sheet1", 0, &stdList)
	//err := UnmarshalXLSX("simple.xlsx", &stdList)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !reflect.DeepEqual(stdList, expectStandardList) {
		fmt.Printf("unexprect std list: %s", convert.MustJsonPrettyString(stdList))
	}

	fmt.Printf(convert.MustJsonString(stdList))
	// output:
	// [{"ID":1,"Name":"Andy","NamePtr":"Andy","Age":1,"Slice":[1,2],"Temp":{"Foo":"Andy"},"WantIgnored":""},{"ID":2,"Name":"Leo","NamePtr":"Leo","Age":2,"Slice":[2,3,4],"Temp":{"Foo":"Leo"},"WantIgnored":""},{"ID":3,"Name":"Ben","NamePtr":"Ben","Age":3,"Slice":[3,4,5,6],"Temp":{"Foo":"Ben"},"WantIgnored":""},{"ID":4,"Name":"Ming","NamePtr":"Ming","Age":4,"Slice":[1],"Temp":{"Foo":"Ming"},"WantIgnored":""}]

}
