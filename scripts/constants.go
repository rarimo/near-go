package scripts

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/rarimo/near-go/common"

	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/bundle"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
)

const (
	targetNetwork            = "Near"
	nftMetadataReference     = "https://bafkreiaociirm4drxdf7ldvkemoxu3fsnvt4z4urtu35ss3xysl3gmw25q.ipfs.nftstorage.link/"
	nftMetadataReferenceHash = "DhIRFnBxuMv1jqojHXpssm1nzPKRnTfZS3fEl7My2uw="
	nftMedia                 = "https://bafkreiblbldzupel5ci36xhaw2kpci4q53yvjnq55ueqawep6nigjggcze.ipfs.nftstorage.link/"
	nftMediaHash             = "KwrHmjyL6JG/XOC2lPEjkO7xVLYd7QkAWI/zUGSYwsk="
)

var MaxGas uint64 = 300000000000000

var nftMetadata = map[bool]*common.NftMetadataView{
	true: {
		Title:         "Wrapped Rarimo Bridge NFT#1",
		Description:   "Wrapped Rarimo Bridge Test Collection NFT#1",
		Media:         nftMedia,
		MediaHash:     mustDecodeBase64(nftMediaHash),
		Reference:     nftMetadataReference,
		ReferenceHash: mustDecodeBase64(nftMetadataReferenceHash),
		Copies:        1,
	},
	false: {
		Title:         "Rarimo Bridge NFT#1",
		Description:   "Rarimo Bridge Test Collection NFT#1",
		Media:         nftMedia,
		MediaHash:     mustDecodeBase64(nftMediaHash),
		Reference:     nftMetadataReference,
		ReferenceHash: mustDecodeBase64(nftMetadataReferenceHash),
		Copies:        1,
	},
}

var (
	content1 = operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash("txHash1").SetOpId("eventId1").SetCurrentNetwork("networkFrom1").Build().GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  targetNetwork,
		TargetContract: getRandSlice(),
		Data:           data.NewTransferDataBuilder().SetId("1").Build().GetContent(),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	content2 = operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash("txHash2").SetOpId("networkFrom2").SetCurrentNetwork("networkFrom2").Build().GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  targetNetwork,
		TargetContract: getRandSlice(),
		Data:           data.NewTransferDataBuilder().SetId("1").Build().GetContent(),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	content3 = operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash("txHash3").SetOpId("networkFrom3").SetCurrentNetwork("networkFrom3").Build().GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  targetNetwork,
		TargetContract: getRandSlice(),
		Data:           data.NewTransferDataBuilder().SetId("1").Build().GetContent(),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	content4 = operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash("txHash4").SetOpId("networkFrom4").SetCurrentNetwork("networkFrom4").Build().GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  targetNetwork,
		TargetContract: getRandSlice(),
		Data:           data.NewTransferDataBuilder().SetId("1").Build().GetContent(),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	content5 = operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash("txHash5").SetOpId("networkFrom5").SetCurrentNetwork("networkFrom5").Build().GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  targetNetwork,
		TargetContract: getRandSlice(),
		Data:           data.NewTransferDataBuilder().SetId("1").Build().GetContent(),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	content6 = operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash("txHash6").SetOpId("networkFrom6").SetCurrentNetwork("networkFrom6").Build().GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  targetNetwork,
		TargetContract: getRandSlice(),
		Data:           data.NewTransferDataBuilder().SetId("1").Build().GetContent(),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	content7 = operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash("txHash7").SetOpId("networkFrom7").SetCurrentNetwork("networkFrom7").Build().GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  targetNetwork,
		TargetContract: getRandSlice(),
		Data:           data.NewTransferDataBuilder().SetId("1").Build().GetContent(),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	content8 = operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash("txHash8").SetOpId("networkFrom8").SetCurrentNetwork("networkFrom8").Build().GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  targetNetwork,
		TargetContract: getRandSlice(),
		Data:           data.NewTransferDataBuilder().SetId("1").Build().GetContent(),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	content9 = operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash("txHash9").SetOpId("networkFrom9").SetCurrentNetwork("networkFrom9").Build().GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  targetNetwork,
		TargetContract: getRandSlice(),
		Data:           data.NewTransferDataBuilder().SetId("1").Build().GetContent(),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}
)

func getRandSlice() []byte {
	var hash [32]byte
	rand.Read(hash[:])
	return hash[:]
}

func mustDecodeBase64(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
