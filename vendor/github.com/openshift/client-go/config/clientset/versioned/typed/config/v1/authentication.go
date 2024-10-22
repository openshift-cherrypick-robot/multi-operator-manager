// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AuthenticationsGetter has a method to return a AuthenticationInterface.
// A group's client should implement this interface.
type AuthenticationsGetter interface {
	Authentications() AuthenticationInterface
}

// AuthenticationInterface has methods to work with Authentication resources.
type AuthenticationInterface interface {
	Create(ctx context.Context, authentication *v1.Authentication, opts metav1.CreateOptions) (*v1.Authentication, error)
	Update(ctx context.Context, authentication *v1.Authentication, opts metav1.UpdateOptions) (*v1.Authentication, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, authentication *v1.Authentication, opts metav1.UpdateOptions) (*v1.Authentication, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Authentication, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AuthenticationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Authentication, err error)
	Apply(ctx context.Context, authentication *configv1.AuthenticationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Authentication, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, authentication *configv1.AuthenticationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Authentication, err error)
	AuthenticationExpansion
}

// authentications implements AuthenticationInterface
type authentications struct {
	*gentype.ClientWithListAndApply[*v1.Authentication, *v1.AuthenticationList, *configv1.AuthenticationApplyConfiguration]
}

// newAuthentications returns a Authentications
func newAuthentications(c *ConfigV1Client) *authentications {
	return &authentications{
		gentype.NewClientWithListAndApply[*v1.Authentication, *v1.AuthenticationList, *configv1.AuthenticationApplyConfiguration](
			"authentications",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.Authentication { return &v1.Authentication{} },
			func() *v1.AuthenticationList { return &v1.AuthenticationList{} }),
	}
}