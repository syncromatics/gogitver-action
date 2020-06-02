package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/src-d/go-git.v4/plumbing"

	gogitver "github.com/syncromatics/gogitver/pkg/git"

	"github.com/pkg/errors"
	"gopkg.in/src-d/go-billy.v4/memfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func main() {
	repo, ok := os.LookupEnv("GITHUB_REPOSITORY")
	if !ok {
		log.Fatal("did not find GITHUB_REPOSITORY env variable")
	}

	ref, ok := os.LookupEnv("GITHUB_REF")
	if !ok {
		log.Fatal("did not find GITHUB_REF env variable")
	}

	remote := fmt.Sprintf("https://github.com/%s", repo)

	fs := memfs.New()
	storage := memory.NewStorage()
	r, err := git.Clone(storage, fs, &git.CloneOptions{
		URL:           remote,
		ReferenceName: plumbing.ReferenceName(ref),
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed cloning git"))
	}

	version, err := gogitver.GetCurrentVersion(r, gogitver.GetDefaultSettings(), &gogitver.BranchSettings{}, false)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed getting current version"))
	}

	fmt.Printf("::set-output name=version::v%s\n", version)
}
