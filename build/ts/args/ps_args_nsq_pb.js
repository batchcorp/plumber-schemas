// source: args/ps_args_nsq.proto
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

goog.exportSymbol('proto.protos.args.NSQConn', null, global);
goog.exportSymbol('proto.protos.args.NSQReadArgs', null, global);
goog.exportSymbol('proto.protos.args.NSQWriteArgs', null, global);
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
proto.protos.args.NSQConn = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.args.NSQConn, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.args.NSQConn.displayName = 'proto.protos.args.NSQConn';
}
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
proto.protos.args.NSQReadArgs = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.args.NSQReadArgs, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.args.NSQReadArgs.displayName = 'proto.protos.args.NSQReadArgs';
}
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
proto.protos.args.NSQWriteArgs = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.args.NSQWriteArgs, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.args.NSQWriteArgs.displayName = 'proto.protos.args.NSQWriteArgs';
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
proto.protos.args.NSQConn.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.args.NSQConn.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.args.NSQConn} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.NSQConn.toObject = function(includeInstance, msg) {
  var f, obj = {
    nsqdAddress: jspb.Message.getFieldWithDefault(msg, 1, ""),
    lookupdAddress: jspb.Message.getFieldWithDefault(msg, 2, ""),
    useTls: jspb.Message.getBooleanFieldWithDefault(msg, 3, false),
    insecureTls: jspb.Message.getBooleanFieldWithDefault(msg, 4, false),
    tlsCaCert: msg.getTlsCaCert_asB64(),
    tlsClientCert: msg.getTlsClientCert_asB64(),
    tlsClientKey: msg.getTlsClientKey_asB64(),
    authSecret: jspb.Message.getFieldWithDefault(msg, 8, ""),
    clientId: jspb.Message.getFieldWithDefault(msg, 9, "")
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
 * @return {!proto.protos.args.NSQConn}
 */
proto.protos.args.NSQConn.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.args.NSQConn;
  return proto.protos.args.NSQConn.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.args.NSQConn} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.args.NSQConn}
 */
proto.protos.args.NSQConn.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setNsqdAddress(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLookupdAddress(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUseTls(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setInsecureTls(value);
      break;
    case 5:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setTlsCaCert(value);
      break;
    case 6:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setTlsClientCert(value);
      break;
    case 7:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setTlsClientKey(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setAuthSecret(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setClientId(value);
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
proto.protos.args.NSQConn.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.args.NSQConn.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.args.NSQConn} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.NSQConn.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNsqdAddress();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getLookupdAddress();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getUseTls();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getInsecureTls();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getTlsCaCert_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      5,
      f
    );
  }
  f = message.getTlsClientCert_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      6,
      f
    );
  }
  f = message.getTlsClientKey_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      7,
      f
    );
  }
  f = message.getAuthSecret();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getClientId();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
};


/**
 * optional string nsqd_address = 1;
 * @return {string}
 */
proto.protos.args.NSQConn.prototype.getNsqdAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.NSQConn} returns this
 */
proto.protos.args.NSQConn.prototype.setNsqdAddress = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string lookupd_address = 2;
 * @return {string}
 */
proto.protos.args.NSQConn.prototype.getLookupdAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.NSQConn} returns this
 */
proto.protos.args.NSQConn.prototype.setLookupdAddress = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional bool use_tls = 3;
 * @return {boolean}
 */
proto.protos.args.NSQConn.prototype.getUseTls = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.args.NSQConn} returns this
 */
proto.protos.args.NSQConn.prototype.setUseTls = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * optional bool insecure_tls = 4;
 * @return {boolean}
 */
proto.protos.args.NSQConn.prototype.getInsecureTls = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 4, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.args.NSQConn} returns this
 */
proto.protos.args.NSQConn.prototype.setInsecureTls = function(value) {
  return jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional bytes tls_ca_cert = 5;
 * @return {!(string|Uint8Array)}
 */
proto.protos.args.NSQConn.prototype.getTlsCaCert = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * optional bytes tls_ca_cert = 5;
 * This is a type-conversion wrapper around `getTlsCaCert()`
 * @return {string}
 */
proto.protos.args.NSQConn.prototype.getTlsCaCert_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getTlsCaCert()));
};


/**
 * optional bytes tls_ca_cert = 5;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getTlsCaCert()`
 * @return {!Uint8Array}
 */
proto.protos.args.NSQConn.prototype.getTlsCaCert_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getTlsCaCert()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.protos.args.NSQConn} returns this
 */
