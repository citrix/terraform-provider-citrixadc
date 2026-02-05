package rewriteaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RewriteactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Can be used to preserve information about this rewrite action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the user-defined rewrite action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite action\" or 'my rewrite action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the rewrite action. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite action\" or 'my rewrite action').",
			},
			"refinesearch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify additional criteria to refine the results of the search. \nAlways starts with the \"extend(m,n)\" operation, where 'm' specifies number of bytes to the left of selected data and 'n' specifies number of bytes to the right of selected data to extend the selected area.\nYou can use refineSearch only on body expressions, and for the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types.\nExample: -refineSearch 'EXTEND(10, 20).REGEX_SELECT(re~0x[0-9a-zA-Z]+~).",
			},
			"search": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Search facility that is used to match multiple strings in the request or response. Used in the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types. The following search types are supported:\n* Text (\"text(string)\") - A literal string. Example: -search text(\"hello\")\n* Regular expression (\"regex(re<delimiter>regular exp<delimiter>)\") - Pattern that is used to match multiple strings in the request or response. The pattern may be a PCRE-format regular expression with a delimiter that consists of any printable ASCII non-alphanumeric character except for the underscore (_) and space ( ) that is not otherwise used in the expression. Example: -search regex(re~^hello*~) The preceding regular expression can use the tilde (~) as the delimiter because that character does not appear in the regular expression itself.\n* XPath (\"xpath(xp<delimiter>xpath expression<delimiter>)\") - An XPath expression to search XML. The delimiter has the same rules as for regex. Example: -search xpath(xp%/a/b%)\n* JSON (\"xpath_json(xp<delimiter>xpath expression<delimiter>)\") - An XPath expression to search JSON. The delimiter has the same rules as for regex. Example: -search xpath_json(xp%/a/b%)\nNOTE: JSON searches use the same syntax as XPath searches, but operate on JSON files instead of standard XML files.\n* HTML (\"xpath_html(xp<delimiter>xpath expression<delimiter>)\") - An XPath expression to search HTML. The delimiter has the same rules as for regex. Example: -search xpath_html(xp%/html/body%)\nNOTE: HTML searches use the same syntax as XPath searches, but operate on HTML files instead of standard XML files; HTML 5 rules for the file syntax are used; HTML 4 and later are supported.\n* Patset (\"patset(patset)\") - A predefined pattern set. Example: -search patset(\"patset1\").\n* Datset (\"dataset(dataset)\") - A predefined dataset. Example: -search dataset(\"dataset1\").\n* AVP (\"avp(avp number)\") - AVP number that is used to match multiple AVPs in a Diameter/Radius Message. Example: -search avp(999)\n\nNote: for all these the TARGET prefix can be used in the replacement expression to specify the text that was selected by the -search parameter, optionally adjusted by the -refineSearch parameter.\nExample: TARGET.BEFORE_STR(\",\")",
			},
			"stringbuilderexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that specifies the content to insert into the request or response at the specified location, or that replaces the specified string.",
			},
			"target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that specifies which part of the request or response to rewrite.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of user-defined rewrite action. The information that you provide for, and the effect of, each type are as follows:: \n* REPLACE <target> <string_builder_expr>. Replaces the string with the string-builder expression.\n* REPLACE_ALL <target> <string_builder_expr> -search <search_expr>. In the request or response specified by <target>, replaces all occurrences of the string defined by <string_builder_expr> with the string defined by <search_expr>.\n* REPLACE_HTTP_RES <string_builder_expr>. Replaces the complete HTTP response with the string defined by the string-builder expression.\n* REPLACE_SIP_RES <target> - Replaces the complete SIP response with the string specified by <target>.\n* INSERT_HTTP_HEADER <header_string_builder_expr> <contents_string_builder_expr>. Inserts the HTTP header specified by <header_string_builder_expr> and header contents specified by <contents_string_builder_expr>.\n* DELETE_HTTP_HEADER <target>. Deletes the HTTP header specified by <target>.\n* CORRUPT_HTTP_HEADER <target>. Replaces the header name of all occurrences of the HTTP header specified by <target> with a corrupted name, so that it will not be recognized by the receiver  Example: MY_HEADER is changed to MHEY_ADER.\n* INSERT_BEFORE <target_expr> <string_builder_expr>. Finds the string specified in <target_expr> and inserts the string in <string_builder_expr> before it.\n* INSERT_BEFORE_ALL <target> <string_builder_expr> -search <search_expr>. In the request or response specified by <target>, locates all occurrences of the string specified in <string_builder_expr> and inserts the string specified in <search_expr> before each.\n* INSERT_AFTER <target_expr> <string_builder_expr>. Finds the string specified in <target_expr>, and inserts the string specified in <string_builder_expr> after it.\n* INSERT_AFTER_ALL <target> <string_builder_expr> -search <search_expr>. In the request or response specified by <target>, locates all occurrences of the string specified by <string_builder_expr> and inserts the string specified by <search_expr> after each.\n* DELETE <target>. Finds and deletes the specified target.\n* DELETE_ALL <target> -search <string_builder_expr>. In the request or response specified by <target>, locates and deletes all occurrences of the string specified by <string_builder_expr>.\n* REPLACE_DIAMETER_HEADER_FIELD <target> <field value>. In the request or response modify the header field specified by <target>. Use Diameter.req.flags.SET(<flag>) or Diameter.req.flags.UNSET<flag> as 'stringbuilderexpression' to set or unset flags.\n* REPLACE_DNS_HEADER_FIELD <target>. In the request or response modify the header field specified by <target>. \n* REPLACE_DNS_ANSWER_SECTION <target>. Replace the DNS answer section in the response. This is currently applicable for A and AAAA records only. Use DNS.NEW_RRSET_A & DNS.NEW_RRSET_AAAA expressions to configure the new answer section.\n* REPLACE_MQTT <target> <string_builder_expr> : Replace MQTT message fields specified in <target_expr> to the value specified in <string_builder_expr>\n* INSERT_MQTT <string_builder_expr> : Insert the string_builder_expr to an appropriate packet field in the MQTT message.\n* INSERT_AFTER_MQTT <target_expr> <string_builder_expr> : Insert a topic specified in <string_builder_expr> in the MQTT Subscribe or Unsubscribe message after the specified target_expr.\n* INSERT_BEFORE_MQTT <target_expr> <string_builder_expr> : Insert a topic specified in <string_builder_expr> in the MQTT Subscribe or Unsubscribe message before the specified target_expr.\n* DELETE_MQTT <target> : Deletes the specified target in the MQTT message.",
			},
		},
	}
}
