package sslservicegroup_ecccurve_binding

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
var _ resource.Resource = &SslservicegroupEcccurveBindingResource{}
var _ resource.ResourceWithConfigure = (*SslservicegroupEcccurveBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslservicegroupEcccurveBindingResource)(nil)

func NewSslservicegroupEcccurveBindingResource() resource.Resource {
	return &SslservicegroupEcccurveBindingResource{}
}

// SslservicegroupEcccurveBindingResource defines the resource implementation.
type SslservicegroupEcccurveBindingResource struct {
	client *service.NitroClient
}

func (r *SslservicegroupEcccurveBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslservicegroupEcccurveBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_ecccurve_binding"
}

func (r *SslservicegroupEcccurveBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslservicegroupEcccurveBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslservicegroupEcccurveBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservicegroup_ecccurve_binding resource")
	sslservicegroup_ecccurve_binding := sslservicegroup_ecccurve_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslservicegroup_ecccurve_binding.Type(), &sslservicegroup_ecccurve_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservicegroup_ecccurve_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslservicegroup_ecccurve_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ecccurvename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ecccurvename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslservicegroupEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupEcccurveBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslservicegroupEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservicegroup_ecccurve_binding resource")

	r.readSslservicegroupEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupEcccurveBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslservicegroupEcccurveBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservicegroup_ecccurve_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslservicegroup_ecccurve_binding := sslservicegroup_ecccurve_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslservicegroup_ecccurve_binding.Type(), &sslservicegroup_ecccurve_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservicegroup_ecccurve_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslservicegroup_ecccurve_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslservicegroup_ecccurve_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslservicegroupEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupEcccurveBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslservicegroupEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservicegroup_ecccurve_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "ecccurvename"}, nil)
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
	if val, ok := idMap["ecccurvename"]; ok && val != "" {
		// URL-encode the arg value (matches SDK v2 url.QueryEscape) so slashy/special
		// ecccurve values are passed correctly in the NITRO args= query string.
		argsMap["ecccurvename"] = url.QueryEscape(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslservicegroup_ecccurve_binding.Type(), servicegroupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslservicegroup_ecccurve_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslservicegroup_ecccurve_binding binding")
}

// Helper function to read sslservicegroup_ecccurve_binding data from API
func (r *SslservicegroupEcccurveBindingResource) readSslservicegroupEcccurveBindingFromApi(ctx context.Context, data *SslservicegroupEcccurveBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "ecccurvename"}, nil)
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
		ResourceType:             service.Sslservicegroup_ecccurve_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_ecccurve_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslservicegroup_ecccurve_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ecccurvename
		if idVal, ok := idMap["ecccurvename"]; ok {
			if val, ok := v["ecccurvename"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ecccurvename"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("sslservicegroup_ecccurve_binding not found with the provided ID attributes"))
		return
	}

	sslservicegroup_ecccurve_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