proto.protos.args.NSQConn.prototype.setTlsCaCert = function(value) {
  return jspb.Message.setProto3BytesField(this, 5, value);
};


/**
 * optional bytes tls_client_cert = 6;
 * @return {!(string|Uint8Array)}
 */
proto.protos.args.NSQConn.prototype.getTlsClientCert = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * optional bytes tls_client_cert = 6;
 * This is a type-conversion wrapper around `getTlsClientCert()`
 * @return {string}
 */
proto.protos.args.NSQConn.prototype.getTlsClientCert_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getTlsClientCert()));
};


/**
 * optional bytes tls_client_cert = 6;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getTlsClientCert()`
 * @return {!Uint8Array}
 */
proto.protos.args.NSQConn.prototype.getTlsClientCert_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getTlsClientCert()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.protos.args.NSQConn} returns this
 */
proto.protos.args.NSQConn.prototype.setTlsClientCert = function(value) {
  return jspb.Message.setProto3BytesField(this, 6, value);
};


/**
 * optional bytes tls_client_key = 7;
 * @return {!(string|Uint8Array)}
 */
proto.protos.args.NSQConn.prototype.getTlsClientKey = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * optional bytes tls_client_key = 7;
 * This is a type-conversion wrapper around `getTlsClientKey()`
 * @return {string}
 */
proto.protos.args.NSQConn.prototype.getTlsClientKey_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getTlsClientKey()));
};


/**
 * optional bytes tls_client_key = 7;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getTlsClientKey()`
 * @return {!Uint8Array}
 */
proto.protos.args.NSQConn.prototype.getTlsClientKey_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getTlsClientKey()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.protos.args.NSQConn} returns this
 */
proto.protos.args.NSQConn.prototype.setTlsClientKey = function(value) {
  return jspb.Message.setProto3BytesField(this, 7, value);
};


/**
 * optional string auth_secret = 8;
 * @return {string}
 */
proto.protos.args.NSQConn.prototype.getAuthSecret = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.NSQConn} returns this
 */
proto.protos.args.NSQConn.prototype.setAuthSecret = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string client_id = 9;
 * @return {string}
 */
proto.protos.args.NSQConn.prototype.getClientId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.NSQConn} returns this
 */
proto.protos.args.NSQConn.prototype.setClientId = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};





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
proto.protos.args.NSQReadArgs.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.args.NSQReadArgs.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.args.NSQReadArgs} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.NSQReadArgs.toObject = function(includeInstance, msg) {
  var f, obj = {
    topic: jspb.Message.getFieldWithDefault(msg, 1, ""),
    channel: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.protos.args.NSQReadArgs}
 */
proto.protos.args.NSQReadArgs.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.args.NSQReadArgs;
  return proto.protos.args.NSQReadArgs.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.args.NSQReadArgs} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.args.NSQReadArgs}
 */
proto.protos.args.NSQReadArgs.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setChannel(value);
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
proto.protos.args.NSQReadArgs.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.args.NSQReadArgs.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.args.NSQReadArgs} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.NSQReadArgs.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTopic();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getChannel();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string topic = 1;
 * @return {string}
 */
proto.protos.args.NSQReadArgs.prototype.getTopic = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.NSQReadArgs} returns this
 */
proto.protos.args.NSQReadArgs.prototype.setTopic = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string channel = 2;
 * @return {string}
 */
proto.protos.args.NSQReadArgs.prototype.getChannel = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.NSQReadArgs} returns this
 */
proto.protos.args.NSQReadArgs.prototype.setChannel = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};





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
proto.protos.args.NSQWriteArgs.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.args.NSQWriteArgs.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.args.NSQWriteArgs} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.NSQWriteArgs.toObject = function(includeInstance, msg) {
  var f, obj = {
    topic: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.protos.args.NSQWriteArgs}
 */
proto.protos.args.NSQWriteArgs.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.args.NSQWriteArgs;
  return proto.protos.args.NSQWriteArgs.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.args.NSQWriteArgs} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.args.NSQWriteArgs}
 */
proto.protos.args.NSQWriteArgs.deserializeBinaryFromReader = function(msg, reader) {
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
proto.protos.args.NSQWriteArgs.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.args.NSQWriteArgs.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.args.NSQWriteArgs} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.args.NSQWriteArgs.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTopic();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string topic = 1;
 * @return {string}
 */
proto.protos.args.NSQWriteArgs.prototype.getTopic = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.args.NSQWriteArgs} returns this
 */
proto.protos.args.NSQWriteArgs.prototype.setTopic = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


goog.object.extend(exports, proto.protos.args);