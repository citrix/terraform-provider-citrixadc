package csvserver_lbvserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsvserverLbvserverBindingResourceModel describes the resource data model.
type CsvserverLbvserverBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Lbvserver     types.String `tfsdk:"lbvserver"`
	Name          types.String `tfsdk:"name"`
	Targetvserver types.String `tfsdk:"targetvserver"`
}

func (r *CsvserverLbvserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csvserver_lbvserver_binding resource.",
			},
			"lbvserver": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the default lb vserver bound. Use this param for Default binding only. For Example: bind cs vserver cs1 -lbvserver lb1",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"targetvserver": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The virtual server name (created with the add lb vserver command) to which content will be switched.",
			},
		},
	}
}

func csvserver_lbvserver_bindingGetThePayloadFromthePlan(ctx context.Context, data *CsvserverLbvserverBindingResourceModel) cs.Csvserverlbvserverbinding {
	tflog.Debug(ctx, "In csvserver_lbvserver_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	csvserver_lbvserver_binding := cs.Csvserverlbvserverbinding{}
	if !data.Lbvserver.IsNull() {
		csvserver_lbvserver_binding.Lbvserver = data.Lbvserver.ValueString()
	}
	if !data.Name.IsNull() {
		csvserver_lbvserver_binding.Name = data.Name.ValueString()
	}
	if !data.Targetvserver.IsNull() {
		csvserver_lbvserver_binding.Targetvserver = data.Targetvserver.ValueString()
	}

	return csvserver_lbvserver_binding
}

func csvserver_lbvserver_bindingSetAttrFromGet(ctx context.Context, data *CsvserverLbvserverBindingResourceModel, getResponseData map[string]interface{}) *CsvserverLbvserverBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_lbvserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["lbvserver"]; ok && val != nil {
		data.Lbvserver = types.StringValue(val.(string))
	} else {
		data.Lbvserver = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["targetvserver"]; ok && val != nil {
		data.Targetvserver = types.StringValue(val.(string))
	} else {
		data.Targetvserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes
	bindingId := fmt.Sprintf("%s,%s", data.Name.ValueString(), data.Lbvserver.ValueString())
	data.Id = types.StringValue(bindingId)

	return data
}
