package mapdomain_mapbmr_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// MapdomainMapbmrBindingResourceModel describes the resource data model.
type MapdomainMapbmrBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Mapbmrname types.String `tfsdk:"mapbmrname"`
	Name       types.String `tfsdk:"name"`
}

func (r *MapdomainMapbmrBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the mapdomain_mapbmr_binding resource.",
			},
			"mapbmrname": schema.StringAttribute{
				Required:    true,
				Description: "Basic Mapping rule name.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the MAP Domain. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Domain is created . The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"add network MapDomain map1\").",
			},
		},
	}
}

func mapdomain_mapbmr_bindingGetThePayloadFromtheConfig(ctx context.Context, data *MapdomainMapbmrBindingResourceModel) network.Mapdomainmapbmrbinding {
	tflog.Debug(ctx, "In mapdomain_mapbmr_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	mapdomain_mapbmr_binding := network.Mapdomainmapbmrbinding{}
	if !data.Mapbmrname.IsNull() {
		mapdomain_mapbmr_binding.Mapbmrname = data.Mapbmrname.ValueString()
	}
	if !data.Name.IsNull() {
		mapdomain_mapbmr_binding.Name = data.Name.ValueString()
	}

	return mapdomain_mapbmr_binding
}

func mapdomain_mapbmr_bindingSetAttrFromGet(ctx context.Context, data *MapdomainMapbmrBindingResourceModel, getResponseData map[string]interface{}) *MapdomainMapbmrBindingResourceModel {
	tflog.Debug(ctx, "In mapdomain_mapbmr_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["mapbmrname"]; ok && val != nil {
		data.Mapbmrname = types.StringValue(val.(string))
	} else {
		data.Mapbmrname = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("mapbmrname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Mapbmrname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
