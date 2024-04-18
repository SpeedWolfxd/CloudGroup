package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// A function that returns a map of status codes
func statusHandler(urls []string) map[string]int {
	//The map that stores the status codes using url as string index
	statusCodes := make(map[string]int)

	for _, url := range urls {
		//Fetches the url
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error while checking status for %s: %v\n", url, err)
			statusCodes[url] = http.StatusServiceUnavailable // Or any other suitable status code
			errorCheck++
			continue
		}
		defer resp.Body.Close()

		//Status code being added
		statusCodes[url] = resp.StatusCode
	}

	return statusCodes
}

/*
	This function displays the specified data for task3, takes in an argument of type time

Displays all APIs with each their own status code
I've only used the main APIs that are presented, not any shortcuts or adding any argument to Countries or
two letter iso code to the language api
*/
func MyStatusHandler(w http.ResponseWriter, r *http.Request, startTime time.Time) {

	errorCheck = 0
	if r.Method == http.MethodGet {
		//List of Apis
		myApis := []string{
			COUNTRIES_URL,
			CURRENCY_URL,
			METEO_URL,
		}

		//Uses function to return the map containing the status codes
		statusCodes := statusHandler(myApis)

		//Uptime that displays how long the service has been up
		uptime := time.Since(startTime).Seconds()

		//Struct containing the results that will be displayed
		type resultStruct struct {
			Meteoapi     string  `json:"meteoapi"`
			Countriesapi string  `json:"countriesapi"`
			Currencyapi  string  `json:"currencyapi"`
			Version      string  `json:"version"`
			Uptime       float64 `json:"uptime"`
		}

		//Adding values to the struct
		result := resultStruct{
			Meteoapi:     fmt.Sprint(statusCodes[METEO_URL]),
			Countriesapi: fmt.Sprint(statusCodes[COUNTRIES_URL]),
			Currencyapi:  fmt.Sprint(statusCodes[CURRENCY_URL]),
			Version:      "v1",
			Uptime:       uptime,
		}

		//Sets header and encodes the result
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		if errorCheck == 0 {
			finalMessage := "Status of statushandling: "
			status := fmt.Sprint(http.StatusOK)
			fmt.Fprintf(w, finalMessage+status)
		}
	} else {
		fmt.Fprintf(w, "Method not set to GET: "+fmt.Sprint(http.StatusMethodNotAllowed))
	}
}
