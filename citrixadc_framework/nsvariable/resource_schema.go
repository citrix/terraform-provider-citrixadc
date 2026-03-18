package nsvariable

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				Default:     stringdefault.StaticString("lru"),
				Description: "Action to perform if an assignment to a map exceeds its configured max-entries:\n   lru - (default) reuse the least recently used entry in the map.\n   undef - force the assignment to return an undefined (Undef) result to the policy executing the assignment.",
			},
			"ifnovalue": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("init"),
				Description: "Action to perform if on a variable reference in an expression if the variable is single-valued and uninitialized\nor if the variable is a map and there is no value for the specified key:\n   init - (default) initialize the single-value variable, or create a map entry for the key and the initial value,\nusing the -init value or its default.\n   undef - force the expression evaluation to return an undefined (Undef) result to the policy executing the expression.",
			},
			"ifvaluetoobig": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("truncate"),
				Description: "Action to perform if an value is assigned to a text variable that exceeds its configured max-size,\nor if a key is used that exceeds its configured max-size:\n   truncate - (default) truncate the text string to the first max-size bytes and proceed.\n   undef - force the assignment or expression evaluation to return an undefined (Undef) result to the policy executing the assignment or expression.",
			},
			"init": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initialization value for this variable, to which a singleton variable or map entry will be set if it is referenced before an assignment action has assigned it a value. If the singleton variable or map entry already has been assigned a value, setting this parameter will have no effect on that variable value. Default: 0 for ulong, NULL for text",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Variable name.  This follows the same syntax rules as other expression entity names:\n   It must begin with an alpha character (A-Z or a-z) or an underscore (_).\n   The rest of the characters must be alpha, numeric (0-9) or underscores.\n   It cannot be re or xp (reserved for regular and XPath expressions).\n   It cannot be an expression reserved word (e.g. SYS or HTTP).\n   It cannot be used for an existing expression object (HTTP callout, patset, dataset, stringmap, or named expression).",
			},
			"scope": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("global"),
				Description: "Scope of the variable:\n   global - (default) one set of values visible across all Packet Engines on a standalone Citrix ADC, an HA pair, or all nodes of a cluster\n   transaction - one value for each request-response transaction (singleton variables only; no expiration)",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Specification of the variable type; one of the following:\n   ulong - singleton variable with an unsigned 64-bit value.\n   text(value-max-size) - singleton variable with a text string value.\n   map(text(key-max-size),ulong,max-entries) - map of text string keys to unsigned 64-bit values.\n   map(text(key-max-size),text(value-max-size),max-entries) - map of text string keys to text string values.\nwhere\n   value-max-size is a positive integer that is the maximum number of bytes in a text string value.\n   key-max-size is a positive integer that is the maximum number of bytes in a text string key.\n   max-entries is a positive integer that is the maximum number of entries in a map variable.\n      For a global singleton text variable, value-max-size <= 64000.\n      For a global map with ulong values, key-max-size <= 64000.\n      For a global map with text values,  key-max-size + value-max-size <= 64000.\n   max-entries is a positive integer that is the maximum number of entries in a map variable. This has a theoretical maximum of 2^64-1, but in actual use will be much smaller, considering the memory available for use by the map.\nExample:\n   map(text(10),text(20),100) specifies a map of text string keys (max size 10 bytes) to text string values (max size 20 bytes), with 100 max entries.",
			},
		},
	}
}

func nsvariableGetThePayloadFromtheConfig(ctx context.Context, data *NsvariableResourceModel) ns.Nsvariable {
	tflog.Debug(ctx, "In nsvariableGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsvariable := ns.Nsvariable{}
	if !data.Comment.IsNull() {
		nsvariable.Comment = data.Comment.ValueString()
	}
	if !data.Expires.IsNull() {
		nsvariable.Expires = utils.IntPtr(int(data.Expires.ValueInt64()))
	}
	if !data.Iffull.IsNull() {
		nsvariable.Iffull = data.Iffull.ValueString()
	}
	if !data.Ifnovalue.IsNull() {
		nsvariable.Ifnovalue = data.Ifnovalue.ValueString()
	}
	if !data.Ifvaluetoobig.IsNull() {
		nsvariable.Ifvaluetoobig = data.Ifvaluetoobig.ValueString()
	}
	if !data.Init.IsNull() {
		nsvariable.Init = data.Init.ValueString()
	}
	if !data.Name.IsNull() {
		nsvariable.Name = data.Name.ValueString()
	}
	if !data.Scope.IsNull() {
		nsvariable.Scope = data.Scope.ValueString()
	}
	if !data.Type.IsNull() {
		nsvariable.Type = data.Type.ValueString()
	}

	return nsvariable
}

func nsvariableSetAttrFromGet(ctx context.Context, data *NsvariableResourceModel, getResponseData map[string]interface{}) *NsvariableResourceModel {
	tflog.Debug(ctx, "In nsvariableSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["expires"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Expires = types.Int64Value(intVal)
		}
	} else {
		data.Expires = types.Int64Null()
	}
	if val, ok := getResponseData["iffull"]; ok && val != nil {
		data.Iffull = types.StringValue(val.(string))
	} else {
		data.Iffull = types.StringNull()
	}
	if val, ok := getResponseData["ifnovalue"]; ok && val != nil {
		data.Ifnovalue = types.StringValue(val.(string))
	} else {
		data.Ifnovalue = types.StringNull()
	}
	if val, ok := getResponseData["ifvaluetoobig"]; ok && val != nil {
		data.Ifvaluetoobig = types.StringValue(val.(string))
	} else {
		data.Ifvaluetoobig = types.StringNull()
	}
	if val, ok := getResponseData["init"]; ok && val != nil {
		data.Init = types.StringValue(val.(string))
	} else {
		data.Init = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["scope"]; ok && val != nil {
		data.Scope = types.StringValue(val.(string))
	} else {
		data.Scope = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
