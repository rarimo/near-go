package action

import "time"

type NftMetadata struct {
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
