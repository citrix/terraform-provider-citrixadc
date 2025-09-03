/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package ns

/**
* Configuration for licenseserver resource.
*/
type Nslicenseserverpool struct {
	/**
	* If this flag is used while doing getinventory, it displays all licenses from licenseserver.
	*/
	Getalllicenses bool `json:"getalllicenses,omitempty"`

	//------- Read only Parameter ---------;

	Instancetotal string `json:"instancetotal,omitempty"`
	Instanceavailable string `json:"instanceavailable,omitempty"`
	Standardbandwidthtotal string `json:"standardbandwidthtotal,omitempty"`
	Standardbandwidthavailable string `json:"standardbandwidthavailable,omitempty"`
	Enterprisebandwidthtotal string `json:"enterprisebandwidthtotal,omitempty"`
	Enterprisebandwidthavailable string `json:"enterprisebandwidthavailable,omitempty"`
	Platinumbandwidthtotal string `json:"platinumbandwidthtotal,omitempty"`
	Platinumbandwidthavailable string `json:"platinumbandwidthavailable,omitempty"`
	Standardcputotal string `json:"standardcputotal,omitempty"`
	Standardcpuavailable string `json:"standardcpuavailable,omitempty"`
	Enterprisecputotal string `json:"enterprisecputotal,omitempty"`
	Enterprisecpuavailable string `json:"enterprisecpuavailable,omitempty"`
	Platinumcputotal string `json:"platinumcputotal,omitempty"`
	Platinumcpuavailable string `json:"platinumcpuavailable,omitempty"`
	Cpxinstancetotal string `json:"cpxinstancetotal,omitempty"`
	Cpxinstanceavailable string `json:"cpxinstanceavailable,omitempty"`
	Vpx1stotal string `json:"vpx1stotal,omitempty"`
	Vpx1savailable string `json:"vpx1savailable,omitempty"`
	Vpx1ptotal string `json:"vpx1ptotal,omitempty"`
	Vpx1pavailable string `json:"vpx1pavailable,omitempty"`
	Vpx5stotal string `json:"vpx5stotal,omitempty"`
	Vpx5savailable string `json:"vpx5savailable,omitempty"`
	Vpx5ptotal string `json:"vpx5ptotal,omitempty"`
	Vpx5pavailable string `json:"vpx5pavailable,omitempty"`
	Vpx10stotal string `json:"vpx10stotal,omitempty"`
	Vpx10savailable string `json:"vpx10savailable,omitempty"`
	Vpx10etotal string `json:"vpx10etotal,omitempty"`
	Vpx10eavailable string `json:"vpx10eavailable,omitempty"`
	Vpx10ptotal string `json:"vpx10ptotal,omitempty"`
	Vpx10pavailable string `json:"vpx10pavailable,omitempty"`
	Vpx25stotal string `json:"vpx25stotal,omitempty"`
	Vpx25savailable string `json:"vpx25savailable,omitempty"`
	Vpx25etotal string `json:"vpx25etotal,omitempty"`
	Vpx25eavailable string `json:"vpx25eavailable,omitempty"`
	Vpx25ptotal string `json:"vpx25ptotal,omitempty"`
	Vpx25pavailable string `json:"vpx25pavailable,omitempty"`
	Vpx50stotal string `json:"vpx50stotal,omitempty"`
	Vpx50savailable string `json:"vpx50savailable,omitempty"`
	Vpx50etotal string `json:"vpx50etotal,omitempty"`
	Vpx50eavailable string `json:"vpx50eavailable,omitempty"`
	Vpx50ptotal string `json:"vpx50ptotal,omitempty"`
	Vpx50pavailable string `json:"vpx50pavailable,omitempty"`
	Vpx100stotal string `json:"vpx100stotal,omitempty"`
	Vpx100savailable string `json:"vpx100savailable,omitempty"`
	Vpx100etotal string `json:"vpx100etotal,omitempty"`
	Vpx100eavailable string `json:"vpx100eavailable,omitempty"`
	Vpx100ptotal string `json:"vpx100ptotal,omitempty"`
	Vpx100pavailable string `json:"vpx100pavailable,omitempty"`
	Vpx200stotal string `json:"vpx200stotal,omitempty"`
	Vpx200savailable string `json:"vpx200savailable,omitempty"`
	Vpx200etotal string `json:"vpx200etotal,omitempty"`
	Vpx200eavailable string `json:"vpx200eavailable,omitempty"`
	Vpx200ptotal string `json:"vpx200ptotal,omitempty"`
	Vpx200pavailable string `json:"vpx200pavailable,omitempty"`
	Vpx500stotal string `json:"vpx500stotal,omitempty"`
	Vpx500savailable string `json:"vpx500savailable,omitempty"`
	Vpx500etotal string `json:"vpx500etotal,omitempty"`
	Vpx500eavailable string `json:"vpx500eavailable,omitempty"`
	Vpx500ptotal string `json:"vpx500ptotal,omitempty"`
	Vpx500pavailable string `json:"vpx500pavailable,omitempty"`
	Vpx1000stotal string `json:"vpx1000stotal,omitempty"`
	Vpx1000savailable string `json:"vpx1000savailable,omitempty"`
	Vpx1000etotal string `json:"vpx1000etotal,omitempty"`
	Vpx1000eavailable string `json:"vpx1000eavailable,omitempty"`
	Vpx1000ptotal string `json:"vpx1000ptotal,omitempty"`
	Vpx1000pavailable string `json:"vpx1000pavailable,omitempty"`
	Vpx2000ptotal string `json:"vpx2000ptotal,omitempty"`
	Vpx2000pavailable string `json:"vpx2000pavailable,omitempty"`
	Vpx3000stotal string `json:"vpx3000stotal,omitempty"`
	Vpx3000savailable string `json:"vpx3000savailable,omitempty"`
	Vpx3000etotal string `json:"vpx3000etotal,omitempty"`
	Vpx3000eavailable string `json:"vpx3000eavailable,omitempty"`
	Vpx3000ptotal string `json:"vpx3000ptotal,omitempty"`
	Vpx3000pavailable string `json:"vpx3000pavailable,omitempty"`
	Vpx4000ptotal string `json:"vpx4000ptotal,omitempty"`
	Vpx4000pavailable string `json:"vpx4000pavailable,omitempty"`
	Vpx5000stotal string `json:"vpx5000stotal,omitempty"`
	Vpx5000savailable string `json:"vpx5000savailable,omitempty"`
	Vpx5000etotal string `json:"vpx5000etotal,omitempty"`
	Vpx5000eavailable string `json:"vpx5000eavailable,omitempty"`
	Vpx5000ptotal string `json:"vpx5000ptotal,omitempty"`
	Vpx5000pavailable string `json:"vpx5000pavailable,omitempty"`
	Vpx8000stotal string `json:"vpx8000stotal,omitempty"`
	Vpx8000savailable string `json:"vpx8000savailable,omitempty"`
	Vpx8000etotal string `json:"vpx8000etotal,omitempty"`
	Vpx8000eavailable string `json:"vpx8000eavailable,omitempty"`
	Vpx8000ptotal string `json:"vpx8000ptotal,omitempty"`
	Vpx8000pavailable string `json:"vpx8000pavailable,omitempty"`
	Vpx10000stotal string `json:"vpx10000stotal,omitempty"`
	Vpx10000savailable string `json:"vpx10000savailable,omitempty"`
	Vpx10000etotal string `json:"vpx10000etotal,omitempty"`
	Vpx10000eavailable string `json:"vpx10000eavailable,omitempty"`
	Vpx10000ptotal string `json:"vpx10000ptotal,omitempty"`
	Vpx10000pavailable string `json:"vpx10000pavailable,omitempty"`
	Vpx15000stotal string `json:"vpx15000stotal,omitempty"`
	Vpx15000savailable string `json:"vpx15000savailable,omitempty"`
	Vpx15000etotal string `json:"vpx15000etotal,omitempty"`
	Vpx15000eavailable string `json:"vpx15000eavailable,omitempty"`
	Vpx15000ptotal string `json:"vpx15000ptotal,omitempty"`
	Vpx15000pavailable string `json:"vpx15000pavailable,omitempty"`
	Vpx25000stotal string `json:"vpx25000stotal,omitempty"`
	Vpx25000savailable string `json:"vpx25000savailable,omitempty"`
	Vpx25000etotal string `json:"vpx25000etotal,omitempty"`
	Vpx25000eavailable string `json:"vpx25000eavailable,omitempty"`
	Vpx25000ptotal string `json:"vpx25000ptotal,omitempty"`
	Vpx25000pavailable string `json:"vpx25000pavailable,omitempty"`
	Vpx40000stotal string `json:"vpx40000stotal,omitempty"`
	Vpx40000savailable string `json:"vpx40000savailable,omitempty"`
	Vpx40000etotal string `json:"vpx40000etotal,omitempty"`
	Vpx40000eavailable string `json:"vpx40000eavailable,omitempty"`
	Vpx40000ptotal string `json:"vpx40000ptotal,omitempty"`
	Vpx40000pavailable string `json:"vpx40000pavailable,omitempty"`
	Vpx100000stotal string `json:"vpx100000stotal,omitempty"`
	Vpx100000savailable string `json:"vpx100000savailable,omitempty"`
	Vpx100000etotal string `json:"vpx100000etotal,omitempty"`
	Vpx100000eavailable string `json:"vpx100000eavailable,omitempty"`
	Vpx100000ptotal string `json:"vpx100000ptotal,omitempty"`
	Vpx100000pavailable string `json:"vpx100000pavailable,omitempty"`
	Licensemode string `json:"licensemode,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
