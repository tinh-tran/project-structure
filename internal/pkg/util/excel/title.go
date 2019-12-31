package excel

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"school_project/internal/pkg/util/convert"
	"school_project/internal/pkg/util/excel/library"
)

type titleRow struct {
	// map[A1]0
	dstMap map[string]int
	// map[0]A1
	srcMap map[int]string

	// sorted titles
	titles []string

	typeFieldMap map[reflect.Type]map[int][]*fieldConfig
}

//TODO remove tiếng việt không dấu
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func newRowAsMap(rd *read) (r *titleRow, err error) {
	defer func() {
		if rc := recover(); rc != nil {
			err = fmt.Errorf("%s", rc)
		}
	}()
	r = &titleRow{
		dstMap: make(map[string]int),
		srcMap: make(map[int]string),
		titles: make([]string, 0),
	}
	tempCell := &xlsxC{}
	for t, err := rd.decoder.Token(); err == nil; t, err = rd.decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			if token.Name.Local == _C {
				tempCell.R = ""
				tempCell.T = ""
				for _, a := range token.Attr {
					switch a.Name.Local {
					case _R:
						tempCell.R = a.Value
					case _T:
						tempCell.T = a.Value
					}
				}
			}
		case xml.EndElement:
			if token.Name.Local == _RowPrefix {
				// end line this one
				r.typeFieldMap = make(map[reflect.Type]map[int][]*fieldConfig)
				return r, nil
			}
		case xml.CharData:
			trimedColumnName := strings.TrimRight(tempCell.R, _AllNumber)
			columnIndex := ToDecimalism(trimedColumnName)
			var str string
			if tempCell.T == _S {
				// get string from shared
				str = rd.connecter.getSharedString(convert.MustInt(string(token)))
			} else {
				str = string(token)
			}
			t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
			strFormat, _, _ := transform.String(t, str)
			strFormat = strings.Replace(strFormat, "(", "", -1)
			strFormat = strings.Trim(strFormat, "~!@#$%^&*)_+}{|\":?></.,][';=-0987654321`\\(")
			r.dstMap[library.UpperCamelCase(strFormat)] = columnIndex
			r.srcMap[columnIndex] = library.UpperCamelCase(strFormat)
			r.titles = append(r.titles, library.UpperCamelCase(strFormat))
		}
	}

	return nil, ErrNoRow
}

// return: a copy of map[ColumnIndex][]*fieldConfig
func (tr *titleRow) MapToFields(s *schema) (rowToFiled map[int][]*fieldConfig, err error) {
	fieldsMap, ok := tr.typeFieldMap[s.Type]
	if !ok {
		fieldsMap = make(map[int][]*fieldConfig)
		for _, field := range s.Fields {
			var cloIndex int
			// Use ColumnName to find index
			if i, ok := tr.dstMap[field.ColumnName]; ok {
				cloIndex = i
			} else if field.IsRequired {
				// Use 26-number-system to find
				// cloIndex = twentysix.ToDecimalism(field.ColumnName)
				return nil, fmt.Errorf("go-excel: column name = \"%s\" is not exist", field.ColumnName)
			} else {
				// continue if is not required.
				continue
			}

			if fAry, ok := fieldsMap[cloIndex]; !ok {
				fieldsMap[cloIndex] = []*fieldConfig{field}
			} else {
				fieldsMap[cloIndex] = append(fAry, field)
			}
		}
		tr.typeFieldMap[s.Type] = fieldsMap
	}
	copyMap := make(map[int][]*fieldConfig)
	for k, v := range fieldsMap {
		copyMap[k] = v
	}
	return copyMap, nil
}

func numOfChar(c rune) int {
	a := c - rune('A')
	return int(a)
}

func charOfNum(n int) rune {
	return rune('A') + rune(n)
}

func ToDecimalism(s string) int {
	res := 0
	ary := []rune(s)

	for i, j := len(ary)-2, 26; i >= 0; i, j = i-1, j*26 {
		c := ary[i]
		// log.Printf("res(%d)=res(%d)+(Num(%d)+1)*%d, c=%s, i=%d, j=%d\n", res+(NumOfChar(c)+1)*j, res, NumOfChar(c), j, string(c), i, j)
		res = res + (numOfChar(c)+1)*j
	}

	return res + numOfChar(ary[len(ary)-1])
}
