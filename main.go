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

type Declaration struct {
	Declaration string `json:"declaration"`
	Reference   string `json:"reference"`
}

func init() {
	declarationsFile = os.Getenv("DECLARATIONS_FILE")
  if len(declarationsFile) == 0 {
    staticDir = "./static/declarations"
  }
	staticDir = os.Getenv("STATIC_PATH")
  if len(staticDir) == 0 {
    staticDir = "./static"
  }
}

func main() {
  fmt.Println("Hi Mom")
	var portStr = ":" + strconv.Itoa(port)

	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.HandleFunc("/api/declaration/random", randomDeclarationFunc)
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
	decl, ref, err := RandomDeclaration(declarationsFile)
	if err != nil {
		fmt.Println("Error in randomDeclarationFunc", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	declaration := &Declaration{
		Declaration: *decl,
		Reference:   *ref,
	}
	byteArray, err := json.Marshal(declaration)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(byteArray)
}

func healthFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	byteArray := []byte(`{"status": "OK"}`)
	w.WriteHeader(http.StatusOK)
	w.Write(byteArray)
}
