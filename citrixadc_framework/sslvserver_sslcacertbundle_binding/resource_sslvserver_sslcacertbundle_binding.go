package sslvserver_sslcacertbundle_binding

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
var _ resource.Resource = &SslvserverSslcacertbundleBindingResource{}
var _ resource.ResourceWithConfigure = (*SslvserverSslcacertbundleBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslvserverSslcacertbundleBindingResource)(nil)

func NewSslvserverSslcacertbundleBindingResource() resource.Resource {
	return &SslvserverSslcacertbundleBindingResource{}
}

// SslvserverSslcacertbundleBindingResource defines the resource implementation.
type SslvserverSslcacertbundleBindingResource struct {
	client *service.NitroClient
}

func (r *SslvserverSslcacertbundleBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslvserverSslcacertbundleBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslcacertbundle_binding"
}

func (r *SslvserverSslcacertbundleBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslvserverSslcacertbundleBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslvserverSslcacertbundleBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslvserver_sslcacertbundle_binding resource")
	sslvserver_sslcacertbundle_binding := sslvserver_sslcacertbundle_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslvserver_sslcacertbundle_binding.Type(), &sslvserver_sslcacertbundle_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslvserver_sslcacertbundle_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslvserver_sslcacertbundle_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("cacertbundlename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Cacertbundlename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslvserverSslcacertbundleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcacertbundleBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslvserverSslcacertbundleBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslvserver_sslcacertbundle_binding resource")

	r.readSslvserverSslcacertbundleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcacertbundleBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslvserverSslcacertbundleBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslvserver_sslcacertbundle_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslvserver_sslcacertbundle_binding := sslvserver_sslcacertbundle_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslvserver_sslcacertbundle_binding.Type(), &sslvserver_sslcacertbundle_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslvserver_sslcacertbundle_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslvserver_sslcacertbundle_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslvserver_sslcacertbundle_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslvserverSslcacertbundleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcacertbundleBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslvserverSslcacertbundleBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslvserver_sslcacertbundle_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vservername_value, ok := idMap["vservername"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vservername' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["cacertbundlename"]; ok && val != "" {
		argsMap["cacertbundlename"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslvserver_sslcacertbundle_binding.Type(), vservername_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslvserver_sslcacertbundle_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslvserver_sslcacertbundle_binding binding")
}

// Helper function to read sslvserver_sslcacertbundle_binding data from API
func (r *SslvserverSslcacertbundleBindingResource) readSslvserverSslcacertbundleBindingFromApi(ctx context.Context, data *SslvserverSslcacertbundleBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	vservername_Name, ok := idMap["vservername"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'vservername' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslvserver_sslcacertbundle_binding.Type(),
		ResourceName:             vservername_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslcacertbundle_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslvserver_sslcacertbundle_binding returned empty array.")
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

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("sslvserver_sslcacertbundle_binding not found with the provided ID attributes"))
		return
	}

	sslvserver_sslcacertbundle_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
