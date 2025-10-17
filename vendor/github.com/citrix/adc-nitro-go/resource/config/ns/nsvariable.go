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
* Configuration for variable resource.
*/
type Nsvariable struct {
	/**
	* Variable name.  This follows the same syntax rules as other expression entity names:
		It must begin with an alpha character (A-Z or a-z) or an underscore (_).
		The rest of the characters must be alpha, numeric (0-9) or underscores.
		It cannot be re or xp (reserved for regular and XPath expressions).
		It cannot be an expression reserved word (e.g. SYS or HTTP).
		It cannot be used for an existing expression object (HTTP callout, patset, dataset, stringmap, or named expression).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Specification of the variable type; one of the following:
		ulong - singleton variable with an unsigned 64-bit value.
		text(value-max-size) - singleton variable with a text string value.
		map(text(key-max-size),ulong,max-entries) - map of text string keys to unsigned 64-bit values.
		map(text(key-max-size),text(value-max-size),max-entries) - map of text string keys to text string values.
		where
		value-max-size is a positive integer that is the maximum number of bytes in a text string value.
		key-max-size is a positive integer that is the maximum number of bytes in a text string key.
		max-entries is a positive integer that is the maximum number of entries in a map variable.
		For a global singleton text variable, value-max-size <= 64000.
		For a global map with ulong values, key-max-size <= 64000.
		For a global map with text values,  key-max-size + value-max-size <= 64000.
		max-entries is a positive integer that is the maximum number of entries in a map variable. This has a theoretical maximum of 2^64-1, but in actual use will be much smaller, considering the memory available for use by the map.
		Example:
		map(text(10),text(20),100) specifies a map of text string keys (max size 10 bytes) to text string values (max size 20 bytes), with 100 max entries.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Scope of the variable:
		global - (default) one set of values visible across all Packet Engines on a standalone Citrix ADC, an HA pair, or all nodes of a cluster
		transaction - one value for each request-response transaction (singleton variables only; no expiration)
	*/
	Scope string `json:"scope,omitempty"`
	/**
	* Action to perform if an assignment to a map exceeds its configured max-entries:
		lru - (default) reuse the least recently used entry in the map.
		undef - force the assignment to return an undefined (Undef) result to the policy executing the assignment.
	*/
	Iffull string `json:"iffull,omitempty"`
	/**
	* Action to perform if an value is assigned to a text variable that exceeds its configured max-size,
		or if a key is used that exceeds its configured max-size:
		truncate - (default) truncate the text string to the first max-size bytes and proceed.
		undef - force the assignment or expression evaluation to return an undefined (Undef) result to the policy executing the assignment or expression.
	*/
	Ifvaluetoobig string `json:"ifvaluetoobig,omitempty"`
	/**
	* Action to perform if on a variable reference in an expression if the variable is single-valued and uninitialized
		or if the variable is a map and there is no value for the specified key:
		init - (default) initialize the single-value variable, or create a map entry for the key and the initial value,
		using the -init value or its default.
		undef - force the expression evaluation to return an undefined (Undef) result to the policy executing the expression.
	*/
	Ifnovalue string `json:"ifnovalue,omitempty"`
	/**
	* Initialization value for this variable, to which a singleton variable or map entry will be set if it is referenced before an assignment action has assigned it a value. If the singleton variable or map entry already has been assigned a value, setting this parameter will have no effect on that variable value. Default: 0 for ulong, NULL for text
	*/
	Init string `json:"init,omitempty"`
	/**
	* Value expiration in seconds. If the value is not referenced within the expiration period it will be deleted. 0 (the default) means no expiration.
	*/
	Expires *int `json:"expires,omitempty"`
	/**
	* Comments associated with this variable.
	*/
	Comment string `json:"comment,omitempty"`

	//------- Read only Parameter ---------;

	Referencecount string `json:"referencecount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
