// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var connect_pb = require('./connect_pb.js');
var read_pb = require('./read_pb.js');
var write_pb = require('./write_pb.js');
var relay_pb = require('./relay_pb.js');

function serialize_protos_CreateConnectionRequest(arg) {
  if (!(arg instanceof connect_pb.CreateConnectionRequest)) {
    throw new Error('Expected argument of type protos.CreateConnectionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_CreateConnectionRequest(buffer_arg) {
  return connect_pb.CreateConnectionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_CreateConnectionResponse(arg) {
  if (!(arg instanceof connect_pb.CreateConnectionResponse)) {
    throw new Error('Expected argument of type protos.CreateConnectionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_CreateConnectionResponse(buffer_arg) {
  return connect_pb.CreateConnectionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_CreateRelayRequest(arg) {
  if (!(arg instanceof relay_pb.CreateRelayRequest)) {
    throw new Error('Expected argument of type protos.CreateRelayRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_CreateRelayRequest(buffer_arg) {
  return relay_pb.CreateRelayRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_CreateRelayResponse(arg) {
  if (!(arg instanceof relay_pb.CreateRelayResponse)) {
    throw new Error('Expected argument of type protos.CreateRelayResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_CreateRelayResponse(buffer_arg) {
  return relay_pb.CreateRelayResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_DeleteConnectionRequest(arg) {
  if (!(arg instanceof connect_pb.DeleteConnectionRequest)) {
    throw new Error('Expected argument of type protos.DeleteConnectionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_DeleteConnectionRequest(buffer_arg) {
  return connect_pb.DeleteConnectionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_DeleteConnectionResponse(arg) {
  if (!(arg instanceof connect_pb.DeleteConnectionResponse)) {
    throw new Error('Expected argument of type protos.DeleteConnectionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_DeleteConnectionResponse(buffer_arg) {
  return connect_pb.DeleteConnectionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_DeleteRelayRequest(arg) {
  if (!(arg instanceof relay_pb.DeleteRelayRequest)) {
    throw new Error('Expected argument of type protos.DeleteRelayRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_DeleteRelayRequest(buffer_arg) {
  return relay_pb.DeleteRelayRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_DeleteRelayResponse(arg) {
  if (!(arg instanceof relay_pb.DeleteRelayResponse)) {
    throw new Error('Expected argument of type protos.DeleteRelayResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_DeleteRelayResponse(buffer_arg) {
  return relay_pb.DeleteRelayResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_GetAllConnectionsRequest(arg) {
  if (!(arg instanceof connect_pb.GetAllConnectionsRequest)) {
    throw new Error('Expected argument of type protos.GetAllConnectionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_GetAllConnectionsRequest(buffer_arg) {
  return connect_pb.GetAllConnectionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_GetAllConnectionsResponse(arg) {
  if (!(arg instanceof connect_pb.GetAllConnectionsResponse)) {
    throw new Error('Expected argument of type protos.GetAllConnectionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_GetAllConnectionsResponse(buffer_arg) {
  return connect_pb.GetAllConnectionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_GetAllReadsRequest(arg) {
  if (!(arg instanceof read_pb.GetAllReadsRequest)) {
    throw new Error('Expected argument of type protos.GetAllReadsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_GetAllReadsRequest(buffer_arg) {
  return read_pb.GetAllReadsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_GetAllReadsResponse(arg) {
  if (!(arg instanceof read_pb.GetAllReadsResponse)) {
    throw new Error('Expected argument of type protos.GetAllReadsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_GetAllReadsResponse(buffer_arg) {
  return read_pb.GetAllReadsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_GetAllRelaysRequest(arg) {
  if (!(arg instanceof relay_pb.GetAllRelaysRequest)) {
    throw new Error('Expected argument of type protos.GetAllRelaysRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_GetAllRelaysRequest(buffer_arg) {
  return relay_pb.GetAllRelaysRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_GetAllRelaysResponse(arg) {
  if (!(arg instanceof relay_pb.GetAllRelaysResponse)) {
    throw new Error('Expected argument of type protos.GetAllRelaysResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_GetAllRelaysResponse(buffer_arg) {
  return relay_pb.GetAllRelaysResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_GetConnectionRequest(arg) {
  if (!(arg instanceof connect_pb.GetConnectionRequest)) {
    throw new Error('Expected argument of type protos.GetConnectionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_GetConnectionRequest(buffer_arg) {
  return connect_pb.GetConnectionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_GetConnectionResponse(arg) {
  if (!(arg instanceof connect_pb.GetConnectionResponse)) {
    throw new Error('Expected argument of type protos.GetConnectionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_GetConnectionResponse(buffer_arg) {
  return connect_pb.GetConnectionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_ResumeRelayRequest(arg) {
  if (!(arg instanceof relay_pb.ResumeRelayRequest)) {
    throw new Error('Expected argument of type protos.ResumeRelayRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_ResumeRelayRequest(buffer_arg) {
  return relay_pb.ResumeRelayRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_ResumeRelayResponse(arg) {
  if (!(arg instanceof relay_pb.ResumeRelayResponse)) {
    throw new Error('Expected argument of type protos.ResumeRelayResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_ResumeRelayResponse(buffer_arg) {
  return relay_pb.ResumeRelayResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_StartReadRequest(arg) {
  if (!(arg instanceof read_pb.StartReadRequest)) {
    throw new Error('Expected argument of type protos.StartReadRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_StartReadRequest(buffer_arg) {
  return read_pb.StartReadRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_StartReadResponse(arg) {
  if (!(arg instanceof read_pb.StartReadResponse)) {
    throw new Error('Expected argument of type protos.StartReadResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_StartReadResponse(buffer_arg) {
  return read_pb.StartReadResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_StopReadRequest(arg) {
  if (!(arg instanceof read_pb.StopReadRequest)) {
    throw new Error('Expected argument of type protos.StopReadRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_StopReadRequest(buffer_arg) {
  return read_pb.StopReadRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_StopReadResponse(arg) {
  if (!(arg instanceof read_pb.StopReadResponse)) {
    throw new Error('Expected argument of type protos.StopReadResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_StopReadResponse(buffer_arg) {
  return read_pb.StopReadResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_StopRelayRequest(arg) {
  if (!(arg instanceof relay_pb.StopRelayRequest)) {
    throw new Error('Expected argument of type protos.StopRelayRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_StopRelayRequest(buffer_arg) {
  return relay_pb.StopRelayRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_StopRelayResponse(arg) {
  if (!(arg instanceof relay_pb.StopRelayResponse)) {
    throw new Error('Expected argument of type protos.StopRelayResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_StopRelayResponse(buffer_arg) {
  return relay_pb.StopRelayResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_StreamReadRequest(arg) {
  if (!(arg instanceof read_pb.StreamReadRequest)) {
    throw new Error('Expected argument of type protos.StreamReadRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_StreamReadRequest(buffer_arg) {
  return read_pb.StreamReadRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_StreamReadResponse(arg) {
  if (!(arg instanceof read_pb.StreamReadResponse)) {
    throw new Error('Expected argument of type protos.StreamReadResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_StreamReadResponse(buffer_arg) {
  return read_pb.StreamReadResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_TestConnectionRequest(arg) {
  if (!(arg instanceof connect_pb.TestConnectionRequest)) {
    throw new Error('Expected argument of type protos.TestConnectionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_TestConnectionRequest(buffer_arg) {
  return connect_pb.TestConnectionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_TestConnectionResponse(arg) {
  if (!(arg instanceof connect_pb.TestConnectionResponse)) {
    throw new Error('Expected argument of type protos.TestConnectionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_TestConnectionResponse(buffer_arg) {
  return connect_pb.TestConnectionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_UpdateConnectionRequest(arg) {
  if (!(arg instanceof connect_pb.UpdateConnectionRequest)) {
    throw new Error('Expected argument of type protos.UpdateConnectionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_UpdateConnectionRequest(buffer_arg) {
  return connect_pb.UpdateConnectionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_UpdateConnectionResponse(arg) {
  if (!(arg instanceof connect_pb.UpdateConnectionResponse)) {
    throw new Error('Expected argument of type protos.UpdateConnectionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_UpdateConnectionResponse(buffer_arg) {
  return connect_pb.UpdateConnectionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_UpdateRelayRequest(arg) {
  if (!(arg instanceof relay_pb.UpdateRelayRequest)) {
    throw new Error('Expected argument of type protos.UpdateRelayRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_UpdateRelayRequest(buffer_arg) {
  return relay_pb.UpdateRelayRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_UpdateRelayResponse(arg) {
  if (!(arg instanceof relay_pb.UpdateRelayResponse)) {
    throw new Error('Expected argument of type protos.UpdateRelayResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_UpdateRelayResponse(buffer_arg) {
  return relay_pb.UpdateRelayResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_WriteRequest(arg) {
  if (!(arg instanceof write_pb.WriteRequest)) {
    throw new Error('Expected argument of type protos.WriteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_WriteRequest(buffer_arg) {
  return write_pb.WriteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protos_WriteResponse(arg) {
  if (!(arg instanceof write_pb.WriteResponse)) {
    throw new Error('Expected argument of type protos.WriteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protos_WriteResponse(buffer_arg) {
  return write_pb.WriteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var PlumberServerService = exports.PlumberServerService = {
  // List configured/known connections
getAllConnections: {
    path: '/protos.PlumberServer/GetAllConnections',
    requestStream: false,
    responseStream: false,
    requestType: connect_pb.GetAllConnectionsRequest,
    responseType: connect_pb.GetAllConnectionsResponse,
    requestSerialize: serialize_protos_GetAllConnectionsRequest,
    requestDeserialize: deserialize_protos_GetAllConnectionsRequest,
    responseSerialize: serialize_protos_GetAllConnectionsResponse,
    responseDeserialize: deserialize_protos_GetAllConnectionsResponse,
  },
  // Fetch a specific connection by ID
getConnection: {
    path: '/protos.PlumberServer/GetConnection',
    requestStream: false,
    responseStream: false,
    requestType: connect_pb.GetConnectionRequest,
    responseType: connect_pb.GetConnectionResponse,
    requestSerialize: serialize_protos_GetConnectionRequest,
    requestDeserialize: deserialize_protos_GetConnectionRequest,
    responseSerialize: serialize_protos_GetConnectionResponse,
    responseDeserialize: deserialize_protos_GetConnectionResponse,
  },
  // Create a connection in plumber
createConnection: {
    path: '/protos.PlumberServer/CreateConnection',
    requestStream: false,
    responseStream: false,
    requestType: connect_pb.CreateConnectionRequest,
    responseType: connect_pb.CreateConnectionResponse,
    requestSerialize: serialize_protos_CreateConnectionRequest,
    requestDeserialize: deserialize_protos_CreateConnectionRequest,
    responseSerialize: serialize_protos_CreateConnectionResponse,
    responseDeserialize: deserialize_protos_CreateConnectionResponse,
  },
  // Test a connection before saving its configuration
testConnection: {
    path: '/protos.PlumberServer/TestConnection',
    requestStream: false,
    responseStream: false,
    requestType: connect_pb.TestConnectionRequest,
    responseType: connect_pb.TestConnectionResponse,
    requestSerialize: serialize_protos_TestConnectionRequest,
    requestDeserialize: deserialize_protos_TestConnectionRequest,
    responseSerialize: serialize_protos_TestConnectionResponse,
    responseDeserialize: deserialize_protos_TestConnectionResponse,
  },
  // Any active connections will be restarted
updateConnection: {
    path: '/protos.PlumberServer/UpdateConnection',
    requestStream: false,
    responseStream: false,
    requestType: connect_pb.UpdateConnectionRequest,
    responseType: connect_pb.UpdateConnectionResponse,
    requestSerialize: serialize_protos_UpdateConnectionRequest,
    requestDeserialize: deserialize_protos_UpdateConnectionRequest,
    responseSerialize: serialize_protos_UpdateConnectionResponse,
    responseDeserialize: deserialize_protos_UpdateConnectionResponse,
  },
  // If there are any active connections, delete will cause them to get closed
deleteConnection: {
    path: '/protos.PlumberServer/DeleteConnection',
    requestStream: false,
    responseStream: false,
    requestType: connect_pb.DeleteConnectionRequest,
    responseType: connect_pb.DeleteConnectionResponse,
    requestSerialize: serialize_protos_DeleteConnectionRequest,
    requestDeserialize: deserialize_protos_DeleteConnectionRequest,
    responseSerialize: serialize_protos_DeleteConnectionResponse,
    responseDeserialize: deserialize_protos_DeleteConnectionResponse,
  },
  // Start reading data from a connection
startRead: {
    path: '/protos.PlumberServer/StartRead',
    requestStream: false,
    responseStream: false,
    requestType: read_pb.StartReadRequest,
    responseType: read_pb.StartReadResponse,
    requestSerialize: serialize_protos_StartReadRequest,
    requestDeserialize: deserialize_protos_StartReadRequest,
    responseSerialize: serialize_protos_StartReadResponse,
    responseDeserialize: deserialize_protos_StartReadResponse,
  },
  // Streams messages received off of a read
streamRead: {
    path: '/protos.PlumberServer/StreamRead',
    requestStream: false,
    responseStream: true,
    requestType: read_pb.StreamReadRequest,
    responseType: read_pb.StreamReadResponse,
    requestSerialize: serialize_protos_StreamReadRequest,
    requestDeserialize: deserialize_protos_StreamReadRequest,
    responseSerialize: serialize_protos_StreamReadResponse,
    responseDeserialize: deserialize_protos_StreamReadResponse,
  },
  // List all reads that have been created
getAllReads: {
    path: '/protos.PlumberServer/GetAllReads',
    requestStream: false,
    responseStream: false,
    requestType: read_pb.GetAllReadsRequest,
    responseType: read_pb.GetAllReadsResponse,
    requestSerialize: serialize_protos_GetAllReadsRequest,
    requestDeserialize: deserialize_protos_GetAllReadsRequest,
    responseSerialize: serialize_protos_GetAllReadsResponse,
    responseDeserialize: deserialize_protos_GetAllReadsResponse,
  },
  // Stop reading data from a connection
stopRead: {
    path: '/protos.PlumberServer/StopRead',
    requestStream: false,
    responseStream: false,
    requestType: read_pb.StopReadRequest,
    responseType: read_pb.StopReadResponse,
    requestSerialize: serialize_protos_StopReadRequest,
    requestDeserialize: deserialize_protos_StopReadRequest,
    responseSerialize: serialize_protos_StopReadResponse,
    responseDeserialize: deserialize_protos_StopReadResponse,
  },
  // Write data to a connection
write: {
    path: '/protos.PlumberServer/Write',
    requestStream: false,
    responseStream: false,
    requestType: write_pb.WriteRequest,
    responseType: write_pb.WriteResponse,
    requestSerialize: serialize_protos_WriteRequest,
    requestDeserialize: deserialize_protos_WriteRequest,
    responseSerialize: serialize_protos_WriteResponse,
    responseDeserialize: deserialize_protos_WriteResponse,
  },
  // Create a data relay from plumber server to the Batch platform
createRelay: {
    path: '/protos.PlumberServer/CreateRelay',
    requestStream: false,
    responseStream: false,
    requestType: relay_pb.CreateRelayRequest,
    responseType: relay_pb.CreateRelayResponse,
    requestSerialize: serialize_protos_CreateRelayRequest,
    requestDeserialize: deserialize_protos_CreateRelayRequest,
    responseSerialize: serialize_protos_CreateRelayResponse,
    responseDeserialize: deserialize_protos_CreateRelayResponse,
  },
  // Update a relay (such as API token) - relay will be interrupted!
updateRelay: {
    path: '/protos.PlumberServer/UpdateRelay',
    requestStream: false,
    responseStream: false,
    requestType: relay_pb.UpdateRelayRequest,
    responseType: relay_pb.UpdateRelayResponse,
    requestSerialize: serialize_protos_UpdateRelayRequest,
    requestDeserialize: deserialize_protos_UpdateRelayRequest,
    responseSerialize: serialize_protos_UpdateRelayResponse,
    responseDeserialize: deserialize_protos_UpdateRelayResponse,
  },
  resumeRelay: {
    path: '/protos.PlumberServer/ResumeRelay',
    requestStream: false,
    responseStream: false,
    requestType: relay_pb.ResumeRelayRequest,
    responseType: relay_pb.ResumeRelayResponse,
    requestSerialize: serialize_protos_ResumeRelayRequest,
    requestDeserialize: deserialize_protos_ResumeRelayRequest,
    responseSerialize: serialize_protos_ResumeRelayResponse,
    responseDeserialize: deserialize_protos_ResumeRelayResponse,
  },
  stopRelay: {
    path: '/protos.PlumberServer/StopRelay',
    requestStream: false,
    responseStream: false,
    requestType: relay_pb.StopRelayRequest,
    responseType: relay_pb.StopRelayResponse,
    requestSerialize: serialize_protos_StopRelayRequest,
    requestDeserialize: deserialize_protos_StopRelayRequest,
    responseSerialize: serialize_protos_StopRelayResponse,
    responseDeserialize: deserialize_protos_StopRelayResponse,
  },
  getAllRelays: {
    path: '/protos.PlumberServer/GetAllRelays',
    requestStream: false,
    responseStream: false,
    requestType: relay_pb.GetAllRelaysRequest,
    responseType: relay_pb.GetAllRelaysResponse,
    requestSerialize: serialize_protos_GetAllRelaysRequest,
    requestDeserialize: deserialize_protos_GetAllRelaysRequest,
    responseSerialize: serialize_protos_GetAllRelaysResponse,
    responseDeserialize: deserialize_protos_GetAllRelaysResponse,
  },
  // Delete an existing relay
deleteRelay: {
    path: '/protos.PlumberServer/DeleteRelay',
    requestStream: false,
    responseStream: false,
    requestType: relay_pb.DeleteRelayRequest,
    responseType: relay_pb.DeleteRelayResponse,
    requestSerialize: serialize_protos_DeleteRelayRequest,
    requestDeserialize: deserialize_protos_DeleteRelayRequest,
    responseSerialize: serialize_protos_DeleteRelayResponse,
    responseDeserialize: deserialize_protos_DeleteRelayResponse,
  },
};

exports.PlumberServerClient = grpc.makeGenericClientConstructor(PlumberServerService);
