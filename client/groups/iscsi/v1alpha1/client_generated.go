// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/kubernetes-csi/csi-proxy/client"
	"github.com/kubernetes-csi/csi-proxy/client/api/iscsi/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"google.golang.org/grpc"
)

const groupName = "iscsi"

var version = apiversion.NewVersionOrPanic("v1alpha1")

type Client struct {
	client     v1alpha1.IscsiClient
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the iscsi API group version v1alpha1.
// It's the caller's responsibility to Close the client when done.
func NewClient() (*Client, error) {
	pipePath := client.PipePath(groupName, version)

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := v1alpha1.NewIscsiClient(connection)
	return &Client{
		client:     client,
		connection: connection,
	}, nil
}

// Close closes the client. It must be called before the client gets GC-ed.
func (w *Client) Close() error {
	return w.connection.Close()
}

// ensures we implement all the required methods
var _ v1alpha1.IscsiClient = &Client{}

func (w *Client) AddTargetPortal(context context.Context, request *v1alpha1.AddTargetPortalRequest, opts ...grpc.CallOption) (*v1alpha1.AddTargetPortalResponse, error) {
	return w.client.AddTargetPortal(context, request, opts...)
}

func (w *Client) ConnectTarget(context context.Context, request *v1alpha1.ConnectTargetRequest, opts ...grpc.CallOption) (*v1alpha1.ConnectTargetResponse, error) {
	return w.client.ConnectTarget(context, request, opts...)
}

func (w *Client) DisconnectTarget(context context.Context, request *v1alpha1.DisconnectTargetRequest, opts ...grpc.CallOption) (*v1alpha1.DisconnectTargetResponse, error) {
	return w.client.DisconnectTarget(context, request, opts...)
}

func (w *Client) DiscoverTargetPortal(context context.Context, request *v1alpha1.DiscoverTargetPortalRequest, opts ...grpc.CallOption) (*v1alpha1.DiscoverTargetPortalResponse, error) {
	return w.client.DiscoverTargetPortal(context, request, opts...)
}

func (w *Client) GetTargetDisks(context context.Context, request *v1alpha1.GetTargetDisksRequest, opts ...grpc.CallOption) (*v1alpha1.GetTargetDisksResponse, error) {
	return w.client.GetTargetDisks(context, request, opts...)
}

func (w *Client) ListTargetPortals(context context.Context, request *v1alpha1.ListTargetPortalsRequest, opts ...grpc.CallOption) (*v1alpha1.ListTargetPortalsResponse, error) {
	return w.client.ListTargetPortals(context, request, opts...)
}

func (w *Client) RemoveTargetPortal(context context.Context, request *v1alpha1.RemoveTargetPortalRequest, opts ...grpc.CallOption) (*v1alpha1.RemoveTargetPortalResponse, error) {
	return w.client.RemoveTargetPortal(context, request, opts...)
}