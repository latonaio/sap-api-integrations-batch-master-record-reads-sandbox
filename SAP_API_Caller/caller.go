package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetBatchMasterRecord(Material, BatchIdentifyingPlant, Batch string) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		c.Batch(Material, BatchIdentifyingPlant, Batch)
		wg.Done()
	}()
	
	wg.Wait()
}

func (c *SAPAPICaller) Batch(Material, BatchIdentifyingPlant, Batch string) {
	res, err := c.callBatchSrvAPIRequirementBatch("Batch", Material, BatchIdentifyingPlant, Batch)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callBatchSrvAPIRequirement(api, Material, BatchIdentifyingPlant, Batch string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BATCH_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Material, BatchIdentifyingPlant, Batch")
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and BatchIdentifyingPlant eq '%s' and Batch eq '%s'", Material, BatchIdentifyingPlant, Batch))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}