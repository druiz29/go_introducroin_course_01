package main

import (
	"github.com/xuri/excelize/v2"
)

func readGameIDsFromExcel(filePath string) (map[string]bool, error) {
    f, err := excelize.OpenFile(filePath)
    if err != nil {
        return nil, err
    }

    ids := make(map[string]bool)
    rows, err := f.GetRows("Sheet1")
    if err != nil {
        return nil, err
    }

    for _, row := range rows {
        if len(row) > 0 {
            ids[row[0]] = true
        }
    }

    return ids, nil
}
