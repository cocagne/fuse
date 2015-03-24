// Copyright 2015 Google Inc. All Rights Reserved.
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

package fuse

import (
	"log"

	"github.com/jacobsa/bazilfuse"
	"github.com/jacobsa/fuse/fuseops"
)

// A connection to the fuse kernel process.
type Connection struct {
	logger *log.Logger
	wrapped
}

func newConnection(wrapped *bazilfuse.Conn) (c *Connection, err error)

// Read the next op from the kernel process. Return io.EOF if the kernel has
// closed the connection.
//
// This function delivers ops in exactly the order they are received from
// /dev/fuse. Be wary of naively calling it concurrently: you probably want
// ordering guarantees between e.g. write ops and flush ops. For example,
// close(2) causes WriteFileOps to be issued before a FlushFileOp, but doesn't
// wait for their response before issuing the latter (cf.
// https://github.com/jacobsa/fuse/issues/3).
func (c *Connection) ReadOp() (op fuseops.Op, err error)
