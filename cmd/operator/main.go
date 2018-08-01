/*
Copyright 2018 Keith Clawson.
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
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/tangentspace/terrazzo/pkg/client"
	"github.com/tangentspace/terrazzo/pkg/config"
)

const (
	version = "0.0.1"
)

var (
	printVersion bool
)

func init() {
	flag.BoolVar(&printVersion, "version", false, "Show version and quit")
	flag.Parse()
}

func main() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	if printVersion {
		fmt.Println("terrazzo-operator Version:", version)
		os.Exit(0)
	}

	log.Infof("terrazzo-operator Version: %v", version)
	config := &config.Config{}
	client.Run(config)
}
