package lbaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbactionResourceModel describes the resource data model.
type LbactionResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Type    types.String `tfsdk:"type"`
	Value   types.List   `tfsdk:"value"`
}

func (r *LbactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbaction resource.",
			},
			"comment": schema.StringAttribute{
				// SDK v2 parity: Optional + Computed, no ForceNew (updateable), no Default.
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this LB action.",
			},
			"name": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the LB action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb action\" or 'my lb action').",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). It drives an
				// in-place rename via Update, so it must NOT force replacement and must
				// NOT be Computed (it is a pure user input, never echoed back by GET).
				Optional:    true,
				Description: "New name for the LB action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb action\" or my lb action').",
			},
			"type": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of an LB action. Available settings function as follows:\n* NOLBACTION - Does not consider LB action in making LB decision.\n* SELECTIONORDER - services bound to vserver with order specified in value parameter is considerd for lb/gslb decision.",
			},
			"value": schema.ListAttribute{
				// SDK v2 parity: TypeList of TypeInt, Optional + Computed, updateable.
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The selection order list used during lb/gslb decision. Preference of services during lb/gslb decision is as follows - services corresponding to first order specified in the sequence is considered first, services corresponding to second order specified in the sequence is considered next and so on. For example, if -value 2 1 3 is specified here and service-1 bound to a vserver with order 1, service-2 bound to a vserver with order 2 and  service-3 bound to a vserver with order 3. Then preference of selecting services in LB decision is as follows: service-2, service-1, service-3.",
			},
		},
	}
}

// lbactionValueToIntList converts the plan/state types.List (Int64) into the []int
// the NITRO struct expects.
func lbactionValueToIntList(ctx context.Context, list types.List) []int {
	var values []int64
	list.ElementsAs(ctx, &values, false)
	intValues := make([]int, len(values))
	for i, v := range values {
		intValues[i] = int(v)
	}
	return intValues
}

// lbactionValueFromGet converts the NITRO GET "value" field (a slice of stringified
// integers) into a types.List of Int64.
func lbactionValueFromGet(getResponseData map[string]interface{}) types.List {
	if val, ok := getResponseData["value"]; ok && val != nil {
		if rawList, ok := val.([]interface{}); ok {
			elems := make([]attr.Value, 0, len(rawList))
			for _, v := range rawList {
				i64, err := utils.ConvertToInt64(v)
				if err != nil {
					continue
				}
				elems = append(elems, types.Int64Value(i64))
			}
			listVal, diags := types.ListValue(types.Int64Type, elems)
			if diags.HasError() {
				return types.ListNull(types.Int64Type)
			}
			return listVal
		}
	}
	return types.ListNull(types.Int64Type)
}

func lbactionGetThePayloadFromthePlan(ctx context.Context, data *LbactionResourceModel) lb.Lbaction {
	tflog.Debug(ctx, "In lbactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbaction := lb.Lbaction{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		lbaction.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		lbaction.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		lbaction.Type = data.Type.ValueString()
	}
	if !data.Value.IsNull() && !data.Value.IsUnknown() {
		lbaction.Value = lbactionValueToIntList(ctx, data.Value)
	}

	return lbaction
}

func lbactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *LbactionResourceModel) lb.Lbaction {
	tflog.Debug(ctx, "In lbactionGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields.
	// name and type are ForceNew (RequiresReplace) and never reach Update.
	lbaction := lb.Lbaction{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		lbaction.Name = data.Name.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		lbaction.Comment = data.Comment.ValueString()
	}
	if !data.Value.IsNull() && !data.Value.IsUnknown() {
		lbaction.Value = lbactionValueToIntList(ctx, data.Value)
	}

	return lbaction
}

func lbactionSetAttrFromGet(ctx context.Context, data *LbactionResourceModel, getResponseData map[string]interface{}) *LbactionResourceModel {
	tflog.Debug(ctx, "In lbactionSetAttrFromGet Function")

	// Convert API response to model (resource path - preserves user-facing key/rename inputs).
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	// name is the user-facing key. Once a rename has happened (via newname), the live
	// object name (tracked by data.Id) diverges from the configured name, and GET
	// returns the live (new) name. Overwriting name from GET would clobber the user's
	// configured value and trigger a spurious RequiresReplace diff. Only adopt the GET
	// value when we don't already have one (e.g. on import, where state carries only
	// the ID); otherwise preserve.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	}
	data.Value = lbactionValueFromGet(getResponseData)

	return data
}

// lbactionSetAttrFromGetForDatasource faithfully copies every field from the GET
// response. The datasource has no prior plan/state to preserve, so it must populate
// the model directly from the API response and set the ID itself.
func lbactionSetAttrFromGetForDatasource(ctx context.Context, data *LbactionResourceModel, getResponseData map[string]interface{}) *LbactionResourceModel {
	tflog.Debug(ctx, "In lbactionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	data.Value = lbactionValueFromGet(getResponseData)

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
