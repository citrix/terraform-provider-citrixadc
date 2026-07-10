package policypatset_pattern_binding

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PolicypatsetPatternBindingResource{}
var _ resource.ResourceWithConfigure = (*PolicypatsetPatternBindingResource)(nil)
var _ resource.ResourceWithImportState = (*PolicypatsetPatternBindingResource)(nil)

func NewPolicypatsetPatternBindingResource() resource.Resource {
	return &PolicypatsetPatternBindingResource{}
}

// PolicypatsetPatternBindingResource defines the resource implementation.
type PolicypatsetPatternBindingResource struct {
	client *service.NitroClient
}

func (r *PolicypatsetPatternBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicypatsetPatternBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policypatset_pattern_binding"
}

func (r *PolicypatsetPatternBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicypatsetPatternBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicypatsetPatternBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policypatset_pattern_binding resource")
	policypatset_pattern_binding := policypatset_pattern_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource (PUT), matching the SDK v2 behaviour
	err := r.client.UpdateUnnamedResource(service.Policypatset_pattern_binding.Type(), &policypatset_pattern_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policypatset_pattern_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created policypatset_pattern_binding resource")

	// Set composite ID (name:string) for the resource before reading state
	data.Id = types.StringValue(policypatset_pattern_bindingComputeId(&data))

	// Read the updated state back
	r.readPolicypatsetPatternBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetPatternBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicypatsetPatternBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policypatset_pattern_binding resource")

	r.readPolicypatsetPatternBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetPatternBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state PolicypatsetPatternBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state. All attributes are RequiresReplace, so Update is a
	// no-op: Terraform recreates the binding rather than calling Update for any change.
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for policypatset_pattern_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readPolicypatsetPatternBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetPatternBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicypatsetPatternBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policypatset_pattern_binding resource")

	// Binding with parent - delete using DeleteResourceWithArgsMap.
	// Parse the ID (handles both the new name:string form and the legacy name,string form).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "string"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	argsMap := make(map[string]string)
	if val, ok := idMap["string"]; ok && val != "" {
		// URL-encode the slashy/special characters that may appear in the pattern string.
		argsMap["String"] = url.QueryEscape(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Policypatset_pattern_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete policypatset_pattern_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted policypatset_pattern_binding binding")
}

// Helper function to read policypatset_pattern_binding data from API
func (r *PolicypatsetPatternBindingResource) readPolicypatsetPatternBindingFromApi(ctx context.Context, data *PolicypatsetPatternBindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID (supports legacy name,string form)
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "string"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}
	stringText, ok := idMap["string"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'string' not found in ID string")
		return
	}

	findParams := service.FindParams{
		ResourceType:             service.Policypatset_pattern_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 2823,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policypatset_pattern_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		tflog.Warn(ctx, fmt.Sprintf("policypatset_pattern_binding %s returned empty array; clearing state", data.Id.ValueString()))
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right String
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["String"].(string); ok && val == stringText {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		tflog.Warn(ctx, fmt.Sprintf("policypatset_pattern_binding %s String not found in array; clearing state", data.Id.ValueString()))
		data.Id = types.StringNull()
		return
	}

	// Ensure the identity keys are populated from the ID before applying the GET response.
	data.Name = types.StringValue(name_Name)
	data.String = types.StringValue(stringText)

	policypatset_pattern_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
