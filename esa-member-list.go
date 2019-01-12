package main

import (
    "fmt"
    "os"
    "net/http"
    "encoding/json"
)

type PageResp struct {
	PrevPage   *uint `json:"prev_page"`
	NextPage   *uint `json:"next_page"`
	TotalCount uint  `json:"total_count"`
	Page       uint  `json:"page"`
	PerPage    uint  `json:"per_page"`
	MaxPerPage uint  `json:"max_per_page"`
}

type MembersResp struct {
	PageResp
	Members []Member `json:"members"`
}

type Member struct {
	Name       string `json:"name,omitempty"`
	ScreenName string `json:"screen_name,omitempty"`
	Icon       string `json:"icon,omitempty"`
	Email      string `json:"email,omitempty"`
	PostsCount int64  `json:"posts_count,omitempty"`
}

func main() {
  var url = "https://api.esa.io/v1/teams/" + os.Getenv("ESA_TEAM_NAME") + "/members?access_token=" + os.Getenv("ESA_ACCESS_TOKEN") + "&per_page=100"

  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    fmt.Println(err)
  }

  client := &http.Client{}
  resp, err := client.Do(req)


  if err != nil {
    fmt.Println(err)
  }

  defer resp.Body.Close()

  var d MembersResp

  // レスポンスのJSONをパースする
  decoder := json.NewDecoder(resp.Body)
  decoder.Decode(&d)

  members := d.Members

  for _, member := range members {
    fmt.Println(member.ScreenName)
  }
}
