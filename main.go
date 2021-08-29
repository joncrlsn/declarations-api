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
	action  string = "help"
	port    int    = 8080
	version string = "0.1.0"
)

type Declaration struct {
	Declaration string `json:"declaration"`
	Reference   string `json:"reference"`
}

func main() {

	var portStr = ":" + strconv.Itoa(port)

	fmt.Println(portStr)

	// Handle HTTP files in the static directory
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/api/declaration/random", randomDeclarationFunc)

	l, err := net.Listen("tcp4", portStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Fatal(http.Serve(l, nil))
	fmt.Println("Done")
}

func randomDeclarationFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decl, ref, err := RandomDeclaration("./static/declarations")
	if err != nil {
		fmt.Println("Error in RandomDeclaration", err)
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
