// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

package inventory_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/opiproject/godpu/mocks"
	pb "github.com/opiproject/opi-api/common/v1/gen/go"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInventory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Inventory Suite")
}

func testCloser() {}

func getterWithResponse(_ grpc.ClientConnInterface) pb.InventorySvcClient {
	mockPb := mocks.NewInventorySvcClient(GinkgoT())
	mockPb.EXPECT().GetInventory(mock.Anything, mock.Anything).
		Return(getTestInvResponse(), nil).Once()
	return mockPb
}

func getterWithError(_ grpc.ClientConnInterface) pb.InventorySvcClient {
	mockPb := mocks.NewInventorySvcClient(GinkgoT())
	mockPb.EXPECT().GetInventory(mock.Anything, mock.Anything).
		Return(nil, errors.New("error getting inventory")).Once()
	return mockPb
}

func getTestInvResponse() *pb.Inventory {
	var resp pb.Inventory
	_ = json.Unmarshal(getTestInvResponseBytes(), &resp)
	return &resp
}

func getTestInvResponseBytes() []byte {
	return []byte(`
		{"bios": {"vendor": "Phoenix Technologies LTD","version": "6.00","date": "11/12/2020"},"system": {"name": "VMware Virtual Platform","vendor": "VMware, Inc.",
"serial_number": "VMware-56 4d 39 01 b5 ae 15 df-71 f0 87 e8 df 81 ee c0","uuid": "01394d56-aeb5-df15-71f0-87e8df81eec0","version": "None"},
"baseboard": {"serial_number": "None","vendor": "Intel Corporation","version": "None","product": "440BX Desktop Reference Platform"},
"chassis": {"asset_tag": "No Asset Tag","serial_number": "None","type": "1","type_description": "Other","vendor": "No Enclosure","version": "N/A"},
"processor": {"total_cores": 4,"total_threads": 4},"memory": {"total_physical_bytes": 10603200512,"total_usable_bytes": 10244431872},
"pci": {"driver": "agpgart-intel","address": "0000:00:00.0","vendor": "Intel Corporation","product": "440BX/ZX/DX - 82443BX/ZX/DX Host bridge",
"revision": "0x01","subsystem": "Virtual Machine Chipset","class": "Bridge","subclass": "Host bridge"},"pci": {"address": "0000:00:01.0",
"vendor": "Intel Corporation","product": "440BX/ZX/DX - 82443BX/ZX/DX AGP bridge","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"address": "0000:00:07.0","vendor": "Intel Corporation","product": "82371AB/EB/MB PIIX4 ISA","revision": "0x08",
"subsystem": "Virtual Machine Chipset","class": "Bridge","subclass": "ISA bridge"},"pci": {"driver": "ata_piix","address": "0000:00:07.1",
"vendor": "Intel Corporation","product": "82371AB/EB/MB PIIX4 IDE","revision": "0x01","subsystem": "Virtual Machine Chipset","class": "Mass storage controller",
"subclass": "IDE interface"},"pci": {"address": "0000:00:07.3","vendor": "Intel Corporation","product": "82371AB/EB/MB PIIX4 ACPI","revision": "0x08",
"subsystem": "Virtual Machine Chipset","class": "Bridge","subclass": "Bridge"},"pci": {"driver": "vmw_vmci","address": "0000:00:07.7","vendor": "VMware",
"product": "Virtual Machine Communication Interface","revision": "0x10","subsystem": "unknown","class": "Generic system peripheral",
"subclass": "System peripheral"},"pci": {"driver": "vmwgfx","address": "0000:00:0f.0","vendor": "VMware","product": "SVGA II Adapter","revision": "0x00",
"subsystem": "unknown","class": "Display controller","subclass": "VGA compatible controller"},"pci": {"driver": "mptspi","address": "0000:00:10.0",
"vendor": "Broadcom / LSI","product": "53c1030 PCI-X Fusion-MPT Dual Ultra320 SCSI","revision": "0x01","subsystem": "LSI Logic Parallel SCSI Controller",
"class": "Mass storage controller","subclass": "SCSI storage controller"},"pci": {"address": "0000:00:11.0","vendor": "VMware","product": "PCI bridge",
"revision": "0x02","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:15.0","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:15.1","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:15.2","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:15.3","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:15.4","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:15.5","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:15.6","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:15.7","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:16.0","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:16.1","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:16.2","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:16.3","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:16.4","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:16.5","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:16.6","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:16.7","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:17.0","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:17.1","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:17.2","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:17.3","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:17.4","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:17.5","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:17.6","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:17.7","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:18.0","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:18.1","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:18.2","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:18.3","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:18.4","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:18.5","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01",
"subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport","address": "0000:00:18.6","vendor": "VMware",
"product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge","subclass": "PCI bridge"},"pci": {"driver": "pcieport",
"address": "0000:00:18.7","vendor": "VMware","product": "PCI Express Root Port","revision": "0x01","subsystem": "unknown","class": "Bridge",
"subclass": "PCI bridge"},"pci": {"driver": "uhci_hcd","address": "0000:02:00.0","vendor": "VMware","product": "USB1.1 UHCI Controller","revision": "0x00",
"subsystem": "unknown","class": "Serial bus controller","subclass": "USB controller"},"pci": {"driver": "e1000","address": "0000:02:01.0",
"vendor": "Intel Corporation","product": "82545EM Gigabit Ethernet Controller (Copper)","revision": "0x01","subsystem": "PRO/1000 MT Single Port Adapter",
"class": "Network controller","subclass": "Ethernet controller"},"pci": {"driver": "snd_ens1371","address": "0000:02:02.0","vendor": "Ensoniq",
"product": "ES1371/ES1373 / Creative Labs CT2518","revision": "0x02","subsystem": "Audio PCI 64V/128/5200 / Creative CT4810/CT5803/CT5806 [Sound Blaster PCI]",
"class": "Multimedia controller","subclass": "Multimedia audio controller"},"pci": {"driver": "ehci-pci","address": "0000:02:03.0","vendor": "VMware",
"product": "USB2 EHCI Controller","revision": "0x00","subsystem": "unknown","class": "Serial bus controller","subclass": "USB controller"},
"pci": {"driver": "ahci","address": "0000:02:04.0","vendor": "VMware","product": "SATA AHCI controller","revision": "0x00","subsystem": "unknown",
"class": "Mass storage controller","subclass": "SATA controller"}}
	`)
}
