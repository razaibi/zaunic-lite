package core

import (
	"fmt"
	"log"
	"os"
	"strings"
	"zaunic-lite/core/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func isEnvVar(envVal string) bool {
	_, present := os.LookupEnv(envVal)
	return present
}

func getCloudSecrets(secrets []models.Secret) map[string]interface{} {
	secretsMap := make(map[string]interface{})
	for _, s := range secrets {
		switch strings.ToUpper(s.Source) {
		case "AWS":
			secretsMap[s.Name] = getAWSSecret(
				s.Name,
				strings.ToUpper(s.Env),
				s.Region,
			)
		}
	}
	return secretsMap
}

func getAWSSecret(
	secretName string,
	env string,
	region string,
) string {

	accessKeyName := fmt.Sprintf(
		"AWS_%s_ACCESS_KEY_ID",
		env,
	)

	secretKeyName := fmt.Sprintf(
		"AWS_%s_SECRET_ACCESS_KEY",
		env,
	)

	var accessKeyId string
	var secretAccessKey string

	if isEnvVar(accessKeyName) && isEnvVar(secretKeyName) {
		accessKeyId = os.Getenv(accessKeyName)
		secretAccessKey = os.Getenv(secretKeyName)
	} else {
		return ""
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKeyId,
			secretAccessKey,
			"",
		),
	}))

	svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))

	result, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &secretName})
	if err != nil {
		log.Fatal(err.Error())
	}

	return *result.SecretString

}
