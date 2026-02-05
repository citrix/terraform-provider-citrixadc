package snmpuser

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpuserResourceModel describes the resource data model.
type SnmpuserResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Authpasswd types.String `tfsdk:"authpasswd"`
	Authtype   types.String `tfsdk:"authtype"`
	Group      types.String `tfsdk:"group"`
	Name       types.String `tfsdk:"name"`
	Privpasswd types.String `tfsdk:"privpasswd"`
	Privtype   types.String `tfsdk:"privtype"`
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
				Computed:    true,
				Description: "Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the pass phrase includes one or more spaces, enclose it in double or single quotation marks (for example, \"my phrase\" or 'my phrase').",
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
				Required:    true,
				Description: "Name for the SNMPv3 user. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my user\" or 'my user').",
			},
			"privpasswd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the key includes one or more spaces, enclose it in double or single quotation marks (for example, \"my key\" or 'my key').",
			},
			"privtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encryption algorithm used by the Citrix ADC and the SNMPv3 user for encrypting the communication between them. You must specify the same encryption algorithm when you configure the SNMPv3 user in the SNMP manager.",
			},
		},
	}
}

func snmpuserGetThePayloadFromtheConfig(ctx context.Context, data *SnmpuserResourceModel) snmp.Snmpuser {
	tflog.Debug(ctx, "In snmpuserGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	snmpuser := snmp.Snmpuser{}
	if !data.Authpasswd.IsNull() {
		snmpuser.Authpasswd = data.Authpasswd.ValueString()
	}
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
	if !data.Privtype.IsNull() {
		snmpuser.Privtype = data.Privtype.ValueString()
	}

	return snmpuser
}

func snmpuserSetAttrFromGet(ctx context.Context, data *SnmpuserResourceModel, getResponseData map[string]interface{}) *SnmpuserResourceModel {
	tflog.Debug(ctx, "In snmpuserSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["authpasswd"]; ok && val != nil {
		data.Authpasswd = types.StringValue(val.(string))
	} else {
		data.Authpasswd = types.StringNull()
	}
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
	if val, ok := getResponseData["privpasswd"]; ok && val != nil {
		data.Privpasswd = types.StringValue(val.(string))
	} else {
		data.Privpasswd = types.StringNull()
	}
	if val, ok := getResponseData["privtype"]; ok && val != nil {
		data.Privtype = types.StringValue(val.(string))
	} else {
		data.Privtype = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
