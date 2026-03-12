package appfwprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileResourceModel describes the resource data model.
type AppfwprofileResourceModel struct {
	Id                                         types.String `tfsdk:"id"`
	Addcookieflags                             types.String `tfsdk:"addcookieflags"`
	Apispec                                    types.String `tfsdk:"apispec"`
	Archivename                                types.String `tfsdk:"archivename"`
	AsProfBypassListEnable                     types.String `tfsdk:"as_prof_bypass_list_enable"`
	AsProfDenyListEnable                       types.String `tfsdk:"as_prof_deny_list_enable"`
	Augment                                    types.Bool   `tfsdk:"augment"`
	Blockkeywordaction                         types.List   `tfsdk:"blockkeywordaction"`
	Bufferoverflowaction                       types.List   `tfsdk:"bufferoverflowaction"`
	Bufferoverflowmaxcookielength              types.Int64  `tfsdk:"bufferoverflowmaxcookielength"`
	Bufferoverflowmaxheaderlength              types.Int64  `tfsdk:"bufferoverflowmaxheaderlength"`
	Bufferoverflowmaxquerylength               types.Int64  `tfsdk:"bufferoverflowmaxquerylength"`
	Bufferoverflowmaxtotalheaderlength         types.Int64  `tfsdk:"bufferoverflowmaxtotalheaderlength"`
	Bufferoverflowmaxurllength                 types.Int64  `tfsdk:"bufferoverflowmaxurllength"`
	Canonicalizehtmlresponse                   types.String `tfsdk:"canonicalizehtmlresponse"`
	Ceflogging                                 types.String `tfsdk:"ceflogging"`
	Checkrequestheaders                        types.String `tfsdk:"checkrequestheaders"`
	Clientipexpression                         types.String `tfsdk:"clientipexpression"`
	Cmdinjectionaction                         types.List   `tfsdk:"cmdinjectionaction"`
	Cmdinjectiongrammar                        types.String `tfsdk:"cmdinjectiongrammar"`
	Cmdinjectiontype                           types.String `tfsdk:"cmdinjectiontype"`
	Comment                                    types.String `tfsdk:"comment"`
	Contenttypeaction                          types.List   `tfsdk:"contenttypeaction"`
	Cookieconsistencyaction                    types.List   `tfsdk:"cookieconsistencyaction"`
	Cookieencryption                           types.String `tfsdk:"cookieencryption"`
	Cookiehijackingaction                      types.List   `tfsdk:"cookiehijackingaction"`
	Cookieproxying                             types.String `tfsdk:"cookieproxying"`
	Cookiesamesiteattribute                    types.String `tfsdk:"cookiesamesiteattribute"`
	Cookietransforms                           types.String `tfsdk:"cookietransforms"`
	Creditcard                                 types.List   `tfsdk:"creditcard"`
	Creditcardaction                           types.List   `tfsdk:"creditcardaction"`
	Creditcardmaxallowed                       types.Int64  `tfsdk:"creditcardmaxallowed"`
	Creditcardxout                             types.String `tfsdk:"creditcardxout"`
	Crosssitescriptingaction                   types.List   `tfsdk:"crosssitescriptingaction"`
	Crosssitescriptingcheckcompleteurls        types.String `tfsdk:"crosssitescriptingcheckcompleteurls"`
	Crosssitescriptingtransformunsafehtml      types.String `tfsdk:"crosssitescriptingtransformunsafehtml"`
	Csrftagaction                              types.List   `tfsdk:"csrftagaction"`
	Customsettings                             types.String `tfsdk:"customsettings"`
	Defaultcharset                             types.String `tfsdk:"defaultcharset"`
	Defaultfieldformatmaxlength                types.Int64  `tfsdk:"defaultfieldformatmaxlength"`
	Defaultfieldformatmaxoccurrences           types.Int64  `tfsdk:"defaultfieldformatmaxoccurrences"`
	Defaultfieldformatminlength                types.Int64  `tfsdk:"defaultfieldformatminlength"`
	Defaultfieldformattype                     types.String `tfsdk:"defaultfieldformattype"`
	Defaults                                   types.String `tfsdk:"defaults"`
	Denyurlaction                              types.List   `tfsdk:"denyurlaction"`
	Dosecurecreditcardlogging                  types.String `tfsdk:"dosecurecreditcardlogging"`
	Dynamiclearning                            types.List   `tfsdk:"dynamiclearning"`
	Enableformtagging                          types.String `tfsdk:"enableformtagging"`
	Errorurl                                   types.String `tfsdk:"errorurl"`
	Excludefileuploadfromchecks                types.String `tfsdk:"excludefileuploadfromchecks"`
	Exemptclosureurlsfromsecuritychecks        types.String `tfsdk:"exemptclosureurlsfromsecuritychecks"`
	Fakeaccountdetection                       types.String `tfsdk:"fakeaccountdetection"`
	Fieldconsistencyaction                     types.List   `tfsdk:"fieldconsistencyaction"`
	Fieldformataction                          types.List   `tfsdk:"fieldformataction"`
	Fieldscan                                  types.String `tfsdk:"fieldscan"`
	Fieldscanlimit                             types.Int64  `tfsdk:"fieldscanlimit"`
	Fileuploadmaxnum                           types.Int64  `tfsdk:"fileuploadmaxnum"`
	Fileuploadtypesaction                      types.List   `tfsdk:"fileuploadtypesaction"`
	Geolocationlogging                         types.String `tfsdk:"geolocationlogging"`
	Grpcaction                                 types.List   `tfsdk:"grpcaction"`
	Htmlerrorobject                            types.String `tfsdk:"htmlerrorobject"`
	Htmlerrorstatuscode                        types.Int64  `tfsdk:"htmlerrorstatuscode"`
	Htmlerrorstatusmessage                     types.String `tfsdk:"htmlerrorstatusmessage"`
	Importprofilename                          types.String `tfsdk:"importprofilename"`
	Infercontenttypexmlpayloadaction           types.List   `tfsdk:"infercontenttypexmlpayloadaction"`
	Insertcookiesamesiteattribute              types.String `tfsdk:"insertcookiesamesiteattribute"`
	Inspectcontenttypes                        types.List   `tfsdk:"inspectcontenttypes"`
	Inspectquerycontenttypes                   types.List   `tfsdk:"inspectquerycontenttypes"`
	Invalidpercenthandling                     types.String `tfsdk:"invalidpercenthandling"`
	Jsonblockkeywordaction                     types.List   `tfsdk:"jsonblockkeywordaction"`
	Jsoncmdinjectionaction                     types.List   `tfsdk:"jsoncmdinjectionaction"`
	Jsoncmdinjectiongrammar                    types.String `tfsdk:"jsoncmdinjectiongrammar"`
	Jsoncmdinjectiontype                       types.String `tfsdk:"jsoncmdinjectiontype"`
	Jsondosaction                              types.List   `tfsdk:"jsondosaction"`
	Jsonerrorobject                            types.String `tfsdk:"jsonerrorobject"`
	Jsonerrorstatuscode                        types.Int64  `tfsdk:"jsonerrorstatuscode"`
	Jsonerrorstatusmessage                     types.String `tfsdk:"jsonerrorstatusmessage"`
	Jsonfieldscan                              types.String `tfsdk:"jsonfieldscan"`
	Jsonfieldscanlimit                         types.Int64  `tfsdk:"jsonfieldscanlimit"`
	Jsonmessagescan                            types.String `tfsdk:"jsonmessagescan"`
	Jsonmessagescanlimit                       types.Int64  `tfsdk:"jsonmessagescanlimit"`
	Jsonsqlinjectionaction                     types.List   `tfsdk:"jsonsqlinjectionaction"`
	Jsonsqlinjectiongrammar                    types.String `tfsdk:"jsonsqlinjectiongrammar"`
	Jsonsqlinjectiontype                       types.String `tfsdk:"jsonsqlinjectiontype"`
	Jsonxssaction                              types.List   `tfsdk:"jsonxssaction"`
	Logeverypolicyhit                          types.String `tfsdk:"logeverypolicyhit"`
	Matchurlstring                             types.String `tfsdk:"matchurlstring"`
	Messagescan                                types.String `tfsdk:"messagescan"`
	Messagescanlimit                           types.Int64  `tfsdk:"messagescanlimit"`
	Messagescanlimitcontenttypes               types.List   `tfsdk:"messagescanlimitcontenttypes"`
	Multipleheaderaction                       types.List   `tfsdk:"multipleheaderaction"`
	Name                                       types.String `tfsdk:"name"`
	Optimizepartialreqs                        types.String `tfsdk:"optimizepartialreqs"`
	Overwrite                                  types.Bool   `tfsdk:"overwrite"`
	Percentdecoderecursively                   types.String `tfsdk:"percentdecoderecursively"`
	Postbodylimit                              types.Int64  `tfsdk:"postbodylimit"`
	Postbodylimitaction                        types.List   `tfsdk:"postbodylimitaction"`
	Postbodylimitsignature                     types.Int64  `tfsdk:"postbodylimitsignature"`
	Protofileobject                            types.String `tfsdk:"protofileobject"`
	Refererheadercheck                         types.String `tfsdk:"refererheadercheck"`
	Relaxationrules                            types.Bool   `tfsdk:"relaxationrules"`
	Replaceurlstring                           types.String `tfsdk:"replaceurlstring"`
	Requestcontenttype                         types.String `tfsdk:"requestcontenttype"`
	Responsecontenttype                        types.String `tfsdk:"responsecontenttype"`
	Restaction                                 types.List   `tfsdk:"restaction"`
	Rfcprofile                                 types.String `tfsdk:"rfcprofile"`
	Semicolonfieldseparator                    types.String `tfsdk:"semicolonfieldseparator"`
	Sessioncookiename                          types.String `tfsdk:"sessioncookiename"`
	Sessionlessfieldconsistency                types.String `tfsdk:"sessionlessfieldconsistency"`
	Sessionlessurlclosure                      types.String `tfsdk:"sessionlessurlclosure"`
	Signatures                                 types.String `tfsdk:"signatures"`
	Sqlinjectionaction                         types.List   `tfsdk:"sqlinjectionaction"`
	Sqlinjectionchecksqlwildchars              types.String `tfsdk:"sqlinjectionchecksqlwildchars"`
	Sqlinjectiongrammar                        types.String `tfsdk:"sqlinjectiongrammar"`
	Sqlinjectiononlycheckfieldswithsqlchars    types.String `tfsdk:"sqlinjectiononlycheckfieldswithsqlchars"`
	Sqlinjectionparsecomments                  types.String `tfsdk:"sqlinjectionparsecomments"`
	Sqlinjectionruletype                       types.String `tfsdk:"sqlinjectionruletype"`
	Sqlinjectiontransformspecialchars          types.String `tfsdk:"sqlinjectiontransformspecialchars"`
	Sqlinjectiontype                           types.String `tfsdk:"sqlinjectiontype"`
	Starturlaction                             types.List   `tfsdk:"starturlaction"`
	Starturlclosure                            types.String `tfsdk:"starturlclosure"`
	Streaming                                  types.String `tfsdk:"streaming"`
	Stripcomments                              types.String `tfsdk:"stripcomments"`
	Striphtmlcomments                          types.String `tfsdk:"striphtmlcomments"`
	Stripxmlcomments                           types.String `tfsdk:"stripxmlcomments"`
	Trace                                      types.String `tfsdk:"trace"`
	Type                                       types.List   `tfsdk:"type"`
	Urldecoderequestcookies                    types.String `tfsdk:"urldecoderequestcookies"`
	Usehtmlerrorobject                         types.String `tfsdk:"usehtmlerrorobject"`
	Verboseloglevel                            types.String `tfsdk:"verboseloglevel"`
	Xmlattachmentaction                        types.List   `tfsdk:"xmlattachmentaction"`
	Xmldosaction                               types.List   `tfsdk:"xmldosaction"`
	Xmlerrorobject                             types.String `tfsdk:"xmlerrorobject"`
	Xmlerrorstatuscode                         types.Int64  `tfsdk:"xmlerrorstatuscode"`
	Xmlerrorstatusmessage                      types.String `tfsdk:"xmlerrorstatusmessage"`
	Xmlformataction                            types.List   `tfsdk:"xmlformataction"`
	Xmlsoapfaultaction                         types.List   `tfsdk:"xmlsoapfaultaction"`
	Xmlsqlinjectionaction                      types.List   `tfsdk:"xmlsqlinjectionaction"`
	Xmlsqlinjectionchecksqlwildchars           types.String `tfsdk:"xmlsqlinjectionchecksqlwildchars"`
	Xmlsqlinjectiononlycheckfieldswithsqlchars types.String `tfsdk:"xmlsqlinjectiononlycheckfieldswithsqlchars"`
	Xmlsqlinjectionparsecomments               types.String `tfsdk:"xmlsqlinjectionparsecomments"`
	Xmlsqlinjectiontype                        types.String `tfsdk:"xmlsqlinjectiontype"`
	Xmlvalidationaction                        types.List   `tfsdk:"xmlvalidationaction"`
	Xmlwsiaction                               types.List   `tfsdk:"xmlwsiaction"`
	Xmlxssaction                               types.List   `tfsdk:"xmlxssaction"`
}

