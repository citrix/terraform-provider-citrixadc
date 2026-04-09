package snmpuser

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpuserResourceModel describes the resource data model.
type SnmpuserResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Authpasswd          types.String `tfsdk:"authpasswd"`
	AuthpasswdWo        types.String `tfsdk:"authpasswd_wo"`
	AuthpasswdWoVersion types.Int64  `tfsdk:"authpasswd_wo_version"`
	Authtype            types.String `tfsdk:"authtype"`
	Group               types.String `tfsdk:"group"`
	Name                types.String `tfsdk:"name"`
	Privpasswd          types.String `tfsdk:"privpasswd"`
	PrivpasswdWo        types.String `tfsdk:"privpasswd_wo"`
	PrivpasswdWoVersion types.Int64  `tfsdk:"privpasswd_wo_version"`
	Privtype            types.String `tfsdk:"privtype"`
}

func (r *SnmpuserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmpuser resource.",
			},
			"authpasswd": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the pass phrase includes one or more spaces, enclose it in double or single quotation marks (for example, \"my phrase\" or 'my phrase').",
			},
			"authpasswd_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the pass phrase includes one or more spaces, enclose it in double or single quotation marks (for example, \"my phrase\" or 'my phrase').",
			},
			"authpasswd_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a authpasswd_wo update.",
			},
			"authtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication algorithm used by the Citrix ADC and the SNMPv3 user for authenticating the communication between them. You must specify the same authentication algorithm when you configure the SNMPv3 user in the SNMP manager.",
			},
			"group": schema.StringAttribute{
				Required:    true,
				Description: "Name of the configured SNMPv3 group to which to bind this SNMPv3 user. The access rights (bound SNMPv3 views) and security level set for this group are assigned to this user.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the SNMPv3 user. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my user\" or 'my user').",
			},
			"privpasswd": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the key includes one or more spaces, enclose it in double or single quotation marks (for example, \"my key\" or 'my key').",
			},
			"privpasswd_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the key includes one or more spaces, enclose it in double or single quotation marks (for example, \"my key\" or 'my key').",
			},
			"privpasswd_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a privpasswd_wo update.",
			},
			"privtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encryption algorithm used by the Citrix ADC and the SNMPv3 user for encrypting the communication between them. You must specify the same encryption algorithm when you configure the SNMPv3 user in the SNMP manager.",
			},
		},
	}
}

func snmpuserGetThePayloadFromthePlan(ctx context.Context, data *SnmpuserResourceModel) snmp.Snmpuser {
	tflog.Debug(ctx, "In snmpuserGetThePayloadFromthePlan Function")

	// Create API request body from the model
	snmpuser := snmp.Snmpuser{}
	if !data.Authpasswd.IsNull() {
		snmpuser.Authpasswd = data.Authpasswd.ValueString()
	}
	// Skip write-only attribute: authpasswd_wo
	// Skip version tracker attribute: authpasswd_wo_version
	if !data.Authtype.IsNull() {
		snmpuser.Authtype = data.Authtype.ValueString()
	}
	if !data.Group.IsNull() {
		snmpuser.Group = data.Group.ValueString()
	}
	if !data.Name.IsNull() {
		snmpuser.Name = data.Name.ValueString()
	}
	if !data.Privpasswd.IsNull() {
		snmpuser.Privpasswd = data.Privpasswd.ValueString()
	}
	// Skip write-only attribute: privpasswd_wo
	// Skip version tracker attribute: privpasswd_wo_version
	if !data.Privtype.IsNull() {
		snmpuser.Privtype = data.Privtype.ValueString()
	}

	return snmpuser
}

func snmpuserGetThePayloadFromtheConfig(ctx context.Context, data *SnmpuserResourceModel, payload *snmp.Snmpuser) {
	tflog.Debug(ctx, "In snmpuserGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: authpasswd_wo -> authpasswd
	if !data.AuthpasswdWo.IsNull() {
		authpasswdWo := data.AuthpasswdWo.ValueString()
		if authpasswdWo != "" {
			payload.Authpasswd = authpasswdWo
		}
	}
	// Handle write-only secret attribute: privpasswd_wo -> privpasswd
	if !data.PrivpasswdWo.IsNull() {
		privpasswdWo := data.PrivpasswdWo.ValueString()
		if privpasswdWo != "" {
			payload.Privpasswd = privpasswdWo
		}
	}
}

func snmpuserSetAttrFromGet(ctx context.Context, data *SnmpuserResourceModel, getResponseData map[string]interface{}) *SnmpuserResourceModel {
	tflog.Debug(ctx, "In snmpuserSetAttrFromGet Function")

	// Convert API response to model
	// authpasswd is not returned by NITRO API (secret/ephemeral) - retain from config
	// authpasswd_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// authpasswd_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["authtype"]; ok && val != nil {
		data.Authtype = types.StringValue(val.(string))
	} else {
		data.Authtype = types.StringNull()
	}
	if val, ok := getResponseData["group"]; ok && val != nil {
		data.Group = types.StringValue(val.(string))
	} else {
		data.Group = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// privpasswd is not returned by NITRO API (secret/ephemeral) - retain from config
	// privpasswd_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// privpasswd_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["privtype"]; ok && val != nil {
		data.Privtype = types.StringValue(val.(string))
	} else {
		data.Privtype = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
