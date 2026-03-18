package linkset

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LinksetResourceModel describes the resource data model.
type LinksetResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Linksetid        types.String `tfsdk:"linkset_id"`
	Interfacebinding types.Set    `tfsdk:"interfacebinding"`
}

func (r *LinksetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the linkset resource.",
			},
			"linkset_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Unique identifier for the linkset. Must be of the form LS/x, where x can be an integer from 1 to 32.",
			},
			"interfacebinding": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Set of interface bindings associated with the linkset.",
			},
		},
	}
}

func linksetGetThePayloadFromtheConfig(ctx context.Context, data *LinksetResourceModel) network.Linkset {
	tflog.Debug(ctx, "In linksetGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	linkset := network.Linkset{}
	if !data.Id.IsNull() {
		linkset.Id = data.Id.ValueString()
	}

	return linkset
}

func linksetSetAttrFromGet(ctx context.Context, data *LinksetResourceModel, getResponseData map[string]interface{}) *LinksetResourceModel {
	tflog.Debug(ctx, "In linksetSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		data.Linksetid = types.StringValue(val.(string))
	} else {
		data.Linksetid = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Linksetid.ValueString())

	return data
}

func (d *LinksetDataSource) readLinksetInterfaceBindings(ctx context.Context, data *LinksetResourceModel, linksetName string) error {
	bindings, err := d.client.FindResourceArray(service.Linkset_interface_binding.Type(), linksetName)
	if err != nil {
		return fmt.Errorf("unable to read linkset interface bindings: %w", err)
	}

	processedBindings := make([]string, 0, len(bindings))
	for _, val := range bindings {
		if ifnum, ok := val["ifnum"].(string); ok {
			processedBindings = append(processedBindings, ifnum)
		}
	}

	// Convert to appropriate Framework type (adjust based on your LinksetResourceModel schema)
	// Example if interfacebinding is types.List or types.Set:
	interfaceSet, diags := types.SetValueFrom(ctx, types.StringType, processedBindings)
	if diags.HasError() {
		return fmt.Errorf("error converting interface bindings to set")
	}
	data.Interfacebinding = interfaceSet

	return nil
}
