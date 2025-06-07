// Copyright 2018 fatedier, fatedier@gmail.com
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

package msg

import (
	"io"

	jsonMsg "github.com/fatedier/golib/msg/json"
)

type Message = jsonMsg.Message

var msgCtl *jsonMsg.MsgCtl

// XOR key for simple encryption
var xorKey = []byte("1.1.1")

func init() {
	msgCtl = jsonMsg.NewMsgCtl()
	for typeByte, msg := range msgTypeMap {
		msgCtl.RegisterMsg(typeByte, msg)
	}
}

func xorBytes(data []byte) {
	for i := range data {
		data[i] ^= xorKey[i%len(xorKey)]
	}
}

func ReadMsg(c io.Reader) (msg Message, err error) {
	msg, err = msgCtl.ReadMsg(c)
	if err != nil {
		return nil, err
	}

	// XOR decrypt the message data
	if data, ok := msg.(interface{ GetData() []byte }); ok {
		xorBytes(data.GetData())
	}
	return msg, nil
}

func ReadMsgInto(c io.Reader, msg Message) (err error) {
	err = msgCtl.ReadMsgInto(c, msg)
	if err != nil {
		return err
	}

	// XOR decrypt the message data
	if data, ok := msg.(interface{ GetData() []byte }); ok {
		xorBytes(data.GetData())
	}
	return nil
}

func WriteMsg(c io.Writer, msg any) (err error) {
	// XOR encrypt the message data before writing
	if data, ok := msg.(interface{ GetData() []byte }); ok {
		xorBytes(data.GetData())
		defer func() {
			// Decrypt back after writing
			xorBytes(data.GetData())
		}()
	}
	return msgCtl.WriteMsg(c, msg)
}
