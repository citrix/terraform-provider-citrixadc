package sslprofile_sslcipher_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslprofileSslcipherBindingResource{}
var _ resource.ResourceWithConfigure = (*SslprofileSslcipherBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslprofileSslcipherBindingResource)(nil)

func NewSslprofileSslcipherBindingResource() resource.Resource {
	return &SslprofileSslcipherBindingResource{}
}

// SslprofileSslcipherBindingResource defines the resource implementation.
type SslprofileSslcipherBindingResource struct {
	client *service.NitroClient
}

func (r *SslprofileSslcipherBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslprofileSslcipherBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile_sslcipher_binding"
}

func (r *SslprofileSslcipherBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslprofileSslcipherBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslprofileSslcipherBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslprofile_sslcipher_binding resource")
	sslprofile_sslcipher_binding := sslprofile_sslcipher_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslprofile_sslcipher_binding.Type(), &sslprofile_sslcipher_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslprofile_sslcipher_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslprofile_sslcipher_binding resource")

	// Set ID for the resource before reading state (new key:UrlEncode(value)
	// format; attribute order matches the legacy resource_id_mapping.json order
	// "name,ciphername").
	data.Id = types.StringValue(sslprofile_sslcipher_bindingComputeId(&data))

	// Read the updated state back
	r.readSslprofileSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslcipherBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslprofileSslcipherBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslprofile_sslcipher_binding resource")

	r.readSslprofileSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslcipherBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslprofileSslcipherBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslprofile_sslcipher_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Ciphername.Equal(state.Ciphername) {
		tflog.Debug(ctx, fmt.Sprintf("ciphername has changed for sslprofile_sslcipher_binding"))
		hasChange = true
	}
	if !data.Cipherpriority.Equal(state.Cipherpriority) {
		tflog.Debug(ctx, fmt.Sprintf("cipherpriority has changed for sslprofile_sslcipher_binding"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		sslprofile_sslcipher_binding := sslprofile_sslcipher_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslprofile_sslcipher_binding.Type(), &sslprofile_sslcipher_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslprofile_sslcipher_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslprofile_sslcipher_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslprofile_sslcipher_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslprofileSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslcipherBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslprofileSslcipherBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslprofile_sslcipher_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "ciphername"}, nil)
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
	if val, ok := idMap["ciphername"]; ok && val != "" {
		argsMap["ciphername"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslprofile_sslcipher_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslprofile_sslcipher_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslprofile_sslcipher_binding binding")
}

// Helper function to read sslprofile_sslcipher_binding data from API
func (r *SslprofileSslcipherBindingResource) readSslprofileSslcipherBindingFromApi(ctx context.Context, data *SslprofileSslcipherBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "ciphername"}, nil)
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
		ResourceType:             service.Sslprofile_sslcipher_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_sslcipher_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslprofile_sslcipher_binding returned empty array.")
		return
	}

	// Iterate through results to find the matching cipher binding. The NITRO
	// GET response exposes the bound cipher's name in "cipheraliasname"
	// (there is no "ciphername" key), so match the ID's ciphername against
	// the response's cipheraliasname.
	foundIndex := -1
	if cipherVal, ok := idMap["ciphername"]; ok {
		for i, v := range dataArr {
			if val, ok := v["cipheraliasname"].(string); ok && val == cipherVal {
				foundIndex = i
				break
			}
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("sslprofile_sslcipher_binding not found with the provided ID attributes"))
		return
	}

	sslprofile_sslcipher_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
