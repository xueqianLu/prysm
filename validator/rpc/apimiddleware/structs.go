package apimiddleware

import "github.com/prysmaticlabs/prysm/v3/api/gateway/apimiddleware"

type listKeystoresResponseJson struct {
	Keystores []*keystoreJson `json:"data"`
}

type keystoreJson struct {
	ValidatingPubkey apimiddleware.HexString `json:"validating_pubkey"`
	DerivationPath   string                  `json:"derivation_path"`
}

type importKeystoresRequestJson struct {
	Keystores          []string `json:"keystores"`
	Passwords          []string `json:"passwords"`
	SlashingProtection string   `json:"slashing_protection"`
}

type importKeystoresResponseJson struct {
	Statuses []*statusJson `json:"data"`
}

type deleteKeystoresRequestJson struct {
	PublicKeys []apimiddleware.HexString `json:"pubkeys"`
}

type statusJson struct {
	Status  string `json:"status" enum:"true"`
	Message string `json:"message"`
}

type deleteKeystoresResponseJson struct {
	Statuses           []*statusJson `json:"data"`
	SlashingProtection string        `json:"slashing_protection"`
}

//remote keymanager api

type listRemoteKeysResponseJson struct {
	Keystores []*remoteKeysListJson `json:"data"`
}

type remoteKeysListJson struct {
	Pubkey   apimiddleware.HexString `json:"pubkey"`
	Url      string                  `json:"url"`
	Readonly bool                    `json:"readonly"`
}

type remoteKeysJson struct {
	Pubkey   apimiddleware.HexString `json:"pubkey"`
	Url      string                  `json:"url"`
	Readonly bool                    `json:"readonly"`
}

type importRemoteKeysRequestJson struct {
	Keystores []*remoteKeysJson `json:"remote_keys"`
}

type importRemoteKeysResponseJson struct {
	Statuses []*statusJson `json:"data"`
}

type deleteRemoteKeysRequestJson struct {
	PublicKeys []apimiddleware.HexString `json:"pubkeys"`
}

type deleteRemoteKeysResponseJson struct {
	Statuses []*statusJson `json:"data"`
}

type feeRecipientJson struct {
	Pubkey     apimiddleware.HexString `json:"pubkey"`
	Ethaddress string                  `json:"ethaddress" address:"true"`
}

type gasLimitJson struct {
	Pubkey   apimiddleware.HexString `json:"pubkey"`
	GasLimit string                  `json:"gas_limit"`
}

type getFeeRecipientByPubkeyResponseJson struct {
	Data *feeRecipientJson `json:"data"`
}

type setFeeRecipientByPubkeyRequestJson struct {
	Ethaddress apimiddleware.HexString `json:"ethaddress"`
}

type getGasLimitResponseJson struct {
	Data *gasLimitJson `json:"data"`
}

type setGasLimitRequestJson struct {
	GasLimit string `json:"gas_limit"`
}

type deleteGasLimitRequestJson struct {
	Pubkey apimiddleware.HexString `json:"pubkey"`
}
