package appfwprofile_fileuploadtype_binding

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
var _ resource.Resource = &AppfwprofileFileuploadtypeBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileFileuploadtypeBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileFileuploadtypeBindingResource)(nil)

func NewAppfwprofileFileuploadtypeBindingResource() resource.Resource {
	return &AppfwprofileFileuploadtypeBindingResource{}
}

// AppfwprofileFileuploadtypeBindingResource defines the resource implementation.
type AppfwprofileFileuploadtypeBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileFileuploadtypeBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fileuploadtype_binding"
}

func (r *AppfwprofileFileuploadtypeBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileFileuploadtypeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_fileuploadtype_binding resource")
	appfwprofile_fileuploadtype_binding := appfwprofile_fileuploadtype_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_fileuploadtype_binding.Type(), &appfwprofile_fileuploadtype_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_fileuploadtype_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_fileuploadtype_binding resource")

	// Set ID for the resource before reading state
	// filetype is a list; it is encoded as a ';'-joined string to match the legacy
	// SDK v2 composite ID format.
	data.Id = types.StringValue(appfwprofile_fileuploadtype_bindingComposeId(&data))

	// Read the updated state back
	r.readAppfwprofileFileuploadtypeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileFileuploadtypeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_fileuploadtype_binding resource")

	r.readAppfwprofileFileuploadtypeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileFileuploadtypeBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_fileuploadtype_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_fileuploadtype_binding := appfwprofile_fileuploadtype_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_fileuploadtype_binding.Type(), &appfwprofile_fileuploadtype_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_fileuploadtype_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_fileuploadtype_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_fileuploadtype_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileFileuploadtypeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileFileuploadtypeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_fileuploadtype_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "fileuploadtype", "as_fileuploadtypes_url", "filetype"}, nil)
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
	if val, ok := idMap["as_fileuploadtypes_url"]; ok && val != "" {
		argsMap["as_fileuploadtypes_url"] = val
	}
	if val, ok := idMap["filetype"]; ok && val != "" {
		argsMap["filetype"] = val
	}
	if val, ok := idMap["fileuploadtype"]; ok && val != "" {
		argsMap["fileuploadtype"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_fileuploadtype_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_fileuploadtype_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_fileuploadtype_binding binding")
}

// Helper function to read appfwprofile_fileuploadtype_binding data from API
func (r *AppfwprofileFileuploadtypeBindingResource) readAppfwprofileFileuploadtypeBindingFromApi(ctx context.Context, data *AppfwprofileFileuploadtypeBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "fileuploadtype", "as_fileuploadtypes_url", "filetype"}, nil)
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
		ResourceType:             service.Appfwprofile_fileuploadtype_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fileuploadtype_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_fileuploadtype_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_fileuploadtypes_url
		if idVal, ok := idMap["as_fileuploadtypes_url"]; ok {
			if val, ok := v["as_fileuploadtypes_url"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_fileuploadtypes_url"].(string); ok {
			match = false
			continue
		}

		// Check filetype. The GET response returns filetype as a list, which the
		// composite ID encodes as a ';'-joined string. Join the response list the
		// same way before comparing.
		if idVal, ok := idMap["filetype"]; ok {
			dataFiletype := ""
			if v["filetype"] != nil {
				if filetypeSlice, ok := v["filetype"].([]interface{}); ok {
					dataFiletype = strings.Join(utils.ToStringList(filetypeSlice), ";")
				}
			}
			if dataFiletype != idVal {
				match = false
				continue
			}
		}

		// Check fileuploadtype
		if idVal, ok := idMap["fileuploadtype"]; ok {
			if val, ok := v["fileuploadtype"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["fileuploadtype"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_fileuploadtype_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_fileuploadtype_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
