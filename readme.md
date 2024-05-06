# Zillow scraper in Go

## Overview
This Go library provides functions to retrieve and parse property details from Zillow.

## Features
- Fetches property details from Zillow using a property URL or ID.
- Parses the retrieved HTML content to extract structured property information.
- Implemented in Go for performance and efficiency.
- Easy to integrate with existing Go projects.

## Usage

The library offers functionalities for retrieving property details using either the property URL or ID. Additionally, it supports specifying a proxy for cases where scraping restrictions might be encountered.

## Installation

```bash
go get -u github.com/johnbalvin/gozillow
```
## Ways of searching

There are 3 ways of searching, same as on the zillow page

```bash
SearchSold //for searching properties sold
SearchForRent //for searching properties to rent
SearchForSale //for searching properties for sale
```

## Examples

### Quick testing

```Go
    package main
    
    import (
        "encoding/json"
        "fmt"
        "github.com/johnbalvin/gozillow"
        "log"
        "os"
    )
    
    func main() {
        //coordinates should be 2 points one from shouth and one from north(if you think it like a square
        //this presents the two points of the diagonal from this square)
        coords := gozillow.CoordinatesInput{
            Ne: gozillow.CoordinatesValues{
                Latitude: 47.76725314073866,
                Longitud: -122.15539952490977,
            },
            Sw: gozillow.CoordinatesValues{
                Latitude: 47.67128302452179,
                Longitud: -122.3442270395582,
            },
        }
        zoomValue := 2
        //pagination is for the list that you see at the right when searching
        //you don't need to iterate over all the pages because zillow sends the whole data on mapresults at once on the first page
        //however the maximum result zillow returns is 500, so if mapResults is 500
        //try playing with the zoom or moving the coordinates, pagination won't help because you will always get at maximum 500 results
        pagination := 1
        _, fullResult, err := gozillow.SearchSold(pagination, zoomValue, coords, nil)
        if err != nil {
            log.Println(err)
            return
        }
        rawJSON, _ := json.MarshalIndent(fullResult, "", "  ")
        fmt.Printf("%s", rawJSON) //in case you don't have write permisions
        if err := os.WriteFile("./fullResult.json", rawJSON, 06444); err != nil {
            log.Println(err)
            return
        }
    }
```

```Go
    package main
    
    import (
        "encoding/json"
        "fmt"
        "github.com/johnbalvin/gozillow"
        "log"
        "os"
    )
    
    func main() {
        //coordinates should be 2 points one from shouth and one from north(if you think it like a square
        //this presents the two points of the diagonal from this square)
        coords := gozillow.CoordinatesInput{
            Ne: gozillow.CoordinatesValues{
                Latitude: 47.76725314073866,
                Longitud: -122.15539952490977,
            },
            Sw: gozillow.CoordinatesValues{
                Latitude: 47.67128302452179,
                Longitud: -122.3442270395582,
            },
        }
        zoomValue := 2
        //pagination is for the list that you see at the right when searching
        //you don't need to iterate over all the pages because zillow sends the whole data on mapresults at once on the first page
        //however the maximum result zillow returns is 500, so if mapResults is 500
        //try playing with the zoom or moving the coordinates, pagination won't help because you will always get at maximum 500 results
        pagination := 1
        _, fullResult, err := gozillow.SearchForRent(pagination, zoomValue, coords, nil)
        if err != nil {
            log.Println(err)
            return
        }
        rawJSON, _ := json.MarshalIndent(fullResult, "", "  ")
        fmt.Printf("%s", rawJSON) //in case you don't have write permisions
        if err := os.WriteFile("./fullResult.json", rawJSON, 06444); err != nil {
            log.Println(err)
            return
        }
    }
```

