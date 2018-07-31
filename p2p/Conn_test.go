/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */
package p2p

import crypto "github.com/libp2p/go-libp2p-crypto"
import mock "github.com/stretchr/testify/mock"
import multiaddr "github.com/multiformats/go-multiaddr"
import net "github.com/libp2p/go-libp2p-net"
import peer "github.com/libp2p/go-libp2p-peer"

// MockConn is an autogenerated mock type for the MockConn type
type MockConn struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockConn) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetStreams provides a mock function with given fields:
func (_m *MockConn) GetStreams() []net.Stream {
	ret := _m.Called()

	var r0 []net.Stream
	if rf, ok := ret.Get(0).(func() []net.Stream); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]net.Stream)
		}
	}

	return r0
}

// LocalMultiaddr provides a mock function with given fields:
func (_m *MockConn) LocalMultiaddr() multiaddr.Multiaddr {
	ret := _m.Called()

	var r0 multiaddr.Multiaddr
	if rf, ok := ret.Get(0).(func() multiaddr.Multiaddr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(multiaddr.Multiaddr)
		}
	}

	return r0
}

// LocalPeer provides a mock function with given fields:
func (_m *MockConn) LocalPeer() peer.ID {
	ret := _m.Called()

	var r0 peer.ID
	if rf, ok := ret.Get(0).(func() peer.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(peer.ID)
	}

	return r0
}

// LocalPrivateKey provides a mock function with given fields:
func (_m *MockConn) LocalPrivateKey() crypto.PrivKey {
	ret := _m.Called()

	var r0 crypto.PrivKey
	if rf, ok := ret.Get(0).(func() crypto.PrivKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivKey)
		}
	}

	return r0
}

// NewStream provides a mock function with given fields:
func (_m *MockConn) NewStream() (net.Stream, error) {
	ret := _m.Called()

	var r0 net.Stream
	if rf, ok := ret.Get(0).(func() net.Stream); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Stream)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteMultiaddr provides a mock function with given fields:
func (_m *MockConn) RemoteMultiaddr() multiaddr.Multiaddr {
	ret := _m.Called()

	var r0 multiaddr.Multiaddr
	if rf, ok := ret.Get(0).(func() multiaddr.Multiaddr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(multiaddr.Multiaddr)
		}
	}

	return r0
}

// RemotePeer provides a mock function with given fields:
func (_m *MockConn) RemotePeer() peer.ID {
	ret := _m.Called()

	var r0 peer.ID
	if rf, ok := ret.Get(0).(func() peer.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(peer.ID)
	}

	return r0
}

// RemotePublicKey provides a mock function with given fields:
func (_m *MockConn) RemotePublicKey() crypto.PubKey {
	ret := _m.Called()

	var r0 crypto.PubKey
	if rf, ok := ret.Get(0).(func() crypto.PubKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PubKey)
		}
	}

	return r0
}
