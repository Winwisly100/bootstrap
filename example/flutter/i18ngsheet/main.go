package main

import (
    "context"
    "time"

    "github.com/kataras/sheets"
)

func main() {
    ctx := context.TODO()
    //                            or .Token(ctx, ...)
    client := sheets.NewClient(sheets.ServiceAccount(ctx, "client_secret.json"))

    var (
        spreadsheetID := "1Ku0YXrcy8Nqmji7ABS8AmLAyxP5duQIRwmaAJAqyMYY"
        dataRange := "NamedRange or selectors like A1:E4 or *"
        records []struct{
            Timestamp time.Time
            Email     string
            Username  string
            IgnoredMe string `sheets:"-"`
        }{}
    )

    // Fill the "records" slice from a spreadsheet of one or more data range.
    err := client.ReadSpreadsheet(ctx, &records, spreadsheetID, dataRange)
    if err != nil {
        panic(err)
    }

    // Update a spreadsheet on specific range.
    updated, err := client.UpdateSpreadsheet(ctx, spreadsheetID, sheets.ValueRange{
        Range: "A2:Z",
        MajorDimension: sheets.Rows,
        Values: [][]interface{}{
            {"updated record value: 1.1", "updated record value: 1.2"},
            {"updated record value: 2.1", "updated record value: 2.2"},
        },
    })

    // Clears record values of a spreadsheet.
    cleared, err := client.ClearSpreadsheet(ctx, spreadsheetID, "A1:E5")

    // [...]
}