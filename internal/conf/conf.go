package conf

import (
	"errors"
	"io/ioutil"

	logrus "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	AwsId       string `yaml:"awsId"`
	AwsSecret   string `yaml:"awsSecret"`
	AwsEndpoint string `yaml:"awsEndpoint"`
	AwsBucket   string `yaml:"awsBucket"`
}

var AppConfig = InitAppConfig()

func InitAppConfig() Config {
	var readConf Config

	readLocation := "config.yaml"
	buf, err := readYaml(readLocation)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("ErrorOnReadingConfigFile")
	}
	err = yaml.Unmarshal(buf, &readConf)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("ErrorOnReadingConfigFile")
	}
	return readConf
}

func readYaml(loc string) ([]byte, error) {
	buf, err := ioutil.ReadFile(loc)
	if err != nil {
		err := errors.New("ErrorOnReadingConfigFile")
		return nil, err
	}
	return buf, nil
}
