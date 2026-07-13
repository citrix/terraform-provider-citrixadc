package clusternodegroup_csvserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ClusternodegroupCsvserverBindingResourceModel describes the resource data model.
type ClusternodegroupCsvserverBindingResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Vserver types.String `tfsdk:"vserver"`
}

func (r *ClusternodegroupCsvserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternodegroup_csvserver_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "vserver that need to be bound to this nodegroup.",
			},
		},
	}
}

func clusternodegroup_csvserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *ClusternodegroupCsvserverBindingResourceModel) cluster.Clusternodegroupcsvserverbinding {
	tflog.Debug(ctx, "In clusternodegroup_csvserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	clusternodegroup_csvserver_binding := cluster.Clusternodegroupcsvserverbinding{}
	if !data.Name.IsNull() {
		clusternodegroup_csvserver_binding.Name = data.Name.ValueString()
	}
	if !data.Vserver.IsNull() {
		clusternodegroup_csvserver_binding.Vserver = data.Vserver.ValueString()
	}

	return clusternodegroup_csvserver_binding
}

func clusternodegroup_csvserver_bindingSetAttrFromGet(ctx context.Context, data *ClusternodegroupCsvserverBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupCsvserverBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_csvserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else {
		data.Vserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vserver:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
