package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Artists   []Artists
	Locations Locations
	Dates     Dates
	Relation  Relation
}

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
	// Next         int
	// Previous     int
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID             int `json:"id"`
		DatesLocations string
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	hometmpl := template.Must(template.ParseFiles("./home.html"))

	// homeID, _ := strconv.Atoi(r.FormValue("home"))

	// type data struct {
	// 	Artists interface{}
	// 	// LocDate  interface{}
	// 	// Relation interface{}
	// }

	// creates variable to access API infotmation via the ID number
	homepage := Response{Artists: ArtistAPI(), Locations: LocationsAPI()[1]}

	if r.URL.Path != "/home" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	hometmpl.Execute(w, homepage)
}

func main() {
	// fileserver := http.FileServer(http.Dir("./home"))

	// http.Handle("/home/", http.StripPrefix("/home/", fileserver))
	// fmt.Println(len(responseObject.Artists))
	// for _, p := range artists {
	// 	fmt.Println("Members:", p.Members[0])
	// }
	// for i := 0; i < len(responseObject.Artists); i++ {
	//     fmt.Println(responseObject.Artists[i].ID.Name)
	// }
	// http.HandleFunc("/", Homepage)

	http.HandleFunc("/home", HomePage)
	// http.HandleFunc("/locations", locationsHandler)
	// http.HandleFunc("/dates", datesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Starting server on port", "8080")
}

func ArtistAPI() []Artists {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// var response Response
	// json.Unmarshal(body, &Response)
	var artists []Artists
	json.Unmarshal([]byte(responseData), &artists)
	// for i := range artists {

	// 	if i == 51 {
	// 		artists[51].Next = 1
	// 	} else {
	// 		artists[i].Next = i + 2
	// 	}
	// }
	// for i := range artists {

	// 	if i == 0 {
	// 		artists[0].Previous = 52
	// 	} else {
	// 		artists[i].Previous = i
	// 	}
	// }
	return artists
}

func LocationsAPI() []Locations {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// var response Response
	// json.Unmarshal(body, &Response)
	var locations []Locations
	json.Unmarshal([]byte(responseData), &locations)
	// for i := range artists {

	// 	if i == 51 {
	// 		artists[51].Next = 1
	// 	} else {
	// 		artists[i].Next = i + 2
	// 	}
	// }
	// for i := range artists {

	// 	if i == 0 {
	// 		artists[0].Previous = 52
	// 	} else {
	// 		artists[i].Previous = i
	// 	}
	// }
	return locations
}

// func groupieTrackerHandler(w http.ResponseWriter, r *http.Request) {
// 	// p := groupieTrackerPage{Title: "Groupie Tracker"}
// 	t, _ := template.ParseFiles("artists.html")

// 	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
// 	if err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(1)
// 	}
// 	defer response.Body.Close()

// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// var response Response
// 	// json.Unmarshal(body, &Response)
// 	var artists []Artists
// 	json.Unmarshal([]byte(responseData), &artists)
// 	for i := range artists {

// 		if i == 51 {
// 			artists[51].Next = 1
// 		} else {
// 			artists[i].Next = i + 2
// 		}
// 	}
// 	for i := range artists {

// 		if i == 0 {
// 			artists[0].Previous = 52
// 		} else {
// 			artists[i].Previous = i
// 		}
// 	}
// 	t.Execute(w, artists)
// }

// func locationsHandler(w http.ResponseWriter, r *http.Request) {
// 	// p := groupieTrackerPage{Title: "Groupie Tracker"}
// 	t, _ := template.ParseFiles("locations.html")

// 	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
// 	if err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(1)
// 	}
// 	defer response.Body.Close()

// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	var locations Locations
// 	json.Unmarshal([]byte(responseData), &locations)
// 	// fmt.Println(locations)
// 	fmt.Print(locations.Index)

// 	t.Execute(w, locations)
// 	// t.Execute(w, locations.Index[1].ID)
// }

// func datesHandler(w http.ResponseWriter, r *http.Request) {
// 	// p := groupieTrackerPage{Title: "Groupie Tracker"}
// 	t, _ := template.ParseFiles("dates.html")

// 	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
// 	if err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(1)
// 	}
// 	defer response.Body.Close()

// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}

// 	// var response Response
// 	// json.Unmarshal(body, &Response)
// 	var dates Dates
// 	json.Unmarshal([]byte(responseData), &dates)
// 	// fmt.Println(locations)
// 	fmt.Print(dates.Index)

// 	t.Execute(w, dates)
// 	// t.Execute(w, locations.Index[1].ID)
// }

// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	input := r.FormValue("input")
// 	banner := r.FormValue("banner")
// 	path := "banners/" + banner + ".txt"
// 	// if input != "" && banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
// 	// 	fmt.Fprintf(w, "%s\n", newline(input, path))
// 	// }
// 	if input == "" {
// 		http.Error(w, "400 Bad Request Error.", http.StatusBadRequest)
// 	} else if !(banner == "standard" || banner == "shadow" || banner == "thinkertoy") {
// 		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
// 	} else {
// 		fmt.Fprintf(w, "%s\n", newline(input, path))
// 	}
// }
