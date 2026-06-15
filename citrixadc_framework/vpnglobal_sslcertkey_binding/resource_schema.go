package vpnglobal_sslcertkey_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalSslcertkeyBindingResourceModel describes the resource data model.
type VpnglobalSslcertkeyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Cacert                 types.String `tfsdk:"cacert"`
	Certkeyname            types.String `tfsdk:"certkeyname"`
	Crlcheck               types.String `tfsdk:"crlcheck"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Ocspcheck              types.String `tfsdk:"ocspcheck"`
	Userdataencryptionkey  types.String `tfsdk:"userdataencryptionkey"`
}

func (r *VpnglobalSslcertkeyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_sslcertkey_binding resource.",
			},
			"certkeyname": schema.StringAttribute{
				// SDK v2 contract: Required + ForceNew
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SSL certkey to use in signing tokens. Only RSA cert key is allowed",
			},
			"cacert": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the CA certificate binding.",
			},
			"crlcheck": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The state of the CRL check parameter (Mandatory/Optional).",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"ocspcheck": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The state of the OCSP check parameter (Mandatory/Optional).",
			},
			"userdataencryptionkey": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Certificate to be used for encrypting user data like KB Question and Answers, Alternate Email Address, etc.",
			},
		},
	}
}

func vpnglobal_sslcertkey_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalSslcertkeyBindingResourceModel) vpn.Vpnglobalsslcertkeybinding {
	tflog.Debug(ctx, "In vpnglobal_sslcertkey_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_sslcertkey_binding := vpn.Vpnglobalsslcertkeybinding{}
	if !data.Cacert.IsNull() && !data.Cacert.IsUnknown() {
		vpnglobal_sslcertkey_binding.Cacert = data.Cacert.ValueString()
	}
	if !data.Certkeyname.IsNull() && !data.Certkeyname.IsUnknown() {
		vpnglobal_sslcertkey_binding.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Crlcheck.IsNull() && !data.Crlcheck.IsUnknown() {
		vpnglobal_sslcertkey_binding.Crlcheck = data.Crlcheck.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_sslcertkey_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Ocspcheck.IsNull() && !data.Ocspcheck.IsUnknown() {
		vpnglobal_sslcertkey_binding.Ocspcheck = data.Ocspcheck.ValueString()
	}
	if !data.Userdataencryptionkey.IsNull() && !data.Userdataencryptionkey.IsUnknown() {
		vpnglobal_sslcertkey_binding.Userdataencryptionkey = data.Userdataencryptionkey.ValueString()
	}

	return vpnglobal_sslcertkey_binding
}

// vpnglobal_sslcertkey_bindingSetAttrFromGet is the RESOURCE-side setter. It preserves
// the prior state/plan values for inputs the NITRO GET response does not echo back
// (e.g. cacert / userdataencryptionkey), so the post-apply state matches config and we
// avoid "inconsistent result after apply" diffs (Pattern 7 / Pattern 13). It does NOT
// recompute data.Id — the ID is the plain certkeyname set once in Create.
func vpnglobal_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_sslcertkey_bindingSetAttrFromGet Function")

	// The optional inputs (cacert, crlcheck, gotopriorityexpression, ocspcheck,
	// userdataencryptionkey) are RequiresReplace and not Computed; the NITRO GET for this
	// binding does not reliably echo them back, and adopting a server value that differs
	// from config would trigger "inconsistent result after apply". Preserve the prior
	// plan/state values (Pattern 7) — only certkeyname (the key) is adopted from the GET
	// response when missing, to support import where state carries only the ID.
	if data.Certkeyname.IsNull() || data.Certkeyname.ValueString() == "" {
		if val, ok := getResponseData["certkeyname"]; ok && val != nil {
			data.Certkeyname = types.StringValue(val.(string))
		}
	}

	return data
}

// vpnglobal_sslcertkey_bindingSetAttrFromGetForDatasource is the DATASOURCE-side setter.
// A datasource has no prior plan/state to preserve, so it faithfully copies every field
// from the GET response (null when absent) and sets data.Id to the plain certkeyname,
// since the datasource never calls Create (Pattern 7 split).
func vpnglobal_sslcertkey_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_sslcertkey_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["cacert"]; ok && val != nil {
		data.Cacert = types.StringValue(val.(string))
	} else {
		data.Cacert = types.StringNull()
	}
	if val, ok := getResponseData["certkeyname"]; ok && val != nil {
		data.Certkeyname = types.StringValue(val.(string))
	} else {
		data.Certkeyname = types.StringNull()
	}
	if val, ok := getResponseData["crlcheck"]; ok && val != nil {
		data.Crlcheck = types.StringValue(val.(string))
	} else {
		data.Crlcheck = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["ocspcheck"]; ok && val != nil {
		data.Ocspcheck = types.StringValue(val.(string))
	} else {
		data.Ocspcheck = types.StringNull()
	}
	if val, ok := getResponseData["userdataencryptionkey"]; ok && val != nil {
		data.Userdataencryptionkey = types.StringValue(val.(string))
	} else {
		data.Userdataencryptionkey = types.StringNull()
	}

	// Plain single-key ID (matches SDK v2 d.SetId(certkeyname))
	data.Id = data.Certkeyname

	return data
}
