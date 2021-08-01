package crawlers

import (
	httpClient "centralize-jobs/utils"
	"log"
)

func GetJobsVulpi() string {
	response, err := httpClient.Get("https://api.vulpi.com.br/v1/board")
	if err != nil {
		log.Fatal(err)
	}
	return string(response)
}
