package responses

type Batch struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Material                  string `json:"Material"`
			BatchIdentifyingPlant     string `json:"BatchIdentifyingPlant"`
			Batch                     string `json:"Batch"`
			Supplier                  string `json:"Supplier"`
			BatchBySupplier           string `json:"BatchBySupplier"`
			CountryOfOrigin           string `json:"CountryOfOrigin"`
			RegionOfOrigin            string `json:"RegionOfOrigin"`
			MatlBatchAvailabilityDate string `json:"MatlBatchAvailabilityDate"`
			ShelfLifeExpirationDate   string `json:"ShelfLifeExpirationDate"`
			ManufactureDate           string `json:"ManufactureDate"`
			CreationDateTime          string `json:"CreationDateTime"`
			LastChangeDateTime        string `json:"LastChangeDateTime"`
			BatchIsMarkedForDeletion  bool   `json:"BatchIsMarkedForDeletion"`
		} `json:"results"`
	} `json:"d"`
}
