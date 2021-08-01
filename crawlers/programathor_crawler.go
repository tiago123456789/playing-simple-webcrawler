package crawlers

import (
	"centralize-jobs/models"
	httpClient "centralize-jobs/utils"

	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetJobsProgramathor() []models.JobProgramathor {
	jobs := []models.JobProgramathor{}
	response, err := httpClient.Get("https://programathor.com.br/jobs")
	if err != nil {
		log.Fatal(err)
	}

	html := string(response)
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	document.Find(".cell-list ").Each(func(i int, s *goquery.Selection) {
		value, _ := s.Find("a").Attr("href")
		logo, _ := s.Find(".logo-list img").Attr("src")
		company := s.Find(".cell-list-content-icon span:nth-child(1)").Text()
		workWay := s.Find(".cell-list-content-icon span:nth-child(2)").Text()
		salary := s.Find(".cell-list-content-icon span:nth-child(3)").Text()
		experienceLevel := s.Find(".cell-list-content-icon span:nth-child(4)").Text()
		laborWay := s.Find(".cell-list-content-icon span:nth-child(5)").Text()
		job := models.JobProgramathor{
			Link:            "https://programathor.com.br" + value,
			Logo:            logo,
			Company:         company,
			WorkWay:         workWay,
			Salary:          salary,
			ExperienceLevel: experienceLevel,
			LaborWay:        laborWay,
		}
		jobs = append(jobs, job)
	})
	return jobs
}
