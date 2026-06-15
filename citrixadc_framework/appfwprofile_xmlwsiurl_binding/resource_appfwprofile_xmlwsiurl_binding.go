package appfwprofile_xmlwsiurl_binding

import (
	"context"
	"fmt"
	"net/url"
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
var _ resource.Resource = &AppfwprofileXmlwsiurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileXmlwsiurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileXmlwsiurlBindingResource)(nil)

func NewAppfwprofileXmlwsiurlBindingResource() resource.Resource {
	return &AppfwprofileXmlwsiurlBindingResource{}
}

// AppfwprofileXmlwsiurlBindingResource defines the resource implementation.
type AppfwprofileXmlwsiurlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileXmlwsiurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_xmlwsiurl_binding"
}

func (r *AppfwprofileXmlwsiurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileXmlwsiurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_xmlwsiurl_binding resource")
	appfwprofile_xmlwsiurl_binding := appfwprofile_xmlwsiurl_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlwsiurl_binding.Type(), &appfwprofile_xmlwsiurl_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_xmlwsiurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_xmlwsiurl_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("xmlwsiurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Xmlwsiurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileXmlwsiurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileXmlwsiurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_xmlwsiurl_binding resource")

	r.readAppfwprofileXmlwsiurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileXmlwsiurlBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_xmlwsiurl_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_xmlwsiurl_binding := appfwprofile_xmlwsiurl_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlwsiurl_binding.Type(), &appfwprofile_xmlwsiurl_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_xmlwsiurl_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_xmlwsiurl_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_xmlwsiurl_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileXmlwsiurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileXmlwsiurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_xmlwsiurl_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "xmlwsiurl"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// Mirror the SDK v2 delete: pass URL-encoded arg values via DeleteResourceWithArgs
	// (idMap values are already URL-decoded, so re-encode them here), and include ruletype.
	args := make([]string, 0)
	if val, ok := idMap["xmlwsiurl"]; ok && val != "" {
		args = append(args, fmt.Sprintf("xmlwsiurl:%s", url.QueryEscape(val)))
	}
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(data.Ruletype.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Appfwprofile_xmlwsiurl_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_xmlwsiurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_xmlwsiurl_binding binding")
}

// Helper function to read appfwprofile_xmlwsiurl_binding data from API
func (r *AppfwprofileXmlwsiurlBindingResource) readAppfwprofileXmlwsiurlBindingFromApi(ctx context.Context, data *AppfwprofileXmlwsiurlBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "xmlwsiurl"}, nil)
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
		ResourceType:             service.Appfwprofile_xmlwsiurl_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_xmlwsiurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_xmlwsiurl_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check xmlwsiurl
		if idVal, ok := idMap["xmlwsiurl"]; ok {
			if val, ok := v["xmlwsiurl"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["xmlwsiurl"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_xmlwsiurl_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_xmlwsiurl_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
