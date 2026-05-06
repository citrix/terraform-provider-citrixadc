package aaakcdaccount

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaakcdaccountResourceModel describes the resource data model.
type AaakcdaccountResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Cacert               types.String `tfsdk:"cacert"`
	Delegateduser        types.String `tfsdk:"delegateduser"`
	Enterpriserealm      types.String `tfsdk:"enterpriserealm"`
	Kcdaccount           types.String `tfsdk:"kcdaccount"`
	Kcdpassword          types.String `tfsdk:"kcdpassword"`
	KcdpasswordWo        types.String `tfsdk:"kcdpassword_wo"`
	KcdpasswordWoVersion types.Int64  `tfsdk:"kcdpassword_wo_version"`
	Keytab               types.String `tfsdk:"keytab"`
	Realmstr             types.String `tfsdk:"realmstr"`
	Servicespn           types.String `tfsdk:"servicespn"`
	Usercert             types.String `tfsdk:"usercert"`
	Userrealm            types.String `tfsdk:"userrealm"`
}

func (r *AaakcdaccountResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaakcdaccount resource.",
			},
			"cacert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA Cert for UserCert or when doing PKINIT backchannel.",
			},
			"delegateduser": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username that can perform kerberos constrained delegation.",
			},
			"enterpriserealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enterprise Realm of the user. This should be given only in certain KDC deployments where KDC expects Enterprise username instead of Principal Name",
			},
			"kcdaccount": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the KCD account.",
			},
			"kcdpassword": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password for Delegated User.",
			},
			"kcdpassword_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Password for Delegated User.",
			},
			"kcdpassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a kcdpassword_wo update.",
			},
			"keytab": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The path to the keytab file. If specified other parameters in this command need not be given",
			},
			"realmstr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Kerberos Realm.",
			},
			"servicespn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service SPN. When specified, this will be used to fetch kerberos tickets. If not specified, Citrix ADC will construct SPN using service fqdn",
			},
			"usercert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL Cert (including private key) for Delegated User.",
			},
			"userrealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Realm of the user",
			},
		},
	}
}

func aaakcdaccountGetThePayloadFromthePlan(ctx context.Context, data *AaakcdaccountResourceModel) aaa.Aaakcdaccount {
	tflog.Debug(ctx, "In aaakcdaccountGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaakcdaccount := aaa.Aaakcdaccount{}
	if !data.Cacert.IsNull() {
		aaakcdaccount.Cacert = data.Cacert.ValueString()
	}
	if !data.Delegateduser.IsNull() {
		aaakcdaccount.Delegateduser = data.Delegateduser.ValueString()
	}
	if !data.Enterpriserealm.IsNull() {
		aaakcdaccount.Enterpriserealm = data.Enterpriserealm.ValueString()
	}
	if !data.Kcdaccount.IsNull() {
		aaakcdaccount.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Kcdpassword.IsNull() {
		aaakcdaccount.Kcdpassword = data.Kcdpassword.ValueString()
	}
	// Skip write-only attribute: kcdpassword_wo
	// Skip version tracker attribute: kcdpassword_wo_version
	if !data.Keytab.IsNull() {
		aaakcdaccount.Keytab = data.Keytab.ValueString()
	}
	if !data.Realmstr.IsNull() {
		aaakcdaccount.Realmstr = data.Realmstr.ValueString()
	}
	if !data.Servicespn.IsNull() {
		aaakcdaccount.Servicespn = data.Servicespn.ValueString()
	}
	if !data.Usercert.IsNull() {
		aaakcdaccount.Usercert = data.Usercert.ValueString()
	}
	if !data.Userrealm.IsNull() {
		aaakcdaccount.Userrealm = data.Userrealm.ValueString()
	}

	return aaakcdaccount
}

func aaakcdaccountGetThePayloadFromtheConfig(ctx context.Context, data *AaakcdaccountResourceModel, payload *aaa.Aaakcdaccount) {
	tflog.Debug(ctx, "In aaakcdaccountGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: kcdpassword_wo -> kcdpassword
	if !data.KcdpasswordWo.IsNull() {
		kcdpasswordWo := data.KcdpasswordWo.ValueString()
		if kcdpasswordWo != "" {
			payload.Kcdpassword = kcdpasswordWo
		}
	}
}

func aaakcdaccountSetAttrFromGet(ctx context.Context, data *AaakcdaccountResourceModel, getResponseData map[string]interface{}) *AaakcdaccountResourceModel {
	tflog.Debug(ctx, "In aaakcdaccountSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacert"]; ok && val != nil {
		data.Cacert = types.StringValue(val.(string))
	} else {
		data.Cacert = types.StringNull()
	}
	if val, ok := getResponseData["delegateduser"]; ok && val != nil {
		data.Delegateduser = types.StringValue(val.(string))
	} else {
		data.Delegateduser = types.StringNull()
	}
	if val, ok := getResponseData["enterpriserealm"]; ok && val != nil {
		data.Enterpriserealm = types.StringValue(val.(string))
	} else {
		data.Enterpriserealm = types.StringNull()
	}
	if val, ok := getResponseData["kcdaccount"]; ok && val != nil {
		data.Kcdaccount = types.StringValue(val.(string))
	} else {
		data.Kcdaccount = types.StringNull()
	}
	// kcdpassword is not returned by NITRO API (secret/ephemeral) - retain from config
	// kcdpassword_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// kcdpassword_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["keytab"]; ok && val != nil {
		data.Keytab = types.StringValue(val.(string))
	} else {
		data.Keytab = types.StringNull()
	}
	if val, ok := getResponseData["realmstr"]; ok && val != nil {
		data.Realmstr = types.StringValue(strings.ToLower(val.(string)))
	} else {
		data.Realmstr = types.StringNull()
	}
	if val, ok := getResponseData["servicespn"]; ok && val != nil {
		data.Servicespn = types.StringValue(val.(string))
	} else {
		data.Servicespn = types.StringNull()
	}
	if val, ok := getResponseData["usercert"]; ok && val != nil {
		data.Usercert = types.StringValue(val.(string))
	} else {
		data.Usercert = types.StringNull()
	}
	if val, ok := getResponseData["userrealm"]; ok && val != nil {
		data.Userrealm = types.StringValue(val.(string))
	} else {
		data.Userrealm = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Kcdaccount.ValueString()))

	return data
}
