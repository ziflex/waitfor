package waitfor

import (
	"context"
	"errors"
	"net/url"
	"strings"
)

type (
	ResourceFactory func(u *url.URL) (Resource, error)

	ResourceConfig struct {
		Scheme  string
		Factory ResourceFactory
	}

	Resource interface {
		Test(ctx context.Context) error
	}

	Registry struct {
		resources map[string]ResourceFactory
	}
)

func newRegistry(configs []ResourceConfig) *Registry {
	resources := make(map[string]ResourceFactory)

	for _, c := range configs {
		if c.Scheme != "" {
			resources[c.Scheme] = c.Factory
		}
	}

	return &Registry{resources}
}

func (r *Registry) Register(scheme string, factory ResourceFactory) error {
	scheme = strings.TrimSpace(scheme)
	_, exists := r.resources[scheme]

	if exists {
		return errors.New("resource is already registered with a given scheme:" + scheme)
	}

	r.resources[scheme] = factory

	return nil
}

func (r *Registry) Resolve(location string) (Resource, error) {
	u, err := url.Parse(location)

	if err != nil {
		return nil, err
	}

	rf, found := r.resources[u.Scheme]

	if !found {
		return nil, errors.New("resource with a given scheme is not found:" + u.Scheme)
	}

	return rf(u)
}

func (r *Registry) List() []string {
	list := make([]string, 0, len(r.resources))

	for k := range r.resources {
		list = append(list, k)
	}

	return list
}
