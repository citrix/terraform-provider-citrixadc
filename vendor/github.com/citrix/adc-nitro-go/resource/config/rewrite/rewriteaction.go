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

package rewrite

/**
* Configuration for rewrite action resource.
*/
type Rewriteaction struct {
	/**
	* Name for the user-defined rewrite action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my rewrite action" or 'my rewrite action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of user-defined rewrite action. The information that you provide for, and the effect of, each type are as follows:: 
		* REPLACE <target> <string_builder_expr>. Replaces the string with the string-builder expression.
		* REPLACE_ALL <target> <string_builder_expr1> -(pattern|search) <string_builder_expr2>. In the request or response specified by <target>, replaces all occurrences of the string defined by <string_builder_expr1> with the string defined by <string_builder_expr2>. You can use a PCRE-format pattern or the search facility to find the strings to be replaced.
		* REPLACE_HTTP_RES <string_builder_expr>. Replaces the complete HTTP response with the string defined by the string-builder expression.
		* REPLACE_SIP_RES <target> - Replaces the complete SIP response with the string specified by <target>.
		* INSERT_HTTP_HEADER <header_string_builder_expr> <contents_string_builder_expr>. Inserts the HTTP header specified by <header_string_builder_expr> and header contents specified by <contents_string_builder_expr>.
		* DELETE_HTTP_HEADER <target>. Deletes the HTTP header specified by <target>.
		* CORRUPT_HTTP_HEADER <target>. Replaces the header name of all occurrences of the HTTP header specified by <target> with a corrupted name, so that it will not be recognized by the receiver  Example: MY_HEADER is changed to MHEY_ADER.
		* INSERT_BEFORE <string_builder_expr1> <string_builder_expr1>. Finds the string specified in <string_builder_expr1> and inserts the string in <string_builder_expr2> before it.
		* INSERT_BEFORE_ALL <target> <string_builder_expr1> -(pattern|search) <string_builder_expr2>. In the request or response specified by <target>, locates all occurrences of the string specified in <string_builder_expr1> and inserts the string specified in <string_builder_expr2> before each. You can use a PCRE-format pattern or the search facility to find the strings.
		* INSERT_AFTER <string_builder_expr1> <string_builder_expr2>. Finds the string specified in <string_builder_expr1>, and inserts the string specified in <string_builder_expr2> after it.
		* INSERT_AFTER_ALL <target> <string_builder_expr1> -(pattern|search) <string_builder_expr>. In the request or response specified by <target>, locates all occurrences of the string specified by <string_builder_expr1> and inserts the string specified by <string_builder_expr2> after each. You can use a PCRE-format pattern or the search facility to find the strings.
		* DELETE <target>. Finds and deletes the specified target.
		* DELETE_ALL <target> -(pattern|search) <string_builder_expr>. In the request or response specified by <target>, locates and deletes all occurrences of the string specified by <string_builder_expr>. You can use a PCRE-format pattern or the search facility to find the strings.
		* REPLACE_DIAMETER_HEADER_FIELD <target> <field value>. In the request or response modify the header field specified by <target>. Use Diameter.req.flags.SET(<flag>) or Diameter.req.flags.UNSET<flag> as 'stringbuilderexpression' to set or unset flags.
		* REPLACE_DNS_HEADER_FIELD <target>. In the request or response modify the header field specified by <target>. 
		* REPLACE_DNS_ANSWER_SECTION <target>. Replace the DNS answer section in the response. This is currently applicable for A and AAAA records only. Use DNS.NEW_RRSET_A & DNS.NEW_RRSET_AAAA expressions to configure the new answer section 
	*/
	Type string `json:"type,omitempty"`
	/**
	* Expression that specifies which part of the request or response to rewrite.
	*/
	Target string `json:"target,omitempty"`
	/**
	* Expression that specifies the content to insert into the request or response at the specified location, or that replaces the specified string.
	*/
	Stringbuilderexpr string `json:"stringbuilderexpr,omitempty"`
	/**
	* DEPRECATED in favor of -search: Pattern that is used to match multiple strings in the request or response. The pattern may be a string literal (without quotes) or a PCRE-format regular expression with a delimiter that consists of any printable ASCII non-alphanumeric character except for the underscore (_) and space ( ) that is not otherwise used in the expression. Example: re~https?://|HTTPS?://~ The preceding regular expression can use the tilde (~) as the delimiter because that character does not appear in the regular expression itself. Used in the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types.
	*/
	Pattern string `json:"pattern,omitempty"`
	/**
	* Search facility that is used to match multiple strings in the request or response. Used in the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types. The following search types are supported:
		* Text ("text(string)") - A literal string. Example: -search text("hello")
		* Regular expression ("regex(re<delimiter>regular exp<delimiter>)") - Pattern that is used to match multiple strings in the request or response. The pattern may be a PCRE-format regular expression with a delimiter that consists of any printable ASCII non-alphanumeric character except for the underscore (_) and space ( ) that is not otherwise used in the expression. Example: -search regex(re~^hello*~) The preceding regular expression can use the tilde (~) as the delimiter because that character does not appear in the regular expression itself.
		* XPath ("xpath(xp<delimiter>xpath expression<delimiter>)") - An XPath expression to search XML. The delimiter has the same rules as for regex. Example: -search xpath(xp%/a/b%)
		* JSON ("xpath_json(xp<delimiter>xpath expression<delimiter>)") - An XPath expression to search JSON. The delimiter has the same rules as for regex. Example: -search xpath_json(xp%/a/b%)
		NOTE: JSON searches use the same syntax as XPath searches, but operate on JSON files instead of standard XML files.
		* HTML ("xpath_html(xp<delimiter>xpath expression<delimiter>)") - An XPath expression to search HTML. The delimiter has the same rules as for regex. Example: -search xpath_html(xp%/html/body%)
		NOTE: HTML searches use the same syntax as XPath searches, but operate on HTML files instead of standard XML files; HTML 5 rules for the file syntax are used; HTML 4 and later are supported.
		* Patset ("patset(patset)") - A predefined pattern set. Example: -search patset("patset1").
		* Datset ("dataset(dataset)") - A predefined dataset. Example: -search dataset("dataset1").
		* AVP ("avp(avp number)") - AVP number that is used to match multiple AVPs in a Diameter/Radius Message. Example: -search avp(999)
		Note: for all these the TARGET prefix can be used in the replacement expression to specify the text that was selected by the -search parameter, optionally adjusted by the -refineSearch parameter.
		Example: TARGET.BEFORE_STR(",")
	*/
	Search string `json:"search,omitempty"`
	/**
	* Bypass the safety check and allow unsafe expressions. An unsafe expression is one that contains references to message elements that might not be present in all messages. If an expression refers to a missing request element, an empty string is used instead.
	*/
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`
	/**
	* Specify additional criteria to refine the results of the search. 
		Always starts with the "extend(m,n)" operation, where 'm' specifies number of bytes to the left of selected data and 'n' specifies number of bytes to the right of selected data to extend the selected area.
		You can use refineSearch only on body expressions, and for the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types.
		Example: -refineSearch 'EXTEND(10, 20).REGEX_SELECT(re~0x[0-9a-zA-Z]+~).
	*/
	Refinesearch string `json:"refinesearch,omitempty"`
	/**
	* Comment. Can be used to preserve information about this rewrite action.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the rewrite action. 
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my rewrite action" or 'my rewrite action').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Description string `json:"description,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
