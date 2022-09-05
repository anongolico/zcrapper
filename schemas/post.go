package schemas

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/anongolico/ib/config"
)

type Post struct {
	Id          string       `json:"id,omitempty"`
	Creacion    time.Time    `json:"creacion,omitempty"`
	Hilo        Hilo         `json:"hilo,omitempty"`
	Comentarios []Comentario `json:"comentarios,omitempty"`
}

// ScanPost retrieves the post with the given id, reads its content
// and marshals the response into a new Post variable if no errors
// occurred during the process
func ScanPost(id string, buffer []byte) (*Post, error) {
	// creates the request
	req, err := http.NewRequest(http.MethodGet, config.MainUrl+id, nil)
	if err != nil {
		return nil, err
	}

	// adds the cookie so it can get over the 'muro'
	req.AddCookie(&config.IdentityCookie)

	var client http.Client

	// makes the request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// since some posts are quite large, the scanner has to have
	// a considerably large buffer. The size can be customized
	// from other functions
	if len(buffer) == 0 {
		buffer = make([]byte, 512*2048)
	}
	scanner := bufio.NewScanner(res.Body)
	scanner.Buffer(buffer, len(buffer))

	// byte array to store the response's body
	var body []byte

	for scanner.Scan() {
		line := scanner.Bytes()
		// usually, all clones have the same variable with the json content
		if bytes.Contains(line, []byte("window.data")) {
			_, body, _ = bytes.Cut(line, []byte("   window.data = "))
			break
		}
	}

	post := new(Post)

	err = json.Unmarshal(body, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
