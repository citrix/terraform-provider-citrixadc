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
* Configuration for timeout resource.
*/
type Nstimeout struct {
	/**
	* Interval, in seconds, at which the Citrix ADC zombie cleanup process must run. This process cleans up inactive TCP connections.
	*/
	Zombie int `json:"zombie,omitempty"`
	/**
	* Client idle timeout (in seconds). If zero, the service-type default value is taken when service is created.
	*/
	Client int `json:"client"`
	/**
	* Server idle timeout (in seconds).  If zero, the service-type default value is taken when service is created.
	*/
	Server int `json:"server"`
	/**
	* Global idle timeout, in seconds, for client connections of HTTP service type. This value is over ridden by the client timeout that is configured on individual entities.
	*/
	Httpclient int `json:"httpclient"`
	/**
	* Global idle timeout, in seconds, for server connections of HTTP service type. This value is over ridden by the server timeout that is configured on individual entities.
	*/
	Httpserver int `json:"httpserver"`
	/**
	* Global idle timeout, in seconds, for non-HTTP client connections of TCP service type. This value is over ridden by the client timeout that is configured on individual entities.
	*/
	Tcpclient int `json:"tcpclient"`
	/**
	* Global idle timeout, in seconds, for non-HTTP server connections of TCP service type. This value is over ridden by the server timeout that is configured on entities.
	*/
	Tcpserver int `json:"tcpserver"`
	/**
	* Global idle timeout, in seconds, for non-TCP client connections. This value is over ridden by the client timeout that is configured on individual entities.
	*/
	Anyclient int `json:"anyclient"`
	/**
	* Global idle timeout, in seconds, for non TCP server connections. This value is over ridden by the server timeout that is configured on individual entities.
	*/
	Anyserver int `json:"anyserver"`
	/**
	* Global idle timeout, in seconds, for TCP client connections. This value takes precedence over  entity level timeout settings (vserver/service). This is applicable only to transport protocol TCP.
	*/
	Anytcpclient int `json:"anytcpclient"`
	/**
	* Global idle timeout, in seconds, for TCP server connections. This value takes precedence over entity level timeout settings ( vserver/service). This is applicable only to transport protocol TCP.
	*/
	Anytcpserver int `json:"anytcpserver"`
	/**
	* Idle timeout, in seconds, for connections that are in TCP half-closed state.
	*/
	Halfclose int `json:"halfclose,omitempty"`
	/**
	* Interval at which the zombie clean-up process for non-TCP connections should run. Inactive IP NAT connections will be cleaned up.
	*/
	Nontcpzombie int `json:"nontcpzombie,omitempty"`
	/**
	* Alternative idle timeout, in seconds, for closed TCP NATPCB connections.
	*/
	Reducedfintimeout int `json:"reducedfintimeout,omitempty"`
	/**
	* Timer interval, in seconds, for abruptly terminated TCP NATPCB connections.
	*/
	Reducedrsttimeout int `json:"reducedrsttimeout"`
	/**
	* Timer interval, in seconds, for new TCP NATPCB connections on which no data was received.
	*/
	Newconnidletimeout int `json:"newconnidletimeout,omitempty"`

}
