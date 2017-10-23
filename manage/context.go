package manage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
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

	d, err := os.Open(filepath.Join(projectPath, "contexts"))
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
	return filepath.Join(sampath, project, "contexts", context), nil
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

func CreateContext(project string, context string) error {
	pe, err := IsProjectExists(project)
	if err != nil {
		return err
	}
	if !pe {
		fmt.Fprintf(os.Stderr, "Project %s doesn't exists\n", project)
		os.Exit(1)
	}

	contextPath, err := contextPath(project, context)
	if err != nil {
		return err
	}

	fi, err := os.Stat(contextPath)
	if !os.IsNotExist(err) && fi.IsDir() {
		fmt.Fprintf(os.Stderr, "Context folder already exists\n")
		os.Exit(1)
	}
	if !os.IsNotExist(err) && !fi.IsDir() {
		fmt.Fprintf(os.Stderr, "Context path ( %s  ) is not a directory \n", filepath.Join(viper.GetString("SAMPATH"), project, context))
		os.Exit(1)
	}

	sampath, _ := getSamPath()

	err = os.Mkdir(filepath.Join(sampath, project, "contexts", context), 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create Context Path ( %s  ): %e \n", filepath.Join(viper.GetString("SAMPATH"), project, context), err)
		os.Exit(1)
	}

	return nil
}
