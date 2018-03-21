package framework

import (
	"time"

	"github.com/appscode/go/crypto/rand"
	api "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	"github.com/appscode/searchlight/pkg/icinga"
	"github.com/appscode/searchlight/test/e2e/matcher"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Invocation) ClusterAlert() *api.ClusterAlert {
	return &api.ClusterAlert{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rand.WithUniqSuffix("clusteralert"),
			Namespace: f.namespace,
			Labels: map[string]string{
				"app": f.app,
			},
		},
		Spec: api.ClusterAlertSpec{
			CheckInterval: metav1.Duration{time.Second * 5},
			Vars:          make(map[string]string),
		},
	}
}

func (f *Framework) CreateClusterAlert(obj *api.ClusterAlert) error {
	_, err := f.extClient.MonitoringV1alpha1().ClusterAlerts(obj.Namespace).Create(obj)
	return err
}

func (f *Framework) GetClusterAlert(meta metav1.ObjectMeta) (*api.ClusterAlert, error) {
	return f.extClient.MonitoringV1alpha1().ClusterAlerts(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
}

func (f *Framework) DeleteClusterAlert(meta metav1.ObjectMeta) error {
	return f.extClient.MonitoringV1alpha1().ClusterAlerts(meta.Namespace).Delete(meta.Name, &metav1.DeleteOptions{})
}

func (f *Framework) getClusterAlertObjects(meta metav1.ObjectMeta, clusterAlertSpec api.ClusterAlertSpec) ([]icinga.IcingaHost, error) {
	objectList := []icinga.IcingaHost{
		{
			Type:           icinga.TypeCluster,
			AlertNamespace: meta.Namespace,
		},
	}
	return objectList, nil
}

func (f *Framework) EventuallyClusterAlertIcingaService(meta metav1.ObjectMeta, nodeAlertSpec api.ClusterAlertSpec) GomegaAsyncAssertion {
	objectList, err := f.getClusterAlertObjects(meta, nodeAlertSpec)
	Expect(err).NotTo(HaveOccurred())

	in := icinga.NewClusterHost(f.icingaClient, "6").
		IcingaServiceSearchQuery(meta.Name, objectList...)

	return Eventually(
		func() matcher.IcingaServiceState {
			var respService icinga.ResponseObject
			status, err := f.icingaClient.Objects().Service("").Get([]string{}, in).Do().Into(&respService)
			if status == 0 {
				return matcher.IcingaServiceState{Unknown: 1.0}
			}
			Expect(err).NotTo(HaveOccurred())

			var icingaServiceState matcher.IcingaServiceState
			for _, service := range respService.Results {
				if service.Attrs.LastState == 0.0 {
					icingaServiceState.OK++
				}
				if service.Attrs.LastState == 1.0 {
					icingaServiceState.Warning++
				}
				if service.Attrs.LastState == 2.0 {
					icingaServiceState.Critical++
				}
				if service.Attrs.LastState == 3.0 {
					icingaServiceState.Unknown++
				}
			}
			return icingaServiceState
		},
		time.Minute*5,
		time.Second*5,
	)
}

func (f *Framework) CleanClusterAlert() {
	caList, err := f.extClient.MonitoringV1alpha1().ClusterAlerts(f.namespace).List(metav1.ListOptions{})
	if err != nil {
		return
	}
	for _, e := range caList.Items {
		f.extClient.MonitoringV1alpha1().ClusterAlerts(f.namespace).Delete(e.Name, &metav1.DeleteOptions{})
	}
}
