package code

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message"
)

//Client Function
func Client() {
	log.Println("Connecting to server...")

	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected")

	defer c.Logout()

	if err := c.Login("trianom4@gmail.com", "akuberuntung"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(mbox.Messages)

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{"\\Seen"}

	uids, err := c.Search(criteria)
	if err != nil {
		log.Println(err)
	}

	seqset := new(imap.SeqSet)
	seqset.AddNum(uids...)

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchRFC822, imap.FetchEnvelope, imap.FetchFlags, imap.FetchInternalDate}, messages)
	}()

	for msg := range messages {
		for _, r := range msg.Body {
			entity, err := message.Read(r)
			if err != nil {
				log.Fatal(err)
			}

			multiPartReader := entity.MultipartReader()

			for e, err := multiPartReader.NextPart(); err != io.EOF; e, err = multiPartReader.NextPart() {
				kind, params, cErr := e.Header.ContentType()
				if cErr != nil {
					log.Fatal(cErr)
				}

				if kind != "application/pdf" {
					continue
				}

				c, rErr := ioutil.ReadAll(e.Body)
				if rErr != nil {
					log.Fatal(rErr)
				}

				log.Printf("Dump file %s", params["name"])

				if fErr := ioutil.WriteFile(params["name"], c, 0777); fErr != nil {
					log.Fatal(fErr)
				}
				var filenams = params["name"]

				UploadFile("UploadDrive", filenams)
				// log.Print(filenams)
			}
		}
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done")

}
