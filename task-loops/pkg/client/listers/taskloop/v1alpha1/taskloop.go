/*
Copyright 2020 The Tekton Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/tektoncd/experimental/task-loops/pkg/apis/taskloop/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TaskLoopLister helps list TaskLoops.
type TaskLoopLister interface {
	// List lists all TaskLoops in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TaskLoop, err error)
	// TaskLoops returns an object that can list and get TaskLoops.
	TaskLoops(namespace string) TaskLoopNamespaceLister
	TaskLoopListerExpansion
}

// taskLoopLister implements the TaskLoopLister interface.
type taskLoopLister struct {
	indexer cache.Indexer
}

// NewTaskLoopLister returns a new TaskLoopLister.
func NewTaskLoopLister(indexer cache.Indexer) TaskLoopLister {
	return &taskLoopLister{indexer: indexer}
}

// List lists all TaskLoops in the indexer.
func (s *taskLoopLister) List(selector labels.Selector) (ret []*v1alpha1.TaskLoop, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TaskLoop))
	})
	return ret, err
}

// TaskLoops returns an object that can list and get TaskLoops.
func (s *taskLoopLister) TaskLoops(namespace string) TaskLoopNamespaceLister {
	return taskLoopNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TaskLoopNamespaceLister helps list and get TaskLoops.
type TaskLoopNamespaceLister interface {
	// List lists all TaskLoops in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TaskLoop, err error)
	// Get retrieves the TaskLoop from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TaskLoop, error)
	TaskLoopNamespaceListerExpansion
}

// taskLoopNamespaceLister implements the TaskLoopNamespaceLister
// interface.
type taskLoopNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TaskLoops in the indexer for a given namespace.
func (s taskLoopNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TaskLoop, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TaskLoop))
	})
	return ret, err
}

// Get retrieves the TaskLoop from the indexer for a given namespace and name.
func (s taskLoopNamespaceLister) Get(name string) (*v1alpha1.TaskLoop, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("taskloop"), name)
	}
	return obj.(*v1alpha1.TaskLoop), nil
}
