package fileform

import(
	"bytes"
	"net/http"
	"mime/multipart"
)

type (
	FileInfo struct {
		FileName string
		Data     string
	}
	Form map[string]interface{}
)

func NewRequestWithForm(method string, route string, f Form) (*http.Request, error){
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)

	for key, value := range f{
		switch value := value.(type){
		case string:
			w.WriteField(key, value)
		case FileInfo:
			field, err := w.CreateFormFile(key, value.FileName)
			if err != nil{
				return nil, err
			}
			field.Write([]byte(value.Data))
		}
	}

	w.Close()
	req, err := http.NewRequest(method, route, b)
	if err != nil{
		return nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	return req, nil
}