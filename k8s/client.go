package k8s

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// "github.com/kubernetes/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	anno_k    = "cfgedit.io/enable"
	anno_v    = "true"
	clientSet *kubernetes.Clientset

	kubeconfig = flag.String("kubeconfig", "config", "(optional) absolute path to the kubeconfig file")
)

func initInCfg() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func initOutCfg() {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func lstCfgs(nss []string) []string {
	r := []string{}
	for _, ns := range nss {
		cfgs, err := clientSet.CoreV1().ConfigMaps(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Println("Get namespace ", ns, " failed")
			return nil
		}
		for _, cfg := range cfgs.Items {
			if cfg.Annotations[anno_k] != anno_v {
				continue
			}
			r = append(r, cfg.Namespace+"/"+cfg.Name)
		}
	}
	return r
}

func getCfg(ns, name string) map[string]string {
	cfg, err := clientSet.CoreV1().ConfigMaps(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Println("Get namespace ", ns, " failed")
		return nil
	}
	if cfg.Annotations[anno_k] != anno_v {
		return nil
	}
	return cfg.Data
}

func updateKvs(ns, name string, kv map[string]string) error {
	cfgClientSet := clientSet.CoreV1().ConfigMaps(ns)
	cfg, err := cfgClientSet.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Println("Get namespace ", ns, " failed")
		return nil
	}
	if cfg.Annotations[anno_k] != anno_v {
		return errors.New("Not allowed!")
	}
	for k, v := range kv {
		cfg.Data[k] = v
	}
	_, err = cfgClientSet.Update(context.TODO(), cfg, metav1.UpdateOptions{})
	fmt.Printf("Update return %p\n", err)
	return err
}
