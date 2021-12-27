package gittools

import (
	_ "embed"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

//go:embed id_rsa
var sshKey []byte

//
var auth *ssh.PublicKeys

func publicKey() (*ssh.PublicKeys, error) {
	publicKey, err := ssh.NewPublicKeys("git", sshKey, "")
	if err != nil {
		return nil, err
	}
	return publicKey, err
}

// Pull 抓取资源
func Pull() error {
	var args struct {
		Url  string
		Name string
	}

	if args.Name == "" {
		args.Name = strings.TrimSuffix(path.Base(args.Url), path.Ext(args.Url))
	}
	dir := path.Join("./public", args.Name)
	//
	os.RemoveAll(dir)
	//
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:  args.Url,
		Auth: auth,
		// ReferenceName: plumbing.ReferenceName("release"),
		SingleBranch: true,
		Progress:     os.Stdout,
	})
	return err
}
