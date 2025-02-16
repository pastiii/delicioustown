package test

import (
	"DeliciousTown/app/util/common"
	"fmt"
	"testing"
)

const defaultSheetName = "Sheet"

func TestExport(t *testing.T) {
	var specialWidth []*common.SpecialWidth
	specialWidth = common.InitSpecialWidth()
	list := common.InitData()
	title := common.InitTitle()
	res := common.Export(title, list, defaultSheetName, specialWidth)
	if res != nil {
		fmt.Print(res.Error())
	}
}
