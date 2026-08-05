package nsvariable

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsvariableResourceModel describes the resource data model.
type NsvariableResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Comment       types.String `tfsdk:"comment"`
	Expires       types.Int64  `tfsdk:"expires"`
	Iffull        types.String `tfsdk:"iffull"`
	Ifnovalue     types.String `tfsdk:"ifnovalue"`
	Ifvaluetoobig types.String `tfsdk:"ifvaluetoobig"`
	Init          types.String `tfsdk:"init"`
	Name          types.String `tfsdk:"name"`
	Scope         types.String `tfsdk:"scope"`
	Type          types.String `tfsdk:"type"`
}

func (r *NsvariableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsvariable resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this variable.",
			},
			"expires": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Value expiration in seconds. If the value is not referenced within the expiration period it will be deleted. 0 (the default) means no expiration.",
			},
			"iffull": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if an assignment to a map exceeds its configured max-entries:\n   lru - (default) reuse the least recently used entry in the map.\n   undef - force the assignment to return an undefined (Undef) result to the policy executing the assignment.",
			},
			"ifnovalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if on a variable reference in an expression if the variable is single-valued and uninitialized\nor if the variable is a map and there is no value for the specified key:\n   init - (default) initialize the single-value variable, or create a map entry for the key and the initial value,\nusing the -init value or its default.\n   undef - force the expression evaluation to return an undefined (Undef) result to the policy executing the expression.",
			},
			"ifvaluetoobig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if an value is assigned to a text variable that exceeds its configured max-size,\nor if a key is used that exceeds its configured max-size:\n   truncate - (default) truncate the text string to the first max-size bytes and proceed.\n   undef - force the assignment or expression evaluation to return an undefined (Undef) result to the policy executing the assignment or expression.",
			},
			"init": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initialization value for this variable, to which a singleton variable or map entry will be set if it is referenced before an assignment action has assigned it a value. If the singleton variable or map entry already has been assigned a value, setting this parameter will have no effect on that variable value. Default: 0 for ulong, NULL for text",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Variable name.  This follows the same syntax rules as other expression entity names:\n   It must begin with an alpha character (A-Z or a-z) or an underscore (_).\n   The rest of the characters must be alpha, numeric (0-9) or underscores.\n   It cannot be re or xp (reserved for regular and XPath expressions).\n   It cannot be an expression reserved word (e.g. SYS or HTTP).\n   It cannot be used for an existing expression object (HTTP callout, patset, dataset, stringmap, or named expression).",
			},
			"scope": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Scope of the variable:\n   global - (default) one set of values visible across all Packet Engines on a standalone Citrix ADC, an HA pair, or all nodes of a cluster\n   transaction - one value for each request-response transaction (singleton variables only; no expiration)",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specification of the variable type; one of the following:\n   ulong - singleton variable with an unsigned 64-bit value.\n   text(value-max-size) - singleton variable with a text string value.\n   map(text(key-max-size),ulong,max-entries) - map of text string keys to unsigned 64-bit values.\n   map(text(key-max-size),text(value-max-size),max-entries) - map of text string keys to text string values.\nwhere\n   value-max-size is a positive integer that is the maximum number of bytes in a text string value.\n   key-max-size is a positive integer that is the maximum number of bytes in a text string key.\n   max-entries is a positive integer that is the maximum number of entries in a map variable.\n      For a global singleton text variable, value-max-size <= 64000.\n      For a global map with ulong values, key-max-size <= 64000.\n      For a global map with text values,  key-max-size + value-max-size <= 64000.\n   max-entries is a positive integer that is the maximum number of entries in a map variable. This has a theoretical maximum of 2^64-1, but in actual use will be much smaller, considering the memory available for use by the map.\nExample:\n   map(text(10),text(20),100) specifies a map of text string keys (max size 10 bytes) to text string values (max size 20 bytes), with 100 max entries.",
			},
		},
	}
}

