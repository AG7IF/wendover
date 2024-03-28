package aws

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/ag7if/wendover/internal/config"
)

func setupLogging(cfg aws.Config) (func() error, error) {
	logGroupName := viper.GetString(config.AWSLogGroupName)
	logStreamName := viper.GetString(config.AWSLogStreamName)

	writer := NewCloudWatchWriter(cfg, logGroupName, logStreamName, 1)
	log.Logger = log.Output(writer)

	return writer.Close, nil
}

func fetchFromParameterStore(cfg aws.Config) (map[string]string, error) {
	client := ssm.NewFromConfig(cfg)

	path := "/wendover"
	recursive := true
	input := &ssm.GetParametersByPathInput{
		Path:      &path,
		Recursive: &recursive,
	}

	output, err := client.GetParametersByPath(context.TODO(), input)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	params := make(map[string]string)
	for _, v := range output.Parameters {
		name := *v.Name
		name = strings.Replace(name, "/wendover/", "", 1)
		name = strings.Replace(name, "/", ".", -1)

		params[name] = *v.Value
	}

	return params, nil
}

func fetchDBCredsFromSecretsManager(cfg aws.Config, secretName string) (map[string]string, error) {
	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	output, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	str := *output.SecretString
	creds := make(map[string]string)
	err = json.Unmarshal([]byte(str), &creds)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return creds, nil
}

func SetupConfigFromParameterStore() (func() error, error) {
	cfg, err := awscfg.LoadDefaultConfig(context.TODO(), awscfg.WithRegion(viper.GetString(config.AWSRegion)))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	closer, err := setupLogging(cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	params, err := fetchFromParameterStore(cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	creds, err := fetchDBCredsFromSecretsManager(cfg, params["database.credentials"])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	params[config.DatabaseUser] = creds["username"]
	params[config.DatabasePassword] = creds["password"]

	for k, v := range params {
		viper.Set(k, v)
	}

	return closer, err
}
