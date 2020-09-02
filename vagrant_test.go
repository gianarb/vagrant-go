package vagrant_test

import (
	"context"
	"testing"

	"github.com/gianarb/vagrant-go"
)

func TestVagrantSetupGuide(t *testing.T) {
	ctx := context.Background()

	machine, err := vagrant.Up(ctx,
		vagrant.WithLogger(t.Logf),
		vagrant.WithMachineName("test"),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := machine.Destroy(ctx)
		if err != nil {
			t.Error(err)
		}
	}()
	_, err = machine.Exec(ctx, "touch", "/tmp/file.text")
	if err != nil {
		t.Error(err)
	}
}