```Go
    package main
    
    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
        "github.com/johnbalvin/gozillow"
    )
    
    func main() {
        // Coordinates defining the search area (southwest and northeast corners)
        coords := gozillow.CoordinatesInput{
            Sw: gozillow.CoordinatesValues{
                Latitude: -1.03866277790021, // Replace with your south latitude
                Longitud: -77.53091734683608, // Replace with your south longitude
            },
            Ne: gozillow.CoordinatesValues{
                Latitude: -1.1225978433925647, // Replace with your north latitude
                Longitud: -77.59713412765507, // Replace with your north longitude
            },
        }
    
        // Zoom level (1-20) for search granularity
        zoomValue := 2
        //pagination is for the list that you see at the right when searching
        //you don't need to iterate over all the pages because zillow sends the whole data on mapresults at once on the first page
        //however the maximum result zillow returns is 500, so if mapResults is 500
        //try playing with the zoom or moving the coordinates, pagination won't help because you will always get at maximum 500 results
        pagination := 1
        // Search for properties within the specified area and zoom level
        _, fullResult, err := gozillow.SearchForSale(pagination, zoomValue, coords, nil) // Optional proxy can be passed here
        if err != nil {
            log.Println("Error:", err)
            return
        }
        rawJSON, _ := json.MarshalIndent(fullResult, "", "  ")
        fmt.Printf("%s\n", rawJSON)
        if err := os.WriteFile("./searchResultAll.json", rawJSON, 0644); err != nil {
            log.Println("Error writing file:", err)
        }
    }
```

### Basic Data Retrieval

#### Get Property Details using URL
This example demonstrates how to retrieve property details using a Zillow property URL:
```Go
    package main
    
    import (
        "encoding/json"
        "fmt"
        "github.com/johnbalvin/gozillow"
        "log"
        "os"
    )
    
    func main() {
        // Property URL (replace with your desired URL)
        propertyURL := "https://www.zillow.com/homedetails/3051-NW-207th-Ter-Miami-Gardens-FL-33056/44063708_zpid/"
    
        // Retrieve property details
        property, err := gozillow.DetailsFromPropertyURL(propertyURL, nil)
        //property, err := details.GetFromPropertyURL(propertyURL, nil)
        if err != nil {
            // Handle different error types (consider using trace.ErrorCode for more specific handling)
            fmt.Println("Error retrieving property details:", err)
            return
        }
    
        rawJSON, _ := json.MarshalIndent(property, "", "  ")
        fmt.Printf("%s", rawJSON) //in case you don't have write permisions
        if err := os.WriteFile("./details.json", rawJSON, 0644); err != nil {
            log.Println(err)
            return
        }
    }
```

#### Get Property Details using Property ID
Alternatively, you can retrieve details using the Zillow property ID (zpid) extracted from the URL:
```Go
    package main
    
    import (
        "encoding/json"
        "fmt"
        "github.com/johnbalvin/gozillow"
        "log"
        "os"
    )
    
    func main() {
        propertyID:=344690910
        property, err := gozillow.DetailsFromPropertyID(propertyID, nil)
        if err != nil {
            fmt.Println("Error retrieving property details:", err)
            return
        }
    
        rawJSON, _ := json.MarshalIndent(property, "", "  ")
        fmt.Printf("%s", rawJSON) //in case you don't have write permisions
        if err := os.WriteFile("./details.json", rawJSON, 0644); err != nil {
            log.Println(err)
            return
        }
    }
```

### Using proxy

In scenarios where Zillow might have scraping restrictions, you can utilize a proxy to anonymize your requests. Here's an example:

```Go
    package main
    
    import (
        "encoding/json"
        "fmt"
        "github.com/johnbalvin/gozillow"
        "github.com/johnbalvin/gozillow/utils"
        "log"
        "os"
    )
    
    func main() {
        proxyURL, err := utils.ParseProxy("http://[IP | domain]:[port]", "username", "password")
        if err != nil {
            log.Println("test:1 -> err: ", err)
            return
        }
    
        // Property URL (replace with your desired URL)
        propertyURL := "https://www.zillow.com/homedetails/858-Shady-Grove-Ln-Harrah-OK-73045/339897685_zpid/"
    
        // Retrieve property details
        property, err := gozillow.DetailsFromPropertyURL(propertyURL, proxyURL)
        if err != nil {
            // Handle different error types (consider using trace.ErrorCode for more specific handling)
            fmt.Println("Error retrieving property details:", err)
            return
        }
    
        rawJSON, _ := json.MarshalIndent(property, "", "  ")
        fmt.Printf("%s", rawJSON) //in case you don't have write permisions
        if err := os.WriteFile("./details.json", rawJSON, 0644); err != nil {
            log.Println(err)
            return
        }
    
    }
```

### Limitations
Currently, the library retrieves data by scraping Zillow's website. This approach might be fragile and susceptible to changes in Zillow's HTML structure.

### Contributing
We welcome contributions to improve this library! Feel free to submit pull requests for bug fixes, new features, or enhancements.

### License
This library is licensed under the MIT License. See the LICENSE file for details.