// nsvariableGetThePayloadFromthePlan builds the full create payload from the plan.
// Each field is guarded on !IsNull() && !IsUnknown() so unconfigured Optional+Computed
// attributes (Unknown in the plan) are not sent - matching the SDK v2 behaviour where
// expires was only sent when present in the raw config.
func nsvariableGetThePayloadFromthePlan(ctx context.Context, data *NsvariableResourceModel) ns.Nsvariable {
	tflog.Debug(ctx, "In nsvariableGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nsvariable := ns.Nsvariable{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		nsvariable.Comment = data.Comment.ValueString()
	}
	if !data.Expires.IsNull() && !data.Expires.IsUnknown() {
		nsvariable.Expires = utils.IntPtr(int(data.Expires.ValueInt64()))
	}
	if !data.Iffull.IsNull() && !data.Iffull.IsUnknown() {
		nsvariable.Iffull = data.Iffull.ValueString()
	}
	if !data.Ifnovalue.IsNull() && !data.Ifnovalue.IsUnknown() {
		nsvariable.Ifnovalue = data.Ifnovalue.ValueString()
	}
	if !data.Ifvaluetoobig.IsNull() && !data.Ifvaluetoobig.IsUnknown() {
		nsvariable.Ifvaluetoobig = data.Ifvaluetoobig.ValueString()
	}
	if !data.Init.IsNull() && !data.Init.IsUnknown() {
		nsvariable.Init = data.Init.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nsvariable.Name = data.Name.ValueString()
	}
	if !data.Scope.IsNull() && !data.Scope.IsUnknown() {
		nsvariable.Scope = data.Scope.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		nsvariable.Type = data.Type.ValueString()
	}

	return nsvariable
}

// nsvariableGetTheUpdatablePayloadFromThePlan builds the update payload, restricted to
// the NITRO-updatable fields. name/type/scope are ForceNew (RequiresReplace) in SDK v2
// and therefore never reach Update; only name is carried as the resource identifier.
func nsvariableGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *NsvariableResourceModel) ns.Nsvariable {
	tflog.Debug(ctx, "In nsvariableGetTheUpdatablePayloadFromThePlan Function")

	nsvariable := ns.Nsvariable{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nsvariable.Name = data.Name.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		nsvariable.Comment = data.Comment.ValueString()
	}
	if !data.Expires.IsNull() && !data.Expires.IsUnknown() {
		nsvariable.Expires = utils.IntPtr(int(data.Expires.ValueInt64()))
	}
	if !data.Iffull.IsNull() && !data.Iffull.IsUnknown() {
		nsvariable.Iffull = data.Iffull.ValueString()
	}
	if !data.Ifnovalue.IsNull() && !data.Ifnovalue.IsUnknown() {
		nsvariable.Ifnovalue = data.Ifnovalue.ValueString()
	}
	if !data.Ifvaluetoobig.IsNull() && !data.Ifvaluetoobig.IsUnknown() {
		nsvariable.Ifvaluetoobig = data.Ifvaluetoobig.ValueString()
	}
	if !data.Init.IsNull() && !data.Init.IsUnknown() {
		nsvariable.Init = data.Init.ValueString()
	}

	return nsvariable
}

func nsvariableSetAttrFromGet(ctx context.Context, data *NsvariableResourceModel, getResponseData map[string]interface{}) *NsvariableResourceModel {
	tflog.Debug(ctx, "In nsvariableSetAttrFromGet Function")

	// Convert API response to model.
	// The else-branches only null a value when it is Unknown (i.e. an unconfigured
	// Optional+Computed attribute resolving after create). A known/configured value that
	// NITRO omits from GET (e.g. expires=0) is preserved, avoiding the omit-on-default
	// "inconsistent result after apply" trap.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["expires"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Expires = types.Int64Value(intVal)
		}
	} else if data.Expires.IsUnknown() {
		data.Expires = types.Int64Null()
	}
	if val, ok := getResponseData["iffull"]; ok && val != nil {
		data.Iffull = types.StringValue(val.(string))
	} else if data.Iffull.IsUnknown() {
		data.Iffull = types.StringNull()
	}
	if val, ok := getResponseData["ifnovalue"]; ok && val != nil {
		data.Ifnovalue = types.StringValue(val.(string))
	} else if data.Ifnovalue.IsUnknown() {
		data.Ifnovalue = types.StringNull()
	}
	if val, ok := getResponseData["ifvaluetoobig"]; ok && val != nil {
		data.Ifvaluetoobig = types.StringValue(val.(string))
	} else if data.Ifvaluetoobig.IsUnknown() {
		data.Ifvaluetoobig = types.StringNull()
	}
	if val, ok := getResponseData["init"]; ok && val != nil {
		data.Init = types.StringValue(val.(string))
	} else if data.Init.IsUnknown() {
		data.Init = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["scope"]; ok && val != nil {
		data.Scope = types.StringValue(val.(string))
	} else if data.Scope.IsUnknown() {
		data.Scope = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