func (r *AppfwprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile resource.",
			},
			"addcookieflags": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Add the specified flags to cookies. Available settings function as follows:\n* None - Do not add flags to cookies.\n* HTTP Only - Add the HTTP Only flag to cookies, which prevents scripts from accessing cookies.\n* Secure - Add Secure flag to cookies.\n* All - Add both HTTPOnly and Secure flags to cookies.",
			},
			"apispec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the API Specification.",
			},
			"archivename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Source for tar archive.",
			},
			"as_prof_bypass_list_enable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable bypass list for the profile.",
			},
			"as_prof_deny_list_enable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable deny list for the profile.",
			},
			"augment": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Augment Relaxation Rules during import",
			},
			"blockkeywordaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Block Keyword action. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -blockKeywordAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -blockKeywordAction none\".",
			},
			"bufferoverflowaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Buffer Overflow actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -bufferOverflowAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -bufferOverflowAction none\".",
			},
			"bufferoverflowmaxcookielength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4096),
				Description: "Maximum length, in characters, for cookies sent to your protected web sites. Requests with longer cookies are blocked.",
			},
			"bufferoverflowmaxheaderlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4096),
				Description: "Maximum length, in characters, for HTTP headers in requests sent to your protected web sites. Requests with longer headers are blocked.",
			},
			"bufferoverflowmaxquerylength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(65535),
				Description: "Maximum length, in bytes, for query string sent to your protected web sites. Requests with longer query strings are blocked.",
			},
			"bufferoverflowmaxtotalheaderlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(65535),
				Description: "Maximum length, in bytes, for the total HTTP header length in requests sent to your protected web sites. The minimum value of this and maxHeaderLen in httpProfile will be used. Requests with longer length are blocked.",
			},
			"bufferoverflowmaxurllength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1024),
				Description: "Maximum length, in characters, for URLs on your protected web sites. Requests with longer URLs are blocked.",
			},
			"canonicalizehtmlresponse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Perform HTML entity encoding for any special characters in responses sent by your protected web sites.",
			},
			"ceflogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable CEF format logs for the profile.",
			},
			"checkrequestheaders": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check request headers as well as web forms for injected SQL and cross-site scripts.",
			},
			"clientipexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression to get the client IP.",
			},
			"cmdinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Command injection action. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -cmdInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -cmdInjectionAction none\".",
			},
			"cmdinjectiongrammar": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for CMD injection using CMD grammar",
			},
			"cmdinjectiontype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("CMDSplCharANDKeyword"),
				Description: "Available CMD injection types.\n-CMDSplChar              : Checks for CMD Special Chars\n-CMDKeyword              : Checks for CMD Keywords\n-CMDSplCharANDKeyword    : Checks for both and blocks if both are found\n-CMDSplCharORKeyword     : Checks for both and blocks if anyone is found,\n-None                    : Disables checking using both CMD Special Char and Keyword",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"contenttypeaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Content-type actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -contentTypeaction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -contentTypeaction none\".",
			},
			"cookieconsistencyaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "One or more Cookie Consistency actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -cookieConsistencyAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -cookieConsistencyAction none\".",
			},
			"cookieencryption": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Type of cookie encryption. Available settings function as follows:\n* None - Do not encrypt cookies.\n* Decrypt Only - Decrypt encrypted cookies, but do not encrypt cookies.\n* Encrypt Session Only - Encrypt session cookies, but not permanent cookies.\n* Encrypt All - Encrypt all cookies.",
			},
			"cookiehijackingaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "One or more actions to prevent cookie hijacking. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\nNOTE: Cookie Hijacking feature is not supported for TLSv1.3\n\nCLI users: To enable one or more actions, type \"set appfw profile -cookieHijackingAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -cookieHijackingAction none\".",
			},
			"cookieproxying": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Cookie proxy setting. Available settings function as follows:\n* None - Do not proxy cookies.\n* Session Only - Proxy session cookies by using the Citrix ADC session ID, but do not proxy permanent cookies.",
			},
			"cookiesamesiteattribute": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("LAX"),
				Description: "Cookie Samesite attribute added to support adding cookie SameSite attribute for all set-cookies including appfw session cookies. Default value will be \"SameSite=Lax\".",
			},
			"cookietransforms": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform the specified type of cookie transformation.\nAvailable settings function as follows:\n* Encryption - Encrypt cookies.\n* Proxying - Mask contents of server cookies by sending proxy cookie to users.\n* Cookie flags - Flag cookies as HTTP only to prevent scripts on user's browser from accessing and possibly modifying them.\nCAUTION: Make sure that this parameter is set to ON if you are configuring any cookie transformations. If it is set to OFF, no cookie transformations are performed regardless of any other settings.",
			},
			"creditcard": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Credit card types that the application firewall should protect.",
			},
			"creditcardaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "One or more Credit Card actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -creditCardAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -creditCardAction none\".",
			},
			"creditcardmaxallowed": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This parameter value is used by the block action. It represents the maximum number of credit card numbers that can appear on a web page served by your protected web sites. Pages that contain more credit card numbers are blocked.",
			},
			"creditcardxout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mask any credit card number detected in a response by replacing each digit, except the digits in the final group, with the letter \"X.\"",
			},
			"crosssitescriptingaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Cross-Site Scripting (XSS) actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -crossSiteScriptingAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -crossSiteScriptingAction none\".",
			},
			"crosssitescriptingcheckcompleteurls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check complete URLs for cross-site scripts, instead of just the query portions of URLs.",
			},
			"crosssitescriptingtransformunsafehtml": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transform cross-site scripts. This setting configures the application firewall to disable dangerous HTML instead of blocking the request.\nCAUTION: Make sure that this parameter is set to ON if you are configuring any cross-site scripting transformations. If it is set to OFF, no cross-site scripting transformations are performed regardless of any other settings.",
			},
			"csrftagaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "One or more Cross-Site Request Forgery (CSRF) Tagging actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -CSRFTagAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -CSRFTagAction none\".",
			},
			"customsettings": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Object name for custom settings.\nThis check is applicable to Profile Type: HTML, XML.",
			},
			"defaultcharset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default character set for protected web pages. Web pages sent by your protected web sites in response to user requests are assigned this character set if the page does not already specify a character set. The character sets supported by the application firewall are:\n* iso-8859-1 (English US)\n* big5 (Chinese Traditional)\n* gb2312 (Chinese Simplified)\n* sjis (Japanese Shift-JIS)\n* euc-jp (Japanese EUC-JP)\n* iso-8859-9 (Turkish)\n* utf-8 (Unicode)\n* euc-kr (Korean)",
			},
			"defaultfieldformatmaxlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(65535),
				Description: "Maximum length, in characters, for data entered into a field that is assigned the default field type.",
			},
			"defaultfieldformatmaxoccurrences": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(65535),
				Description: "Maxiumum allowed occurrences of the form field name in a request.",
			},
			"defaultfieldformatminlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum length, in characters, for data entered into a field that is assigned the default field type.\nTo disable the minimum and maximum length settings and allow data of any length to be entered into the field, set this parameter to zero (0).",
			},
			"defaultfieldformattype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Designate a default field type to be applied to web form fields that do not have a field type explicitly assigned to them.",
			},
			"defaults": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Default configuration to apply to the profile. Basic defaults are intended for standard content that requires little further configuration, such as static web site content. Advanced defaults are intended for specialized content that requires significant specialized configuration, such as heavily scripted or dynamic content.\n\nCLI users: When adding an application firewall profile, you can set either the defaults or the type, but not both. To set both options, create the profile by using the add appfw profile command, and then use the set appfw profile command to configure the other option.",
			},
			"denyurlaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Deny URL actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nNOTE: The Deny URL check takes precedence over the Start URL check. If you enable blocking for the Deny URL check, the application firewall blocks any URL that is explicitly blocked by a Deny URL, even if the same URL would otherwise be allowed by the Start URL check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -denyURLaction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -denyURLaction none\".",
			},
			"dosecurecreditcardlogging": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Setting this option logs credit card numbers in the response when the match is found.",
			},
			"dynamiclearning": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more security checks. Available options are as follows:\n* SQLInjection - Enable dynamic learning for SQLInjection security check.\n* CrossSiteScripting - Enable dynamic learning for CrossSiteScripting security check.\n* fieldFormat - Enable dynamic learning for  fieldFormat security check.\n* None - Disable security checks for all security checks.\n\nCLI users: To enable dynamic learning on one or more security checks, type \"set appfw profile -dynamicLearning\" followed by the security checks to be enabled. To turn off dynamic learning on all security checks, type \"set appfw profile -dynamicLearning none\".",
			},
			"enableformtagging": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Enable tagging of web form fields for use by the Form Field Consistency and CSRF Form Tagging checks.",
			},
			"errorurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL that application firewall uses as the Error URL.",
			},
			"excludefileuploadfromchecks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exclude uploaded files from Form checks.",
			},
			"exemptclosureurlsfromsecuritychecks": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Exempt URLs that pass the Start URL closure check from SQL injection, cross-site script, field format and field consistency security checks at locations other than headers.",
			},
			"fakeaccountdetection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fake account detection flag : ON/OFF. If set to ON fake account detection in enabled on ADC, if set to OFF fake account detection is disabled.",
			},
			"fieldconsistencyaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "One or more Form Field Consistency actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -fieldConsistencyaction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -fieldConsistencyAction none\".",
			},
			"fieldformataction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Field Format actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of suggested web form fields and field format assignments.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -fieldFormatAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -fieldFormatAction none\".",
			},
			"fieldscan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check if formfield limit scan is ON or OFF.",
			},
			"fieldscanlimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2048),
				Description: "Field scan limit value for HTML",
			},
			"fileuploadmaxnum": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(65535),
				Description: "Maximum allowed number of file uploads per form-submission request. The maximum setting (65535) allows an unlimited number of uploads.",
			},
			"fileuploadtypesaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more file upload types actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -fileUploadTypeAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -fileUploadTypeAction none\".",
			},
			"geolocationlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable Geo-Location Logging in CEF format logs for the profile.",
			},
			"grpcaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "gRPC validation",
			},
			"htmlerrorobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to assign to the HTML Error Object.\nMust begin with a letter, number, or the underscore character \\(_\\), and must contain only letters, numbers, and the hyphen \\(-\\), period \\(.\\) pound \\(\\#\\), space \\( \\), at (@), equals \\(=\\), colon \\(:\\), and underscore characters. Cannot be changed after the HTML error object is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my HTML error object\" or 'my HTML error object'\\).",
			},
			"htmlerrorstatuscode": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(200),
				Description: "Response status code associated with HTML error page. Non-empty HTML error object must be imported to the application firewall profile for the status code.",
			},
			"htmlerrorstatusmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Response status message associated with HTML error page",
			},
			"importprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the profile which will be created/updated to associate the relaxation rules",
			},
			"infercontenttypexmlpayloadaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more infer content type payload actions. Available settings function as follows:\n* Block - Block connections that have mismatch in content-type header and payload.\n* Log - Log connections that have mismatch in content-type header and payload. The mismatched content-type in HTTP request header will be logged for the request.\n* Stats - Generate statistics when there is mismatch in content-type header and payload.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -inferContentTypeXMLPayloadAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -inferContentTypeXMLPayloadAction none\". Please note \"none\" action cannot be used with any other action type.",
			},
			"insertcookiesamesiteattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure whether application firewall should add samesite attribute for set-cookies",
			},
			"inspectcontenttypes": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more InspectContentType lists.\n* application/x-www-form-urlencoded\n* multipart/form-data\n* text/x-gwt-rpc\n\nCLI users: To enable, type \"set appfw profile -InspectContentTypes\" followed by the content types to be inspected.",
			},
			"inspectquerycontenttypes": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Inspect request query as well as web forms for injected SQL and cross-site scripts for following content types.",
			},
			"invalidpercenthandling": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("secure_mode"),
				Description: "Configure the method that the application firewall uses to handle percent-encoded names and values. Available settings function as follows:\n* asp_mode - Microsoft ASP format.\n* secure_mode - Secure format.",
			},
			"jsonblockkeywordaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "JSON Block Keyword action. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONBlockKeywordAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONBlockKeywordAction none\".",
			},
			"jsoncmdinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more JSON CMD Injection actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONCMDInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONCMDInjectionAction none\".",
			},
			"jsoncmdinjectiongrammar": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for CMD injection using CMD grammar in JSON",
			},
			"jsoncmdinjectiontype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("CMDSplCharANDKeyword"),
				Description: "Available CMD injection types.\n-CMDSplChar              : Checks for CMD Special Chars\n-CMDKeyword              : Checks for CMD Keywords\n-CMDSplCharANDKeyword    : Checks for both and blocks if both are found\n-CMDSplCharORKeyword     : Checks for both and blocks if anyone is found,\n-None                    : Disables checking using both SQL Special Char and Keyword",
			},
			"jsondosaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more JSON Denial-of-Service (JsonDoS) actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONDoSAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONDoSAction none\".",
			},
			"jsonerrorobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to the imported JSON Error Object to be set on application firewall profile.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my JSON error object\" or 'my JSON error object'\\).",
			},
			"jsonerrorstatuscode": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(200),
				Description: "Response status code associated with JSON error page. Non-empty JSON error object must be imported to the application firewall profile for the status code.",
			},
			"jsonerrorstatusmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Response status message associated with JSON error page",
			},
			"jsonfieldscan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check if JSON field limit scan is ON or OFF.",
			},
			"jsonfieldscanlimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2048),
				Description: "Field scan limit value for JSON",
			},
			"jsonmessagescan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check if JSON message limit scan is ON or OFF",
			},
			"jsonmessagescanlimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1000000),
				Description: "Message scan limit value for JSON",
			},
			"jsonsqlinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more JSON SQL Injection actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONSQLInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONSQLInjectionAction none\".",
			},
			"jsonsqlinjectiongrammar": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for SQL injection using SQL grammar in JSON",
			},
			"jsonsqlinjectiontype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SQLSplCharANDKeyword"),
				Description: "Available SQL injection types.\n-SQLSplChar              : Checks for SQL Special Chars\n-SQLKeyword              : Checks for SQL Keywords\n-SQLSplCharANDKeyword    : Checks for both and blocks if both are found\n-SQLSplCharORKeyword     : Checks for both and blocks if anyone is found,\n-None                    : Disables checking using both SQL Special Char and Keyword",
			},
			"jsonxssaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more JSON Cross-Site Scripting actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -JSONXssAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -JSONXssAction none\".",
			},
			"logeverypolicyhit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log every profile match, regardless of security checks results.",
			},
			"matchurlstring": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Match this action url in archived Relaxation Rules to replace.",
			},
			"messagescan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check if HTML message limit scan is ON or OFF",
			},
			"messagescanlimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1000000),
				Description: "Message scan limit value for HTML",
			},
			"messagescanlimitcontenttypes": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Enable Message Scan Limit for following content types.",
			},
			"multipleheaderaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more multiple header actions. Available settings function as follows:\n* Block - Block connections that have multiple headers.\n* Log - Log connections that have multiple headers.\n* KeepLast - Keep only last header when multiple headers are present.\n\nRequest headers inspected:\n* Accept-Encoding\n* Content-Encoding\n* Content-Range\n* Content-Type\n* Host\n* Range\n* Referer\n\nCLI users: To enable one or more actions, type \"set appfw profile -multipleHeaderAction\" followed by the actions to be enabled.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"optimizepartialreqs": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Optimize handle of HTTP partial requests i.e. those with range headers.\nAvailable settings are as follows:\n* ON  - Partial requests by the client result in partial requests to the backend server in most cases.\n* OFF - Partial requests by the client are changed to full requests to the backend server",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Purge existing Relaxation Rules and replace during import",
			},
			"percentdecoderecursively": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Configure whether the application firewall should use percentage recursive decoding",
			},
			"postbodylimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(20000000),
				Description: "Maximum allowed HTTP post body size, in bytes. Maximum supported value is 10GB. Citrix recommends enabling streaming option for large values of post body limit (>20MB).",
			},
			"postbodylimitaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Post Body Limit actions. Available settings function as follows:\n* Block - Block connections that violate this security check. Must always be set.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -PostBodyLimitAction block\" followed by the other actions to be enabled.",
			},
			"postbodylimitsignature": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2048),
				Description: "Maximum allowed HTTP post body size for signature inspection for location HTTP_POST_BODY in the signatures, in bytes. Note that the changes in value could impact CPU and latency profile.",
			},
			"protofileobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the imported proto file.",
			},
			"refererheadercheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable validation of Referer headers.\nReferer validation ensures that a web form that a user sends to your web site originally came from your web site, not an outside attacker.\nAlthough this parameter is part of the Start URL check, referer validation protects against cross-site request forgery (CSRF) attacks, not Start URL attacks.",
			},
			"relaxationrules": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Import all appfw relaxation rules",
			},
			"replaceurlstring": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Replace matched url string with this action url string while restoring Relaxation Rules",
			},
			"requestcontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default Content-Type header for requests.\nA Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.",
			},
			"responsecontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default Content-Type header for responses.\nA Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.",
			},
			"restaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "rest validation",
			},
			"rfcprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Object name of the rfc profile.",
			},
			"semicolonfieldseparator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow ';' as a form field separator in URL queries and POST form bodies.",
			},
			"sessioncookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the session cookie that the application firewall uses to track user sessions.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"sessionlessfieldconsistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform sessionless Field Consistency Checks.",
			},
			"sessionlessurlclosure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable session less URL Closure Checks.\nThis check is applicable to Profile Type: HTML.",
			},
			"signatures": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Object name for signatures.\nThis check is applicable to Profile Type: HTML, XML.",
			},
			"sqlinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more HTML SQL Injection actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -SQLInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -SQLInjectionAction none\".",
			},
			"sqlinjectionchecksqlwildchars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for form fields that contain SQL wild chars .",
			},
			"sqlinjectiongrammar": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for SQL injection using SQL grammar",
			},
			"sqlinjectiononlycheckfieldswithsqlchars": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Check only form fields that contain SQL special strings (characters) for injected SQL code.\nMost SQL servers require a special string to activate an SQL request, so SQL code without a special string is harmless to most SQL servers.",
			},
			"sqlinjectionparsecomments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parse HTML comments and exempt them from the HTML SQL Injection check. You must specify the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows:\n* Check all - Check all content.\n* ANSI - Exempt content that is part of an ANSI (Mozilla-style) comment.\n* Nested - Exempt content that is part of a nested (Microsoft-style) comment.\n* ANSI Nested - Exempt content that is part of any type of comment.",
			},
			"sqlinjectionruletype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ALLOW"),
				Description: "Specifies SQL Injection rule type: ALLOW/DENY. If ALLOW rule type is configured then allow list rules are used, if DENY rule type is configured then deny rules are used.",
			},
			"sqlinjectiontransformspecialchars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transform injected SQL code. This setting configures the application firewall to disable SQL special strings instead of blocking the request. Since most SQL servers require a special string to activate an SQL keyword, in most cases a request that contains injected SQL code is safe if special strings are disabled.\nCAUTION: Make sure that this parameter is set to ON if you are configuring any SQL injection transformations. If it is set to OFF, no SQL injection transformations are performed regardless of any other settings.",
			},
			"sqlinjectiontype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SQLSplCharANDKeyword"),
				Description: "Available SQL injection types.\n-SQLSplChar              : Checks for SQL Special Chars\n-SQLKeyword		 : Checks for SQL Keywords\n-SQLSplCharANDKeyword    : Checks for both and blocks if both are found\n-SQLSplCharORKeyword     : Checks for both and blocks if anyone is found\n-None                    : Disables checking using both SQL Special Char and Keyword",
			},
			"starturlaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Start URL actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -startURLaction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -startURLaction none\".",
			},
			"starturlclosure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Toggle  the state of Start URL Closure.",
			},
			"streaming": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting this option converts content-length form submission requests (requests with content-type \"application/x-www-form-urlencoded\" or \"multipart/form-data\") to chunked requests when atleast one of the following protections : Signatures, SQL injection protection, XSS protection, form field consistency protection, starturl closure, CSRF tagging, JSON SQL, JSON XSS, JSON DOS is enabled. Please make sure that the backend server accepts chunked requests before enabling this option. Citrix recommends enabling this option for large request sizes(>20MB).",
			},
			"stripcomments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Strip HTML comments.\nThis check is applicable to Profile Type: HTML.",
			},
			"striphtmlcomments": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Strip HTML comments before forwarding a web page sent by a protected web site in response to a user request.",
			},
			"stripxmlcomments": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Strip XML comments before forwarding a web page sent by a protected web site in response to a user request.",
			},
			"trace": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Toggle  the state of trace",
			},
			"type": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Application firewall profile type, which controls which security checks and settings are applied to content that is filtered with the profile. Available settings function as follows:\n* HTML - HTML-based web sites.\n* XML -  XML-based web sites and services.\n* JSON - JSON-based web sites and services.\n* HTML XML (Web 2.0) - Sites that contain both HTML and XML content, such as ATOM feeds, blogs, and RSS feeds.\n* HTML JSON  - Sites that contain both HTML and JSON content.\n* XML JSON   - Sites that contain both XML and JSON content.\n* HTML XML JSON   - Sites that contain HTML, XML and JSON content.",
			},
			"urldecoderequestcookies": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL Decode request cookies before subjecting them to SQL and cross-site scripting checks.",
			},
			"usehtmlerrorobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send an imported HTML Error object to a user when a request is blocked, instead of redirecting the user to the designated Error URL.",
			},
			"verboseloglevel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("pattern"),
				Description: "Detailed Logging Verbose Log Level.",
			},
			"xmlattachmentaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Attachment actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLAttachmentAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLAttachmentAction none\".",
			},
			"xmldosaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Denial-of-Service (XDoS) actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLDoSAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLDoSAction none\".",
			},
			"xmlerrorobject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to assign to the XML Error Object, which the application firewall displays when a user request is blocked.\nMust begin with a letter, number, or the underscore character \\(_\\), and must contain only letters, numbers, and the hyphen \\(-\\), period \\(.\\) pound \\(\\#\\), space \\( \\), at (@), equals \\(=\\), colon \\(:\\), and underscore characters. Cannot be changed after the XML error object is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my XML error object\" or 'my XML error object'\\).",
			},
			"xmlerrorstatuscode": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(200),
				Description: "Response status code associated with XML error page. Non-empty XML error object must be imported to the application firewall profile for the status code.",
			},
			"xmlerrorstatusmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Response status message associated with XML error page",
			},
			"xmlformataction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Format actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLFormatAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLFormatAction none\".",
			},
			"xmlsoapfaultaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML SOAP Fault Filtering actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n* Remove - Remove all violations for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLSOAPFaultAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLSOAPFaultAction none\".",
			},
			"xmlsqlinjectionaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML SQL Injection actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLSQLInjectionAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLSQLInjectionAction none\".",
			},
			"xmlsqlinjectionchecksqlwildchars": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Check for form fields that contain SQL wild chars .",
			},
			"xmlsqlinjectiononlycheckfieldswithsqlchars": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Check only form fields that contain SQL special characters, which most SQL servers require before accepting an SQL command, for injected SQL.",
			},
			"xmlsqlinjectionparsecomments": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("checkall"),
				Description: "Parse comments in XML Data and exempt those sections of the request that are from the XML SQL Injection check. You must configure the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows:\n* Check all - Check all content.\n* ANSI - Exempt content that is part of an ANSI (Mozilla-style) comment.\n* Nested - Exempt content that is part of a nested (Microsoft-style) comment.\n* ANSI Nested - Exempt content that is part of any type of comment.",
			},
			"xmlsqlinjectiontype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SQLSplCharANDKeyword"),
				Description: "Available SQL injection types.\n-SQLSplChar              : Checks for SQL Special Chars\n-SQLKeyword              : Checks for SQL Keywords\n-SQLSplCharANDKeyword    : Checks for both and blocks if both are found\n-SQLSplCharORKeyword     : Checks for both and blocks if anyone is found",
			},
			"xmlvalidationaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Validation actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLValidationAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLValidationAction none\".",
			},
			"xmlwsiaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more Web Services Interoperability (WSI) actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Learn - Use the learning engine to generate a list of exceptions to this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLWSIAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLWSIAction none\".",
			},
			"xmlxssaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more XML Cross-Site Scripting actions. Available settings function as follows:\n* Block - Block connections that violate this security check.\n* Log - Log violations of this security check.\n* Stats - Generate statistics for this security check.\n* None - Disable all actions for this security check.\n\nCLI users: To enable one or more actions, type \"set appfw profile -XMLXSSAction\" followed by the actions to be enabled. To turn off all actions, type \"set appfw profile -XMLXSSAction none\".",
			},
		},
	}
}

func appfwprofileGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileResourceModel) appfw.Appfwprofile {
	tflog.Debug(ctx, "In appfwprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile := appfw.Appfwprofile{}
	if !data.Addcookieflags.IsNull() {
		appfwprofile.Addcookieflags = data.Addcookieflags.ValueString()
	}
	if !data.Apispec.IsNull() {
		appfwprofile.Apispec = data.Apispec.ValueString()
	}
	if !data.Archivename.IsNull() {
		appfwprofile.Archivename = data.Archivename.ValueString()
	}
	if !data.AsProfBypassListEnable.IsNull() {
		appfwprofile.Asprofbypasslistenable = data.AsProfBypassListEnable.ValueString()
	}
	if !data.AsProfDenyListEnable.IsNull() {
		appfwprofile.Asprofdenylistenable = data.AsProfDenyListEnable.ValueString()
	}
	if !data.Augment.IsNull() {
		appfwprofile.Augment = data.Augment.ValueBool()
	}
	if !data.Bufferoverflowmaxcookielength.IsNull() {
		appfwprofile.Bufferoverflowmaxcookielength = utils.IntPtr(int(data.Bufferoverflowmaxcookielength.ValueInt64()))
	}
	if !data.Bufferoverflowmaxheaderlength.IsNull() {
		appfwprofile.Bufferoverflowmaxheaderlength = utils.IntPtr(int(data.Bufferoverflowmaxheaderlength.ValueInt64()))
	}
	if !data.Bufferoverflowmaxquerylength.IsNull() {
		appfwprofile.Bufferoverflowmaxquerylength = utils.IntPtr(int(data.Bufferoverflowmaxquerylength.ValueInt64()))
	}
	if !data.Bufferoverflowmaxtotalheaderlength.IsNull() {
		appfwprofile.Bufferoverflowmaxtotalheaderlength = utils.IntPtr(int(data.Bufferoverflowmaxtotalheaderlength.ValueInt64()))
	}
	if !data.Bufferoverflowmaxurllength.IsNull() {
		appfwprofile.Bufferoverflowmaxurllength = utils.IntPtr(int(data.Bufferoverflowmaxurllength.ValueInt64()))
	}
	if !data.Canonicalizehtmlresponse.IsNull() {
		appfwprofile.Canonicalizehtmlresponse = data.Canonicalizehtmlresponse.ValueString()
	}
	if !data.Ceflogging.IsNull() {
		appfwprofile.Ceflogging = data.Ceflogging.ValueString()
	}
	if !data.Checkrequestheaders.IsNull() {
		appfwprofile.Checkrequestheaders = data.Checkrequestheaders.ValueString()
	}
	if !data.Clientipexpression.IsNull() {
		appfwprofile.Clientipexpression = data.Clientipexpression.ValueString()
	}
	if !data.Cmdinjectiongrammar.IsNull() {
		appfwprofile.Cmdinjectiongrammar = data.Cmdinjectiongrammar.ValueString()
	}
	if !data.Cmdinjectiontype.IsNull() {
		appfwprofile.Cmdinjectiontype = data.Cmdinjectiontype.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile.Comment = data.Comment.ValueString()
	}
	if !data.Cookieencryption.IsNull() {
		appfwprofile.Cookieencryption = data.Cookieencryption.ValueString()
	}
	if !data.Cookieproxying.IsNull() {
		appfwprofile.Cookieproxying = data.Cookieproxying.ValueString()
	}
	if !data.Cookiesamesiteattribute.IsNull() {
		appfwprofile.Cookiesamesiteattribute = data.Cookiesamesiteattribute.ValueString()
	}
	if !data.Cookietransforms.IsNull() {
		appfwprofile.Cookietransforms = data.Cookietransforms.ValueString()
	}
	if !data.Creditcardmaxallowed.IsNull() {
		appfwprofile.Creditcardmaxallowed = utils.IntPtr(int(data.Creditcardmaxallowed.ValueInt64()))
	}
	if !data.Creditcardxout.IsNull() {
		appfwprofile.Creditcardxout = data.Creditcardxout.ValueString()
	}
	if !data.Crosssitescriptingcheckcompleteurls.IsNull() {
		appfwprofile.Crosssitescriptingcheckcompleteurls = data.Crosssitescriptingcheckcompleteurls.ValueString()
	}
	if !data.Crosssitescriptingtransformunsafehtml.IsNull() {
		appfwprofile.Crosssitescriptingtransformunsafehtml = data.Crosssitescriptingtransformunsafehtml.ValueString()
	}
	if !data.Customsettings.IsNull() {
		appfwprofile.Customsettings = data.Customsettings.ValueString()
	}
	if !data.Defaultcharset.IsNull() {
		appfwprofile.Defaultcharset = data.Defaultcharset.ValueString()
	}
	if !data.Defaultfieldformatmaxlength.IsNull() {
		appfwprofile.Defaultfieldformatmaxlength = utils.IntPtr(int(data.Defaultfieldformatmaxlength.ValueInt64()))
	}
	if !data.Defaultfieldformatmaxoccurrences.IsNull() {
		appfwprofile.Defaultfieldformatmaxoccurrences = utils.IntPtr(int(data.Defaultfieldformatmaxoccurrences.ValueInt64()))
	}
	if !data.Defaultfieldformatminlength.IsNull() {
		appfwprofile.Defaultfieldformatminlength = utils.IntPtr(int(data.Defaultfieldformatminlength.ValueInt64()))
	}
	if !data.Defaultfieldformattype.IsNull() {
		appfwprofile.Defaultfieldformattype = data.Defaultfieldformattype.ValueString()
	}
	if !data.Defaults.IsNull() {
		appfwprofile.Defaults = data.Defaults.ValueString()
	}
	if !data.Dosecurecreditcardlogging.IsNull() {
		appfwprofile.Dosecurecreditcardlogging = data.Dosecurecreditcardlogging.ValueString()
	}
	if !data.Enableformtagging.IsNull() {
		appfwprofile.Enableformtagging = data.Enableformtagging.ValueString()
	}
	if !data.Errorurl.IsNull() {
		appfwprofile.Errorurl = data.Errorurl.ValueString()
	}
	if !data.Excludefileuploadfromchecks.IsNull() {
		appfwprofile.Excludefileuploadfromchecks = data.Excludefileuploadfromchecks.ValueString()
	}
	if !data.Exemptclosureurlsfromsecuritychecks.IsNull() {
		appfwprofile.Exemptclosureurlsfromsecuritychecks = data.Exemptclosureurlsfromsecuritychecks.ValueString()
	}
	if !data.Fakeaccountdetection.IsNull() {
		appfwprofile.Fakeaccountdetection = data.Fakeaccountdetection.ValueString()
	}
	if !data.Fieldscan.IsNull() {
		appfwprofile.Fieldscan = data.Fieldscan.ValueString()
	}
	if !data.Fieldscanlimit.IsNull() {
		appfwprofile.Fieldscanlimit = utils.IntPtr(int(data.Fieldscanlimit.ValueInt64()))
	}
	if !data.Fileuploadmaxnum.IsNull() {
		appfwprofile.Fileuploadmaxnum = utils.IntPtr(int(data.Fileuploadmaxnum.ValueInt64()))
	}
	if !data.Geolocationlogging.IsNull() {
		appfwprofile.Geolocationlogging = data.Geolocationlogging.ValueString()
	}
	if !data.Htmlerrorobject.IsNull() {
		appfwprofile.Htmlerrorobject = data.Htmlerrorobject.ValueString()
	}
	if !data.Htmlerrorstatuscode.IsNull() {
		appfwprofile.Htmlerrorstatuscode = utils.IntPtr(int(data.Htmlerrorstatuscode.ValueInt64()))
	}
	if !data.Htmlerrorstatusmessage.IsNull() {
		appfwprofile.Htmlerrorstatusmessage = data.Htmlerrorstatusmessage.ValueString()
	}
	if !data.Importprofilename.IsNull() {
		appfwprofile.Importprofilename = data.Importprofilename.ValueString()
	}
	if !data.Insertcookiesamesiteattribute.IsNull() {
		appfwprofile.Insertcookiesamesiteattribute = data.Insertcookiesamesiteattribute.ValueString()
	}
	if !data.Invalidpercenthandling.IsNull() {
		appfwprofile.Invalidpercenthandling = data.Invalidpercenthandling.ValueString()
	}
	if !data.Jsoncmdinjectiongrammar.IsNull() {
		appfwprofile.Jsoncmdinjectiongrammar = data.Jsoncmdinjectiongrammar.ValueString()
	}
	if !data.Jsoncmdinjectiontype.IsNull() {
		appfwprofile.Jsoncmdinjectiontype = data.Jsoncmdinjectiontype.ValueString()
	}
	if !data.Jsonerrorobject.IsNull() {
		appfwprofile.Jsonerrorobject = data.Jsonerrorobject.ValueString()
	}
	if !data.Jsonerrorstatuscode.IsNull() {
		appfwprofile.Jsonerrorstatuscode = utils.IntPtr(int(data.Jsonerrorstatuscode.ValueInt64()))
	}
	if !data.Jsonerrorstatusmessage.IsNull() {
		appfwprofile.Jsonerrorstatusmessage = data.Jsonerrorstatusmessage.ValueString()
	}
	if !data.Jsonfieldscan.IsNull() {
		appfwprofile.Jsonfieldscan = data.Jsonfieldscan.ValueString()
	}
	if !data.Jsonfieldscanlimit.IsNull() {
		appfwprofile.Jsonfieldscanlimit = utils.IntPtr(int(data.Jsonfieldscanlimit.ValueInt64()))
	}
	if !data.Jsonmessagescan.IsNull() {
		appfwprofile.Jsonmessagescan = data.Jsonmessagescan.ValueString()
	}
	if !data.Jsonmessagescanlimit.IsNull() {
		appfwprofile.Jsonmessagescanlimit = utils.IntPtr(int(data.Jsonmessagescanlimit.ValueInt64()))
	}
	if !data.Jsonsqlinjectiongrammar.IsNull() {
		appfwprofile.Jsonsqlinjectiongrammar = data.Jsonsqlinjectiongrammar.ValueString()
	}
	if !data.Jsonsqlinjectiontype.IsNull() {
		appfwprofile.Jsonsqlinjectiontype = data.Jsonsqlinjectiontype.ValueString()
	}
	if !data.Logeverypolicyhit.IsNull() {
		appfwprofile.Logeverypolicyhit = data.Logeverypolicyhit.ValueString()
	}
	if !data.Matchurlstring.IsNull() {
		appfwprofile.Matchurlstring = data.Matchurlstring.ValueString()
	}
	if !data.Messagescan.IsNull() {
		appfwprofile.Messagescan = data.Messagescan.ValueString()
	}
	if !data.Messagescanlimit.IsNull() {
		appfwprofile.Messagescanlimit = utils.IntPtr(int(data.Messagescanlimit.ValueInt64()))
	}
	if !data.Name.IsNull() {
		appfwprofile.Name = data.Name.ValueString()
	}
	if !data.Optimizepartialreqs.IsNull() {
		appfwprofile.Optimizepartialreqs = data.Optimizepartialreqs.ValueString()
	}
	if !data.Overwrite.IsNull() {
		appfwprofile.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Percentdecoderecursively.IsNull() {
		appfwprofile.Percentdecoderecursively = data.Percentdecoderecursively.ValueString()
	}
	if !data.Postbodylimit.IsNull() {
		appfwprofile.Postbodylimit = utils.IntPtr(int(data.Postbodylimit.ValueInt64()))
	}
	if !data.Postbodylimitsignature.IsNull() {
		appfwprofile.Postbodylimitsignature = utils.IntPtr(int(data.Postbodylimitsignature.ValueInt64()))
	}
	if !data.Protofileobject.IsNull() {
		appfwprofile.Protofileobject = data.Protofileobject.ValueString()
	}
	if !data.Refererheadercheck.IsNull() {
		appfwprofile.Refererheadercheck = data.Refererheadercheck.ValueString()
	}
	if !data.Relaxationrules.IsNull() {
		appfwprofile.Relaxationrules = data.Relaxationrules.ValueBool()
	}
	if !data.Replaceurlstring.IsNull() {
		appfwprofile.Replaceurlstring = data.Replaceurlstring.ValueString()
	}
	if !data.Requestcontenttype.IsNull() {
		appfwprofile.Requestcontenttype = data.Requestcontenttype.ValueString()
	}
	if !data.Responsecontenttype.IsNull() {
		appfwprofile.Responsecontenttype = data.Responsecontenttype.ValueString()
	}
	if !data.Rfcprofile.IsNull() {
		appfwprofile.Rfcprofile = data.Rfcprofile.ValueString()
	}
	if !data.Semicolonfieldseparator.IsNull() {
		appfwprofile.Semicolonfieldseparator = data.Semicolonfieldseparator.ValueString()
	}
	if !data.Sessioncookiename.IsNull() {
		appfwprofile.Sessioncookiename = data.Sessioncookiename.ValueString()
	}
	if !data.Sessionlessfieldconsistency.IsNull() {
		appfwprofile.Sessionlessfieldconsistency = data.Sessionlessfieldconsistency.ValueString()
	}
	if !data.Sessionlessurlclosure.IsNull() {
		appfwprofile.Sessionlessurlclosure = data.Sessionlessurlclosure.ValueString()
	}
	if !data.Signatures.IsNull() {
		appfwprofile.Signatures = data.Signatures.ValueString()
	}
	if !data.Sqlinjectionchecksqlwildchars.IsNull() {
		appfwprofile.Sqlinjectionchecksqlwildchars = data.Sqlinjectionchecksqlwildchars.ValueString()
	}
	if !data.Sqlinjectiongrammar.IsNull() {
		appfwprofile.Sqlinjectiongrammar = data.Sqlinjectiongrammar.ValueString()
	}
	if !data.Sqlinjectiononlycheckfieldswithsqlchars.IsNull() {
		appfwprofile.Sqlinjectiononlycheckfieldswithsqlchars = data.Sqlinjectiononlycheckfieldswithsqlchars.ValueString()
	}
	if !data.Sqlinjectionparsecomments.IsNull() {
		appfwprofile.Sqlinjectionparsecomments = data.Sqlinjectionparsecomments.ValueString()
	}
	if !data.Sqlinjectionruletype.IsNull() {
		appfwprofile.Sqlinjectionruletype = data.Sqlinjectionruletype.ValueString()
	}
	if !data.Sqlinjectiontransformspecialchars.IsNull() {
		appfwprofile.Sqlinjectiontransformspecialchars = data.Sqlinjectiontransformspecialchars.ValueString()
	}
	if !data.Sqlinjectiontype.IsNull() {
		appfwprofile.Sqlinjectiontype = data.Sqlinjectiontype.ValueString()
	}
	if !data.Starturlclosure.IsNull() {
		appfwprofile.Starturlclosure = data.Starturlclosure.ValueString()
	}
	if !data.Streaming.IsNull() {
		appfwprofile.Streaming = data.Streaming.ValueString()
	}
	if !data.Stripcomments.IsNull() {
		appfwprofile.Stripcomments = data.Stripcomments.ValueString()
	}
	if !data.Striphtmlcomments.IsNull() {
		appfwprofile.Striphtmlcomments = data.Striphtmlcomments.ValueString()
	}
	if !data.Stripxmlcomments.IsNull() {
		appfwprofile.Stripxmlcomments = data.Stripxmlcomments.ValueString()
	}
	if !data.Trace.IsNull() {
		appfwprofile.Trace = data.Trace.ValueString()
	}
	if !data.Urldecoderequestcookies.IsNull() {
		appfwprofile.Urldecoderequestcookies = data.Urldecoderequestcookies.ValueString()
	}
	if !data.Usehtmlerrorobject.IsNull() {
		appfwprofile.Usehtmlerrorobject = data.Usehtmlerrorobject.ValueString()
	}
	if !data.Verboseloglevel.IsNull() {
		appfwprofile.Verboseloglevel = data.Verboseloglevel.ValueString()
	}
	if !data.Xmlerrorobject.IsNull() {
		appfwprofile.Xmlerrorobject = data.Xmlerrorobject.ValueString()
	}
	if !data.Xmlerrorstatuscode.IsNull() {
		appfwprofile.Xmlerrorstatuscode = utils.IntPtr(int(data.Xmlerrorstatuscode.ValueInt64()))
	}
	if !data.Xmlerrorstatusmessage.IsNull() {
		appfwprofile.Xmlerrorstatusmessage = data.Xmlerrorstatusmessage.ValueString()
	}
	if !data.Xmlsqlinjectionchecksqlwildchars.IsNull() {
		appfwprofile.Xmlsqlinjectionchecksqlwildchars = data.Xmlsqlinjectionchecksqlwildchars.ValueString()
	}
	if !data.Xmlsqlinjectiononlycheckfieldswithsqlchars.IsNull() {
		appfwprofile.Xmlsqlinjectiononlycheckfieldswithsqlchars = data.Xmlsqlinjectiononlycheckfieldswithsqlchars.ValueString()
	}
	if !data.Xmlsqlinjectionparsecomments.IsNull() {
		appfwprofile.Xmlsqlinjectionparsecomments = data.Xmlsqlinjectionparsecomments.ValueString()
	}
	if !data.Xmlsqlinjectiontype.IsNull() {
		appfwprofile.Xmlsqlinjectiontype = data.Xmlsqlinjectiontype.ValueString()
	}

	return appfwprofile
}

func appfwprofileSetAttrFromGet(ctx context.Context, data *AppfwprofileResourceModel, getResponseData map[string]interface{}) *AppfwprofileResourceModel {
	tflog.Debug(ctx, "In appfwprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["addcookieflags"]; ok && val != nil {
		data.Addcookieflags = types.StringValue(val.(string))
	} else {
		data.Addcookieflags = types.StringNull()
	}
	if val, ok := getResponseData["apispec"]; ok && val != nil {
		data.Apispec = types.StringValue(val.(string))
	} else {
		data.Apispec = types.StringNull()
	}
	if val, ok := getResponseData["archivename"]; ok && val != nil {
		data.Archivename = types.StringValue(val.(string))
	} else {
		data.Archivename = types.StringNull()
	}
	if val, ok := getResponseData["as_prof_bypass_list_enable"]; ok && val != nil {
		data.AsProfBypassListEnable = types.StringValue(val.(string))
	} else {
		data.AsProfBypassListEnable = types.StringNull()
	}
	if val, ok := getResponseData["as_prof_deny_list_enable"]; ok && val != nil {
		data.AsProfDenyListEnable = types.StringValue(val.(string))
	} else {
		data.AsProfDenyListEnable = types.StringNull()
	}
	if val, ok := getResponseData["augment"]; ok && val != nil {
		data.Augment = types.BoolValue(val.(bool))
	} else {
		data.Augment = types.BoolNull()
	}
	if val, ok := getResponseData["bufferoverflowmaxcookielength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bufferoverflowmaxcookielength = types.Int64Value(intVal)
		}
	} else {
		data.Bufferoverflowmaxcookielength = types.Int64Null()
	}
	if val, ok := getResponseData["bufferoverflowmaxheaderlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bufferoverflowmaxheaderlength = types.Int64Value(intVal)
		}
	} else {
		data.Bufferoverflowmaxheaderlength = types.Int64Null()
	}
	if val, ok := getResponseData["bufferoverflowmaxquerylength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bufferoverflowmaxquerylength = types.Int64Value(intVal)
		}
	} else {
		data.Bufferoverflowmaxquerylength = types.Int64Null()
	}
	if val, ok := getResponseData["bufferoverflowmaxtotalheaderlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bufferoverflowmaxtotalheaderlength = types.Int64Value(intVal)
		}
	} else {
		data.Bufferoverflowmaxtotalheaderlength = types.Int64Null()
	}
	if val, ok := getResponseData["bufferoverflowmaxurllength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bufferoverflowmaxurllength = types.Int64Value(intVal)
		}
	} else {
		data.Bufferoverflowmaxurllength = types.Int64Null()
	}
	if val, ok := getResponseData["canonicalizehtmlresponse"]; ok && val != nil {
		data.Canonicalizehtmlresponse = types.StringValue(val.(string))
	} else {
		data.Canonicalizehtmlresponse = types.StringNull()
	}
	if val, ok := getResponseData["ceflogging"]; ok && val != nil {
		data.Ceflogging = types.StringValue(val.(string))
	} else {
		data.Ceflogging = types.StringNull()
	}
	if val, ok := getResponseData["checkrequestheaders"]; ok && val != nil {
		data.Checkrequestheaders = types.StringValue(val.(string))
	} else {
		data.Checkrequestheaders = types.StringNull()
	}
	if val, ok := getResponseData["clientipexpression"]; ok && val != nil {
		data.Clientipexpression = types.StringValue(val.(string))
	} else {
		data.Clientipexpression = types.StringNull()
	}
	if val, ok := getResponseData["cmdinjectiongrammar"]; ok && val != nil {
		data.Cmdinjectiongrammar = types.StringValue(val.(string))
	} else {
		data.Cmdinjectiongrammar = types.StringNull()
	}
	if val, ok := getResponseData["cmdinjectiontype"]; ok && val != nil {
		data.Cmdinjectiontype = types.StringValue(val.(string))
	} else {
		data.Cmdinjectiontype = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["cookieencryption"]; ok && val != nil {
		data.Cookieencryption = types.StringValue(val.(string))
	} else {
		data.Cookieencryption = types.StringNull()
	}
	if val, ok := getResponseData["cookieproxying"]; ok && val != nil {
		data.Cookieproxying = types.StringValue(val.(string))
	} else {
		data.Cookieproxying = types.StringNull()
	}
	if val, ok := getResponseData["cookiesamesiteattribute"]; ok && val != nil {
		data.Cookiesamesiteattribute = types.StringValue(val.(string))
	} else {
		data.Cookiesamesiteattribute = types.StringNull()
	}
	if val, ok := getResponseData["cookietransforms"]; ok && val != nil {
		data.Cookietransforms = types.StringValue(val.(string))
	} else {
		data.Cookietransforms = types.StringNull()
	}
	if val, ok := getResponseData["creditcardmaxallowed"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Creditcardmaxallowed = types.Int64Value(intVal)
		}
	} else {
		data.Creditcardmaxallowed = types.Int64Null()
	}
	if val, ok := getResponseData["creditcardxout"]; ok && val != nil {
		data.Creditcardxout = types.StringValue(val.(string))
	} else {
		data.Creditcardxout = types.StringNull()
	}
	if val, ok := getResponseData["crosssitescriptingcheckcompleteurls"]; ok && val != nil {
		data.Crosssitescriptingcheckcompleteurls = types.StringValue(val.(string))
	} else {
		data.Crosssitescriptingcheckcompleteurls = types.StringNull()
	}
	if val, ok := getResponseData["crosssitescriptingtransformunsafehtml"]; ok && val != nil {
		data.Crosssitescriptingtransformunsafehtml = types.StringValue(val.(string))
	} else {
		data.Crosssitescriptingtransformunsafehtml = types.StringNull()
	}
	if val, ok := getResponseData["customsettings"]; ok && val != nil {
		data.Customsettings = types.StringValue(val.(string))
	} else {
		data.Customsettings = types.StringNull()
	}
	if val, ok := getResponseData["defaultcharset"]; ok && val != nil {
		data.Defaultcharset = types.StringValue(val.(string))
	} else {
		data.Defaultcharset = types.StringNull()
	}
	if val, ok := getResponseData["defaultfieldformatmaxlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Defaultfieldformatmaxlength = types.Int64Value(intVal)
		}
	} else {
		data.Defaultfieldformatmaxlength = types.Int64Null()
	}
	if val, ok := getResponseData["defaultfieldformatmaxoccurrences"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Defaultfieldformatmaxoccurrences = types.Int64Value(intVal)
		}
	} else {
		data.Defaultfieldformatmaxoccurrences = types.Int64Null()
	}
	if val, ok := getResponseData["defaultfieldformatminlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Defaultfieldformatminlength = types.Int64Value(intVal)
		}
	} else {
		data.Defaultfieldformatminlength = types.Int64Null()
	}
	if val, ok := getResponseData["defaultfieldformattype"]; ok && val != nil {
		data.Defaultfieldformattype = types.StringValue(val.(string))
	} else {
		data.Defaultfieldformattype = types.StringNull()
	}
	if val, ok := getResponseData["defaults"]; ok && val != nil {
		data.Defaults = types.StringValue(val.(string))
	} else {
		data.Defaults = types.StringNull()
	}
	if val, ok := getResponseData["dosecurecreditcardlogging"]; ok && val != nil {
		data.Dosecurecreditcardlogging = types.StringValue(val.(string))
	} else {
		data.Dosecurecreditcardlogging = types.StringNull()
	}
	if val, ok := getResponseData["enableformtagging"]; ok && val != nil {
		data.Enableformtagging = types.StringValue(val.(string))
	} else {
		data.Enableformtagging = types.StringNull()
	}
	if val, ok := getResponseData["errorurl"]; ok && val != nil {
		data.Errorurl = types.StringValue(val.(string))
	} else {
		data.Errorurl = types.StringNull()
	}
	if val, ok := getResponseData["excludefileuploadfromchecks"]; ok && val != nil {
		data.Excludefileuploadfromchecks = types.StringValue(val.(string))
	} else {
		data.Excludefileuploadfromchecks = types.StringNull()
	}
	if val, ok := getResponseData["exemptclosureurlsfromsecuritychecks"]; ok && val != nil {
		data.Exemptclosureurlsfromsecuritychecks = types.StringValue(val.(string))
	} else {
		data.Exemptclosureurlsfromsecuritychecks = types.StringNull()
	}
	if val, ok := getResponseData["fakeaccountdetection"]; ok && val != nil {
		data.Fakeaccountdetection = types.StringValue(val.(string))
	} else {
		data.Fakeaccountdetection = types.StringNull()
	}
	if val, ok := getResponseData["fieldscan"]; ok && val != nil {
		data.Fieldscan = types.StringValue(val.(string))
	} else {
		data.Fieldscan = types.StringNull()
	}
	if val, ok := getResponseData["fieldscanlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fieldscanlimit = types.Int64Value(intVal)
		}
	} else {
		data.Fieldscanlimit = types.Int64Null()
	}
	if val, ok := getResponseData["fileuploadmaxnum"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Fileuploadmaxnum = types.Int64Value(intVal)
		}
	} else {
		data.Fileuploadmaxnum = types.Int64Null()
	}
	if val, ok := getResponseData["geolocationlogging"]; ok && val != nil {
		data.Geolocationlogging = types.StringValue(val.(string))
	} else {
		data.Geolocationlogging = types.StringNull()
	}
	if val, ok := getResponseData["htmlerrorobject"]; ok && val != nil {
		data.Htmlerrorobject = types.StringValue(val.(string))
	} else {
		data.Htmlerrorobject = types.StringNull()
	}
	if val, ok := getResponseData["htmlerrorstatuscode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Htmlerrorstatuscode = types.Int64Value(intVal)
		}
	} else {
		data.Htmlerrorstatuscode = types.Int64Null()
	}
	if val, ok := getResponseData["htmlerrorstatusmessage"]; ok && val != nil {
		data.Htmlerrorstatusmessage = types.StringValue(val.(string))
	} else {
		data.Htmlerrorstatusmessage = types.StringNull()
	}
	if val, ok := getResponseData["importprofilename"]; ok && val != nil {
		data.Importprofilename = types.StringValue(val.(string))
	} else {
		data.Importprofilename = types.StringNull()
	}
	if val, ok := getResponseData["insertcookiesamesiteattribute"]; ok && val != nil {
		data.Insertcookiesamesiteattribute = types.StringValue(val.(string))
	} else {
		data.Insertcookiesamesiteattribute = types.StringNull()
	}
	if val, ok := getResponseData["invalidpercenthandling"]; ok && val != nil {
		data.Invalidpercenthandling = types.StringValue(val.(string))
	} else {
		data.Invalidpercenthandling = types.StringNull()
	}
	if val, ok := getResponseData["jsoncmdinjectiongrammar"]; ok && val != nil {
		data.Jsoncmdinjectiongrammar = types.StringValue(val.(string))
	} else {
		data.Jsoncmdinjectiongrammar = types.StringNull()
	}
	if val, ok := getResponseData["jsoncmdinjectiontype"]; ok && val != nil {
		data.Jsoncmdinjectiontype = types.StringValue(val.(string))
	} else {
		data.Jsoncmdinjectiontype = types.StringNull()
	}
	if val, ok := getResponseData["jsonerrorobject"]; ok && val != nil {
		data.Jsonerrorobject = types.StringValue(val.(string))
	} else {
		data.Jsonerrorobject = types.StringNull()
	}
	if val, ok := getResponseData["jsonerrorstatuscode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsonerrorstatuscode = types.Int64Value(intVal)
		}
	} else {
		data.Jsonerrorstatuscode = types.Int64Null()
	}
	if val, ok := getResponseData["jsonerrorstatusmessage"]; ok && val != nil {
		data.Jsonerrorstatusmessage = types.StringValue(val.(string))
	} else {
		data.Jsonerrorstatusmessage = types.StringNull()
	}
	if val, ok := getResponseData["jsonfieldscan"]; ok && val != nil {
		data.Jsonfieldscan = types.StringValue(val.(string))
	} else {
		data.Jsonfieldscan = types.StringNull()
	}
	if val, ok := getResponseData["jsonfieldscanlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsonfieldscanlimit = types.Int64Value(intVal)
		}
	} else {
		data.Jsonfieldscanlimit = types.Int64Null()
	}
	if val, ok := getResponseData["jsonmessagescan"]; ok && val != nil {
		data.Jsonmessagescan = types.StringValue(val.(string))
	} else {
		data.Jsonmessagescan = types.StringNull()
	}
	if val, ok := getResponseData["jsonmessagescanlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsonmessagescanlimit = types.Int64Value(intVal)
		}
	} else {
		data.Jsonmessagescanlimit = types.Int64Null()
	}
	if val, ok := getResponseData["jsonsqlinjectiongrammar"]; ok && val != nil {
		data.Jsonsqlinjectiongrammar = types.StringValue(val.(string))
	} else {
		data.Jsonsqlinjectiongrammar = types.StringNull()
	}
	if val, ok := getResponseData["jsonsqlinjectiontype"]; ok && val != nil {
		data.Jsonsqlinjectiontype = types.StringValue(val.(string))
	} else {
		data.Jsonsqlinjectiontype = types.StringNull()
	}
	if val, ok := getResponseData["logeverypolicyhit"]; ok && val != nil {
		data.Logeverypolicyhit = types.StringValue(val.(string))
	} else {
		data.Logeverypolicyhit = types.StringNull()
	}
	if val, ok := getResponseData["matchurlstring"]; ok && val != nil {
		data.Matchurlstring = types.StringValue(val.(string))
	} else {
		data.Matchurlstring = types.StringNull()
	}
	if val, ok := getResponseData["messagescan"]; ok && val != nil {
		data.Messagescan = types.StringValue(val.(string))
	} else {
		data.Messagescan = types.StringNull()
	}
	if val, ok := getResponseData["messagescanlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Messagescanlimit = types.Int64Value(intVal)
		}
	} else {
		data.Messagescanlimit = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["optimizepartialreqs"]; ok && val != nil {
		data.Optimizepartialreqs = types.StringValue(val.(string))
	} else {
		data.Optimizepartialreqs = types.StringNull()
	}
	if val, ok := getResponseData["overwrite"]; ok && val != nil {
		data.Overwrite = types.BoolValue(val.(bool))
	} else {
		data.Overwrite = types.BoolNull()
	}
	if val, ok := getResponseData["percentdecoderecursively"]; ok && val != nil {
		data.Percentdecoderecursively = types.StringValue(val.(string))
	} else {
		data.Percentdecoderecursively = types.StringNull()
	}
	if val, ok := getResponseData["postbodylimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Postbodylimit = types.Int64Value(intVal)
		}
	} else {
		data.Postbodylimit = types.Int64Null()
	}
	if val, ok := getResponseData["postbodylimitsignature"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Postbodylimitsignature = types.Int64Value(intVal)
		}
	} else {
		data.Postbodylimitsignature = types.Int64Null()
	}
	if val, ok := getResponseData["protofileobject"]; ok && val != nil {
		data.Protofileobject = types.StringValue(val.(string))
	} else {
		data.Protofileobject = types.StringNull()
	}
	if val, ok := getResponseData["refererheadercheck"]; ok && val != nil {
		data.Refererheadercheck = types.StringValue(val.(string))
	} else {
		data.Refererheadercheck = types.StringNull()
	}
	if val, ok := getResponseData["relaxationrules"]; ok && val != nil {
		data.Relaxationrules = types.BoolValue(val.(bool))
	} else {
		data.Relaxationrules = types.BoolNull()
	}
	if val, ok := getResponseData["replaceurlstring"]; ok && val != nil {
		data.Replaceurlstring = types.StringValue(val.(string))
	} else {
		data.Replaceurlstring = types.StringNull()
	}
	if val, ok := getResponseData["requestcontenttype"]; ok && val != nil {
		data.Requestcontenttype = types.StringValue(val.(string))
	} else {
		data.Requestcontenttype = types.StringNull()
	}
	if val, ok := getResponseData["responsecontenttype"]; ok && val != nil {
		data.Responsecontenttype = types.StringValue(val.(string))
	} else {
		data.Responsecontenttype = types.StringNull()
	}
	if val, ok := getResponseData["rfcprofile"]; ok && val != nil {
		data.Rfcprofile = types.StringValue(val.(string))
	} else {
		data.Rfcprofile = types.StringNull()
	}
	if val, ok := getResponseData["semicolonfieldseparator"]; ok && val != nil {
		data.Semicolonfieldseparator = types.StringValue(val.(string))
	} else {
		data.Semicolonfieldseparator = types.StringNull()
	}
	if val, ok := getResponseData["sessioncookiename"]; ok && val != nil {
		data.Sessioncookiename = types.StringValue(val.(string))
	} else {
		data.Sessioncookiename = types.StringNull()
	}
	if val, ok := getResponseData["sessionlessfieldconsistency"]; ok && val != nil {
		data.Sessionlessfieldconsistency = types.StringValue(val.(string))
	} else {
		data.Sessionlessfieldconsistency = types.StringNull()
	}
	if val, ok := getResponseData["sessionlessurlclosure"]; ok && val != nil {
		data.Sessionlessurlclosure = types.StringValue(val.(string))
	} else {
		data.Sessionlessurlclosure = types.StringNull()
	}
	if val, ok := getResponseData["signatures"]; ok && val != nil {
		data.Signatures = types.StringValue(val.(string))
	} else {
		data.Signatures = types.StringNull()
	}
	if val, ok := getResponseData["sqlinjectionchecksqlwildchars"]; ok && val != nil {
		data.Sqlinjectionchecksqlwildchars = types.StringValue(val.(string))
	} else {
		data.Sqlinjectionchecksqlwildchars = types.StringNull()
	}
	if val, ok := getResponseData["sqlinjectiongrammar"]; ok && val != nil {
		data.Sqlinjectiongrammar = types.StringValue(val.(string))
	} else {
		data.Sqlinjectiongrammar = types.StringNull()
	}
	if val, ok := getResponseData["sqlinjectiononlycheckfieldswithsqlchars"]; ok && val != nil {
		data.Sqlinjectiononlycheckfieldswithsqlchars = types.StringValue(val.(string))
	} else {
		data.Sqlinjectiononlycheckfieldswithsqlchars = types.StringNull()
	}
	if val, ok := getResponseData["sqlinjectionparsecomments"]; ok && val != nil {
		data.Sqlinjectionparsecomments = types.StringValue(val.(string))
	} else {
		data.Sqlinjectionparsecomments = types.StringNull()
	}
	if val, ok := getResponseData["sqlinjectionruletype"]; ok && val != nil {
		data.Sqlinjectionruletype = types.StringValue(val.(string))
	} else {
		data.Sqlinjectionruletype = types.StringNull()
	}
	if val, ok := getResponseData["sqlinjectiontransformspecialchars"]; ok && val != nil {
		data.Sqlinjectiontransformspecialchars = types.StringValue(val.(string))
	} else {
		data.Sqlinjectiontransformspecialchars = types.StringNull()
	}
	if val, ok := getResponseData["sqlinjectiontype"]; ok && val != nil {
		data.Sqlinjectiontype = types.StringValue(val.(string))
	} else {
		data.Sqlinjectiontype = types.StringNull()
	}
	if val, ok := getResponseData["starturlclosure"]; ok && val != nil {
		data.Starturlclosure = types.StringValue(val.(string))
	} else {
		data.Starturlclosure = types.StringNull()
	}
	if val, ok := getResponseData["streaming"]; ok && val != nil {
		data.Streaming = types.StringValue(val.(string))
	} else {
		data.Streaming = types.StringNull()
	}
	if val, ok := getResponseData["stripcomments"]; ok && val != nil {
		data.Stripcomments = types.StringValue(val.(string))
	} else {
		data.Stripcomments = types.StringNull()
	}
	if val, ok := getResponseData["striphtmlcomments"]; ok && val != nil {
		data.Striphtmlcomments = types.StringValue(val.(string))
	} else {
		data.Striphtmlcomments = types.StringNull()
	}
	if val, ok := getResponseData["stripxmlcomments"]; ok && val != nil {
		data.Stripxmlcomments = types.StringValue(val.(string))
	} else {
		data.Stripxmlcomments = types.StringNull()
	}
	if val, ok := getResponseData["trace"]; ok && val != nil {
		data.Trace = types.StringValue(val.(string))
	} else {
		data.Trace = types.StringNull()
	}
	if val, ok := getResponseData["urldecoderequestcookies"]; ok && val != nil {
		data.Urldecoderequestcookies = types.StringValue(val.(string))
	} else {
		data.Urldecoderequestcookies = types.StringNull()
	}
	if val, ok := getResponseData["usehtmlerrorobject"]; ok && val != nil {
		data.Usehtmlerrorobject = types.StringValue(val.(string))
	} else {
		data.Usehtmlerrorobject = types.StringNull()
	}
	if val, ok := getResponseData["verboseloglevel"]; ok && val != nil {
		data.Verboseloglevel = types.StringValue(val.(string))
	} else {
		data.Verboseloglevel = types.StringNull()
	}
	if val, ok := getResponseData["xmlerrorobject"]; ok && val != nil {
		data.Xmlerrorobject = types.StringValue(val.(string))
	} else {
		data.Xmlerrorobject = types.StringNull()
	}
	if val, ok := getResponseData["xmlerrorstatuscode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Xmlerrorstatuscode = types.Int64Value(intVal)
		}
	} else {
		data.Xmlerrorstatuscode = types.Int64Null()
	}
	if val, ok := getResponseData["xmlerrorstatusmessage"]; ok && val != nil {
		data.Xmlerrorstatusmessage = types.StringValue(val.(string))
	} else {
		data.Xmlerrorstatusmessage = types.StringNull()
	}
	if val, ok := getResponseData["xmlsqlinjectionchecksqlwildchars"]; ok && val != nil {
		data.Xmlsqlinjectionchecksqlwildchars = types.StringValue(val.(string))
	} else {
		data.Xmlsqlinjectionchecksqlwildchars = types.StringNull()
	}
	if val, ok := getResponseData["xmlsqlinjectiononlycheckfieldswithsqlchars"]; ok && val != nil {
		data.Xmlsqlinjectiononlycheckfieldswithsqlchars = types.StringValue(val.(string))
	} else {
		data.Xmlsqlinjectiononlycheckfieldswithsqlchars = types.StringNull()
	}
	if val, ok := getResponseData["xmlsqlinjectionparsecomments"]; ok && val != nil {
		data.Xmlsqlinjectionparsecomments = types.StringValue(val.(string))
	} else {
		data.Xmlsqlinjectionparsecomments = types.StringNull()
	}
	if val, ok := getResponseData["xmlsqlinjectiontype"]; ok && val != nil {
		data.Xmlsqlinjectiontype = types.StringValue(val.(string))
	} else {
		data.Xmlsqlinjectiontype = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
