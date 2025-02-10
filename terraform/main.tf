resource "cicd_example" "creating" {
  working_directory = "/Users/andriizachepilo/terraform-provider-cicd"
  build             = "npm run build"
  test              = "npm test"
  docker_build      = "docker build -t test:tag ."
}
