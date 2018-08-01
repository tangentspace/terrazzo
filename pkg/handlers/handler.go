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

package handlers

import (
	log "github.com/sirupsen/logrus"
	"github.com/tangentspace/terrazzo/pkg/config"
)

type Handler interface {
	Init(c *config.Config) error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(oldObj, newObj interface{})
}

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init(c *config.Config) error {
	return nil
}

func (d *Default) ObjectCreated(obj interface{}) {
	log.Info("Executing create handler")
}

func (d *Default) ObjectDeleted(obj interface{}) {
	log.Info("Executing delete handler")
}

func (d *Default) ObjectUpdated(oldObj, newObj interface{}) {
	log.Info("Executing update handler")
}
