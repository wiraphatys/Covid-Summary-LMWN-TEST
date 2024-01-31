package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

type COVIDCase struct {
    Age       *int    `json:"Age"`
    Province  string  `json:"Province"`
}

type COVIDData struct {
    Data []COVIDCase `json:"Data"`
}

func main() {
    router := gin.Default()
    router.GET("/covid/summary", covidSummaryHandler)
    router.Run(":8080")
}

func covidSummaryHandler(c *gin.Context) {
    resp, err := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch data"})
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read response"})
        return
    }

    var data COVIDData
    if err := json.Unmarshal(body, &data); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse data"})
        return
    }

    summary := processCovidData(data)
    c.JSON(http.StatusOK, summary)
}

func processCovidData(data COVIDData) map[string]interface{} {
    provinceCount := make(map[string]int)
    ageGroupCount := map[string]int{"0-30": 0, "31-60": 0, "61+": 0, "N/A": 0}

    for _, c := range data.Data {
        provinceCount[c.Province]++

        if c.Age == nil {
            ageGroupCount["N/A"]++
            continue
        }

        switch {
        case *c.Age <= 30:
            ageGroupCount["0-30"]++
        case *c.Age <= 60:
            ageGroupCount["31-60"]++
        default:
            ageGroupCount["61+"]++
        }
    }

    return map[string]interface{}{
        "Province":  provinceCount,
        "AgeGroup":  ageGroupCount,
    }
}
