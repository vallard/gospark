package models

import "time"

/*
{"object_kind":"push","before":"6e4b82f739347966d45c35997ba3039042e85e5a","after":"04edcc0632e4a74646dc563a9bf4a1b7cb61866e","ref":"refs/heads/master","checkout_sha":"04edcc0632e4a74646dc563a9bf4a1b7cb61866e","message":null,"user_id":1,"user_name":"Administrator","user_email":"admin@example.com","project_id":1,"repository":{"name":"demo1","url":"ssh://git@localhost:10022/root/demo1.git","description":"Demo one!","homepage":"http://localhost:10080/root/demo1","git_http_url":"http://localhost:10080/root/demo1.git","git_ssh_url":"ssh://git@localhost:10022/root/demo1.git","visibility_level":20},"commits":[{"id":"04edcc0632e4a74646dc563a9bf4a1b7cb61866e","message":"updates to color\n","timestamp":"2016-02-08T12:37:45-08:00","url":"http://localhost:10080/root/demo1/commit/04edcc0632e4a74646dc563a9bf4a1b7cb61866e","author":{"name":"vallard","email":"vallard@benincosa.com"},"added":[],"modified":["html/css/app.css"],"removed":[]},{"id":"325129c6686d97670455364be6cfd11fdf77db2a","message":"updates\n","timestamp":"2016-02-05T13:30:33-08:00","url":"http://localhost:10080/root/demo1/commit/325129c6686d97670455364be6cfd11fdf77db2a","author":{"name":"vallard","email":"vallard@benincosa.com"},"added":[],"modified":["Dockerfile"],"removed":[]},{"id":"6e4b82f739347966d45c35997ba3039042e85e5a","message":"Fixed Docker file\n","timestamp":"2016-02-05T13:21:55-08:00","url":"http://localhost:10080/root/demo1/commit/6e4b82f739347966d45c35997ba3039042e85e5a","author":{"name":"vallard","email":"vallard@benincosa.com"},"added":[],"modified":["Dockerfile"],"removed":[]}],"total_commits_count":3}

*/

type CommitMessage struct {
	Repo    Repository
	Commits []Commit
}

type Repository struct {
	Name        string
	Url         string
	Description string
	Homepage    string
}

type Commit struct {
	Id        string
	Message   string
	Timestamp time.Time
	Author    Author
}

type Author struct {
	Name  string
	Email string
}
