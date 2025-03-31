// Copyright (c) 2025 The bel2 developers

package config

type Config struct {
	Network string

	Signer   bool
	Listener bool

	Http                                 string
	ESCStartHeight                       uint64
	ESCTransactionManagerContractAddress string
	ESCArbiterManagerContractAddress     string
	ESCConfigManagerContractAddress      string
	ESCOrderManagerContractAddress       string
	ESCArbiterAddress                    string

	DataDir            string
	EscKeyFilePath     string
	ArbiterKeyFilePath string

	// loan signed path
	LoanSignedEventPath string
	// loan need sign path
	LoanNeedSignReqPath string
	// loan failed path
	LoanNeedSignFailedPath string
	// loan signed path
	LoanNeedSignSignedPath string
	// loan manually confirmed path
	LoanManuallyConfirmedPath string
	// loan logs path
	LoanLogPath string

	// bitcoin node rpc
	Proxy string
}
