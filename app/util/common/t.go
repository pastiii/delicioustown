package common

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"regexp"
	"sort"
)

// 定义正则表达式模式
const (
	ExcelizeTagKey = "excelize"
	Pattern        = "title:(.*?);index:(.*?);"
)

type ExcelizeTag struct {
	Value interface{}
	Title string
	Index int
}

// ExcelizeImportData 导入数据
// ** 需要在传入的结构体中的字段加上tag：excelize:"title:列头名称;index:列下标(从0开始);"
// f 获取到的excel对象
// dst 导入目标对象【传指针】
// sheetName Sheet名称
// startRow 头行行数（从第startRow+1行开始扫）
func ExcelizeImportData(f *excelize.File, dst interface{}, sheetName string, startRow int) (err error) {
	// 获取所有行
	rows, err := f.GetRows(sheetName)
	if err != nil {
		err = errors.New(sheetName + "工作表不存在")
		return
	}

	// 取目标对象的元素类型、字段类型和 tag
	dataValue := reflect.ValueOf(dst)
	// 判断数据的类型
	if dataValue.Kind() != reflect.Ptr || dataValue.Elem().Kind() != reflect.Slice {
		err = errors.New("Invalid data type")
	}

	// 获取导入目标对象的类型信息
	dataType := dataValue.Elem().Type().Elem()

	// 遍历行，解析数据并填充到目标对象中
	for rowIndex, row := range rows {
		// 跳过头行
		if rowIndex < startRow {
			continue
		}

		// 创建新的目标对象
		newData := reflect.New(dataType).Elem()

		// 遍历目标对象的字段
		for i := 0; i < dataType.NumField(); i++ {
			var excelizeTag ExcelizeTag

			// 获取字段信息和tag
			field := dataType.Field(i)
			tag := field.Tag.Get(ExcelizeTagKey)
			// 如果tag不存在，则跳过
			//if IsEmpty(tag) {
			//	continue
			//}
			excelizeTag, err = getTag(tag)
			if err != nil {
				return
			}

			// 解析tag的值
			excelizeIndex := excelizeTag.Index
			// 防止下标越界
			if excelizeIndex >= len(row) {
				continue
			}

			// 获取单元格的值
			cellValue := row[excelizeIndex]

			// 根据字段类型设置值
			switch field.Type.Kind() {
			case reflect.Int:
				//newData.Field(i).SetInt(ToInt64(cellValue))
			case reflect.String:
				newData.Field(i).SetString(cellValue)
			}
		}

		// 将新的目标对象添加到导入目标对象的slice中
		dataValue.Elem().Set(reflect.Append(dataValue.Elem(), newData))
	}
	return
}

// ExcelizeExportData 导出excel
// ** 需要在传入的结构体中的字段加上tag：excelize:"title:列头名称;index:列下标(从0开始);"
// list 需要导出的对象数组
// sheet sheet名称
func ExcelizeExportData(list interface{}, sheet string) (file *excelize.File, err error) {
	var (
		titleStyle, style int // 列头、数据行样式
	)
	// excel构建
	file = excelize.NewFile()

	// 列头行样式
	titleStyle, err = getTitleRowStyle(file)
	if err != nil {
		return
	}

	// 数据行样式
	style, err = getDataRowStyle(file)
	if err != nil {
		return
	}

	// 构造excel表格
	err = exportBuildExcel(file, list, sheet, titleStyle, style)
	if err != nil {
		return
	}

	return
}

// 读取字段tag值
// tag 字段的tag
func getTag(tag string) (excelizeTag ExcelizeTag, err error) {
	// 编译正则表达式
	re := regexp.MustCompile(Pattern)
	// 提取匹配的子字符串
	values := re.FindStringSubmatch(tag)
	if len(values) > 1 {
		excelizeTag.Title = values[1]
		//excelizeTag.Index = ToInt(values[2])

		return
	} else {
		err = errors.New("未匹配到值")
		return
	}
	return
}

