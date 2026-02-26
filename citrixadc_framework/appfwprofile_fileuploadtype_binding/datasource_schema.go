package appfwprofile_fileuploadtype_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AppfwprofileFileuploadtypeBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"alertonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send SNMP alert?",
			},
			"as_fileuploadtypes_url": schema.StringAttribute{
				Required:    true,
				Description: "FileUploadTypes action URL.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"filetype": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "FileUploadTypes file types.",
			},
			"fileuploadtype": schema.StringAttribute{
				Required:    true,
				Description: "FileUploadTypes to allow/deny.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isnameregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is field name a regular expression?",
			},
			"isregex_fileuploadtypes_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is a regular expression?",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A \"id\" that identifies the rule.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
		},
	}
}
