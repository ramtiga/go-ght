package ght

import (
        "github.com/PuerkitoBio/goquery"
)

type RepoInf struct {
        Name            string
        Description     string
        Url             string
}

var baseUrl string = "https://github.com/trending"
var repoInf []RepoInf


func GetRepoInf(lang string) []RepoInf {
        url := getGithubUrl(lang)
        var repoInf []RepoInf
        repoInf = make([]RepoInf, 25)

        doc, _ := goquery.NewDocument(url)
        doc.Find(".leaderboard-list-content").Each(func(i int, s *goquery.Selection) {
                repoInf[i].Name = s.Find("a[class='repository-name']").Text()
                repoInf[i].Description = s.Find("p[class='repo-leaderboard-description']").Text()
                repoInf[i].Url = s.Find("a[class='repository-name']").Text()

        })
        return repoInf
}

func getGithubUrl(lang string) string {
        if lang == "" {
                return baseUrl
        } else {
                return baseUrl + "?l=" + lang
        }
}
