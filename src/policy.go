package pike

import (
	"bytes"
	_ "embed" // required for embed
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/JamesWoolfenden/arn"
	"github.com/rs/zerolog/log"
)

//go:embed terraform.policy.template
var policyTemplate []byte

//go:embed aws_iam_role.tf
var roleTemplate []byte

// NewAWSPolicy constructor.
func NewAWSPolicy(actions []string, resources bool) (Policy, error) {
	something := Policy{Version: "2012-10-17"}

	sort.Strings(actions)

	categories := make([]string, len(actions))

	for index, action := range actions {
		categories[index] = strings.Split(action, ":")[0]
	}

	myArn := new(arn.AwsArn)

	if resources {
		myArn.Builder()
	}

	sections := Unique(categories)
	var statements []Statement

	for count, section := range sections {
		var myActions []string
		myResource := []string{"*"}

		resource := "*"

		for _, action := range actions {
			mySection := section + ":"
			if strings.Contains(action, mySection) {
				myActions = append(myActions, action)
			}
		}

		if myActions == nil {
			return something, fmt.Errorf("failed to find any action")
		}

		//todo expand with new plan function
		if resources {
			myArn.Service = section
			myArn.Resource = &resource
			myResource = myArn.Builder()
		}

		state := Statement{
			Sid: "VisualEditor" + strconv.Itoa(count), Effect: "Allow", Action: myActions, Resource: myResource,
		}

		statements = append(statements, state)
	}

	something.Statements = statements

	return something, nil
}

// GetPolicy creates new iam polices from a list of Permissions
func GetPolicy(actions Sorted, resources bool) (OutputPolicy, error) {
	var (
		OutPolicy OutputPolicy
		Empty     bool
	)

	Empty = true

	v := reflect.ValueOf(actions)
	typeOfV := v.Type()
	values := make([]interface{}, v.NumField())

	var err error

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()

		switch typeOfV.Field(i).Name {
		case "AWS":
			if actions.AWS == nil {
				continue
			}

			Empty = false
			// dedupe
			AWSPermissions := Unique(actions.AWS)
			OutPolicy.AWS, err = AWSPolicy(AWSPermissions, resources)

			if err != nil {
				log.Print(err)
				continue
			}

		case "GCP":
			if actions.GCP == nil {
				continue
			}

			Empty = false
			// dedupe
			GCPPermissions := Unique(actions.GCP)
			OutPolicy.GCP, err = GCPPolicy(GCPPermissions)

			if err != nil {
				log.Print(err)
				continue
			}

		case "AZURE":
			if actions.AZURE == nil {
				continue
			}

			Empty = false
			// dedupe
			AZUREPermissions := Unique(actions.AZURE)
			OutPolicy.AZURE, err = AZUREPolicy(AZUREPermissions)
			if err != nil {
				log.Print(err)
				continue
			}
		}

	}

	if Empty {
		return OutPolicy, errors.New("no permissions found")
	}

	return OutPolicy, nil
}

// AWSPolicy create an IAM policy
func AWSPolicy(Permissions []string, Resources bool) (AwsOutput, error) {
	var OutPolicy AwsOutput

	Policy, err := NewAWSPolicy(Permissions, Resources)
	if err != nil {
		return OutPolicy, err
	}

	b, err := json.MarshalIndent(Policy, "", "    ")
	if err != nil {
		fmt.Println(err)

		return OutPolicy, err
	}

	type policyDetails struct {
		Policy      string
		Name        string
		Path        string
		Description string
	}

	PolicyName := "terraform_pike"
	theDetails := policyDetails{string(b), PolicyName, "/", "Pike Autogenerated policy from IAC"}

	var output bytes.Buffer

	tmpl, err := template.New("test").Parse(string(policyTemplate))

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&output, theDetails)

	if err != nil {
		panic(err)
	}

	OutPolicy.Terraform = output.String()
	OutPolicy.JSONOut = string(b) + "\n"

	return OutPolicy, nil
}

// Unique make slice unique
func Unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string

	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}

	sort.Strings(result)

	return result
}
