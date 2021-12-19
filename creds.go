package cred

import (
	"github.com/Versent/unicreds"
)

func GetCred(name string) (string, error) {
	tableName := "credential-store"
	region := "us-east-1"
	unicreds.SetAwsConfig(&region, nil, nil)
	d, err := unicreds.GetHighestVersionSecret(&tableName, name, unicreds.NewEncryptionContextValue())

	if err != nil {
		return "", err
	}
	return d.Secret, nil
}
