/*
Copyright © 2019 NAME HERE ramonmaciasgarcia@gmail.com

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
	"log"

	"github.com/ramonmacias/gophercises/task/internal/db"
)

func main() {
	// cmd.Execute()
	db.Start()

	log.Printf("%+v", db.GetClient().Info())

	db.GetClient().Close()
}