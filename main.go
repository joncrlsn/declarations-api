//
// Copyright (c) 2021 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//

package main

//
// Runs an API server that returns a random declaration about myself from the Bible.
//

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var (
	port             int    = 8080
	version          string = "0.1.0"
	declarationsFile string = ""
	staticDir        string = ""
)

func init() {
	declarationsFile = os.Getenv("DECLARATIONS_FILE")
	if len(declarationsFile) == 0 {
		declarationsFile = "./static/declarations"
	}
	staticDir = os.Getenv("STATIC_PATH")
	if len(staticDir) == 0 {
		staticDir = "./static"
	}
}

func main() {
	var portStr = ":" + strconv.Itoa(port)

	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.HandleFunc("/api/declarations/random", randomDeclarationFunc)
	http.HandleFunc("/api/declarations/search/", declarationsSearchFunc)
	http.HandleFunc("/health", healthFunc)

	l, err := net.Listen("tcp4", portStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Listening on port %s\n", portStr)
	fmt.Printf("Running as user %d\n", os.Getuid())
	log.Fatal(http.Serve(l, nil))
	fmt.Println("Done")
}

func randomDeclarationFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	declaration, err := RandomDeclaration(declarationsFile)
	if err != nil {
		fmt.Println("Error in randomDeclarationFunc", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	byteArray, err := json.Marshal(declaration)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(byteArray)
}

func declarationsSearchFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("RequestURI: %s\n", req.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	declarationChannel, err := GrepDeclarations(declarationsFile, "abOve")
	if err != nil {
		fmt.Println("Error in declarationsFunc", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("["))

	firstDeclaration := true
	for decl := range declarationChannel {
		if !firstDeclaration {
			w.Write([]byte(","))
		} else {
			firstDeclaration = false
		}
		byteArray, err := json.Marshal(decl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(byteArray)
	}

	w.Write([]byte("]"))
}

func healthFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	byteArray := []byte(`{"status": "OK"}`)
	w.WriteHeader(http.StatusOK)
	w.Write(byteArray)
}
