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
* Configuration for Integrated Cache content group resource.
*/
type Cachecontentgroup struct {
	/**
	* Name for the content group.  Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the content group is created.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Relative expiry time, in seconds, for expiring positive responses with response codes between 200 and 399. Cannot be used in combination with other Expiry attributes. Similar to -relExpiry but has lower precedence.
	*/
	Weakposrelexpiry int `json:"weakposrelexpiry,omitempty"`
	/**
	* Heuristic expiry time, in percent of the duration, since the object was last modified.
	*/
	Heurexpiryparam int `json:"heurexpiryparam,omitempty"`
	/**
	* Relative expiry time, in seconds, after which to expire an object cached in this content group.
	*/
	Relexpiry int `json:"relexpiry,omitempty"`
	/**
	* Relative expiry time, in milliseconds, after which to expire an object cached in this content group.
	*/
	Relexpirymillisec int `json:"relexpirymillisec,omitempty"`
	/**
	* Local time, up to 4 times a day, at which all objects in the content group must expire. 
		CLI Users:
		For example, to specify that the objects in the content group should expire by 11:00 PM, type the following command: add cache contentgroup <contentgroup name> -absexpiry 23:00 
		To specify that the objects in the content group should expire at 10:00 AM, 3 PM, 6 PM, and 11:00 PM, type: add cache contentgroup <contentgroup name> -absexpiry 10:00 15:00 18:00 23:00
	*/
	Absexpiry []string `json:"absexpiry,omitempty"`
	/**
	* Coordinated Universal Time (GMT), up to 4 times a day, when all objects in the content group must expire.
	*/
	Absexpirygmt []string `json:"absexpirygmt,omitempty"`
	/**
	* Relative expiry time, in seconds, for expiring negative responses. This value is used only if the expiry time cannot be determined from any other source. It is applicable only to the following status codes: 307, 403, 404, and 410.
	*/
	Weaknegrelexpiry int `json:"weaknegrelexpiry,omitempty"`
	/**
	* Parameters to use for parameterized hit evaluation of an object. Up to 128 parameters can be specified. Mutually exclusive with the Hit Selector parameter.
	*/
	Hitparams []string `json:"hitparams,omitempty"`
	/**
	* Parameters for parameterized invalidation of an object. You can specify up to 8 parameters. Mutually exclusive with invalSelector.
	*/
	Invalparams []string `json:"invalparams,omitempty"`
	/**
	* Ignore case when comparing parameter values during parameterized hit evaluation. (Parameter value case is ignored by default during parameterized invalidation.)
	*/
	Ignoreparamvaluecase string `json:"ignoreparamvaluecase,omitempty"`
	/**
	* Evaluate for parameters in the cookie header also.
	*/
	Matchcookies string `json:"matchcookies,omitempty"`
	/**
	* Take the host header into account during parameterized invalidation.
	*/
	Invalrestrictedtohost string `json:"invalrestrictedtohost,omitempty"`
	/**
	* Always poll for the objects in this content group. That is, retrieve the objects from the origin server whenever they are requested.
	*/
	Polleverytime string `json:"polleverytime,omitempty"`
	/**
	* Ignore any request to reload a cached object from the origin server.
		To guard against Denial of Service attacks, set this parameter to YES. For RFC-compliant behavior, set it to NO.
	*/
	Ignorereloadreq string `json:"ignorereloadreq,omitempty"`
	/**
	* Remove cookies from responses.
	*/
	Removecookies string `json:"removecookies,omitempty"`
	/**
	* Attempt to refresh objects that are about to go stale.
	*/
	Prefetch string `json:"prefetch,omitempty"`
	/**
	* Time period, in seconds before an object's calculated expiry time, during which to attempt prefetch.
	*/
	Prefetchperiod int `json:"prefetchperiod,omitempty"`
	/**
	* Time period, in milliseconds before an object's calculated expiry time, during which to attempt prefetch.
	*/
	Prefetchperiodmillisec int `json:"prefetchperiodmillisec,omitempty"`
	/**
	* Maximum number of outstanding prefetches that can be queued for the content group.
	*/
	Prefetchmaxpending int `json:"prefetchmaxpending,omitempty"`
	/**
	* Perform flash cache. Mutually exclusive with Poll Every Time (PET) on the same content group.
	*/
	Flashcache string `json:"flashcache,omitempty"`
	/**
	* Force expiration of the content immediately after the response is downloaded (upon receipt of the last byte of the response body). Applicable only to positive responses.
	*/
	Expireatlastbyte string `json:"expireatlastbyte,omitempty"`
	/**
	* Insert a Via header into the response.
	*/
	Insertvia string `json:"insertvia,omitempty"`
	/**
	* Insert an Age header into the response. An Age header contains information about the age of the object, in seconds, as calculated by the integrated cache.
	*/
	Insertage string `json:"insertage,omitempty"`
	/**
	* Insert an ETag header in the response. With ETag header insertion, the integrated cache does not serve full responses on repeat requests.
	*/
	Insertetag string `json:"insertetag,omitempty"`
	/**
	* Insert a Cache-Control header into the response.
	*/
	Cachecontrol string `json:"cachecontrol,omitempty"`
	/**
	* If the size of an object that is being downloaded is less than or equal to the quick abort value, and a client aborts during the download, the cache stops downloading the response. If the object is larger than the quick abort size, the cache continues to download the response.
	*/
	Quickabortsize int `json:"quickabortsize,omitempty"`
	/**
	* Minimum size of a response that can be cached in this content group.
		Default minimum response size is 0.
	*/
	Minressize int `json:"minressize,omitempty"`
	/**
	* Maximum size of a response that can be cached in this content group.
	*/
	Maxressize int `json:"maxressize,omitempty"`
	/**
	* Maximum amount of memory that the cache can use. The effective limit is based on the available memory of the Citrix ADC.
	*/
	Memlimit int `json:"memlimit,omitempty"`
	/**
	* Ignore Cache-Control and Pragma headers in the incoming request.
	*/
	Ignorereqcachinghdrs string `json:"ignorereqcachinghdrs,omitempty"`
	/**
	* Number of hits that qualifies a response for storage in this content group.
	*/
	Minhits int `json:"minhits,omitempty"`
	/**
	* Force policy evaluation for each response arriving from the origin server. Cannot be set to YES if the Prefetch parameter is also set to YES.
	*/
	Alwaysevalpolicies string `json:"alwaysevalpolicies,omitempty"`
	/**
	* Setting persistHA to YES causes IC to save objects in contentgroup to Secondary node in HA deployment.
	*/
	Persistha string `json:"persistha,omitempty"`
	/**
	* Do not flush objects from this content group under memory pressure.
	*/
	Pinned string `json:"pinned,omitempty"`
	/**
	* Perform DNS resolution for responses only if the destination IP address in the request does not match the destination IP address of the cached response.
	*/
	Lazydnsresolve string `json:"lazydnsresolve,omitempty"`
	/**
	* Selector for evaluating whether an object gets stored in a particular content group. A selector is an abstraction for a collection of PIXL expressions.
	*/
	Hitselector string `json:"hitselector,omitempty"`
	/**
	* Selector for invalidating objects in the content group. A selector is an abstraction for a collection of PIXL expressions.
	*/
	Invalselector string `json:"invalselector,omitempty"`
	/**
	* The type of the content group.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Query string specifying individual objects to flush from this group by using parameterized invalidation. If this parameter is not set, all objects are flushed from the group.
	*/
	Query string `json:"query,omitempty"`
	/**
	* Flush only objects that belong to the specified host. Do not use except with parameterized invalidation. Also, the Invalidation Restricted to Host parameter for the group must be set to YES.
	*/
	Host string `json:"host,omitempty"`
	/**
	* Value of the selector to be used for flushing objects from the content group. Requires that an invalidation selector be configured for the content group.
	*/
	Selectorvalue string `json:"selectorvalue,omitempty"`
	/**
	* content group whose objects are to be sent to secondary.
	*/
	Tosecondary string `json:"tosecondary,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`
	Prefetchcur string `json:"prefetchcur,omitempty"`
	Memusage string `json:"memusage,omitempty"`
	Memdusage string `json:"memdusage,omitempty"`
	Disklimit string `json:"disklimit,omitempty"`
	Cachenon304hits string `json:"cachenon304hits,omitempty"`
	Cache304hits string `json:"cache304hits,omitempty"`
	Cachecells string `json:"cachecells,omitempty"`
	Cachegroupincarnation string `json:"cachegroupincarnation,omitempty"`
	Persist string `json:"persist,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Cachenuminvalpolicy string `json:"cachenuminvalpolicy,omitempty"`
	Markercells string `json:"markercells,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