// getExcelColumnName 根据列数生成 Excel 列名
func getExcelColumnName(columnNumber int) string {
	columnName := ""
	for columnNumber > 0 {
		columnNumber--
		columnName = string(rune('A'+columnNumber%26)) + columnName
		columnNumber /= 26
	}
	return columnName
}

// 列头行样式
func getTitleRowStyle(file *excelize.File) (titleStyle int, err error) {
	return file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E6E6E6"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 16,
		},
	})
}

// 数据行样式
func getDataRowStyle(file *excelize.File) (titleStyle int, err error) {
	return file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Size: 16,
		},
	})
}

// 构造excel表格
func exportBuildExcel(file *excelize.File, data interface{}, sheet string, titleStyle, style int) (err error) {
	// 取目标对象的元素类型、字段类型和 tag
	dataValue := reflect.ValueOf(data)
	// 判断数据的类型
	if dataValue.Kind() != reflect.Slice {
		err = errors.New("invalid data type")
		return
	}

	// 获取导入目标对象的类型信息
	dataType := dataValue.Type().Elem()
	// 遍历目标对象的字段
	var exportTitle []ExcelizeTag
	for i := 0; i < dataType.NumField(); i++ {
		var excelizeTag ExcelizeTag
		// 获取字段信息和tag
		field := dataType.Field(i)
		tag := field.Tag.Get(ExcelizeTagKey)
		// 如果非导出则跳过
		//if IsEmpty(tag) {
		//	continue
		//}
		excelizeTag, err = getTag(tag)
		if err != nil {
			return
		}
		exportTitle = append(exportTitle, excelizeTag)
	}
	// 排序
	sort.Slice(exportTitle, func(i, j int) bool {
		return exportTitle[i].Index < exportTitle[j].Index
	})
	// 列头行
	var titleRowData []interface{}
	for _, colTitle := range exportTitle {
		titleRowData = append(titleRowData, colTitle.Title)
	}
	// 根据列数生成 Excel 列名
	endColName := getExcelColumnName(len(titleRowData))
	_ = file.SetRowHeight(sheet, 1, float64(25))
	_ = file.SetCellStyle(sheet, "A1", endColName+"1", titleStyle)
	_ = file.SetColWidth(sheet, "A", endColName, 25)
	if err = file.SetSheetRow(sheet, "A1", &titleRowData); err != nil {
		return
	}
	row := 1
	//实时写入数据
	for i := 0; i < dataValue.Len(); i++ {
		row++
		startCol := fmt.Sprintf("A%d", row)
		endCol := fmt.Sprintf("%s%d", endColName, row)

		item := dataValue.Index(i)
		typ := item.Type()
		num := item.NumField()

		var exportRow []ExcelizeTag
		//遍历结构体的所有字段
		for j := 0; j < num; j++ {
			//获取到struct标签，需要通过reflect.Type来获取tag标签的值
			dataField := typ.Field(j)
			tagVal := dataField.Tag.Get(ExcelizeTagKey)
			// 如果非导出则跳过
			//if IsEmpty(tagVal) {
			//	continue
			//}
			var dataCol ExcelizeTag
			dataCol, err = getTag(tagVal)
			// 取字段值
			fieldData := item.FieldByName(dataField.Name)
			dataCol.Value = fieldData
			if err != nil {
				return
			}
			exportRow = append(exportRow, dataCol)
		}

		// 排序
		sort.Slice(exportRow, func(i, j int) bool {
			return exportRow[i].Index < exportRow[j].Index
		})

		// 数据列
		var rowData []interface{}
		for _, colTitle := range exportRow {
			rowData = append(rowData, colTitle.Value)
		}

		_ = file.SetCellStyle(sheet, startCol, endCol, style)
		_ = file.SetRowHeight(sheet, row, float64(20))
		if err = file.SetSheetRow(sheet, startCol, &rowData); err != nil {
			return
		}
	}
	return
}

