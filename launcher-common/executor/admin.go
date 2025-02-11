package executor

import (
	"crypto/x509"
	"encoding/base64"
	"github.com/luskaner/ageLANServer/common"
	"github.com/luskaner/ageLANServer/launcher-common/executor/exec"
	"net"
)

func RunSetUp(gameId string, IPs []net.IP, certificate *x509.Certificate, CDN bool) (result *exec.Result) {
	args := make([]string, 0)
	args = append(args, "setup", "-e", gameId)
	if IPs != nil {
		for _, ip := range IPs {
			args = append(args, "-i")
			args = append(args, ip.String())
		}
	}
	if certificate != nil {
		args = append(args, "-l")
		args = append(args, base64.StdEncoding.EncodeToString(certificate.Raw))
	}
	if CDN {
		args = append(args, "-c")
	}
	result = exec.Options{File: common.GetExeFileName(true, common.LauncherConfigAdmin), AsAdmin: true, Wait: true, ExitCode: true, Args: args}.Exec()
	return
}

func RunRevert(gameId string, IPs bool, certificate bool, CDN bool, failfast bool) (result *exec.Result) {
	args := make([]string, 0)
	args = append(args, "revert", "-e", gameId)
	if failfast {
		if IPs {
			args = append(args, "-i")
		}
		if certificate {
			args = append(args, "-l")
		}
		if CDN {
			args = append(args, "-c")
		}
	} else {
		args = append(args, "-a")
	}
	result = exec.Options{File: common.GetExeFileName(true, common.LauncherConfigAdmin), AsAdmin: true, Wait: true, ExitCode: true, Args: args}.Exec()
	return
}
