package iso

import (
	"context"

	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
	"github.com/jetbrains-infra/packer-builder-vsphere/driver"
)

type RemoveCDRomConfig struct {
	RemoveCdrom bool `mapstructure:"remove_cdrom"`
}
type StepRemoveCDRom struct {
	Config *RemoveCDRomConfig
}

func (s *StepRemoveCDRom) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	vm := state.Get("vm").(*driver.VirtualMachine)

	ui.Say("Eject CD-ROM drives...")
	err := vm.EjectCdroms()
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	if s.Config.RemoveCdrom == true {
		ui.Say("Deleting CD-ROM drives...")
		err := vm.RemoveCdroms()
		if err != nil {
			state.Put("error", err)
			return multistep.ActionHalt
		}
	}
	return multistep.ActionContinue
}

func (s *StepRemoveCDRom) Cleanup(state multistep.StateBag) {}
