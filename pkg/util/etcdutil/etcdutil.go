package etcdutil

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned"
)

func Migrate() error {
	time.Sleep(30 * time.Second)
	// TODO: poll if TPR ready?
	fmt.Println("etcd TPR is ready ===")

	kubecli, err := unversioned.New(&restclient.Config{
		Host: "http://127.0.0.1:8080", // TODO: parameter
	})
	if err != nil {
		return err
	}

	ip := "10.240.0.10" // TODO: how to find ip?

	b := []byte(fmt.Sprintf(`{
  "apiVersion": "coreos.com/v1",
  "kind": "EtcdCluster",
  "metadata": {
    "name": "etcd-cluster",
    "namespace": "kube-system"
  },
  "spec": {
    "size": 1,
    "version": "v3.1.0-alpha.1",
    "seed": {
      "MemberClientEndpoints": [
        "http://%s:2379"
      ],
      "RemoveDelay": 60
    }
  }
}`, ip))

	resp, err := kubecli.Client.Post(
		"http://127.0.0.1:8080/apis/coreos.com/v1/namespaces/kube-system/etcdclusters",
		"application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status: %v", resp.Status)
	}

	// TODO: how to know when it's cool?
	time.Sleep(600 * time.Second)

	return nil
}
