package policydataset

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicydatasetResourceModel describes the resource data model.
type PolicydatasetResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Dynamic     types.String `tfsdk:"dynamic"`
	Dynamiconly types.Bool   `tfsdk:"dynamiconly"`
	Name        types.String `tfsdk:"name"`
	Patsetfile  types.String `tfsdk:"patsetfile"`
	Type        types.String `tfsdk:"type"`
}

func (r *PolicydatasetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policydataset resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this dataset or a data bound to this dataset.",
			},
			"dynamic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is used to populate internal dataset information so that the dataset can also be used dynamically in an expression. Here dynamically means the dataset name can also be derived using an expression. For example for a given dataset name \"allow_test\" it can be used dynamically as client.ip.src.equals_any(\"allow_\" + http.req.url.path.get(1)). This cannot be used with default datasets.",
			},
			"dynamiconly": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Shows only dynamic datasets when set true.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dataset. Must not exceed 127 characters.",
			},
			"patsetfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "File which contains list of patterns that needs to be bound to the dataset. A patsetfile cannot be associated with multiple datasets.",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of value to bind to the dataset.",
			},
		},
	}
}

func policydatasetGetThePayloadFromtheConfig(ctx context.Context, data *PolicydatasetResourceModel) policy.Policydataset {
	tflog.Debug(ctx, "In policydatasetGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policydataset := policy.Policydataset{}
	if !data.Comment.IsNull() {
		policydataset.Comment = data.Comment.ValueString()
	}
	if !data.Dynamic.IsNull() {
		policydataset.Dynamic = data.Dynamic.ValueString()
	}
	if !data.Dynamiconly.IsNull() {
		policydataset.Dynamiconly = data.Dynamiconly.ValueBool()
	}
	if !data.Name.IsNull() {
		policydataset.Name = data.Name.ValueString()
	}
	if !data.Patsetfile.IsNull() {
		policydataset.Patsetfile = data.Patsetfile.ValueString()
	}
	if !data.Type.IsNull() {
		policydataset.Type = data.Type.ValueString()
	}

	return policydataset
}

func policydatasetSetAttrFromGet(ctx context.Context, data *PolicydatasetResourceModel, getResponseData map[string]interface{}) *PolicydatasetResourceModel {
	tflog.Debug(ctx, "In policydatasetSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["dynamic"]; ok && val != nil {
		data.Dynamic = types.StringValue(val.(string))
	} else {
		data.Dynamic = types.StringNull()
	}
	if val, ok := getResponseData["dynamiconly"]; ok && val != nil {
		data.Dynamiconly = types.BoolValue(val.(bool))
	} else {
		data.Dynamiconly = types.BoolNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["patsetfile"]; ok && val != nil {
		data.Patsetfile = types.StringValue(val.(string))
	} else {
		data.Patsetfile = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Name.ValueString()))

	return data
}
