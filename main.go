package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// configure the media directory name and port
	const mediaDir = "media"
	port := 8080

	// add a handler for the media files
	http.Handle("/", serveHandler(http.FileServer(http.Dir(mediaDir))))
	http.Handle("/media-data", mediaDataHandler())
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", mediaDir, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func serveHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func mediaDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "POST"{
			streamInfo := retrieveStreamInfo(r)
			fmt.Println(streamInfo)
		} else {
			fmt.Println("HI!!")
		}
	}
}

func retrieveStreamInfo(r *http.Request) string {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		return ""
	}
	videosrc := r.FormValue("videosrc")

	res, err := http.Get(videosrc)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body) //read file
	str := string(data)

	str = strings.Join(strings.Split(str, "#EXT-X-STREAM-INF")[1:], "#EXT-X-STREAM-INF") //remove lines preceding quality level info
	l := strings.Split(str, "\n")
	var lines []string
	for _, line := range l {
		if line != "\n" && line != "" { //save only non-blank lines
			lines = append(lines, line)
		}
	}

	streamInfo := ""
	for i := 0; i < len(lines); i += 2 {
		resolution, _ := GetStringInBetweenTwoString(lines[i], "RESOLUTION=", ",")
		source := lines[i+1]
		streamInfo += resolution + ": " + source + "\n"
	}
	return streamInfo
}

func GetStringInBetweenTwoString(str string, startS string, endS string) (result string,found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result,false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result,false
	}
	result = newS[:e]
	return result,true
}