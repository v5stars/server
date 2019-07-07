package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	repo = flag.String("repo", "", "Repo for protobuf generation, such as v2ray.com/core")
)

func main() {
	flag.Parse()
	wd, _ := os.Getwd()

	protofiles := make(map[string][]string)
	protoc := "protoc"
	gosrc := filepath.Join(wd, "../grpc")
	reporoot := filepath.Join(wd, *repo)

	filepath.Walk(reporoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		dir := filepath.Dir(path)
		filename := filepath.Base(path)
		if strings.HasSuffix(filename, ".proto") {
			protofiles[dir] = append(protofiles[dir], path)
		}

		return nil
	})

	for _, files := range protofiles {
		args := []string{"--proto_path", reporoot, "--go_out", "plugins=grpc:" + gosrc}
		args = append(args, files...)
		cmd := exec.Command(protoc, args...)
		cmd.Env = append(cmd.Env, os.Environ()...)
		output, err := cmd.CombinedOutput()
		if len(output) > 0 {
			fmt.Println(string(output))
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	must(filepath.Walk(reporoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".pb.go") {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		content = bytes.Replace(content, []byte("\"golang.org/x/net/context\""), []byte("\"context\""), 1)

		pos := bytes.Index(content, []byte("\npackage"))
		if pos > 0 {
			content = content[pos+1:]
		}

		return ioutil.WriteFile(path, content, info.Mode())
	}))
}
