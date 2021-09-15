// source: encoding/options.proto
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

goog.exportSymbol('proto.protos.encoding.AvroSettings', null, global);
goog.exportSymbol('proto.protos.encoding.DecodeOptions', null, global);
goog.exportSymbol('proto.protos.encoding.DecodeType', null, global);
goog.exportSymbol('proto.protos.encoding.EncodeOptions', null, global);
goog.exportSymbol('proto.protos.encoding.EncodeType', null, global);
goog.exportSymbol('proto.protos.encoding.ProtobufSettings', null, global);
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
proto.protos.encoding.ProtobufSettings = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.protos.encoding.ProtobufSettings.repeatedFields_, null);
};
goog.inherits(proto.protos.encoding.ProtobufSettings, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.encoding.ProtobufSettings.displayName = 'proto.protos.encoding.ProtobufSettings';
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
proto.protos.encoding.AvroSettings = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.encoding.AvroSettings, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.encoding.AvroSettings.displayName = 'proto.protos.encoding.AvroSettings';
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
proto.protos.encoding.EncodeOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.encoding.EncodeOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.encoding.EncodeOptions.displayName = 'proto.protos.encoding.EncodeOptions';
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
proto.protos.encoding.DecodeOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.encoding.DecodeOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.encoding.DecodeOptions.displayName = 'proto.protos.encoding.DecodeOptions';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.protos.encoding.ProtobufSettings.repeatedFields_ = [2];



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
proto.protos.encoding.ProtobufSettings.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.encoding.ProtobufSettings.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.encoding.ProtobufSettings} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.encoding.ProtobufSettings.toObject = function(includeInstance, msg) {
  var f, obj = {
    protobufRootMessage: jspb.Message.getFieldWithDefault(msg, 1, ""),
    protobufDirsList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    archive: msg.getArchive_asB64()
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
 * @return {!proto.protos.encoding.ProtobufSettings}
 */
proto.protos.encoding.ProtobufSettings.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.encoding.ProtobufSettings;
  return proto.protos.encoding.ProtobufSettings.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.encoding.ProtobufSettings} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.encoding.ProtobufSettings}
 */
proto.protos.encoding.ProtobufSettings.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProtobufRootMessage(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addProtobufDirs(value);
      break;
    case 3:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setArchive(value);
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
proto.protos.encoding.ProtobufSettings.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.encoding.ProtobufSettings.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.encoding.ProtobufSettings} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.encoding.ProtobufSettings.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProtobufRootMessage();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getProtobufDirsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getArchive_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      3,
      f
    );
  }
};


/**
 * optional string protobuf_root_message = 1;
 * @return {string}
 */
proto.protos.encoding.ProtobufSettings.prototype.getProtobufRootMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.encoding.ProtobufSettings} returns this
 */
proto.protos.encoding.ProtobufSettings.prototype.setProtobufRootMessage = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string protobuf_dirs = 2;
 * @return {!Array<string>}
 */
proto.protos.encoding.ProtobufSettings.prototype.getProtobufDirsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.protos.encoding.ProtobufSettings} returns this
 */
proto.protos.encoding.ProtobufSettings.prototype.setProtobufDirsList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.protos.encoding.ProtobufSettings} returns this
 */
proto.protos.encoding.ProtobufSettings.prototype.addProtobufDirs = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.protos.encoding.ProtobufSettings} returns this
 */
proto.protos.encoding.ProtobufSettings.prototype.clearProtobufDirsList = function() {
  return this.setProtobufDirsList([]);
};


/**
 * optional bytes archive = 3;
 * @return {!(string|Uint8Array)}
 */
proto.protos.encoding.ProtobufSettings.prototype.getArchive = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * optional bytes archive = 3;
 * This is a type-conversion wrapper around `getArchive()`
 * @return {string}
 */
proto.protos.encoding.ProtobufSettings.prototype.getArchive_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getArchive()));
};


