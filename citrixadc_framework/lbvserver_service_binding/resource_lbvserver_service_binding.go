package lbvserver_service_binding

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
var _ resource.Resource = &LbvserverServiceBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverServiceBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverServiceBindingResource)(nil)

func NewLbvserverServiceBindingResource() resource.Resource {
	return &LbvserverServiceBindingResource{}
}

// LbvserverServiceBindingResource defines the resource implementation.
type LbvserverServiceBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverServiceBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverServiceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_service_binding"
}

func (r *LbvserverServiceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverServiceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverServiceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_service_binding resource")
	lbvserver_service_binding := lbvserver_service_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call.
	// NITRO binding "add" is POST (idempotent) here, matching the SDK v2 resource.
	// UpdateUnnamedResource (PUT) returns errorcode 273 "Resource already exists" on
	// re-bind, so use AddResource (Pattern 1).
	_, err := r.client.AddResource(service.Lbvserver_service_binding.Type(), "", &lbvserver_service_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_service_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbvserver_service_binding resource")

	// Set ID for the resource before reading state.
	// ID identity matches the SDK v2 contract and resource_id_mapping.json: "name,servicename".
	// servicegroupname is a valid bind target attribute but is NOT part of the resource identity.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLbvserverServiceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverServiceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverServiceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_service_binding resource")

	r.readLbvserverServiceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverServiceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbvserverServiceBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for lbvserver_service_binding: NITRO does not allow a binding
	// to be modified in place (errorcode 273 on re-bind), so every attribute is
	// RequiresReplace and Terraform handles changes via destroy + create (unbind +
	// rebind), matching the SDK v2 ForceNew contract (Pattern 5).
	tflog.Debug(ctx, "Update is a no-op for lbvserver_service_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readLbvserverServiceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverServiceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverServiceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_service_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "servicename"}, nil)
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
	if val, ok := idMap["servicename"]; ok && val != "" {
		// URL-encode in case the servicename contains slashy/special characters.
		argsMap["servicename"] = url.QueryEscape(val)
	}
	// servicegroupname is not part of the identity (ID), but is a valid alternate
	// bind target; include it from state when present so the delete targets the
	// correct binding.
	if !data.Servicegroupname.IsNull() && data.Servicegroupname.ValueString() != "" {
		argsMap["servicegroupname"] = url.QueryEscape(data.Servicegroupname.ValueString())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lbvserver_service_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbvserver_service_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbvserver_service_binding binding")
}

// Helper function to read lbvserver_service_binding data from API
func (r *LbvserverServiceBindingResource) readLbvserverServiceBindingFromApi(ctx context.Context, data *LbvserverServiceBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "servicename"}, nil)
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
		ResourceType:             service.Lbvserver_service_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_service_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lbvserver_service_binding returned empty array.")
		return
	}

	// Iterate through results to find the one matching the binding identity.
	// Identity is name + servicename (per SDK v2 contract / resource_id_mapping.json).
	servicename_id := idMap["servicename"]
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["servicename"].(string); ok {
			if val == servicename_id {
				foundIndex = i
				break
			}
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("lbvserver_service_binding not found with the provided ID attributes"))
		return
	}

	lbvserver_service_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
