package streamidentifier_analyticsprofile_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StreamidentifierAnalyticsprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*StreamidentifierAnalyticsprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*StreamidentifierAnalyticsprofileBindingResource)(nil)

func NewStreamidentifierAnalyticsprofileBindingResource() resource.Resource {
	return &StreamidentifierAnalyticsprofileBindingResource{}
}

// StreamidentifierAnalyticsprofileBindingResource defines the resource implementation.
type StreamidentifierAnalyticsprofileBindingResource struct {
	client *service.NitroClient
}

func (r *StreamidentifierAnalyticsprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *StreamidentifierAnalyticsprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_streamidentifier_analyticsprofile_binding"
}

func (r *StreamidentifierAnalyticsprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *StreamidentifierAnalyticsprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StreamidentifierAnalyticsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating streamidentifier_analyticsprofile_binding resource")
	streamidentifier_analyticsprofile_binding := streamidentifier_analyticsprofile_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Streamidentifier_analyticsprofile_binding.Type(), &streamidentifier_analyticsprofile_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create streamidentifier_analyticsprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created streamidentifier_analyticsprofile_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("analyticsprofile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Analyticsprofile.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readStreamidentifierAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamidentifierAnalyticsprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StreamidentifierAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading streamidentifier_analyticsprofile_binding resource")

	r.readStreamidentifierAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource has been deleted out-of-band - remove from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamidentifierAnalyticsprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state StreamidentifierAnalyticsprofileBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating streamidentifier_analyticsprofile_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		streamidentifier_analyticsprofile_binding := streamidentifier_analyticsprofile_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Streamidentifier_analyticsprofile_binding.Type(), &streamidentifier_analyticsprofile_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update streamidentifier_analyticsprofile_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated streamidentifier_analyticsprofile_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for streamidentifier_analyticsprofile_binding resource, skipping update")
	}

	// Read the updated state back
	r.readStreamidentifierAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamidentifierAnalyticsprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StreamidentifierAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting streamidentifier_analyticsprofile_binding resource")
	// Binding delete: the parent name (streamidentifier) must go in the URL path
	// so DeleteResourceWithArgsMap's pre-delete existence GET hits
	// /streamidentifier_analyticsprofile_binding/<name> (a nameless GET returns
	// errorcode 1095 and the client would wrongly skip the DELETE). The remaining
	// discriminator (analyticsprofile) is passed via the args= query parameter.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name := idMap["name"]

	argsMap := make(map[string]string)
	if val, ok := idMap["analyticsprofile"]; ok && val != "" {
		argsMap["analyticsprofile"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Streamidentifier_analyticsprofile_binding.Type(), name, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete streamidentifier_analyticsprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted streamidentifier_analyticsprofile_binding binding")
}

// Helper function to read streamidentifier_analyticsprofile_binding data from API
func (r *StreamidentifierAnalyticsprofileBindingResource) readStreamidentifierAnalyticsprofileBindingFromApi(ctx context.Context, data *StreamidentifierAnalyticsprofileBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	// Binding GET requires the parent name in the URL path
	// (GET /streamidentifier_analyticsprofile_binding/<name>); a nameless GET
	// returns errorcode 1095 "Name argument required for binding object".
	findParams := service.FindParams{
		ResourceType:             service.Streamidentifier_analyticsprofile_binding.Type(),
		ResourceName:             idMap["name"],
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read streamidentifier_analyticsprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing - signal removal from state
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check analyticsprofile
		if idVal, ok := idMap["analyticsprofile"]; ok {
			if val, ok := v["analyticsprofile"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["analyticsprofile"].(string); ok {
			match = false
			continue
		}

		// Check name
		if idVal, ok := idMap["name"]; ok {
			if val, ok := v["name"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["name"].(string); ok {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing - signal removal from state
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	streamidentifier_analyticsprofile_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
