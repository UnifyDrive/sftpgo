// Copyright (C) 2019-2023 Nicola Murino
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package sftpd implements the SSH File Transfer Protocol as described in https://tools.ietf.org/html/draft-ietf-secsh-filexfer-02.
// It uses pkg/sftp library:
// https://github.com/pkg/sftp
package sftpd

import (
	"strings"
	"time"
)

const (
	logSender        = "sftpd"
	handshakeTimeout = 2 * time.Minute
)

var (
	supportedSSHCommands = []string{"scp", "md5sum", "sha1sum", "sha256sum", "sha384sum", "sha512sum", "cd", "pwd",
		"git-receive-pack", "git-upload-pack", "git-upload-archive", "rsync", "sftpgo-copy", "sftpgo-remove"}
	defaultSSHCommands = []string{"md5sum", "sha1sum", "sha256sum", "cd", "pwd", "scp"}
	sshHashCommands    = []string{"md5sum", "sha1sum", "sha256sum", "sha384sum", "sha512sum"}
	systemCommands     = []string{"git-receive-pack", "git-upload-pack", "git-upload-archive", "rsync"}
	serviceStatus      ServiceStatus
)

type sshSubsystemExitStatus struct {
	Status uint32
}

type sshSubsystemExecMsg struct {
	Command string
}

// HostKey defines the details for a used host key
type HostKey struct {
	Path        string `json:"path"`
	Fingerprint string `json:"fingerprint"`
}

// ServiceStatus defines the service status
type ServiceStatus struct {
	IsActive        bool      `json:"is_active"`
	Bindings        []Binding `json:"bindings"`
	SSHCommands     []string  `json:"ssh_commands"`
	HostKeys        []HostKey `json:"host_keys"`
	Authentications []string  `json:"authentications"`
	HostKeyAlgos    []string  `json:"host_key_algos"`
	MACs            []string  `json:"macs"`
	KexAlgorithms   []string  `json:"kex_algorithms"`
	Ciphers         []string  `json:"ciphers"`
}

// GetSSHCommandsAsString returns enabled SSH commands as comma separated string
func (s *ServiceStatus) GetSSHCommandsAsString() string {
	return strings.Join(s.SSHCommands, ", ")
}

// GetSupportedAuthsAsString returns the supported authentications as comma separated string
func (s *ServiceStatus) GetSupportedAuthsAsString() string {
	return strings.Join(s.Authentications, ", ")
}

// GetHostKeyAlgosAsString returns the enabled host keys algorithms as comma separated string
func (s *ServiceStatus) GetHostKeyAlgosAsString() string {
	return strings.Join(s.HostKeyAlgos, ", ")
}

// GetMACsAsString returns the enabled MAC algorithms as comma separated string
func (s *ServiceStatus) GetMACsAsString() string {
	return strings.Join(s.MACs, ", ")
}

// GetKEXsAsString returns the enabled KEX algorithms as comma separated string
func (s *ServiceStatus) GetKEXsAsString() string {
	return strings.Join(s.KexAlgorithms, ", ")
}

// GetCiphersAsString returns the enabled ciphers as comma separated string
func (s *ServiceStatus) GetCiphersAsString() string {
	return strings.Join(s.Ciphers, ", ")
}

// GetStatus returns the server status
func GetStatus() ServiceStatus {
	return serviceStatus
}

// GetDefaultSSHCommands returns the SSH commands enabled as default
func GetDefaultSSHCommands() []string {
	result := make([]string, len(defaultSSHCommands))
	copy(result, defaultSSHCommands)
	return result
}

// GetSupportedSSHCommands returns the supported SSH commands
func GetSupportedSSHCommands() []string {
	result := make([]string, len(supportedSSHCommands))
	copy(result, supportedSSHCommands)
	return result
}
