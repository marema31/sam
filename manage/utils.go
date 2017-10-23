package manage

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func getSamPath() (string, error) {
	sampath, err := homedir.Expand(viper.GetString("SAMPATH"))
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(sampath)
	if os.IsNotExist(err) || !fi.IsDir() {
		fmt.Fprintf(os.Stderr, "SAM folder doesn't exists\n\n Please use: sam init\n")
		os.Exit(1)
	}
	return sampath, nil
}

func CreateSamPath() error {
	sampath, err := homedir.Expand(viper.GetString("SAMPATH"))
	if err != nil {
		return err
	}
	fi, err := os.Stat(sampath)
	if !os.IsNotExist(err) && fi.IsDir() {
		fmt.Fprintf(os.Stderr, "SAM folder already exists\n")
		os.Exit(1)
	}
	if !os.IsNotExist(err) && !fi.IsDir() {
		fmt.Fprintf(os.Stderr, "SAM Path ( %s  ) is not a directory \n\n Please modify your SAMPATH configuration\n", viper.GetString("SAMPATH"))
		os.Exit(1)
	}

	err = os.Mkdir(sampath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create SAM Path ( %s  ): %e \n", viper.GetString("SAMPATH"), err)
		os.Exit(1)
	}

	return nil
}
