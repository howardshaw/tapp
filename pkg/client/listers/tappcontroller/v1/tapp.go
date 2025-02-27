/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "tkestack.io/tapp/pkg/apis/tappcontroller/v1"
)

// TAppLister helps list TApps.
type TAppLister interface {
	// List lists all TApps in the indexer.
	List(selector labels.Selector) (ret []*v1.TApp, err error)
	// TApps returns an object that can list and get TApps.
	TApps(namespace string) TAppNamespaceLister
	TAppListerExpansion
}

// tAppLister implements the TAppLister interface.
type tAppLister struct {
	indexer cache.Indexer
}

// NewTAppLister returns a new TAppLister.
func NewTAppLister(indexer cache.Indexer) TAppLister {
	return &tAppLister{indexer: indexer}
}

// List lists all TApps in the indexer.
func (s *tAppLister) List(selector labels.Selector) (ret []*v1.TApp, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TApp))
	})
	return ret, err
}

// TApps returns an object that can list and get TApps.
func (s *tAppLister) TApps(namespace string) TAppNamespaceLister {
	return tAppNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TAppNamespaceLister helps list and get TApps.
type TAppNamespaceLister interface {
	// List lists all TApps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TApp, err error)
	// Get retrieves the TApp from the indexer for a given namespace and name.
	Get(name string) (*v1.TApp, error)
	TAppNamespaceListerExpansion
}

// tAppNamespaceLister implements the TAppNamespaceLister
// interface.
type tAppNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TApps in the indexer for a given namespace.
func (s tAppNamespaceLister) List(selector labels.Selector) (ret []*v1.TApp, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TApp))
	})
	return ret, err
}

// Get retrieves the TApp from the indexer for a given namespace and name.
func (s tAppNamespaceLister) Get(name string) (*v1.TApp, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tapp"), name)
	}
	return obj.(*v1.TApp), nil
}
