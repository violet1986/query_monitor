// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package models

import (
	"container/list"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Query struct {
	SampleTime string `json:"sample_time"`
	//	Tmid       int64    `json:"tmid"`
	//	Ssid       int64    `json:"ssid"`
	//	Ccnt       int64    `json:"ccnt"`
	Username string `json:"username"`
	Db       string `json:"db"`
	//	Cost       int64    `json:"cost"`
	Tsubmit string `json:"tsubmit"`
	Tstart  string `json:"tstart"`
	Tfinish string `json:"tfinish"`
	//	ExecTime   string   `json:"exec_time"`
	//	WaitTime   string   `json:"wait_time"`
	//	Status     string   `json:"status"`
	//	RowsOut    int64    `json:"rows_out"`
	//	CpuElapsed int64    `json:"cpu_elapsed"`
	CpuCurrpct string `json:"cpu_currpct"`
	//	SkewCpu    string   `json:"skew_cpu"`
	//	SkewRows   string   `json:"skew_rows"`
	//	QueryHash  int64    `json:"query_hash"`
	//	Priority   string   `json:"priority"`
	//	Queuename  string   `json:"queuename"`
	//	Id         string   `json:"id"`
	QueryText string `json:"query_text,omitempty"`
	//	Explain    []string `json:"explain,omitempty"`
	Timestamp int
}

const archiveSize = 20

// Event archives.
var archive = list.New()

// NewArchive saves new event to archive list.
func NewArchive(q Query) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(q)
}

// GetEvents returns all events after lastReceived.
func GetQueries(lastReceived int) []Query {
	queries := make([]Query, 0, archive.Len())
	for q := archive.Front(); q != nil; q = q.Next() {
		e := q.Value.(Query)
		if e.Timestamp > int(lastReceived) {
			queries = append(queries, e)
		}
	}
	return queries
}
