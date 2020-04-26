package dgit

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
)

// SSHKeyGet ...
func SSHKeyGet() *ssh.PublicKeys {
	var publicKey *ssh.PublicKeys
	sshPath := os.Getenv("HOME") + "/.ssh/id_rsa"
	sshKey, _ := ioutil.ReadFile(sshPath)
	publicKey, keyError := ssh.NewPublicKeys("git", []byte(sshKey), "")
	if keyError != nil {
		fmt.Println(keyError)
	}
	return publicKey
}

// BranchClone ...
// * https://github.com/go-git/go-git/blob/2637279338ca9736c22b5e15fca916121ae8d089/options.go
func BranchClone(cloneOptions map[string]string) error {
	branch := plumbing.ReferenceName("refs/heads/" + cloneOptions["Branch"])
	directory := cloneOptions["Directory"]
	// key := SSHKeyGet()
	err := os.RemoveAll(directory)
	if err != nil {
		return err
	}
	fmt.Println("branch:", branch)
	fmt.Println("URL:", cloneOptions["URL"])
	_, err = git.PlainClone(directory, false, &git.CloneOptions{
		URL:           cloneOptions["URL"],
		RemoteName:    "origin",
		ReferenceName: branch,
		Depth: 1,
		SingleBranch: true,
		// Auth: key,
	})
	if err != nil {
		return err
	}
	return nil
}

// TagsGet Retrieve remote tags without cloning repository
func TagsGet() ([]string, error) {

	var tags []string

	// Create the remote with repository URL
	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{"https://github.com/Zenika/MARCEL"},
	})

	// We can then use every Remote functions to retrieve wanted information
	refs, err := rem.List(&git.ListOptions{})
	if err != nil {
		return tags, err
	}

	// Filters the references list and only keeps tags

	for _, ref := range refs {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		}
	}

	if len(tags) == 0 {
		return tags, errors.New("no tag found")
	}

	return tags, nil
}

// BranchesGet ...
func BranchesGet(url string) ([]string, error) {

	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})

	var branches []string

	// We can then use every Remote functions to retrieve wanted information
	refs, err := rem.List(&git.ListOptions{})
	if err != nil {
		return branches, err
	}

	// Filters the references list and only keeps branches

	for _, ref := range refs {
		if ref.Name().IsBranch() {
			branches = append(branches, ref.Name().Short())
		}
	}

	if len(branches) == 0 {
		return branches, errors.New("no tag found")
	}

	return branches, nil
}
