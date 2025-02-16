package common

import (
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

type Data struct {
	Id      int
	OrderNo string
	State   int
	Account string
	Amount  float64
	Time    time.Time
}

type SpecialWidth struct {
	Col   string
	Width float64
}

// 默认宽度
const defaultWidth = 15

func InitData() []*Data {
	payTime, _ := time.Parse("2006-01-02 15:04:05", "2023-08-04 12:00:00")
	return []*Data{
		{1, "OrderNo-20230803000000", 1, "xiaoming", 12.00, payTime},
		{2, "OrderNo-20230803000001", 2, "xiaoming", 0.01, payTime},
		{3, "OrderNo-20230803000002", 1, "xiaohong", 12.00, payTime},
		{4, "OrderNo-20230803000003", 1, "xiaohong", 15.00, payTime},
	}
}

func InitTitle() []string {
	return []string{"ID", "订单号", "状态", "用户", "金额", "时间"}
}

func InitSpecialWidth() []*SpecialWidth {
	return []*SpecialWidth{
		{"B", 30},
		{"F", 20},
	}
}

// 获取表头样式
func GetTitleStyle(file *excelize.File) (titleStyle int, err error) {
	return file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "FF99CC", Style: 2},
			{Type: "top", Color: "FF99CC", Style: 2},
			{Type: "bottom", Color: "FF99CC", Style: 2},
			{Type: "right", Color: "FF99CC", Style: 2},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFCC"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 16,
		},
	})
}

// 获取数据样式
func GetDataStyle(file *excelize.File) (titleStyle int, err error) {
	return file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "FFFFFF", Style: 9},
			{Type: "top", Color: "FFFFFF", Style: 9},
			{Type: "bottom", Color: "FFFFFF", Style: 9},
			{Type: "right", Color: "FFFFFF", Style: 9},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#CCCCCC"},
			Pattern: 3,
		},
		Font: &excelize.Font{
			Size: 13,
		},
	})
}

// 设置表头及表头样式
func SetTitleWithStyle(eFile *excelize.File, sheetName string, title []string, titleLen int) error {
	style, err := GetTitleStyle(eFile)
	if err != nil {
		return err
	}

	hCell, _ := excelize.CoordinatesToCellName(1, 1)
	vCell, _ := excelize.CoordinatesToCellName(titleLen, 1)
	err = eFile.SetCellStyle(sheetName, hCell, vCell, style)
	if err != nil {
		return err
	}

	for key, val := range title {
		col := key + 1
		cellName, _ := excelize.CoordinatesToCellName(col, 1)
		err = eFile.SetCellValue(sheetName, cellName, val)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

// 设置数据样式
func SetListStyle(eFile *excelize.File, sheetName string, dataLen int, titleLen int) error {
	style, err := GetDataStyle(eFile)
	if err != nil {
		return err
	}

	hCell, _ := excelize.CoordinatesToCellName(1, 2)
	vCell, _ := excelize.CoordinatesToCellName(titleLen, dataLen+1)
	err = eFile.SetCellStyle(sheetName, hCell, vCell, style)
	if err != nil {
		return err
	}

	return nil
}

// 设置列宽
func SetWidth(eFile *excelize.File, sheetName string, cellNum int, specialWidth []*SpecialWidth) {
	startCol, _ := excelize.ColumnNumberToName(1)
	endCol, _ := excelize.ColumnNumberToName(cellNum)
	_ = eFile.SetColWidth(sheetName, startCol, endCol, defaultWidth)
	if len(specialWidth) > 0 {
		for _, val := range specialWidth {
			_ = eFile.SetColWidth(sheetName, val.Col, val.Col, val.Width)
		}
	}
}

// 数据写入
func SetData(eFile *excelize.File, list []*Data, sheetName string)  {
	for key, val := range list {
		row := key + 2
		cell := "A" + strconv.Itoa(row)
		cells := []interface{}{val.Id, val.OrderNo, val.State, val.Account, val.Amount, val.Time}
		err := eFile.SetSheetRow(sheetName, cell, &cells)
		if err != nil {
			panic(err)
		}
	}
}

func Export(title []string, list []*Data, sheetName string, specialWidth []*SpecialWidth) error {
	dataLen := len(list)
	titleLen := len(title)
	eFile := excelize.NewFile()
	sheet, err := eFile.NewSheet(sheetName)
	if err != nil {
		return err
	}

	eFile.SetActiveSheet(sheet)
	SetWidth(eFile, sheetName, titleLen, specialWidth)
	err = SetTitleWithStyle(eFile, sheetName, title, titleLen)
	if err != nil {
		return err
	}

	err = SetListStyle(eFile, sheetName, dataLen, titleLen)
	if err != nil {
		return err
	}

	SetData(eFile, list, sheetName)
	//if sheetName != "Sheet1" {
	//	_ = eFile.DeleteSheet("Sheet1")
	//}

	err = eFile.SaveAs("test.xlsx")
	if err != nil {
		return err
	}

	return nil
}
