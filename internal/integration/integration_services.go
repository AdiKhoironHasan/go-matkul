package integration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AdiKhoironHasan/matkul/pkg/dto"
	util "github.com/AdiKhoironHasan/matkul/pkg/utils"
)

type integService struct {
}

func NewService() IntegServices {
	return &integService{}
}

func (s *integService) GetDosenID(req *dto.MatkulReqDTO) error {
	var response *dto.GetDosenIDResDTO
	var url string

	url = util.GetIntegURL("mahasiswaserv", "getDosenByID")
	url = fmt.Sprintf(url, req.ID)

	getReq, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return fmt.Errorf("failed to create http request mahasiswaserv getDosenID: %v", err)
	}
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(getReq)
	if err != nil {
		return fmt.Errorf("failed to create http request mahasiswaserv getDosenID: %v", err)
	}
	log.Println("Success execute  : ", url)
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode mahasiswaserv getDosenID response: %v", err)
	}

	if response.Data == nil {
		return fmt.Errorf("failed to decode mahasiswaserv getDosenID response: data is NULL")
	}

	if response.Data[0].IdDosen <= 0 {
		return fmt.Errorf("failed to decode mahasiswaserv getDosenID response: data is NULL")
	}

	return nil
}
