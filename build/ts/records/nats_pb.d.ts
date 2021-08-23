// package: protos.records
// file: records/nats.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Nats extends jspb.Message { 
    getSubject(): string;
    setSubject(value: string): Nats;
    getReplay(): string;
    setReplay(value: string): Nats;
    getQueue(): string;
    setQueue(value: string): Nats;
    getValue(): Uint8Array | string;
    getValue_asU8(): Uint8Array;
    getValue_asB64(): string;
    setValue(value: Uint8Array | string): Nats;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Nats.AsObject;
    static toObject(includeInstance: boolean, msg: Nats): Nats.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Nats, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Nats;
    static deserializeBinaryFromReader(message: Nats, reader: jspb.BinaryReader): Nats;
}

export namespace Nats {
    export type AsObject = {
        subject: string,
        replay: string,
        queue: string,
        value: Uint8Array | string,
    }
}
