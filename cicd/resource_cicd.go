package cicd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCICD() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCICDCreate,
		ReadContext:   resourceCICDRead,
		UpdateContext: resourceCICDUpdate,
		DeleteContext: resourceCICDDelete,
		CustomizeDiff: customizeDiff,

		Schema: map[string]*schema.Schema{
			"build": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"test": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"working_directory": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"build_and_test": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"docker_build": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"docker_push": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_registry": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"dockerfile_directory": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_registry_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func customizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	d.SetNewComputed("timestamp")
	return nil
}

func executeCommand(command string, workDir string) (error, string) {
	if workDir == "" {
		return fmt.Errorf("working diretory is not specified"), "."
	} else {
		err := os.Chdir(workDir)
		if err != nil {
			return fmt.Errorf("failed to change directory: %v", err), "."
		}
		output, err := exec.Command("sh", "-c", command).CombinedOutput()
		if err != nil {
			dependenciesErr := ""
			if strings.Contains(string(output), "command not found") {
				dependenciesErr = "* Make sure all necessary dependencies are installed *"
			}
			return fmt.Errorf("command failed with error: %v\nOutput: %s\n%s", err, string(output), dependenciesErr), "."
		}
		return nil, string(output)
	}

}

func resourceCICDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	working_directory := d.Get("working_directory").(string)
	build := d.Get("build").(string)
	test := d.Get("test").(string)
	build_and_test := d.Get("build_and_test").(string)
	dockerfile_dir := d.Get("dockerfile_directory").(string)
	docker_build := d.Get("docker_build").(string)
	cr_url := d.Get("container_registry").(string)
	cr_pass := d.Get("container_registry_password").(string)
	// docker_push := d.Get("docker_push").(string)

	warning := func(processName, output string) diag.Diagnostics {
		return diag.Diagnostics{
			{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("Step %v completed!\n%v", processName, output),
			},
		}
	}

	steps := []struct {
		name string
		cmd  string
	}{
		{"build", build},
		{"test", test},
		{"build_and_test", build_and_test},
		{"docker_build", docker_build},
	}

	for _, step := range steps {
		if step.cmd != "" {
			var err error
			var output string

			if dockerfile_dir != "" && step.name == "docker_build" {
				err, output = executeCommand(step.cmd, dockerfile_dir)
			} else {
				err, output = executeCommand(step.cmd, working_directory)
			}

			if err != nil {
				return diag.FromErr(err)
			}
			diags = append(diags, warning(step.name, output)...)
		}
	}
	if cr_url != "" {
		var input string
		if strings.Contains(cr_url, "amazonaws.com") {
			check_url_format := regexp.MustCompile(`^\d{12}\.dkr\.ecr\.[a-z0-9-]+\.amazonaws\.com$`).MatchString(cr_url)
			if check_url_format {
				err := exec.Command("aws", "sts", "get-caller-identity").Run()
				if err == nil {
					regionRegex := regexp.MustCompile(`(?:[^\.]*\.){3}([^\.]*)`).FindStringSubmatch(cr_url)[1]
					input = fmt.Sprintf("aws ecr get-login-password --region %s | docker login --username AWS --password-stdin %s", regionRegex, cr_url)
				} else {
					return diag.FromErr(fmt.Errorf("unable to locate credentials. You can configure credentials by running 'aws configure' "))
				}
			} else {
				return diag.FromErr(fmt.Errorf("invalid ECR URL format. Expected format: <aws_account_id>.dkr.ecr.<region>.amazonaws.com"))
			}
		} else if strings.Contains(cr_url, "azurecr.io") {
			check_url_format := regexp.MustCompile(`^[a-zA-Z0-9]+\.azurecr\.io$`).MatchString(cr_url)
			if check_url_format {
				err := exec.Command("az", "account", "show").Run()
				if err == nil {
					cr_name := regexp.MustCompile(`[^.]+`).FindStringSubmatch(cr_url)[1]
					input = fmt.Sprintf("az acr login --name %s", cr_name)
				}
			}
		}
		err := exec.Command("sh", "-c", input).Run()
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Set the ID and timestamp for the resource
	d.SetId(fmt.Sprintf("%s-%s", strings.ToLower(build), strings.ToLower(test)))
	d.Set("timestamp", time.Now().Format(time.RFC3339))

	return diags

}
