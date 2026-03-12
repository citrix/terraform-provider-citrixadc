package arpparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ArpparamResourceModel describes the resource data model.
type ArpparamResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Spoofvalidation types.String `tfsdk:"spoofvalidation"`
	Timeout         types.Int64  `tfsdk:"timeout"`
}

func (r *ArpparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the arpparam resource.",
			},
			"spoofvalidation": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "enable/disable arp spoofing validation",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1200),
				Description: "Time-out value (aging time) for the dynamically learned ARP entries, in seconds. The new value applies only to ARP entries that are dynamically learned after the new value is set. Previously existing ARP entries expire after the previously configured aging time.",
			},
		},
	}
}

func arpparamGetThePayloadFromtheConfig(ctx context.Context, data *ArpparamResourceModel) network.Arpparam {
	tflog.Debug(ctx, "In arpparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	arpparam := network.Arpparam{}
	if !data.Spoofvalidation.IsNull() {
		arpparam.Spoofvalidation = data.Spoofvalidation.ValueString()
	}
	if !data.Timeout.IsNull() {
		arpparam.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}

	return arpparam
}

func arpparamSetAttrFromGet(ctx context.Context, data *ArpparamResourceModel, getResponseData map[string]interface{}) *ArpparamResourceModel {
	tflog.Debug(ctx, "In arpparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["spoofvalidation"]; ok && val != nil {
		data.Spoofvalidation = types.StringValue(val.(string))
	} else {
		data.Spoofvalidation = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("arpparam-config")

	return data
}
