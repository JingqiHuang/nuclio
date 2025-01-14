/*
Copyright 2023 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	"github.com/nuclio/nuclio/cmd/nuctl/app"

	"github.com/nuclio/errors"
	"github.com/nuclio/nuclio-sdk-go"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:5054", nil))
	}()
	if err := app.Run(); err != nil {
		if errWithCode, ok := err.(*nuclio.ErrorWithStatusCode); ok && errWithCode != nil {
			os.Stdout.WriteString(errWithCode.Error())
		} else {
			errors.PrintErrorStack(os.Stderr, err, 5)
		}
		os.Exit(1)
	}

	os.Exit(0)
}
