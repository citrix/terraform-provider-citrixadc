package linkset

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
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
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
					setplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Set of interfaces to be bound to the linkset. Changing this forces recreation of the linkset (matches SDK v2 ForceNew behavior).",
			},
		},
	}
}

func linksetGetThePayloadFromthePlan(ctx context.Context, data *LinksetResourceModel) network.Linkset {
	tflog.Debug(ctx, "In linksetGetThePayloadFromthePlan Function")

	// Create API request body from the model
	linkset := network.Linkset{}
	if !data.Linksetid.IsNull() && !data.Linksetid.IsUnknown() {
		linkset.Id = data.Linksetid.ValueString()
	}

	return linkset
}

func linksetSetAttrFromGet(ctx context.Context, data *LinksetResourceModel, getResponseData map[string]interface{}) *LinksetResourceModel {
	tflog.Debug(ctx, "In linksetSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		data.Linksetid = types.StringValue(val.(string))
		// Set ID for the resource
		// Case 2: Single unique attribute - use plain value as ID
		data.Id = types.StringValue(val.(string))
	}

	return data
}

// linksetReadInterfaceBindings reads the interface bindings for the linkset and
// populates the interfacebinding set on the model. It is shared by the resource
// and datasource Read paths. Matching the SDK v2 behavior, a "not found" / error
// while listing bindings is treated as "no bindings" (empty set) rather than a
// hard failure, so a linkset with zero interface bindings reads back cleanly.
func linksetReadInterfaceBindings(ctx context.Context, client *service.NitroClient, data *LinksetResourceModel, linksetName string) error {
	tflog.Debug(ctx, "In linksetReadInterfaceBindings Function")

	bindings, err := client.FindResourceArray(service.Linkset_interface_binding.Type(), linksetName)
	if err != nil {
		// SDK v2 ignores this error and treats it as an empty binding set.
		bindings = []map[string]interface{}{}
	}

	processedBindings := make([]string, 0, len(bindings))
	for _, val := range bindings {
		if ifnum, ok := val["ifnum"].(string); ok {
			processedBindings = append(processedBindings, ifnum)
		}
	}

	interfaceSet, diags := types.SetValueFrom(ctx, types.StringType, processedBindings)
	if diags.HasError() {
		return fmt.Errorf("error converting interface bindings to set")
	}
	data.Interfacebinding = interfaceSet

	return nil
}
