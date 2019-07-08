package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/json-iterator/go"
	"go.etcd.io/etcd/clientv3"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"time"
)

func main() {
	etcdHost := flag.String("etcdHost", "127.0.0.1:2379", "etcd host")
	etcdWatchKey := flag.String("etcdWatchKey", "/registry/pods", "etcd key to watch")

	flag.Parse()

	fmt.Println("connecting to etcd " + *etcdHost)

	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + *etcdHost},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to etcd " + *etcdHost)

	defer etcd.Close()

	watchChan := etcd.Watch(context.Background(), *etcdWatchKey, clientv3.WithPrefix())
	fmt.Println("set WATCH on " + *etcdWatchKey)

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {

			if *etcdWatchKey == "/registry/pods" && event.Type.String() == "PUT" {
				codec := scheme.Codecs.UniversalDecoder(scheme.Scheme.PreferredVersionAllGroups()...)
				obj, err := runtime.Decode(codec, event.Kv.Value)
				if err != nil {
					fmt.Printf("decode err is %s", err)
				}
				pod, ok := obj.(*v1.Pod)
				if ok {
					if pod.Annotations["zk.controller"] == "true" && pod.DeletionTimestamp == nil {
						for _, status := range pod.Status.ContainerStatuses {
							if status.Ready != true {
								for _, container := range pod.Spec.Containers {

									if container.Name == status.Name && status.State.Running.StartedAt != nil {

										if container.ReadinessProbe != nil && (int32(time.Since(status.State.Running.StartedAt.Time).Seconds()) >= int32(container.ReadinessProbe.InitialDelaySeconds+container.ReadinessProbe.TimeoutSeconds)) {
											// call zk
											json, err := jsoniter.Marshal(pod)
											fmt.Println(string(json))

										} else {
											// call zk
											json, err := jsoniter.Marshal(pod)
											fmt.Println(string(json))
										}

									}
								}

							}
						}
					}
				}
			} else {
				fmt.Printf("Event received! %s executed on %q create time %d mode time %d  with value %q\n", event.Type, event.Kv.Key, event.Kv.CreateRevision, event.Kv.ModRevision, event.Kv.Value)
			}

		}
	}
}
