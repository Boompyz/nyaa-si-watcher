package announcer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Boompyz/nyaa-si-watcher/common"
	"github.com/Boompyz/nyaa-si-watcher/torrentoptions"
)

// MailAnnouncer announces when something new has been added
// by sending emails to the specified addresses.
type MailAnnouncer struct {
	targetMails []string
}

// NewMailAnnouncer creates a new MailAnnouncer with
// prepared bunch of emails (read from file).
func NewMailAnnouncer(confDir string) *MailAnnouncer {
	targetMails := common.GetLines(confDir + "/announcemails")
	return &MailAnnouncer{targetMails}
}

// Announce sends an announcement about the given options.
func (m *MailAnnouncer) Announce(options []torrentoptions.TorrentOption) {
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

	for _, addr := range m.targetMails {
		sendMessage(addr, "Found new!", message)
	}
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
