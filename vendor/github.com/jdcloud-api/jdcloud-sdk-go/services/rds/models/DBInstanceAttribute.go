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

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type DBInstanceAttribute struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称，具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Cloud-Database-and-Cache/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例类型，例如主实例，只读实例等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceType string `json:"instanceType"`

    /* 实例引擎类型，如MySQL或SQL Server等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    Engine string `json:"engine"`

    /* 实例引擎版本，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    EngineVersion string `json:"engineVersion"`

    /* 实例规格代码 (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* 磁盘，单位GB (Optional) */
    InstanceStorageGB int `json:"instanceStorageGB"`

    /* CPU核数 (Optional) */
    InstanceCPU int `json:"instanceCPU"`

    /* 内存大小，单位MB (Optional) */
    InstanceMemoryMB int `json:"instanceMemoryMB"`

    /* 地域ID，参见[地域及可用区对照表](../Enum-Definitions/Regions-AZ.md) (Optional) */
    RegionId string `json:"regionId"`

    /* 可用区ID，第一个为主实例在的可用区，参见[地域及可用区对照表](../Enum-Definitions/Regions-AZ.md) (Optional) */
    AzId []string `json:"azId"`

    /* VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 实例公网域名 (Optional) */
    InternalDomainName string `json:"internalDomainName"`

    /* 实例内网域名 (Optional) */
    PublicDomainName string `json:"publicDomainName"`

    /* 应用访问端口 (Optional) */
    InstancePort string `json:"instancePort"`

    /* 访问模式，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    ConnectionMode string `json:"connectionMode"`

    /* 审计状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    AuditStatus string `json:"auditStatus"`

    /* 实例状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 实例创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 计费配置 (Optional) */
    Charge charge.Charge `json:"charge"`
}
