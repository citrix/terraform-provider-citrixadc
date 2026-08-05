package snmpview

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

// SnmpviewResourceModel describes the resource data model.
type SnmpviewResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Subtree types.String `tfsdk:"subtree"`
	Type    types.String `tfsdk:"type"`
}

func (r *SnmpviewResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmpview resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				// SDK v2: Required + ForceNew
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the SNMPv3 view. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a name that helps identify the SNMPv3 view.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my view\" or 'my view').",
			},
			"subtree": schema.StringAttribute{
				Required: true,
				// SDK v2: Required + ForceNew
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A particular branch (subtree) of the MIB tree that you want to associate with this SNMPv3 view. You must specify the subtree as an SNMP OID.",
			},
			"type": schema.StringAttribute{
				// SDK v2: Optional + Computed (no default)
				Optional:    true,
				Computed:    true,
				Description: "Include or exclude the subtree, specified by the subtree parameter, in or from this view. This setting can be useful when you have included a subtree, such as A, in an SNMPv3 view and you want to exclude a specific subtree of A, such as B, from the SNMPv3 view.",
			},
		},
	}
}

func snmpviewGetThePayloadFromthePlan(ctx context.Context, data *SnmpviewResourceModel) snmp.Snmpview {
	tflog.Debug(ctx, "In snmpviewGetThePayloadFromthePlan Function")

	// Create API request body from the model
	snmpview := snmp.Snmpview{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		snmpview.Name = data.Name.ValueString()
	}
	if !data.Subtree.IsNull() && !data.Subtree.IsUnknown() {
		snmpview.Subtree = data.Subtree.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		snmpview.Type = data.Type.ValueString()
	}

	return snmpview
}

func snmpviewSetAttrFromGet(ctx context.Context, data *SnmpviewResourceModel, getResponseData map[string]interface{}) *SnmpviewResourceModel {
	tflog.Debug(ctx, "In snmpviewSetAttrFromGet Function")

	// Convert API response to model. Guard the else-branches so we only null a
	// value that is currently unknown -- never clobber a known configured value
	// that NITRO happens to omit from GET.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["subtree"]; ok && val != nil {
		data.Subtree = types.StringValue(val.(string))
	} else if data.Subtree.IsUnknown() {
		data.Subtree = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}

	// Set ID for the resource.
	// Named resource keyed by name (matches SDK v2 d.SetId(name) and
	// resource_id_mapping.json "snmpview": "name").
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
