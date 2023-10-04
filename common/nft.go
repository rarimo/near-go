package common

import (
	"time"
)

type NftView struct {
	TokenID            string               `json:"token_id"`
	OwnerID            string               `json:"owner_id"`
	Metadata           *NftMetadataView     `json:"metadata"`
	ApprovedAccountIds *map[AccountID]int64 `json:"approved_account_ids"`
}

type NftMetadataView struct {
	// ex. "Arch Nemesis: Mail Carrier" or "Parcel #5055"
	Title string `json:"title,omitempty"`
	// free-form description
	Description string `json:"description,omitempty"`
	// URL to associated media, preferably to decentralized, content-addressed storage
	Media string `json:"media,omitempty"`
	// Base64-encoded sha256 hash of content referenced by the `media` field. Required if `media` is included.
	MediaHash []byte `json:"media_hash,omitempty"`
	// number of copies of this set of metadata in existence when token was minted.
	Copies uint64 `json:"copies,omitempty"`
	// ISO 8601 datetime when token was issued or minted
	IssuedAt time.Time `json:"issued_at,omitempty"`
	// ISO 8601 datetime when token expires
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// ISO 8601 datetime when token starts being valid
	StartsAt time.Time `json:"starts_at,omitempty"`
	// ISO 8601 datetime when token was last updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// anything extra the NFT wants to store on-chain. Can be stringified JSON.
	Extra string `json:"extra,omitempty"`
	// URL to an off-chain JSON file with more info.
	Reference string `json:"reference,omitempty"`
	// Base64-encoded sha256 hash of JSON from reference field. Required if `reference` is included.
	ReferenceHash []byte `json:"reference_hash,omitempty"`
}

type NftContractMetadataView struct {
	// required, essentially a version like "nft-1.0.0"
	Spec string `json:"spec"`
	// required, ex. "Mosaics"
	Name string `json:"name"`
	// required, ex. "MOSIAC"
	Symbol string `json:"symbol"`
	// Data URL
	Icon *string `json:"icon"`
	// Centralized gateway known to have reliable access to decentralized storage assets referenced by `reference` or `media` URLs
	BaseURI *string `json:"base_uri,omitempty"`
	// URL to an off-chain JSON file with more info.
	Reference *string `json:"reference,omitempty"`
	// Base64-encoded sha256 hash of JSON from reference field. Required if `reference` is included.
	ReferenceHash []byte `json:"reference_hash,omitempty"`
}
