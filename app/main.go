package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Pull Request Leaderboard</title>
</head>
<body>
	<h1>Pull Request Leaderboard</h1>
	<div>
		<table>
			<thead>
				<tr>
					<th>Username</th>
					<th>Opened</th>
					<th>Closed</th>
					<th>Comments</th>
				</tr>
			</thead>
			<tbody>
			{{range .Items}}
				<tr>
					<td>{{.Username}}</td>
					<td>{{.PullRequestsOpened}}</td>
					<td>{{.PullRequestsClosed}}</td>
					<td>{{.PullRequestComments}}</td>
				</tr>
			{{else}}
				<tr><td>no data</td></tr>
			{{end}}
			</tbody>
		</table>
	</div>
</body>
</html>
`

// PullRequestData that is returned from github
type PullRequestData struct {
	Username            string
	PullRequestsOpened  int
	PullRequestsClosed  int
	PullRequestComments int
}

const rootAddr = "https://api.github.com"

var client = &http.Client{
	Timeout: time.Second * 5,
}

var port = flag.Int("port", 8080, "the port the server will listen on")
var githubKey = flag.String("github-key", "", "the key that should be used to query the github APIs")
var repos = flag.String("github-repos", "", "the repos that should be interrogated")

func main() {
	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", *port),
		Handler:      r,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Println("server listening on port", *port)
	log.Fatal(server.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index").Parse(tpl)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	items := getPullRequestData(*githubKey, strings.Split(*repos, ","))

	data := struct {
		Items []PullRequestData
	}{
		Items: items,
	}

	err = t.Execute(w, data)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
}

func getPullRequestData(key string, repos []string) []PullRequestData {
	if key == "" {
		log.Println("WARN: no key")
		return []PullRequestData{}
	}

	data := make([]PullRequestData, len(repos))

	for i, r := range repos {
		x := getPullrequestDataForRepo(key, r)

		if x != nil {
			data[i] = *x
		}
	}

	return data
}

func getPullrequestDataForRepo(key, repo string) *PullRequestData {
	url := fmt.Sprintf("%v/repos/%v/pulls", rootAddr, repo)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
		return nil
	}

	req.Header.Set("Authorization", "token "+key)

	resp, err := client.Get(url)

	if err != nil {
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()

	item := &PullRequestData{}
	bytes, err := ioutil.ReadAll(resp.Body)

	log.Println(string(bytes))

	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(bytes, item)

	if err != nil {
		log.Println(err)
		return nil
	}

	return item
}
