package chroot

import (
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
)

// StepChrootProvision provisions the instance within a chroot.
type StepChrootProvision struct {
	mounts []string
}

func (s *StepChrootProvision) Run(state multistep.StateBag) multistep.StepAction {
	hook := state.Get("hook").(packer.Hook)
	mountPath := state.Get("mount_path").(string)
	chrootCommand := state.Get("chroot_command").(string)
	copyCommand := state.Get("copy_command").(string)
	ui := state.Get("ui").(packer.Ui)

	// Create our communicator
	comm := &Communicator{
		Chroot:        mountPath,
		ChrootCommand: chrootCommand,
		CopyCommand:   copyCommand,
	}

	// Provision
	log.Println("Running the provision hook")
	if err := hook.Run(packer.HookProvision, ui, comm, nil); err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepChrootProvision) Cleanup(state multistep.StateBag) {}