/**
 * optional bytes archive = 3;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getArchive()`
 * @return {!Uint8Array}
 */
proto.protos.encoding.ProtobufSettings.prototype.getArchive_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getArchive()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.protos.encoding.ProtobufSettings} returns this
 */
proto.protos.encoding.ProtobufSettings.prototype.setArchive = function(value) {
  return jspb.Message.setProto3BytesField(this, 3, value);
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
proto.protos.encoding.AvroSettings.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.encoding.AvroSettings.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.encoding.AvroSettings} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.encoding.AvroSettings.toObject = function(includeInstance, msg) {
  var f, obj = {
    avroSchemaFile: jspb.Message.getFieldWithDefault(msg, 1, ""),
    schema: msg.getSchema_asB64()
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
 * @return {!proto.protos.encoding.AvroSettings}
 */
proto.protos.encoding.AvroSettings.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.encoding.AvroSettings;
  return proto.protos.encoding.AvroSettings.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.encoding.AvroSettings} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.encoding.AvroSettings}
 */
proto.protos.encoding.AvroSettings.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAvroSchemaFile(value);
      break;
    case 2:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setSchema(value);
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
proto.protos.encoding.AvroSettings.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.encoding.AvroSettings.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.encoding.AvroSettings} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.encoding.AvroSettings.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAvroSchemaFile();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSchema_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      2,
      f
    );
  }
};


/**
 * optional string avro_schema_file = 1;
 * @return {string}
 */
proto.protos.encoding.AvroSettings.prototype.getAvroSchemaFile = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.encoding.AvroSettings} returns this
 */
proto.protos.encoding.AvroSettings.prototype.setAvroSchemaFile = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional bytes schema = 2;
 * @return {!(string|Uint8Array)}
 */
proto.protos.encoding.AvroSettings.prototype.getSchema = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes schema = 2;
 * This is a type-conversion wrapper around `getSchema()`
 * @return {string}
 */
proto.protos.encoding.AvroSettings.prototype.getSchema_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getSchema()));
};


/**
 * optional bytes schema = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getSchema()`
 * @return {!Uint8Array}
 */
proto.protos.encoding.AvroSettings.prototype.getSchema_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getSchema()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.protos.encoding.AvroSettings} returns this
 */
proto.protos.encoding.AvroSettings.prototype.setSchema = function(value) {
  return jspb.Message.setProto3BytesField(this, 2, value);
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
proto.protos.encoding.EncodeOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.encoding.EncodeOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.encoding.EncodeOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.encoding.EncodeOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    schemaId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    encodeType: jspb.Message.getFieldWithDefault(msg, 2, 0),
    protobufSettings: (f = msg.getProtobufSettings()) && proto.protos.encoding.ProtobufSettings.toObject(includeInstance, f),
    avroSettings: (f = msg.getAvroSettings()) && proto.protos.encoding.AvroSettings.toObject(includeInstance, f)
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
 * @return {!proto.protos.encoding.EncodeOptions}
 */
proto.protos.encoding.EncodeOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.encoding.EncodeOptions;
  return proto.protos.encoding.EncodeOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.encoding.EncodeOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.encoding.EncodeOptions}
 */
proto.protos.encoding.EncodeOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setSchemaId(value);
      break;
    case 2:
      var value = /** @type {!proto.protos.encoding.EncodeType} */ (reader.readEnum());
      msg.setEncodeType(value);
      break;
    case 3:
      var value = new proto.protos.encoding.ProtobufSettings;
      reader.readMessage(value,proto.protos.encoding.ProtobufSettings.deserializeBinaryFromReader);
      msg.setProtobufSettings(value);
      break;
    case 4:
      var value = new proto.protos.encoding.AvroSettings;
      reader.readMessage(value,proto.protos.encoding.AvroSettings.deserializeBinaryFromReader);
      msg.setAvroSettings(value);
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
proto.protos.encoding.EncodeOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.encoding.EncodeOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.encoding.EncodeOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.encoding.EncodeOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSchemaId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getEncodeType();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getProtobufSettings();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.protos.encoding.ProtobufSettings.serializeBinaryToWriter
    );
  }
  f = message.getAvroSettings();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.protos.encoding.AvroSettings.serializeBinaryToWriter
    );
  }
};


