package analyzer

import (
	"testing"
)

func TestParseImageNameAndSource(t *testing.T) {
	tests := []struct {
		name           string
		fullImageName  string
		expectedTag    string
		expectedSource string
	}{
		{
			name:           "DockerHub image with tag",
			fullImageName:  "nginx:latest",
			expectedTag:    "nginx:latest",
			expectedSource: "",
		},
		{
			name:           "DockerHub image without tag",
			fullImageName:  "nginx",
			expectedTag:    "nginx",
			expectedSource: "",
		},
		{
			name:           "DockerHub image with docker.io prefix",
			fullImageName:  "docker.io/nginx:latest",
			expectedTag:    "nginx:latest",
			expectedSource: "",
		},
		{
			name:           "DockerHub official library image",
			fullImageName:  "library/nginx:latest",
			expectedTag:    "library/nginx:latest",
			expectedSource: "",
		},
		{
			name:           "DockerHub user image",
			fullImageName:  "user/myapp:v1.0",
			expectedTag:    "user/myapp:v1.0",
			expectedSource: "",
		},
		{
			name:           "GitLab registry image",
			fullImageName:  "registry.gitlab.com/yumbrands/phus/web/web2-frontend/web2-app-image:v2.10.3154",
			expectedTag:    "web2-app-image:v2.10.3154",
			expectedSource: "registry.gitlab.com/yumbrands/phus/web/web2-frontend",
		},
		{
			name:           "GitLab registry image without tag",
			fullImageName:  "registry.gitlab.com/user/project/image",
			expectedTag:    "image",
			expectedSource: "registry.gitlab.com/user/project",
		},
		{
			name:           "AWS ECR image",
			fullImageName:  "123456789012.dkr.ecr.us-east-1.amazonaws.com/my-app:latest",
			expectedTag:    "my-app:latest",
			expectedSource: "123456789012.dkr.ecr.us-east-1.amazonaws.com",
		},
		{
			name:           "Google Container Registry",
			fullImageName:  "gcr.io/project-id/image-name:tag",
			expectedTag:    "image-name:tag",
			expectedSource: "gcr.io/project-id",
		},
		{
			name:           "Harbor registry",
			fullImageName:  "harbor.company.com/project/app:v1.0.0",
			expectedTag:    "app:v1.0.0",
			expectedSource: "harbor.company.com/project",
		},
		{
			name:           "Complex path with multiple levels",
			fullImageName:  "registry.example.com/org/team/project/service/app:v1.2.3",
			expectedTag:    "app:v1.2.3",
			expectedSource: "registry.example.com/org/team/project/service",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTag, gotSource := parseImageNameAndSource(tt.fullImageName)
			if gotTag != tt.expectedTag {
				t.Errorf("parseImageNameAndSource() gotTag = %v, want %v", gotTag, tt.expectedTag)
			}
			if gotSource != tt.expectedSource {
				t.Errorf("parseImageNameAndSource() gotSource = %v, want %v", gotSource, tt.expectedSource)
			}
		})
	}
}
