// source: ps_vc_client.proto
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

var common_ps_common_auth_pb = require('./common/ps_common_auth_pb.js');
goog.object.extend(proto, common_ps_common_auth_pb);
goog.exportSymbol('proto.protos.GetVCEventsRequest', null, global);
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
proto.protos.GetVCEventsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 500, null, null);
};
goog.inherits(proto.protos.GetVCEventsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.GetVCEventsRequest.displayName = 'proto.protos.GetVCEventsRequest';
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
proto.protos.GetVCEventsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.GetVCEventsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.GetVCEventsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.GetVCEventsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    auth: (f = msg.getAuth()) && common_ps_common_auth_pb.Auth.toObject(includeInstance, f)
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
 * @return {!proto.protos.GetVCEventsRequest}
 */
proto.protos.GetVCEventsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.GetVCEventsRequest;
  return proto.protos.GetVCEventsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.GetVCEventsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.GetVCEventsRequest}
 */
proto.protos.GetVCEventsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 9999:
      var value = new common_ps_common_auth_pb.Auth;
      reader.readMessage(value,common_ps_common_auth_pb.Auth.deserializeBinaryFromReader);
      msg.setAuth(value);
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
proto.protos.GetVCEventsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.GetVCEventsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.GetVCEventsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.GetVCEventsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAuth();
  if (f != null) {
    writer.writeMessage(
      9999,
      f,
      common_ps_common_auth_pb.Auth.serializeBinaryToWriter
    );
  }
};


/**
 * optional common.Auth auth = 9999;
 * @return {?proto.protos.common.Auth}
 */
proto.protos.GetVCEventsRequest.prototype.getAuth = function() {
  return /** @type{?proto.protos.common.Auth} */ (
    jspb.Message.getWrapperField(this, common_ps_common_auth_pb.Auth, 9999));
};


/**
 * @param {?proto.protos.common.Auth|undefined} value
 * @return {!proto.protos.GetVCEventsRequest} returns this
*/
proto.protos.GetVCEventsRequest.prototype.setAuth = function(value) {
  return jspb.Message.setWrapperField(this, 9999, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.GetVCEventsRequest} returns this
 */
proto.protos.GetVCEventsRequest.prototype.clearAuth = function() {
  return this.setAuth(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.GetVCEventsRequest.prototype.hasAuth = function() {
  return jspb.Message.getField(this, 9999) != null;
};


goog.object.extend(exports, proto.protos);