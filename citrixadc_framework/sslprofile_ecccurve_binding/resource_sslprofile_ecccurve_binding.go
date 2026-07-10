package sslprofile_ecccurve_binding

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslprofileEcccurveBindingResource{}
var _ resource.ResourceWithConfigure = (*SslprofileEcccurveBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslprofileEcccurveBindingResource)(nil)

func NewSslprofileEcccurveBindingResource() resource.Resource {
	return &SslprofileEcccurveBindingResource{}
}

// SslprofileEcccurveBindingResource defines the resource implementation.
type SslprofileEcccurveBindingResource struct {
	client *service.NitroClient
}

func (r *SslprofileEcccurveBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslprofileEcccurveBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile_ecccurve_binding"
}

func (r *SslprofileEcccurveBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ecccurveNamesFromList extracts the configured ECC curve names from the model list.
func ecccurveNamesFromList(ctx context.Context, data *SslprofileEcccurveBindingResourceModel, diags *diag.Diagnostics) []string {
	var names []string
	if data.Ecccurvename.IsNull() || data.Ecccurvename.IsUnknown() {
		return names
	}
	diags.Append(data.Ecccurvename.ElementsAs(ctx, &names, false)...)
	return names
}

func (r *SslprofileEcccurveBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslprofileEcccurveBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslprofile_ecccurve_binding resource")

	name := data.Name.ValueString()

	// Optionally clear pre-existing (default) ecccurve bindings on the profile.
	if data.RemoveExistingEcccurveBinding.ValueBool() {
		tflog.Debug(ctx, fmt.Sprintf("Removing all existing sslprofile_ecccurve_binding from %s", name))
		existing, err := r.getExistingEcccurveBindings(name)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read existing sslprofile_ecccurve_binding, got error: %s", err))
			return
		}
		for _, curve := range existing {
			if err := r.deleteSingleEcccurveBinding(name, curve); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to remove existing ecccurve binding %s, got error: %s", curve, err))
				return
			}
		}
	}

	curves := ecccurveNamesFromList(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, curve := range curves {
		payload := ssl.Sslprofileecccurvebinding{
			Name:         name,
			Ecccurvename: curve,
		}
		// Binding resource - SDK v2 used AddResource (POST). Keep the same verb.
		if _, err := r.client.AddResource(service.Sslprofile_ecccurve_binding.Type(), name, &payload); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslprofile_ecccurve_binding, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "Created sslprofile_ecccurve_binding resource")

	// ID is the SSL profile name (the parent that groups all curve bindings).
	data.Id = types.StringValue(name)

	// Read the updated state back
	r.readSslprofileEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileEcccurveBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslprofileEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslprofile_ecccurve_binding resource")

	r.readSslprofileEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		// Binding was removed out-of-band; drop from state.
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileEcccurveBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All attributes are RequiresReplace, so Update is a documented no-op.
	var data, state SslprofileEcccurveBindingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslprofile_ecccurve_binding; all attributes are RequiresReplace")

	r.readSslprofileEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileEcccurveBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslprofileEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslprofile_ecccurve_binding resource")

	name := data.Name.ValueString()
	if name == "" {
		// Fall back to the ID (the profile name) for imported state.
		name = data.Id.ValueString()
	}

	curves := ecccurveNamesFromList(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, curve := range curves {
		if err := r.deleteSingleEcccurveBinding(name, curve); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslprofile_ecccurve_binding, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "Deleted sslprofile_ecccurve_binding binding")
}

// deleteSingleEcccurveBinding removes one ecccurve binding from the named SSL profile.
// The ecccurvename arg value is URL-encoded to handle any special characters.
func (r *SslprofileEcccurveBindingResource) deleteSingleEcccurveBinding(name, ecccurvename string) error {
	args := []string{fmt.Sprintf("ecccurvename:%s", url.QueryEscape(ecccurvename))}
	return r.client.DeleteResourceWithArgs(service.Sslprofile_ecccurve_binding.Type(), name, args)
}

// getExistingEcccurveBindings returns the ecccurve names currently bound to the profile.
func (r *SslprofileEcccurveBindingResource) getExistingEcccurveBindings(name string) ([]string, error) {
	findParams := service.FindParams{
		ResourceType:             service.Sslprofile_ecccurve_binding.Type(),
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}
	var curves []string
	for _, v := range dataArr {
		if val, ok := v["ecccurvename"].(string); ok {
			curves = append(curves, val)
		}
	}
	return curves, nil
}

// readSslprofileEcccurveBindingFromApi reads the live ecccurve bindings for the profile
// and stores the curve-name list back into the model. If no bindings remain, the ID is
// cleared so callers can drop the resource from state.
func (r *SslprofileEcccurveBindingResource) readSslprofileEcccurveBindingFromApi(ctx context.Context, data *SslprofileEcccurveBindingResourceModel, diags *diag.Diagnostics) {
	name := data.Name.ValueString()
	if name == "" {
		name = data.Id.ValueString()
	}

	curves, err := r.getExistingEcccurveBindings(name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_ecccurve_binding, got error: %s", err))
		return
	}

	if len(curves) == 0 {
		// No bindings left on the profile.
		data.Id = types.StringNull()
		return
	}

	listVal, d := types.ListValueFrom(ctx, types.StringType, curves)
	diags.Append(d...)
	if diags.HasError() {
		return
	}
	data.Ecccurvename = listVal
	data.Name = types.StringValue(name)
	data.Id = types.StringValue(name)
}
