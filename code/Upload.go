package code

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

// ------------------------------------------------------------------------------------------------------
// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "config/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

//-----------------------------------------------------------------------------------------------------
// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getService() (*drive.Service, error) {
	b, err := ioutil.ReadFile("config/credentials.json")
	if err != nil {
		fmt.Printf("Unable to read credential.json file. Err: %v\n", err)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)

	if err != nil {
		return nil, err
	}

	client := getClient(config)

	service, err := drive.New(client)

	if err != nil {
		fmt.Printf("Cannot create the Google Drive service: %v\n", err)
		return nil, err
	}

	return service, err
}

//-----------------------------------------------------------------------------------------------------
//  create directory  in drive
func createDir(service *drive.Service, name string, parentID string) (*drive.File, error) {
	d := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentID},
	}

	file, err := service.Files.Create(d).Do()

	if err != nil {
		log.Println("Could not create dir: " + err.Error())
		return nil, err
	}

	return file, nil
}

//------------------------------------------------------------------------------------------------------
//  create file to upload to drive
func createFile(service *drive.Service, name string, mimeType string, content io.Reader, parentID string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentID},
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}

//UploadFile function
func UploadFile() {

	var files []string
	fileInfo, err := ioutil.ReadDir("./attachment")
	if err != nil {
		fmt.Println(files, err)
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}

	// ===================================
	// Step 2. Get the Google Drive service
	service, err := getService()

	// Step 3. Create the directory
	dir, err := createDir(service, "Attachment", "root")

	if err != nil {
		panic(fmt.Sprintf("Could not create dir: %v\n", err))
	}
	// ------------------------------------------------------
	// a := []string{"Foo", "Bar"}
	for i, s := range files {
		fmt.Println(i, s)

		// Step 1. Open the file
		f, err := os.Open("./attachment/" + s)

		if err != nil {
			panic(fmt.Sprintf("cannot open file: %v", err))
		}

		defer f.Close()

		// Step 2. Get the Google Drive service
		service, err := getService()

		// Step 3. Create the directory
		// dir, err := createDir(service, folderName, "root")

		// if err != nil {
		// 	panic(fmt.Sprintf("Could not create dir: %v\n", err))
		// }

		//  contentType, err := GetFileContentType(f)
		//  if err != nil {
		// 	 panic(err)
		//  }

		// Step 4. Create the file and upload its content
		file, err := createFile(service, s, "application/pdf", f, dir.Id)

		if err != nil {
			panic(fmt.Sprintf("Could not create file: %v\n", err))
		}

		fmt.Printf("File '%s' successfully uploaded in '%s' directory", file.Name, dir.Name)
	}
}
