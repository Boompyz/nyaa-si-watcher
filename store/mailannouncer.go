package store

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// MailAnnouncer announces when something new has been added
// by sending emails to the specified addresses.
type MailAnnouncer struct {
	TargetMails []string `json:"addresses"`
}

// NewMailAnnouncer creates a new MailAnnouncer with no emails
func NewMailAnnouncer() *MailAnnouncer {
	return &MailAnnouncer{make([]string, 0)}
}

// Announce sends an announcement about the given options.
func (m *MailAnnouncer) Announce(options []TorrentOption) {
	if len(options) == 0 {
		return
	}

	names := make([]string, 0, len(options))
	for _, option := range options {
		names = append(names, option.Title)
	}

	message := `Hello,
I have noticed that new files have appeared that are on our wanted list.
They are: 

` + strings.Join(names, "\n") + `

Regards,
nyaa-si-watcher
`

	for _, addr := range m.TargetMails {
		sendMessage(addr, "Found new!", message)
	}
	fmt.Println("Finished announcing.")
}

func sendMessage(target, topic, body string) {
	command := exec.Command("mail", "-s", topic, target)
	writerCloser, err := command.StdinPipe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	fmt.Println("Starting process.")
	command.Start()
	_, err = writerCloser.Write([]byte(body))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = writerCloser.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Closed writer.")
	command.Wait()
	fmt.Println("Closed process.")
}
