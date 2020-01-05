package model

// Project represents a grouping of child Projects and Repos
type Project struct {
	Name     string
	Repos    []Repo
	Projects []Project
}
