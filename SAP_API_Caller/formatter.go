package sap_api_caller

type BatchMasterRecord struct {
	 ConnectionKey string `json:"connection_key"`
	 Result        bool   `json:"result"`
	 RedisKey      string `json:"redis_key"`
	 Filepath      string `json:"filepath"`
	 APISchema     string `json:"api_schema"`
	 MaterialCode  string `json:"material_code"`
	 Plant         string `json:"plant"`
	 Batch         string `json:"batch_code"`
	 Deleted       string `json:"deleted"`
}
	
type Batch struct {
     Material                 string `json:"Material"`
     BatchIdentifyingPlant    string `json:"BatchIdentifyingPlant"`
     Batch                    string `json:"Batch"`
     Supplier                 string `json:"Supplier"`
     BatchBySupplier          string `json:"BatchBySupplier"`
     CountryOfOrigin          string `json:"CountryOfOrigin"`
     RegionOfOrigin           string `json:"RegionOfOrigin"`
     MatlBatchAvailabilityDate string `json:"MatlBatchAvailabilityDate"`
     ShelfLifeExpirationDate  string `json:"ShelfLifeExpirationDate"`
     ManufactureDate          string `json:"ManufactureDate"`
	 CreationDateTime         string `json:"CreationDateTime"`
     LastChangeDateTime       string `json:"LastChangeDateTime"`
	 BatchIsMarkedForDeletion string `json:"BatchIsMarkedForDeletion"`
}
