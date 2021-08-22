// source: server.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.protos.ServerConfig', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protos.ServerConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.ServerConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.ServerConfig.displayName = 'proto.protos.ServerConfig';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protos.ServerConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.ServerConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.ServerConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.ServerConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    nodeId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    grpcListenAddress: jspb.Message.getFieldWithDefault(msg, 2, ""),
    authToken: jspb.Message.getFieldWithDefault(msg, 3, ""),
    initialCluster: jspb.Message.getFieldWithDefault(msg, 4, ""),
    advertisePeerUrl: jspb.Message.getFieldWithDefault(msg, 5, ""),
    advertiseClientUrl: jspb.Message.getFieldWithDefault(msg, 6, ""),
    listenerPeerUrl: jspb.Message.getFieldWithDefault(msg, 7, ""),
    listenerClientUrl: jspb.Message.getFieldWithDefault(msg, 8, ""),
    peerToken: jspb.Message.getFieldWithDefault(msg, 9, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protos.ServerConfig}
 */
proto.protos.ServerConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.ServerConfig;
  return proto.protos.ServerConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.ServerConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.ServerConfig}
 */
proto.protos.ServerConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setNodeId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setGrpcListenAddress(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setAuthToken(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setInitialCluster(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setAdvertisePeerUrl(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setAdvertiseClientUrl(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setListenerPeerUrl(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setListenerClientUrl(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setPeerToken(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protos.ServerConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.ServerConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.ServerConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.ServerConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNodeId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getGrpcListenAddress();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getAuthToken();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getInitialCluster();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getAdvertisePeerUrl();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getAdvertiseClientUrl();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getListenerPeerUrl();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getListenerClientUrl();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getPeerToken();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
};


/**
 * optional string node_id = 1;
 * @return {string}
 */
proto.protos.ServerConfig.prototype.getNodeId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.ServerConfig} returns this
 */
proto.protos.ServerConfig.prototype.setNodeId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string grpc_listen_address = 2;
 * @return {string}
 */
proto.protos.ServerConfig.prototype.getGrpcListenAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.ServerConfig} returns this
 */
proto.protos.ServerConfig.prototype.setGrpcListenAddress = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string auth_token = 3;
 * @return {string}
 */
proto.protos.ServerConfig.prototype.getAuthToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.ServerConfig} returns this
 */
proto.protos.ServerConfig.prototype.setAuthToken = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string initial_cluster = 4;
 * @return {string}
 */
proto.protos.ServerConfig.prototype.getInitialCluster = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.ServerConfig} returns this
 */
proto.protos.ServerConfig.prototype.setInitialCluster = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string advertise_peer_url = 5;
 * @return {string}
 */
proto.protos.ServerConfig.prototype.getAdvertisePeerUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.ServerConfig} returns this
 */
proto.protos.ServerConfig.prototype.setAdvertisePeerUrl = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string advertise_client_url = 6;
 * @return {string}
 */
proto.protos.ServerConfig.prototype.getAdvertiseClientUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.ServerConfig} returns this
 */
proto.protos.ServerConfig.prototype.setAdvertiseClientUrl = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string listener_peer_url = 7;
 * @return {string}
 */
proto.protos.ServerConfig.prototype.getListenerPeerUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.ServerConfig} returns this
 */
proto.protos.ServerConfig.prototype.setListenerPeerUrl = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string listener_client_url = 8;
 * @return {string}
 */
proto.protos.ServerConfig.prototype.getListenerClientUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.ServerConfig} returns this
 */
proto.protos.ServerConfig.prototype.setListenerClientUrl = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string peer_token = 9;
 * @return {string}
 */
proto.protos.ServerConfig.prototype.getPeerToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.ServerConfig} returns this
 */
proto.protos.ServerConfig.prototype.setPeerToken = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


goog.object.extend(exports, proto.protos);