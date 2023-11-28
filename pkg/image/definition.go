package image

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	TypeISO = "iso"
	TypeRAW = "raw"
)

type Definition struct {
	APIVersion      string          `yaml:"apiVersion"`
	Image           Image           `yaml:"image"`
	OperatingSystem OperatingSystem `yaml:"operatingSystem"`
}

type Image struct {
	ImageType       string `yaml:"imageType"`
	BaseImage       string `yaml:"baseImage"`
	OutputImageName string `yaml:"outputImageName"`
}

type OperatingSystem struct {
	KernelArgs []string              `yaml:"kernelArgs"`
	Users      []OperatingSystemUser `yaml:"users"`
}

type OperatingSystemUser struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSHKey   string `yaml:"sshKey"`
}

func ParseDefinition(data []byte) (*Definition, error) {
	var definition Definition

	if err := yaml.Unmarshal(data, &definition); err != nil {
		return nil, fmt.Errorf("could not parse the image definition: %w", err)
	}

	return &definition, nil
}