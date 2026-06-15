package gslbservice_dnsview_binding

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
var _ resource.Resource = &GslbserviceDnsviewBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbserviceDnsviewBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbserviceDnsviewBindingResource)(nil)

func NewGslbserviceDnsviewBindingResource() resource.Resource {
	return &GslbserviceDnsviewBindingResource{}
}

// GslbserviceDnsviewBindingResource defines the resource implementation.
type GslbserviceDnsviewBindingResource struct {
	client *service.NitroClient
}

func (r *GslbserviceDnsviewBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbserviceDnsviewBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservice_dnsview_binding"
}

func (r *GslbserviceDnsviewBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbserviceDnsviewBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbserviceDnsviewBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservice_dnsview_binding resource")
	gslbservice_dnsview_binding := gslbservice_dnsview_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Gslbservice_dnsview_binding.Type(), &gslbservice_dnsview_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservice_dnsview_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created gslbservice_dnsview_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("viewname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Viewname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readGslbserviceDnsviewBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceDnsviewBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbserviceDnsviewBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservice_dnsview_binding resource")

	r.readGslbserviceDnsviewBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceDnsviewBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state GslbserviceDnsviewBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No-op: every schema attribute (servicename, viewip, viewname) is RequiresReplace,
	// matching the SDK v2 resource which had no Update (all ForceNew). The NITRO binding
	// has no in-place update endpoint, so Terraform forces recreation on any change and
	// Update is never reached with an actual diff.
	tflog.Debug(ctx, "Update is a no-op for gslbservice_dnsview_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readGslbserviceDnsviewBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceDnsviewBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbserviceDnsviewBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservice_dnsview_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicename", "viewname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicename_value, ok := idMap["servicename"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicename' not found in ID")
		return
	}

	// Build delete args; URL-encode the viewname value so slashy/special characters
	// survive in the DELETE ?args=viewname:<value> query (the NITRO client does not
	// encode arg values, only the resource name).
	args := make([]string, 0)
	if val, ok := idMap["viewname"]; ok && val != "" {
		args = append(args, fmt.Sprintf("viewname:%s", url.QueryEscape(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Gslbservice_dnsview_binding.Type(), servicename_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete gslbservice_dnsview_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted gslbservice_dnsview_binding binding")
}

// Helper function to read gslbservice_dnsview_binding data from API
func (r *GslbserviceDnsviewBindingResource) readGslbserviceDnsviewBindingFromApi(ctx context.Context, data *GslbserviceDnsviewBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicename", "viewname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicename_Name, ok := idMap["servicename"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicename' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Gslbservice_dnsview_binding.Type(),
		ResourceName:             servicename_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservice_dnsview_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "gslbservice_dnsview_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check viewname
		if idVal, ok := idMap["viewname"]; ok {
			if val, ok := v["viewname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["viewname"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("gslbservice_dnsview_binding not found with the provided ID attributes"))
		return
	}

	gslbservice_dnsview_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
