package apiprofile_apispec_binding

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
var _ resource.Resource = &ApiprofileApispecBindingResource{}
var _ resource.ResourceWithConfigure = (*ApiprofileApispecBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ApiprofileApispecBindingResource)(nil)

func NewApiprofileApispecBindingResource() resource.Resource {
	return &ApiprofileApispecBindingResource{}
}

// ApiprofileApispecBindingResource defines the resource implementation.
type ApiprofileApispecBindingResource struct {
	client *service.NitroClient
}

func (r *ApiprofileApispecBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ApiprofileApispecBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiprofile_apispec_binding"
}

func (r *ApiprofileApispecBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ApiprofileApispecBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApiprofileApispecBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating apiprofile_apispec_binding resource")
	apiprofile_apispec_binding := apiprofile_apispec_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Apiprofile_apispec_binding.Type(), &apiprofile_apispec_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create apiprofile_apispec_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created apiprofile_apispec_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("apispec:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Apispec.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readApiprofileApispecBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiprofileApispecBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApiprofileApispecBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading apiprofile_apispec_binding resource")

	r.readApiprofileApispecBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiprofileApispecBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ApiprofileApispecBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for apiprofile_apispec_binding: NITRO exposes only add/delete
	// (bind/unbind) and no update endpoint. Both name and apispec are RequiresReplace,
	// so Terraform never reaches Update for an attribute change.
	tflog.Debug(ctx, "Update is a no-op for apiprofile_apispec_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readApiprofileApispecBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiprofileApispecBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApiprofileApispecBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting apiprofile_apispec_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["apispec"]; ok && val != "" {
		// apispec is a required delete arg (NITRO delete URL: args=apispec:<value>).
		// UrlEncode defensively since apispec is a user-supplied string (maxlen 127).
		argsMap["apispec"] = utils.UrlEncode(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Apiprofile_apispec_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete apiprofile_apispec_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted apiprofile_apispec_binding binding")
}

// Helper function to read apiprofile_apispec_binding data from API
func (r *ApiprofileApispecBindingResource) readApiprofileApispecBindingFromApi(ctx context.Context, data *ApiprofileApispecBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Apiprofile_apispec_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read apiprofile_apispec_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "apiprofile_apispec_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check apispec
		if idVal, ok := idMap["apispec"]; ok {
			if val, ok := v["apispec"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["apispec"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("apiprofile_apispec_binding not found with the provided ID attributes"))
		return
	}

	apiprofile_apispec_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}