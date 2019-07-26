// Copyright (c) 2019 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/mosuka/blast/dispatcher"
	"github.com/urfave/cli"
)

func distributorGet(c *cli.Context) error {
	grpcAddr := c.String("grpc-address")
	id := c.Args().Get(0)
	if id == "" {
		err := errors.New("arguments are not correct")
		return err
	}

	client, err := dispatcher.NewGRPCClient(grpcAddr)
	if err != nil {
		return err
	}
	defer func() {
		err := client.Close()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}()

	fields, err := client.GetDocument(id)
	if err != nil {
		return err
	}

	fieldsBytes, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%v", string(fieldsBytes)))

	return nil
}