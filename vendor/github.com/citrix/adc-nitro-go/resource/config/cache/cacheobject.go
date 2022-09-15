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

package cache

/**
* Configuration for cache object resource.
*/
type Cacheobject struct {
	/**
	* URL of the particular object whose details is required. Parameter "host" must be specified along with the URL.
	*/
	Url string `json:"url,omitempty"`
	/**
	* ID of the cached object.
	*/
	Locator int `json:"locator,omitempty"`
	/**
	* HTTP status of the object.
	*/
	Httpstatus int `json:"httpstatus,omitempty"`
	/**
	* Host name of the object. Parameter "url" must be specified.
	*/
	Host string `json:"host,omitempty"`
	/**
	* Host port of the object. You must also set the Host parameter.
	*/
	Port int `json:"port,omitempty"`
	/**
	* Name of the content group to which the object belongs. It will display only the objects belonging to the specified content group. You must also set the Host parameter.
	*/
	Groupname string `json:"groupname,omitempty"`
	/**
	* HTTP request method that caused the object to be stored.
	*/
	Httpmethod string `json:"httpmethod,omitempty"`
	/**
	* Name of the content group whose objects should be listed.
	*/
	Group string `json:"group,omitempty"`
	/**
	* Ignore marker objects. Marker objects are created when a response exceeds the maximum or minimum response size for the content group or has not yet received the minimum number of hits for the content group.
	*/
	Ignoremarkerobjects string `json:"ignoremarkerobjects,omitempty"`
	/**
	* Include responses that have not yet reached a minimum number of hits before being cached.
	*/
	Includenotreadyobjects string `json:"includenotreadyobjects,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`
	/**
	* Object will be saved onto Secondary.
	*/
	Tosecondary string `json:"tosecondary,omitempty"`

	//------- Read only Parameter ---------;

	Cacheressize string `json:"cacheressize,omitempty"`
	Cachereshdrsize string `json:"cachereshdrsize,omitempty"`
	Cacheetag string `json:"cacheetag,omitempty"`
	Httpstatusoutput string `json:"httpstatusoutput,omitempty"`
	Cachereslastmod string `json:"cachereslastmod,omitempty"`
	Cachecontrol string `json:"cachecontrol,omitempty"`
	Cacheresdate string `json:"cacheresdate,omitempty"`
	Contentgroup string `json:"contentgroup,omitempty"`
	Destipv46 string `json:"destipv46,omitempty"`
	Destport string `json:"destport,omitempty"`
	Cachecellcomplex string `json:"cachecellcomplex,omitempty"`
	Hitparams string `json:"hitparams,omitempty"`
	Hitvalues string `json:"hitvalues,omitempty"`
	Cachecellreqtime string `json:"cachecellreqtime,omitempty"`
	Cachecellrestime string `json:"cachecellrestime,omitempty"`
	Cachecurage string `json:"cachecurage,omitempty"`
	Cachecellexpires string `json:"cachecellexpires,omitempty"`
	Cachecellexpiresmillisec string `json:"cachecellexpiresmillisec,omitempty"`
	Flushed string `json:"flushed,omitempty"`
	Prefetch string `json:"prefetch,omitempty"`
	Prefetchperiod string `json:"prefetchperiod,omitempty"`
	Prefetchperiodmillisec string `json:"prefetchperiodmillisec,omitempty"`
	Cachecellcurreaders string `json:"cachecellcurreaders,omitempty"`
	Cachecellcurmisses string `json:"cachecellcurmisses,omitempty"`
	Cachecellhits string `json:"cachecellhits,omitempty"`
	Cachecellmisses string `json:"cachecellmisses,omitempty"`
	Cachecelldhits string `json:"cachecelldhits,omitempty"`
	Cachecellcompressionformat string `json:"cachecellcompressionformat,omitempty"`
	Cachecellappfwmetadataexists string `json:"cachecellappfwmetadataexists,omitempty"`
	Cachecellhttp11 string `json:"cachecellhttp11,omitempty"`
	Cachecellweaketag string `json:"cachecellweaketag,omitempty"`
	Cachecellresbadsize string `json:"cachecellresbadsize,omitempty"`
	Markerreason string `json:"markerreason,omitempty"`
	Cachecellpolleverytime string `json:"cachecellpolleverytime,omitempty"`
	Cachecelletaginserted string `json:"cachecelletaginserted,omitempty"`
	Cachecellreadywithlastbyte string `json:"cachecellreadywithlastbyte,omitempty"`
	Cacheinmemory string `json:"cacheinmemory,omitempty"`
	Cacheindisk string `json:"cacheindisk,omitempty"`
	Cacheinsecondary string `json:"cacheinsecondary,omitempty"`
	Cachedirname string `json:"cachedirname,omitempty"`
	Cachefilename string `json:"cachefilename,omitempty"`
	Cachecelldestipverified string `json:"cachecelldestipverified,omitempty"`
	Cachecellfwpxyobj string `json:"cachecellfwpxyobj,omitempty"`
	Cachecellbasefile string `json:"cachecellbasefile,omitempty"`
	Cachecellminhitflag string `json:"cachecellminhitflag,omitempty"`
	Cachecellminhit string `json:"cachecellminhit,omitempty"`
	Policy string `json:"policy,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Selectorname string `json:"selectorname,omitempty"`
	Rule string `json:"rule,omitempty"`
	Selectorvalue string `json:"selectorvalue,omitempty"`
	Cacheurls string `json:"cacheurls,omitempty"`
	Warnbucketskip string `json:"warnbucketskip,omitempty"`
	Totalobjs string `json:"totalobjs,omitempty"`
	Httpcalloutcell string `json:"httpcalloutcell,omitempty"`
	Httpcalloutname string `json:"httpcalloutname,omitempty"`
	Returntype string `json:"returntype,omitempty"`
	Httpcalloutresult string `json:"httpcalloutresult,omitempty"`
	Locatorshow string `json:"locatorshow,omitempty"`
	Ceflags string `json:"ceflags,omitempty"`

}
