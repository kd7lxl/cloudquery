package compute

import (
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/gorilla/mux"
)

func createVirtualMachineExtensions(router *mux.Router) error {
	var item armcompute.VirtualMachineExtensionsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Compute/virtualMachines/test string/extensions", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}