/**
 * optional string schema_id = 1;
 * @return {string}
 */
proto.protos.encoding.EncodeOptions.prototype.getSchemaId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.encoding.EncodeOptions} returns this
 */
proto.protos.encoding.EncodeOptions.prototype.setSchemaId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional EncodeType encode_type = 2;
 * @return {!proto.protos.encoding.EncodeType}
 */
proto.protos.encoding.EncodeOptions.prototype.getEncodeType = function() {
  return /** @type {!proto.protos.encoding.EncodeType} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.protos.encoding.EncodeType} value
 * @return {!proto.protos.encoding.EncodeOptions} returns this
 */
proto.protos.encoding.EncodeOptions.prototype.setEncodeType = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional ProtobufSettings protobuf_settings = 3;
 * @return {?proto.protos.encoding.ProtobufSettings}
 */
proto.protos.encoding.EncodeOptions.prototype.getProtobufSettings = function() {
  return /** @type{?proto.protos.encoding.ProtobufSettings} */ (
    jspb.Message.getWrapperField(this, proto.protos.encoding.ProtobufSettings, 3));
};


/**
 * @param {?proto.protos.encoding.ProtobufSettings|undefined} value
 * @return {!proto.protos.encoding.EncodeOptions} returns this
*/
proto.protos.encoding.EncodeOptions.prototype.setProtobufSettings = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.encoding.EncodeOptions} returns this
 */
proto.protos.encoding.EncodeOptions.prototype.clearProtobufSettings = function() {
  return this.setProtobufSettings(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.encoding.EncodeOptions.prototype.hasProtobufSettings = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional AvroSettings avro_settings = 4;
 * @return {?proto.protos.encoding.AvroSettings}
 */
proto.protos.encoding.EncodeOptions.prototype.getAvroSettings = function() {
  return /** @type{?proto.protos.encoding.AvroSettings} */ (
    jspb.Message.getWrapperField(this, proto.protos.encoding.AvroSettings, 4));
};


/**
 * @param {?proto.protos.encoding.AvroSettings|undefined} value
 * @return {!proto.protos.encoding.EncodeOptions} returns this
*/
proto.protos.encoding.EncodeOptions.prototype.setAvroSettings = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.encoding.EncodeOptions} returns this
 */
proto.protos.encoding.EncodeOptions.prototype.clearAvroSettings = function() {
  return this.setAvroSettings(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.encoding.EncodeOptions.prototype.hasAvroSettings = function() {
  return jspb.Message.getField(this, 4) != null;
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
proto.protos.encoding.DecodeOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.encoding.DecodeOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.encoding.DecodeOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.encoding.DecodeOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    schemaId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    decodeType: jspb.Message.getFieldWithDefault(msg, 2, 0),
    protobufSettings: (f = msg.getProtobufSettings()) && proto.protos.encoding.ProtobufSettings.toObject(includeInstance, f),
    avroSettings: (f = msg.getAvroSettings()) && proto.protos.encoding.AvroSettings.toObject(includeInstance, f)
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
 * @return {!proto.protos.encoding.DecodeOptions}
 */
proto.protos.encoding.DecodeOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.encoding.DecodeOptions;
  return proto.protos.encoding.DecodeOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.encoding.DecodeOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.encoding.DecodeOptions}
 */
proto.protos.encoding.DecodeOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setSchemaId(value);
      break;
    case 2:
      var value = /** @type {!proto.protos.encoding.DecodeType} */ (reader.readEnum());
      msg.setDecodeType(value);
      break;
    case 3:
      var value = new proto.protos.encoding.ProtobufSettings;
      reader.readMessage(value,proto.protos.encoding.ProtobufSettings.deserializeBinaryFromReader);
      msg.setProtobufSettings(value);
      break;
    case 4:
      var value = new proto.protos.encoding.AvroSettings;
      reader.readMessage(value,proto.protos.encoding.AvroSettings.deserializeBinaryFromReader);
      msg.setAvroSettings(value);
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
proto.protos.encoding.DecodeOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.encoding.DecodeOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.encoding.DecodeOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.encoding.DecodeOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSchemaId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDecodeType();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getProtobufSettings();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.protos.encoding.ProtobufSettings.serializeBinaryToWriter
    );
  }
  f = message.getAvroSettings();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.protos.encoding.AvroSettings.serializeBinaryToWriter
    );
  }
};


