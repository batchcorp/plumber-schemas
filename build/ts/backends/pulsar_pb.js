// source: backends/pulsar.proto
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

goog.exportSymbol('proto.protos.backends.Pulsar', null, global);
goog.exportSymbol('proto.protos.backends.SubscriptionType', null, global);
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
proto.protos.backends.Pulsar = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.backends.Pulsar, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.backends.Pulsar.displayName = 'proto.protos.backends.Pulsar';
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
proto.protos.backends.Pulsar.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.backends.Pulsar.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.backends.Pulsar} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.Pulsar.toObject = function(includeInstance, msg) {
  var f, obj = {
    topic: jspb.Message.getFieldWithDefault(msg, 1, ""),
    subscriptionName: jspb.Message.getFieldWithDefault(msg, 2, ""),
    subscriptionType: jspb.Message.getFieldWithDefault(msg, 3, 0)
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
 * @return {!proto.protos.backends.Pulsar}
 */
proto.protos.backends.Pulsar.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.backends.Pulsar;
  return proto.protos.backends.Pulsar.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.backends.Pulsar} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.backends.Pulsar}
 */
proto.protos.backends.Pulsar.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setTopic(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setSubscriptionName(value);
      break;
    case 3:
      var value = /** @type {!proto.protos.backends.SubscriptionType} */ (reader.readEnum());
      msg.setSubscriptionType(value);
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
proto.protos.backends.Pulsar.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.backends.Pulsar.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.backends.Pulsar} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.Pulsar.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTopic();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSubscriptionName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getSubscriptionType();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
};


/**
 * optional string topic = 1;
 * @return {string}
 */
proto.protos.backends.Pulsar.prototype.getTopic = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.Pulsar} returns this
 */
proto.protos.backends.Pulsar.prototype.setTopic = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string subscription_name = 2;
 * @return {string}
 */
proto.protos.backends.Pulsar.prototype.getSubscriptionName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.Pulsar} returns this
 */
proto.protos.backends.Pulsar.prototype.setSubscriptionName = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional SubscriptionType subscription_type = 3;
 * @return {!proto.protos.backends.SubscriptionType}
 */
proto.protos.backends.Pulsar.prototype.getSubscriptionType = function() {
  return /** @type {!proto.protos.backends.SubscriptionType} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.protos.backends.SubscriptionType} value
 * @return {!proto.protos.backends.Pulsar} returns this
 */
proto.protos.backends.Pulsar.prototype.setSubscriptionType = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * @enum {number}
 */
proto.protos.backends.SubscriptionType = {
  UNSET: 0,
  SHARED: 1,
  EXCLUSIVE: 2,
  FAILOVER: 3,
  KEYSHARED: 4
};

goog.object.extend(exports, proto.protos.backends);