package main

import (
    "os"
    "fmt"
    "github.com/Luxurioust/excelize"
)


func main1() {
    xlsx := excelize.CreateFile()
    // Create a new sheet.
    xlsx.NewSheet(2, "Sheet2")
    // Set value of a cell.
    xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
    xlsx.SetCellValue("Sheet1", "B2", 100)
    // Set active sheet of the workbook.
    xlsx.SetActiveSheet(2)
    // Save xlsx file by the given path.
    err := xlsx.WriteTo("Workbook.xlsx")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


func main() {
    xlsx, err := excelize.OpenFile("Workbook.xlsx")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // Get value from cell by given sheet index and axis.
    cell := xlsx.GetCellValue("Sheet1", "B2")
    fmt.Println(cell)
    // Get all the rows in a sheet.
    rows := xlsx.GetRows("Sheet2")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}
