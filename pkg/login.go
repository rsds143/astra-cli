//  Copyright 2021 Ryan Svihla
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Package pkg is the top level package for shared libraries
package pkg

import (
	"fmt"
	"os"

	"github.com/rsds143/astra-cli/pkg/env"
	"github.com/rsds143/astra-devops-sdk-go/astraops"
)

// LoginService provides interface to implement logins and produce an Client
type LoginService interface {
	Login() (Client, error)
}

// Client is the abstraction for client interactions. Allows alternative db management clients
type Client interface {
	CreateDb(astraops.CreateDb) (astraops.Database, error)
	Terminate(string, bool) error
	FindDb(string) (astraops.Database, error)
	ListDb(string, string, string, int32) ([]astraops.Database, error)
	Park(string) error
	Unpark(string) error
	Resize(string, int32) error
	GetSecureBundle(string) (astraops.SecureBundle, error)
	GetTierInfo() ([]astraops.TierInfo, error)
}

// Creds knows how handle and store credentials
type Creds struct {
	GetHomeFunc func() (string, error) // optional. If not specified os.UserHomeDir is used for log base directory to find creds
}

// Login logs into the Astra DevOps API using the local configuration provided by the 'astra-cli login' command
func (c *Creds) Login() (Client, error) {
	getHome := c.GetHomeFunc
	if getHome == nil {
		getHome = os.UserHomeDir
	}
	confDir, confFile, err := GetHome(getHome)
	if err != nil {
		return &astraops.AuthenticatedClient{}, fmt.Errorf("unable to read conf dir with error '%v'", err)
	}
	hasToken, err := confFile.HasToken()
	if err != nil {
		return &astraops.AuthenticatedClient{}, fmt.Errorf("unable to read token file '%v' with error '%v'", confFile.TokenPath, err)
	}
	var client *astraops.AuthenticatedClient
	if hasToken {
		token, err := ReadToken(confFile.TokenPath)
		if err != nil {
			return &astraops.AuthenticatedClient{}, fmt.Errorf("found token at '%v' but unable to read token with error '%v'", confFile.TokenPath, err)
		}
		return astraops.AuthenticateToken(token, env.Verbose), nil
	}
	hasSa, err := confFile.HasServiceAccount()
	if err != nil {
		return &astraops.AuthenticatedClient{}, fmt.Errorf("unable to read service account file '%v' with error '%v'", confFile.SaPath, err)
	}
	if !hasSa {
		return &astraops.AuthenticatedClient{}, fmt.Errorf("unable to access any file for directory `%v`, run astra-cli login first", confDir)
	}
	clientInfo, err := ReadLogin(confFile.SaPath)
	if err != nil {
		return &astraops.AuthenticatedClient{}, err
	}
	client, err = astraops.Authenticate(clientInfo, env.Verbose)
	if err != nil {
		return &astraops.AuthenticatedClient{}, fmt.Errorf("authenticate failed with error %v", err)
	}
	return client, nil
}
