package main

import "fmt"

// Should be an array
//
type Spreadsheet struct{
	Properties SpreadsheetProperties
	Sheets []Sheet
}

type SpreadsheetProperties struct{
	Title string
}

type Sheet struct {
	Properties SheetProperties
	Data []GridData
}


type SheetProperties struct{
	Title string
	GridProperties GridProperties
}

type GridProperties struct{
	RowCount int
	ColumnCount int
}



type Data struct{
	startRow int
	startColumn int
	rowData RowData
}

type RowData struct{
	values CellData
}

type CellData struct{
	UserEnteredDatea ExtendedValue
	EffectiveData ExtendedValue
}

type ExtendedValue struct{
	  // Union field value can be only one of the following:
  "numberValue": number,
  "stringValue": string,
  "boolValue": boolean,
  "formulaValue": string,
}

