package crawlers

import (
	"centralize-jobs/models"
	httpClient "centralize-jobs/utils"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetJobsGithubBackendBr() []models.JobsGithubBackendBr {
	response, err := httpClient.Get("https://github.com/backend-br/vagas/issues")
	if err != nil {
		log.Fatal(err)
	}

	html := string(response)
	jobs := []models.JobsGithubBackendBr{}
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	document.Find(".js-issue-row").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".Link--primary").Text()
		link, _ := s.Find(".Link--primary").Attr("href")
		link = "https: //github.com" + link
		workWaySplited := strings.Split(title, " ")
		workWay := workWaySplited[0]
		technologies := make([]string, 0)

		s.Find(".labels").Each(func(i int, s *goquery.Selection) {
			technologies = append(technologies, s.Find("a").Text())
		})

		jobs = append(jobs, models.JobsGithubBackendBr{
			Title:        title,
			Link:         link,
			WorkWay:      workWay,
			Technologies: technologies,
		})
	})

	return jobs
}
