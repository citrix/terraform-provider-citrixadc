package interfacepair

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
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
			// Maps to the NITRO "id" attribute. Named "interface_id" for backward
			// compatibility with the SDK v2 resource (Required + ForceNew).
			"interface_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The Interface pair id",
			},
			// SDK v2 marked ifnum ForceNew -> RequiresReplace here.
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "The constituent interfaces in the interface pair",
			},
		},
	}
}

func interfacepairGetThePayloadFromthePlan(ctx context.Context, data *InterfacepairResourceModel, diags *diag.Diagnostics) network.Interfacepair {
	tflog.Debug(ctx, "In interfacepairGetThePayloadFromthePlan Function")

	// Create API request body from the model
	interfacepair := network.Interfacepair{}
	if !data.Interfaceid.IsNull() && !data.Interfaceid.IsUnknown() {
		interfacepair.Id = utils.IntPtr(int(data.Interfaceid.ValueInt64()))
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		ifnumList := make([]string, 0, len(data.Ifnum.Elements()))
		diags.Append(data.Ifnum.ElementsAs(ctx, &ifnumList, false)...)
		interfacepair.Ifnum = ifnumList
	}

	return interfacepair
}

// interfacepairSetAttrFromGet populates resource state from a NITRO GET response.
// The configured "ifnum" (Required + RequiresReplace) is intentionally not
// overwritten from the GET response so the resource preserves the user's
// configured value and avoids "inconsistent result after apply" churn (mirrors
// the SDK v2 resource, which did not re-set ifnum on read). The datasource setter
// reads ifnum instead.
func interfacepairSetAttrFromGet(ctx context.Context, data *InterfacepairResourceModel, getResponseData map[string]interface{}) *InterfacepairResourceModel {
	tflog.Debug(ctx, "In interfacepairSetAttrFromGet Function")

	// Convert API response to model. The NITRO key attribute is "id".
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interfaceid = types.Int64Value(intVal)
		}
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Interfaceid.ValueInt64()))

	return data
}

// interfacepairSetAttrFromGetForDatasource populates all readable attributes
// (including "ifnum") from a NITRO GET response for the datasource.
func interfacepairSetAttrFromGetForDatasource(ctx context.Context, data *InterfacepairResourceModel, getResponseData map[string]interface{}) *InterfacepairResourceModel {
	tflog.Debug(ctx, "In interfacepairSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interfaceid = types.Int64Value(intVal)
		}
	} else {
		data.Interfaceid = types.Int64Null()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		if listVal, listOk := val.([]interface{}); listOk {
			strs := make([]string, 0, len(listVal))
			for _, item := range listVal {
				strs = append(strs, fmt.Sprintf("%v", item))
			}
			listValue, d := types.ListValueFrom(ctx, types.StringType, strs)
			if !d.HasError() {
				data.Ifnum = listValue
			}
		} else if strVal, strOk := val.(string); strOk {
			listValue, d := types.ListValueFrom(ctx, types.StringType, []string{strVal})
			if !d.HasError() {
				data.Ifnum = listValue
			}
		}
	} else {
		data.Ifnum = types.ListNull(types.StringType)
	}

	// Set ID for the resource - single unique attribute (interface_id), plain value.
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Interfaceid.ValueInt64()))

	return data
}
