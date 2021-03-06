// Copyright 2018 JDCLOUD.COM
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
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type QueryPerformanceResult struct {

    /* sql语句 (Optional) */
    Sql string `json:"sql"`

    /* 上次执行时间，格式为YYYY-MM-DD hh:mm:ss (Optional) */
    LastExecutionTime string `json:"lastExecutionTime"`

    /* 平均执行时长，单位毫秒(ms) (Optional) */
    ElapsedTime int `json:"elapsedTime"`

    /* 执行次数 (Optional) */
    ExecutionCount int `json:"executionCount"`

    /* 平均CPU使用时间，单位毫秒(ms) (Optional) */
    WorkerTime int `json:"workerTime"`

    /* 平均逻辑读次数 (Optional) */
    LogicalReads int `json:"logicalReads"`

    /* 平均逻辑写次数 (Optional) */
    LogicalWrites int `json:"logicalWrites"`

    /* 平均物理读次数 (Optional) */
    PhysicalReads int `json:"physicalReads"`

    /* 上次返回记录数 (Optional) */
    LastRows int `json:"lastRows"`
}
