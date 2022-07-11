/*
Copyright 2022 The OpenFunction Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"

	clientset "github.com/openfunction/pkg/client/clientset/versioned"
	corev1alpha2 "github.com/openfunction/pkg/client/clientset/versioned/typed/core/v1alpha2"
	fakecorev1alpha2 "github.com/openfunction/pkg/client/clientset/versioned/typed/core/v1alpha2/fake"
	corev1beta1 "github.com/openfunction/pkg/client/clientset/versioned/typed/core/v1beta1"
	fakecorev1beta1 "github.com/openfunction/pkg/client/clientset/versioned/typed/core/v1beta1/fake"
	eventsv1alpha1 "github.com/openfunction/pkg/client/clientset/versioned/typed/events/v1alpha1"
	fakeeventsv1alpha1 "github.com/openfunction/pkg/client/clientset/versioned/typed/events/v1alpha1/fake"
	networkingv1alpha1 "github.com/openfunction/pkg/client/clientset/versioned/typed/networking/v1alpha1"
	fakenetworkingv1alpha1 "github.com/openfunction/pkg/client/clientset/versioned/typed/networking/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// CoreV1alpha2 retrieves the CoreV1alpha2Client
func (c *Clientset) CoreV1alpha2() corev1alpha2.CoreV1alpha2Interface {
	return &fakecorev1alpha2.FakeCoreV1alpha2{Fake: &c.Fake}
}

// CoreV1beta1 retrieves the CoreV1beta1Client
func (c *Clientset) CoreV1beta1() corev1beta1.CoreV1beta1Interface {
	return &fakecorev1beta1.FakeCoreV1beta1{Fake: &c.Fake}
}

// EventsV1alpha1 retrieves the EventsV1alpha1Client
func (c *Clientset) EventsV1alpha1() eventsv1alpha1.EventsV1alpha1Interface {
	return &fakeeventsv1alpha1.FakeEventsV1alpha1{Fake: &c.Fake}
}

// NetworkingV1alpha1 retrieves the NetworkingV1alpha1Client
func (c *Clientset) NetworkingV1alpha1() networkingv1alpha1.NetworkingV1alpha1Interface {
	return &fakenetworkingv1alpha1.FakeNetworkingV1alpha1{Fake: &c.Fake}
}
