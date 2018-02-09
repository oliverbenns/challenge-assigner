package main

import (
  "log"
  "net/http"
  "github.com/google/go-github/github"
  "context"
  "encoding/json"
  "golang.org/x/oauth2"
  "os"
  "io"
  "io/ioutil"
  "time"
)

func CreateChallenge(w http.ResponseWriter, r *http.Request) {
  var assignee Assignee
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  // Read POST body
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

  if err != nil {
    w.WriteHeader(500)
  }

  if err := r.Body.Close(); err != nil {
    w.WriteHeader(500)
  }

  if err := json.Unmarshal(body, &assignee); err != nil {
    w.WriteHeader(422) // unprocessable entity
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
  }

  accessToken := os.Getenv("ACCESS_TOKEN")
  ghUser := os.Getenv("GH_USER")

  ctx := context.Background()
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: accessToken},
  )
  tc := oauth2.NewClient(ctx, ts)

  client := github.NewClient(tc)

  repoName := assignee.Username + "-" + time.Now().Format("20060102150405")

  repo := &github.Repository {
    Name: github.String(repoName),
    // Private: github.Bool(true),
  }

  // Create the repo.
  repo, _, cErr := client.Repositories.Create(ctx, "", repo)

  // Add readme file
  introduction, _ := ioutil.ReadFile("./challenges/ruby-senior/README.md")
  commitMessage := "Add Introduction"
  branch := "master"
  fileOptions := github.RepositoryContentFileOptions{ Message: &commitMessage, Content: introduction, Branch: &branch }
  client.Repositories.CreateFile(ctx, ghUser, repoName, "README.md", &fileOptions)

  // Add user as collaborator
  collabOptions := github.RepositoryAddCollaboratorOptions{ Permission: "push" }
  client.Repositories.AddCollaborator(ctx, ghUser, repoName, assignee.Username, &collabOptions)
  // log.Print("COLERR")
  // log.Print(colErr)

  if cErr != nil {
    w.WriteHeader(500)
    log.Print("ERR")
    log.Print(cErr)
    // @TODO: Give json error msg.
    json.NewEncoder(w).Encode(cErr)
  } else {
    w.WriteHeader(200)
    log.Print("ORGS")
    log.Print(repo)
    json.NewEncoder(w).Encode(repo)
  }
}

func StartChallenge(w http.ResponseWriter, r *http.Request) {
  var assignedChallenge AssignedChallenge
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  // Read POST body
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

  if err != nil {
    w.WriteHeader(500)
  }

  if err := r.Body.Close(); err != nil {
    w.WriteHeader(500)
  }

  if err := json.Unmarshal(body, &assignedChallenge); err != nil {
    w.WriteHeader(422) // unprocessable entity
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
  }

  accessToken := os.Getenv("ACCESS_TOKEN")
  ghUser := os.Getenv("GH_USER")

  ctx := context.Background()
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: accessToken},
  )
  tc := oauth2.NewClient(ctx, ts)

  client := github.NewClient(tc)

  // Add readme file
  introduction, _ := ioutil.ReadFile("./challenges/ruby-senior/BRIEF.md")
  commitMessage := "Add Brief"
  branch := "master"
  fileOptions := github.RepositoryContentFileOptions{ Message: &commitMessage, Content: introduction, Branch: &branch }
  commit, _, cErr  := client.Repositories.CreateFile(ctx, ghUser, assignedChallenge.RepoName, "BRIEF.md", &fileOptions)

  if cErr != nil {
    w.WriteHeader(500)
    log.Print("ERR")
    log.Print(cErr)
    // @TODO: Give json error msg.
    json.NewEncoder(w).Encode(cErr)
  } else {
    w.WriteHeader(200)
    log.Print("ORGS")
    log.Print(commit)
    json.NewEncoder(w).Encode(commit)
  }

}
