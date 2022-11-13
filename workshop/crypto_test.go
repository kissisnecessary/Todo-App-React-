package workshop_test

import (
	"github.com/Hyperledger-TWGC/fabric-gm-plugins/workshop"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var source, priv, pub workshop.SM2
var SM4Key  workshop.SM4
var priFile = "priv.pem"
var pubFile = "pub.pem"
var msg = []byte("2021-07-03 13:44:10")
var err error
var _ = Describe("Crypto", func() {
	// generate keypair file via tj or ccs
	It("it should able to generate keypair and store in file", func() {
		source, err = workshop.GenerateSM2Instance(workshop.TJ)
		Expect(err).NotTo(HaveOccurred())
		err = source.SaveFile(priFile, pubFile)
		Expect(err).NotTo(HaveOccurred())
		SM4Key, err = workshop.GenerateSM4Instance(workshop.TJ)
	})
	// general function for degist hash, sign
	It("it should able to complete degist via sm3 then sign, verify", func() {
		priv, err = workshop.LoadFromPriPem(priFile)
		Expect(err).NotTo(HaveOccurred())
		pub, e