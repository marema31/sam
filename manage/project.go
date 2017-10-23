package manage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func ListProjects() ([]string, error) {
	var projects []string

	sampath, err := getSamPath()
	if err != nil {
		return nil, err
	}

	d, err := os.Open(sampath)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	fi, err := d.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, fi := range fi {
		if fi.IsDir() {
			projects = append(projects, fi.Name())
		}
	}

	return projects, nil
}

func projectPath(project string) (string, error) {
	sampath, err := getSamPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(sampath, project), nil
}

func IsProjectExists(project string) (bool, error) {
	projectPath, err := projectPath(project)
	if err != nil {
		return false, err
	}
	fi, err := os.Stat(projectPath)
	if os.IsNotExist(err) || !fi.IsDir() {
		return false, nil
	}
	return true, nil
}

func CreateProject(project string) error {
	projectPath, err := projectPath(project)
	if err != nil {
		return err
	}

	fi, err := os.Stat(projectPath)
	if !os.IsNotExist(err) && fi.IsDir() {
		fmt.Fprintf(os.Stderr, "Project folder already exists\n")
		os.Exit(1)
	}
	if !os.IsNotExist(err) && !fi.IsDir() {
		fmt.Fprintf(os.Stderr, "Project path ( %s  ) is not a directory \n", filepath.Join(viper.GetString("SAMPATH"), project))
		os.Exit(1)
	}

	sampath, _ := getSamPath()

	folders := [...]string{project, filepath.Join(project, "contexts"), filepath.Join(project, "options")}
	for _, folder := range folders {
		err = os.Mkdir(filepath.Join(sampath, folder), 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't create Project Path ( %s  ): %e \n", filepath.Join(viper.GetString("SAMPATH"), folder), err)
			os.Exit(1)
		}
	}

	return nil
}