/**
 * optional string schema_id = 1;
 * @return {string}
 */
proto.protos.encoding.DecodeOptions.prototype.getSchemaId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.encoding.DecodeOptions} returns this
 */
proto.protos.encoding.DecodeOptions.prototype.setSchemaId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional DecodeType decode_type = 2;
 * @return {!proto.protos.encoding.DecodeType}
 */
proto.protos.encoding.DecodeOptions.prototype.getDecodeType = function() {
  return /** @type {!proto.protos.encoding.DecodeType} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.protos.encoding.DecodeType} value
 * @return {!proto.protos.encoding.DecodeOptions} returns this
 */
proto.protos.encoding.DecodeOptions.prototype.setDecodeType = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional ProtobufSettings protobuf_settings = 3;
 * @return {?proto.protos.encoding.ProtobufSettings}
 */
proto.protos.encoding.DecodeOptions.prototype.getProtobufSettings = function() {
  return /** @type{?proto.protos.encoding.ProtobufSettings} */ (
    jspb.Message.getWrapperField(this, proto.protos.encoding.ProtobufSettings, 3));
};


/**
 * @param {?proto.protos.encoding.ProtobufSettings|undefined} value
 * @return {!proto.protos.encoding.DecodeOptions} returns this
*/
proto.protos.encoding.DecodeOptions.prototype.setProtobufSettings = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.encoding.DecodeOptions} returns this
 */
proto.protos.encoding.DecodeOptions.prototype.clearProtobufSettings = function() {
  return this.setProtobufSettings(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.encoding.DecodeOptions.prototype.hasProtobufSettings = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional AvroSettings avro_settings = 4;
 * @return {?proto.protos.encoding.AvroSettings}
 */
proto.protos.encoding.DecodeOptions.prototype.getAvroSettings = function() {
  return /** @type{?proto.protos.encoding.AvroSettings} */ (
    jspb.Message.getWrapperField(this, proto.protos.encoding.AvroSettings, 4));
};


/**
 * @param {?proto.protos.encoding.AvroSettings|undefined} value
 * @return {!proto.protos.encoding.DecodeOptions} returns this
*/
proto.protos.encoding.DecodeOptions.prototype.setAvroSettings = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.encoding.DecodeOptions} returns this
 */
proto.protos.encoding.DecodeOptions.prototype.clearAvroSettings = function() {
  return this.setAvroSettings(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.encoding.DecodeOptions.prototype.hasAvroSettings = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * @enum {number}
 */
proto.protos.encoding.EncodeType = {
  ENCODE_TYPE_UNSET: 0,
  ENCODE_TYPE_JSONPB: 1,
  ENCODE_TYPE_AVRO: 2
};

/**
 * @enum {number}
 */
proto.protos.encoding.DecodeType = {
  DECODE_TYPE_UNSET: 0,
  DECODE_TYPE_JSONPB: 1,
  DECODE_TYPE_PROTOBUF: 2,
  DECODE_TYPE_AVRO: 3,
  DECODE_TYPE_THRIFT: 4,
  DECODE_TYPE_FLATBUFFER: 5
};

goog.object.extend(exports, proto.protos.encoding);
