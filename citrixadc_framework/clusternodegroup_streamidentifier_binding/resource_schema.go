package clusternodegroup_streamidentifier_binding

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

// ClusternodegroupStreamidentifierBindingResourceModel describes the resource data model.
type ClusternodegroupStreamidentifierBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Identifiername types.String `tfsdk:"identifiername"`
	Name           types.String `tfsdk:"name"`
}

func (r *ClusternodegroupStreamidentifierBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternodegroup_streamidentifier_binding resource.",
			},
			"identifiername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "stream identifier  and rate limit identifier that need to be bound to this nodegroup.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup to which you want to bind a cluster node or an entity.",
			},
		},
	}
}

func clusternodegroup_streamidentifier_bindingGetThePayloadFromtheConfig(ctx context.Context, data *ClusternodegroupStreamidentifierBindingResourceModel) cluster.Clusternodegroupstreamidentifierbinding {
	tflog.Debug(ctx, "In clusternodegroup_streamidentifier_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	clusternodegroup_streamidentifier_binding := cluster.Clusternodegroupstreamidentifierbinding{}
	if !data.Identifiername.IsNull() {
		clusternodegroup_streamidentifier_binding.Identifiername = data.Identifiername.ValueString()
	}
	if !data.Name.IsNull() {
		clusternodegroup_streamidentifier_binding.Name = data.Name.ValueString()
	}

	return clusternodegroup_streamidentifier_binding
}

func clusternodegroup_streamidentifier_bindingSetAttrFromGet(ctx context.Context, data *ClusternodegroupStreamidentifierBindingResourceModel, getResponseData map[string]interface{}) *ClusternodegroupStreamidentifierBindingResourceModel {
	tflog.Debug(ctx, "In clusternodegroup_streamidentifier_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["identifiername"]; ok && val != nil {
		data.Identifiername = types.StringValue(val.(string))
	} else {
		data.Identifiername = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("identifiername:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Identifiername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
