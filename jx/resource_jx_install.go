package jx

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/terraform/helper/schema"
)

//https://github.com/hashicorp/terraform/blob/master/builtin/provisioners/local-exec/resource_provisioner.go

func resourceInstall() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstallCreate,
		Read:   resourceInstallRead,
		Update: resourceInstallUpdate,
		Delete: resourceInstallDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"admin_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"jx_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"git_provider_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"git_owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"git_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"git_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceInstallCreate(d *schema.ResourceData, m interface{}) error {

	// adminPassword := "1q2w3e4a"
	// jxProvider := "kubernetes"
	// gitProviderUrl := "https://github.com"
	// gitOwner := "opstricks"
	// gitUser := "opstricks"
	// gitToken := "fcfa8bb880640e4e347ffed18ae7d88cb3de07db"

	adminPassword := d.Get("admin_password").(string)
	jxProvider := d.Get("jx_provider").(string)
	gitProviderUrl := d.Get("git_provider_url").(string)
	gitOwner := d.Get("git_owner").(string)
	gitUser := d.Get("git_user").(string)
	gitToken := d.Get("git_token").(string)

	arg := fmt.Sprintf("jx install --provider %s --batch-mode true --default-admin-password %s --no-default-environments true --recreate-existing-draft-repos true --verbose true  --environment-git-owner %s --git-username %s  --git-provider-url %s --git-api-token %s", jxProvider, adminPassword, gitOwner, gitUser, gitProviderUrl, gitToken)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", arg)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	return errStdout

}

func resourceInstallRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInstallUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInstallDelete(d *schema.ResourceData, m interface{}) error {

	// del1 := fmt.Sprintf("helm delete jenkins-x --purge")
	// del2 := fmt.Sprintf("helm delete jxing --purge")
	// del3 := fmt.Sprintf("helm reset -force")

	// var stdoutBuf, stderrBuf bytes.Buffer
	// cmd1 := exec.Command("/bin/sh", "-c", del1)
	// cmd2 := exec.Command("/bin/sh", "-c", del2)
	// cmd3 := exec.Command("/bin/sh", "-c", del3)

	// var errStdout, errStderr error
	// stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	// stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	// err1 := cmd1.Start()
	// if err1 != nil {
	// 	log.Fatalf("cmd.Start() failed with '%s'\n", err1)
	// }
	// err2 := cmd2.Start()
	// if err2 != nil {
	// 	log.Fatalf("cmd.Start() failed with '%s'\n", err2)
	// }
	// err3 := cmd3.Start()
	// if err3 != nil {
	// 	log.Fatalf("cmd.Start() failed with '%s'\n", err3)
	// }

	return nil
}
