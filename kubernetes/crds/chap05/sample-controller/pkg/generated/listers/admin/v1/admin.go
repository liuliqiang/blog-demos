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
	v1 "github.com/liuliqiang/admin-controller/pkg/apis/admin/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AdminLister helps list Admins.
type AdminLister interface {
	// List lists all Admins in the indexer.
	List(selector labels.Selector) (ret []*v1.Admin, err error)
	// Admins returns an object that can list and get Admins.
	Admins(namespace string) AdminNamespaceLister
	AdminListerExpansion
}

// adminLister implements the AdminLister interface.
type adminLister struct {
	indexer cache.Indexer
}

// NewAdminLister returns a new AdminLister.
func NewAdminLister(indexer cache.Indexer) AdminLister {
	return &adminLister{indexer: indexer}
}

// List lists all Admins in the indexer.
func (s *adminLister) List(selector labels.Selector) (ret []*v1.Admin, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Admin))
	})
	return ret, err
}

// Admins returns an object that can list and get Admins.
func (s *adminLister) Admins(namespace string) AdminNamespaceLister {
	return adminNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AdminNamespaceLister helps list and get Admins.
type AdminNamespaceLister interface {
	// List lists all Admins in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Admin, err error)
	// Get retrieves the Admin from the indexer for a given namespace and name.
	Get(name string) (*v1.Admin, error)
	AdminNamespaceListerExpansion
}

// adminNamespaceLister implements the AdminNamespaceLister
// interface.
type adminNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Admins in the indexer for a given namespace.
func (s adminNamespaceLister) List(selector labels.Selector) (ret []*v1.Admin, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Admin))
	})
	return ret, err
}

// Get retrieves the Admin from the indexer for a given namespace and name.
func (s adminNamespaceLister) Get(name string) (*v1.Admin, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("admin"), name)
	}
	return obj.(*v1.Admin), nil
}
