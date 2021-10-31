package configuration

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xeipuuv/gojsonschema"
)

type Config interface {
	GetString(key string) string
	Get(key string) interface{}
	GetInt(key string) int
	GetStringArr(key string) []string
	Init(filePath, schemaPath string) (Config, error)
	GetOsEnvString(key string) string
}

type viperConfig struct {
	log *logrus.Entry
}

func (v *viperConfig) Init(filePath, schemaPath string) (Config, error) {

	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		v.log.Errorf("failed to read configuration from: %v, err: %+v", filePath, err)
		return nil, err
	}

	v.log.Infof("successfully read configuration from: %v", filePath)
	// if err := v.validateConfiguration(filePath, schemaPath); err != nil {
	// 	v.log.Errorf("Invalid configuration: %v, err: %+v", filePath, err)
	// 	return nil, err
	// }

	return v, nil
}

func (v *viperConfig) validateConfiguration(configFile, schemaPath string) error {
	sl := gojsonschema.NewSchemaLoader()
	sl.Draft = gojsonschema.Draft7
	sl.AutoDetect = false

	viper.SetConfigFile(configFile)

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	documentLoader := gojsonschema.NewReferenceLoader("file://" + configFile)

	res, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		v.log.Errorf("validateConfiguration - %s", err.Error())
		return err
	}

	if !res.Valid() {
		v.log.Errorf("validateConfiguration - %s", err.Error())
		return errors.New("invalid schema")
	}

	return nil
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetStringArr(key string) []string {
	return viper.GetStringSlice(key)
}

func (v *viperConfig) Get(key string) interface{} {
	return viper.Get(key)
}

func (v *viperConfig) GetOsEnvString(key string) string {
	return os.Getenv(key)
}

func NewViperConfig(log *logrus.Entry) *viperConfig {
	v := &viperConfig{
		log: log,
	}
	return v
}
