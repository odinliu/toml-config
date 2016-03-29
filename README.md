# toml-config
[![GoDoc](https://godoc.org/github.com/odinliu/toml-config?status.png)](https://godoc.org/github.com/odinliu/toml-config)

A library to bind toml configuration file with golang's struct

# Go version
test on go v1.5.1

# Installation
go get github.com/odinliu/toml-config

# Usage
See godoc [here](https://godoc.org/github.com/odinliu/toml-config)

# Dependencies
[go-toml](https://github.com/pelletier/go-toml)

# Example
```
package main

import (
	"fmt"

	"github.com/odinliu/toml-config"
)

type GeneralConfig struct {
	LogLevel         int    `toml:"general.log_level"`
	LogFile          string `toml:"general.log_file"`
	UserAgent        string `toml:"general.user_agent"`
	EnableCoolie     bool   `toml:"general.enable_cookie"`
	CrawlerCount     int    `toml:"general.crawler_count"`
	HbaseWriterCount uint64 `toml:"general.hbase_writer_count"`
	MaxChannelCount  int    `toml:"general.max_channel_count"`
}

type HbaseConfig struct {
	ZkQuorum []string `toml:"hbase.zkQuorum"`
	ZkRoot   string   `toml:"hbase.zkRoot"`
	Table    string   `toml:"hbase.table"`
}

type SpiderConfig struct {
	General GeneralConfig
	Hbase   HbaseConfig
}

func main() {
	config := SpiderConfig{}
	tmlcfg.BindFile("spider.toml", &config)
	fmt.Println(config.General.UserAgent)
	fmt.Println(config.General.MaxChannelCount)
	fmt.Println(config.General.EnableCoolie)
	fmt.Println(config.Hbase.Table)
	fmt.Println(config.Hbase.ZkQuorum)
	fmt.Println(config.General.HbaseWriterCount)
}

```
# License
Copyright (c) 2016, Yiding Liu. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the author nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
