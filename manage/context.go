package manage

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListContexts(project string) ([]string, error) {
	var contexts []string

	pe, err := IsProjectExists(project)
	if err != nil {
		return nil, err
	}
	if !pe {
		fmt.Fprintf(os.Stderr, "Project %s doesn't exists\n", project)
		os.Exit(1)
	}

	projectPath, err := projectPath(project)
	if err != nil {
		return nil, err
	}

	d, err := os.Open(projectPath)
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
			contexts = append(contexts, fi.Name())
		}
	}

	return contexts, nil
}

func contextPath(project string, context string) (string, error) {
	sampath, err := getSamPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(sampath, project, context), nil
}

// TODO : A TESTER
func IsContextExists(project string, context string) (bool, error) {
	contextPath, err := contextPath(project, context)
	if err != nil {
		return false, err
	}
	fi, err := os.Stat(contextPath)
	if os.IsNotExist(err) || !fi.IsDir() {
		return false, nil
	}
	return true, nil
}
