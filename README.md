# cicd terraform provider

Project Overview
This project introduces a custom Terraform provider aimed at simplifying and automating CI/CD workflows. The provider empowers developers to manage application builds, run tests, create and push Docker images, and deploy Kubernetes manifests using Terraform configuration files. With built-in Git push event triggers, it streamlines the entire development pipeline, from source code changes to automated deployments.

Key Features
Application Build and Test: Execute build and test commands to validate applications before deployment.
Docker Image Creation: Build Docker images from application directories and ensure they are ready for deployment.
Container Registry Integration: Push Docker images to user-defined container registries with support for authentication.
Git Push Trigger: Automatically initiate the build, test, and deployment process upon each Git push event.
Kubernetes Manifest Redeployment: Automatically redeploy Kubernetes manifests using the latest Docker images for seamless and consistent application updates.
Concept
The provider brings the power of Terraform to CI/CD pipelines, helping teams automate complex software delivery workflows without extensive manual intervention. By leveraging Git push triggers and integrating Docker and Kubernetes operations, it delivers a complete end-to-end solution for modern DevOps practices.

This project simplifies cloud-native software delivery, reducing errors and increasing efficiency through automation and infrastructure as code principles.