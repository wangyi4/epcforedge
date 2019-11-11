// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net/http"
	"time"
)

// Connectivity constants
const (
	NEFServerPort = "80"
)

func main() {

	unusedlint()
	NEFRouter := NewNEFRouter()
	s := &http.Server{
		Addr:           ":" + NEFServerPort,
		Handler:        NEFRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("NEF listening on", s.Addr)
	log.Fatal(s.ListenAndServe())
}

func unusedlint() {
	/* For unused variables lint warning to be  removed later */
	ti := TrafficInfluSub{}
	_ = ti

	tis := TrafficInfluSubPatch{}
	_ = tis

	ac := AppSessionContext{}
	_ = ac

	tid := TrafficInfluData{}
	_ = tid

	tids := TrafficInfluDataPatch{}
	_ = tids

	upc := UpPathChange
	_ = upc

	pd := ProblemDetails{}
	_ = pd

	var ipx Ipv6Prefix = " "
	_ = ipx

	var pt Port = 8080
	_ = pt

}
