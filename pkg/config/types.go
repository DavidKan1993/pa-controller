/*
Copyright © 2018 inwinSTACK Inc

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

package config

// Config contains the operator config
type Config struct {
	Threads        int
	SyncSec        int
	Retry          int
	Host           string
	Username       string
	Password       string
	APIKey         string
	MoveType       int
	MoveRule       string
	Vsys           string
	CommitWaitTime int
	Admins         []string
	DaNPartial     bool
	PaOPartial     bool
	Force          bool
	Sync           bool
}
