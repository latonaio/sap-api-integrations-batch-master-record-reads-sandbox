package main

import (
	sap_api_caller "sap-api-integrations-batch-master-record-reads/SAP_API_Caller"
	"sap-api-integrations-batch-master-record-reads/file_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := file_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Batch_Master_Record.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetBatchMasterRecord(
		inoutSDC.Material.BatchIdentifyingPlant.Batch,
	)
}
