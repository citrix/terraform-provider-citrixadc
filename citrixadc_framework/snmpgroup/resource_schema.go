package snmpgroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpgroupResourceModel describes the resource data model.
type SnmpgroupResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Readviewname  types.String `tfsdk:"readviewname"`
	Securitylevel types.String `tfsdk:"securitylevel"`
}

func (r *SnmpgroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmpgroup resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the SNMPv3 group. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the SNMPv3 group. \n            \nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"readviewname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the configured SNMPv3 view that you want to bind to this SNMPv3 group. An SNMPv3 user bound to this group can access the subtrees that are bound to this SNMPv3 view as type INCLUDED, but cannot access the ones that are type EXCLUDED. If the Citrix ADC has multiple SNMPv3 view entries with the same name, all such entries are associated with the SNMPv3 group.",
			},
			"securitylevel": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Security level required for communication between the Citrix ADC and the SNMPv3 users who belong to the group. Specify one of the following options:\nnoAuthNoPriv. Require neither authentication nor encryption.\nauthNoPriv. Require authentication but no encryption.\nauthPriv. Require authentication and encryption.\nNote: If you specify authentication, you must specify an encryption algorithm when you assign an SNMPv3 user to the group. If you also specify encryption, you must assign both an authentication and an encryption algorithm for each group member.",
			},
		},
	}
}

func snmpgroupGetThePayloadFromtheConfig(ctx context.Context, data *SnmpgroupResourceModel) snmp.Snmpgroup {
	tflog.Debug(ctx, "In snmpgroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	snmpgroup := snmp.Snmpgroup{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		snmpgroup.Name = data.Name.ValueString()
	}
	if !data.Readviewname.IsNull() && !data.Readviewname.IsUnknown() {
		snmpgroup.Readviewname = data.Readviewname.ValueString()
	}
	if !data.Securitylevel.IsNull() && !data.Securitylevel.IsUnknown() {
		snmpgroup.Securitylevel = data.Securitylevel.ValueString()
	}

	return snmpgroup
}

func snmpgroupSetAttrFromGet(ctx context.Context, data *SnmpgroupResourceModel, getResponseData map[string]interface{}) *SnmpgroupResourceModel {
	tflog.Debug(ctx, "In snmpgroupSetAttrFromGet Function")

	// Convert API response to model. Only overwrite when NITRO returns a value;
	// never null a known/configured value that the GET may omit.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["readviewname"]; ok && val != nil {
		data.Readviewname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["securitylevel"]; ok && val != nil {
		data.Securitylevel = types.StringValue(val.(string))
	}

	// Set ID for the resource.
	// Single unique attribute (name) - use plain value as ID, matching the
	// SDK v2 resource (d.SetId(name)) and resource_id_mapping.json ("name").
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
