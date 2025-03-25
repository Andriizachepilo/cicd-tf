package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource = (*resourceCICD)(nil)
)

func cicdResource() resource.Resource {
	return &resourceCICD{}
}

type resourceCICD struct{}

func (n *resourceCICD) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (n *resourceCICD) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"build": schema.StringAttribute{
				MarkdownDescription: "Build command to be executed.",
				Optional:            true,
			},
			"test": schema.StringAttribute{
				MarkdownDescription: "Test command to be executed.",
				Optional:            true,
			},
			"working_directory": schema.StringAttribute{
				MarkdownDescription: "Working directory for the commands.",
				Optional:            true,
				Validators: []validator.String{

				},
			},
			"build_and_test": schema.StringAttribute{
				MarkdownDescription: "Build and test command to be executed.",
				Optional:            true,
			},
			"timestamp": schema.StringAttribute{
				MarkdownDescription: "Timestamp of the last update.",
				Computed:            true,
			},
		},
	}
}

func (n *resourceCICD) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {



}

func (r *resourceCICD) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state cicdModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *resourceCICD) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan cicdModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceCICD) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

type cicdModel struct {
	ID                   types.String `tfsdk:"id"`
	Build                types.String `tfsdk:"build"`
	Test                 types.String `tfsdk:"test"`
	WorkingDirectory     types.String `tfsdk:"working_directory"`
	BuildAndTest         types.String `tfsdk:"build_and_test"`
	DockerBuild          types.String `tfsdk:"docker_build"`
	DockerPush           types.String `tfsdk:"docker_push"`
	ContainerRegistryURL types.String `tfsdk:"container_registry_url"`
	DockerfileDirectory  types.String `tfsdk:"dockerfile_directory"`
	Timestamp            types.String `tfsdk:"timestamp"`
}
