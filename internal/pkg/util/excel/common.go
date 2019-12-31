// Package excel provide a simple and light reader to read `*.xlsx` as a relate-db-like table.
// See `ReadMe.md` or `Examples` for more usage.
package excel

// NewConnecter make a new connecter to connect to a exist xlsx file.
func NewConnecter() Connecter {
	return &connect{}
}

// UnmarshalXLSX unmarshal a sheet of XLSX file into a slice container.
// The sheet name will be inferred from element of container
// If container implement the function of GetXLSXSheetName()string, the return string will used.
// Oterwise will use the reflect struct name.
func UnmarshalXLSX(filePath string, container interface{}) error {
	conn := NewConnecter()
	err := conn.Open(filePath)
	if err != nil {
		return err
	}

	rd, err := conn.NewReader(container)
	if err != nil {
		conn.Close()
		return err
	}

	err = rd.ReadAll(container)
	if err != nil {
		conn.Close()
		rd.Close()
		return err
	}
	conn.Close()
	rd.Close()
	return nil
}

func UnmarshalXLSXCustom(filePath, sheetName string, startRow int, container interface{}) error {
	conn := NewConnecter()
	err := conn.Open(filePath)
	if err != nil {
		return err
	}

	rd, err := conn.NewReaderByConfig(&Config{
		// Sheet name as string or sheet model as object or as slice of objecg.
		Sheet: sheetName,
		// Use the index row as title, every row before title-row will be ignore, default is 0.
		TitleRowIndex: startRow,
		// Skip n row after title, default is 0 (not skip), empty row is not counted.
		Skip: 0,
		// Auto prefix to sheet name.
		Prefix: "",
		// Auto suffix to sheet name.
		Suffix: "",
	})

	if err != nil {
		conn.Close()
		return err
	}

	err = rd.ReadAll(container)
	if err != nil {
		conn.Close()
		rd.Close()
		return err
	}
	conn.Close()
	rd.Close()
	return nil
}
