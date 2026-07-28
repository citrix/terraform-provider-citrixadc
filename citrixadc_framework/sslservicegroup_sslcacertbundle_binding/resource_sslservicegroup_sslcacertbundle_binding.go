package sslservicegroup_sslcacertbundle_binding

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
var _ resource.Resource = &SslservicegroupSslcacertbundleBindingResource{}
var _ resource.ResourceWithConfigure = (*SslservicegroupSslcacertbundleBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslservicegroupSslcacertbundleBindingResource)(nil)

func NewSslservicegroupSslcacertbundleBindingResource() resource.Resource {
	return &SslservicegroupSslcacertbundleBindingResource{}
}

// SslservicegroupSslcacertbundleBindingResource defines the resource implementation.
type SslservicegroupSslcacertbundleBindingResource struct {
	client *service.NitroClient
}

func (r *SslservicegroupSslcacertbundleBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslservicegroupSslcacertbundleBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_sslcacertbundle_binding"
}

func (r *SslservicegroupSslcacertbundleBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslservicegroupSslcacertbundleBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslservicegroupSslcacertbundleBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservicegroup_sslcacertbundle_binding resource")
	sslservicegroup_sslcacertbundle_binding := sslservicegroup_sslcacertbundle_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslcacertbundle_binding.Type(), &sslservicegroup_sslcacertbundle_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservicegroup_sslcacertbundle_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslservicegroup_sslcacertbundle_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("cacertbundlename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Cacertbundlename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslservicegroupSslcacertbundleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslcacertbundleBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslservicegroupSslcacertbundleBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservicegroup_sslcacertbundle_binding resource")

	r.readSslservicegroupSslcacertbundleBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the binding is gone out-of-band, remove it from state so a subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslcacertbundleBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslservicegroupSslcacertbundleBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservicegroup_sslcacertbundle_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslservicegroup_sslcacertbundle_binding := sslservicegroup_sslcacertbundle_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslcacertbundle_binding.Type(), &sslservicegroup_sslcacertbundle_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservicegroup_sslcacertbundle_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslservicegroup_sslcacertbundle_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslservicegroup_sslcacertbundle_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslservicegroupSslcacertbundleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslcacertbundleBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslservicegroupSslcacertbundleBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservicegroup_sslcacertbundle_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicegroupname_value, ok := idMap["servicegroupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicegroupname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["cacertbundlename"]; ok && val != "" {
		argsMap["cacertbundlename"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslservicegroup_sslcacertbundle_binding.Type(), servicegroupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslservicegroup_sslcacertbundle_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslservicegroup_sslcacertbundle_binding binding")
}

// Helper function to read sslservicegroup_sslcacertbundle_binding data from API
func (r *SslservicegroupSslcacertbundleBindingResource) readSslservicegroupSslcacertbundleBindingFromApi(ctx context.Context, data *SslservicegroupSslcacertbundleBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicegroupname_Name, ok := idMap["servicegroupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicegroupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslservicegroup_sslcacertbundle_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_sslcacertbundle_binding, got error: %s", err))
		return
	}

	// Resource is missing - deleted out-of-band; signal "gone" so Read removes it from state.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check cacertbundlename
		if idVal, ok := idMap["cacertbundlename"]; ok {
			if val, ok := v["cacertbundlename"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["cacertbundlename"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing - deleted out-of-band; signal "gone" so Read removes it from state.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	sslservicegroup_sslcacertbundle_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
