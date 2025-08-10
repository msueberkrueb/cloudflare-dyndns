package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/dns"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

type RecordDetails struct {
	recordId string
	content  string
}

func GetRecordDetails(client cloudflare.Client, zoneId string, name string) (RecordDetails, error) {
	var recordDetails RecordDetails

	res, err := client.DNS.Records.List(
		context.TODO(),
		dns.RecordListParams{
			ZoneID: cloudflare.F(zoneId),
			Name: cloudflare.F(dns.RecordListParamsName{
				Exact: cloudflare.F(name),
			}),
		},
	)

	if err != nil {
		log.Fatal(err)
		return recordDetails, errors.New("CloudFlareError: Could not query dns records")
	}

	if len(res.Result) == 0 {
		return recordDetails, fmt.Errorf("CloudFlareError: Could not find record with name \"%s\" in zone \"%s\"", name, zoneId)
	}

	recordDetails.recordId = res.Result[0].ID
	recordDetails.content = res.Result[0].Content

	return recordDetails, nil
}

func UpdateRecords(config Config, content string) error {
	for name, recordConfig := range config.Records {
		var body dns.ARecordParam

		var apiToken string
		if recordConfig.ApiToken != "" {
			apiToken = recordConfig.ApiToken
		} else if config.ApiToken != "" {
			apiToken = config.ApiToken
		} else {
			return errors.New("ConfigError: neither default nor record api_token has been set")
		}

		var zoneId string
		if recordConfig.ZoneID != "" {
			zoneId = recordConfig.ZoneID
		} else if config.ZoneID != "" {
			zoneId = config.ZoneID
		} else {
			return errors.New("ConfigError: neither default nor record zone_id has been set")
		}

		client := cloudflare.NewClient(
			option.WithAPIToken(apiToken),
		)

		body.Name = cloudflare.F(name)
		body.Content = cloudflare.F(content)

		body.Proxied = cloudflare.F(recordConfig.Proxied)
		if recordConfig.TTL != 0 {
			body.TTL = cloudflare.F(recordConfig.TTL)
		} else if config.TTL != 0 {
			body.TTL = cloudflare.F(config.TTL)
		}
		if recordConfig.Comment != "" {
			body.Comment = cloudflare.F(recordConfig.Comment)
		} else if config.Comment != "" {
			body.Comment = cloudflare.F(config.Comment)
		}

		recordDetails, err := GetRecordDetails(*client, zoneId, name)

		if err != nil {
			return err
		}

		if recordDetails.content == content {
			log.Printf("\"%s\" was already set to \"%s\"", name, content)
			continue
		}

		_, err = client.DNS.Records.Edit(
			context.TODO(),
			recordDetails.recordId,
			dns.RecordEditParams{
				ZoneID: cloudflare.F(zoneId),
				Body:   body,
			},
		)

		if err != nil {
			log.Fatal(err)
			return fmt.Errorf("CloudFlareError: Could not update record \"%s\"", name)
		}

		log.Printf("Updated \"%s\" to \"%s\"", name, content)
	}

	return nil
}
