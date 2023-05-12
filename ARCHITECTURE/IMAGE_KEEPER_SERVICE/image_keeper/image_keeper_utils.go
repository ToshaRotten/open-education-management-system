package image_keeper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getServiceName(path string) string {
	splited := strings.Split(path, "/")
	return strings.ToLower(splited[len(splited)-2])
}

func searchDockerFilePaths(root string) []string {
	var paths []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Name() == "dockerfile" {
			fmt.Println("Found Dockerfile at", path)
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error searching for Dockerfiles:", err)
	}
	return paths
}

func buildDockerImages(root string) {
	paths := searchDockerFilePaths(root)
	for _, path := range paths {
		imageName := getServiceName(path)
		buildDockerImage(path, imageName)
	}
}

func buildDockerImage(dockerfilePath string, imageName string) {
	cmd := exec.Command("docker", "build", "-t", imageName, "-f", dockerfilePath, ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error building Docker image:", err)
		return
	}

	cmd = exec.Command("docker", "save", imageName, ">", ".")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error building Docker image:", err)
		return
	}

	fmt.Println("Docker image built successfully!")
	fmt.Println(string(output))
}

func readZip(FilePath string, DestDir string) {
	r, err := zip.OpenReader(FilePath)
	if err != nil {
		fmt.Println("Error opening zip file:", err)
		return
	}
	defer r.Close()

	err = os.MkdirAll(DestDir, 0755)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			fmt.Println("Error opening file inside zip:", err)
		}
		defer rc.Close()

		path := DestDir + "/" + f.Name

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
			continue
		}

		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		defer file.Close()

		_, err = io.Copy(file, rc)
		if err != nil {
			fmt.Println("Error copying file contents:", err)
		}
	}
	fmt.Println("Zip file unpacked successfully!")
}
