package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/go-openapi/runtime"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_connections"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIVpnConnectionClient ...
type IBMPIVpnConnectionClient struct {
	session         *ibmpisession.IBMPISession
	cloudInstanceID string
	authInfo        runtime.ClientAuthInfoWriter
	ctx             context.Context
}

// NewIBMPIVpnConnectionClient ...
func NewIBMPIVpnConnectionClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVpnConnectionClient {
	authInfo := ibmpisession.NewAuth(sess, cloudInstanceID)
	return &IBMPIVpnConnectionClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		authInfo:        authInfo,
		ctx:             ctx,
	}
}

// Get information about a VPN connection
func (f *IBMPIVpnConnectionClient) Get(id string) (*models.VPNConnection, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVpnConnectionID(id)
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.GetVPNConnectionOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get VPN Connection %s", id)
	}
	return resp.Payload, nil
}

// Create a VPN connection
func (f *IBMPIVpnConnectionClient) Create(body *models.VPNConnectionCreate) (*models.VPNConnectionCreateResponse, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postaccepted, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPost(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVPNConnectionOperationFailed, f.cloudInstanceID, err)
	}
	if postaccepted != nil && postaccepted.Payload != nil {
		return postaccepted.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create VPN Connection")
}

// Update a VPN connection
func (f *IBMPIVpnConnectionClient) Update(id string, body *models.VPNConnectionUpdate) (*models.VPNConnection, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVpnConnectionID(id).
		WithBody(body)
	putok, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPut(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVPNConnectionOperationFailed, id, err)
	}
	if putok != nil && putok.Payload != nil {
		return putok.Payload, nil
	}
	return nil, fmt.Errorf("failed to Update VPN Connection %s", id)
}

// Get all VPN connections
func (f *IBMPIVpnConnectionClient) GetAll() (*models.VPNConnections, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsGetall(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Get all VPN Connections: %v", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all VPN Connections")
	}
	return resp.Payload, nil
}

// Delete a VPN connection
func (f *IBMPIVpnConnectionClient) Delete(id string) (*models.JobReference, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVpnConnectionID(id)
	delaccepted, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsDelete(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.DeleteVPNConnectionOperationFailed, id, err)
	}
	if delaccepted != nil && delaccepted.Payload != nil {
		return delaccepted.Payload, nil
	}
	return nil, nil
}

// Network get
func (f *IBMPIVpnConnectionClient) GetNetwork(id string) (*models.NetworkIds, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsNetworksGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVpnConnectionID(id)
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsNetworksGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Get Networks for VPN Connection %s: %v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Networks for VPN Connection %s", id)
	}
	return resp.Payload, nil
}

// Network attach
func (f *IBMPIVpnConnectionClient) AddNetwork(id, networkID string) (*models.JobReference, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsNetworksPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVpnConnectionID(id).
		WithBody(&models.NetworkID{NetworkID: &networkID})
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsNetworksPut(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Add Network %s to VPN Connection %s: %v", networkID, id, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, nil
}

// Network detach
func (f *IBMPIVpnConnectionClient) DeleteNetwork(id, networkID string) (*models.JobReference, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsNetworksDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVpnConnectionID(id).
		WithBody(&models.NetworkID{NetworkID: &networkID})
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsNetworksDelete(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Delete Network %s from VPN Connection %s: %v", networkID, id, err)
	}
	if resp != nil && resp.Payload != nil {
		return resp.Payload, nil
	}
	return nil, nil
}

// Subnet get
func (f *IBMPIVpnConnectionClient) GetSubnet(id string) (*models.PeerSubnets, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPeersubnetsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVpnConnectionID(id)
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPeersubnetsGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Get Subnets from VPN Connection %s: %v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Subnets from VPN Connection %s", id)
	}
	return resp.Payload, nil
}

// Subnet attach
func (f *IBMPIVpnConnectionClient) AddSubnet(id, subnet string) (*models.PeerSubnets, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPeersubnetsPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVpnConnectionID(id).
		WithBody(&models.PeerSubnetUpdate{Cidr: &subnet})
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPeersubnetsPut(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Add Subnets to VPN Connection %s: %v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Add Subnets to VPN Connection %s", id)
	}
	return resp.Payload, nil
}

// Subnet detach
func (f *IBMPIVpnConnectionClient) DeleteSubnet(id, subnet string) (*models.PeerSubnets, error) {
	params := p_cloud_v_p_n_connections.NewPcloudVpnconnectionsPeersubnetsDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVpnConnectionID(id).
		WithBody(&models.PeerSubnetUpdate{Cidr: &subnet})
	resp, err := f.session.Power.PCloudVPNConnections.PcloudVpnconnectionsPeersubnetsDelete(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Delete Subnet from VPN Connection %s: %v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Delete Subnet from VPN Connection %s", id)
	}
	return resp.Payload, nil
}
