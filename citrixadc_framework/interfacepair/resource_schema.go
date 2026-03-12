package interfacepair

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// InterfacepairResourceModel describes the resource data model.
type InterfacepairResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Interfaceid types.Int64  `tfsdk:"interface_id"`
	Ifnum       types.List   `tfsdk:"ifnum"`
}

func (r *InterfacepairResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the interfacepair resource.",
			},
			"interface_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The Interface pair id",
			},
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "The constituent interfaces in the interface pair",
			},
		},
	}
}

func interfacepairGetThePayloadFromtheConfig(ctx context.Context, data *InterfacepairResourceModel) network.Interfacepair {
	tflog.Debug(ctx, "In interfacepairGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	interfacepair := network.Interfacepair{}
	if !data.Interfaceid.IsNull() {
		interfacepair.Id = utils.IntPtr(int(data.Interfaceid.ValueInt64()))
	}

	return interfacepair
}

func interfacepairSetAttrFromGet(ctx context.Context, data *InterfacepairResourceModel, getResponseData map[string]interface{}) *InterfacepairResourceModel {
	tflog.Debug(ctx, "In interfacepairSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interfaceid = types.Int64Value(intVal)
		}
	} else {
		data.Interfaceid = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Interfaceid.ValueInt64()))

	return data
}
