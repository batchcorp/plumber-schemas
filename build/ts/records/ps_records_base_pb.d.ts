// package: protos.records
// file: records/ps_records_base.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as records_ps_records_kafka_pb from "../records/ps_records_kafka_pb";

export class ReadRecord extends jspb.Message { 
    getMessageId(): string;
    setMessageId(value: string): ReadRecord;
    getNum(): number;
    setNum(value: number): ReadRecord;

    getMetadataMap(): jspb.Map<string, string>;
    clearMetadataMap(): void;
    getReceivedAtUnixTsUtc(): number;
    setReceivedAtUnixTsUtc(value: number): ReadRecord;
    getPayload(): Uint8Array | string;
    getPayload_asU8(): Uint8Array;
    getPayload_asB64(): string;
    setPayload(value: Uint8Array | string): ReadRecord;

    hasKafka(): boolean;
    clearKafka(): void;
    getKafka(): records_ps_records_kafka_pb.Kafka | undefined;
    setKafka(value?: records_ps_records_kafka_pb.Kafka): ReadRecord;
    getRaw(): Uint8Array | string;
    getRaw_asU8(): Uint8Array;
    getRaw_asB64(): string;
    setRaw(value: Uint8Array | string): ReadRecord;
    getPlumberId(): string;
    setPlumberId(value: string): ReadRecord;

    getRecordCase(): ReadRecord.RecordCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ReadRecord.AsObject;
    static toObject(includeInstance: boolean, msg: ReadRecord): ReadRecord.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ReadRecord, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ReadRecord;
    static deserializeBinaryFromReader(message: ReadRecord, reader: jspb.BinaryReader): ReadRecord;
}

export namespace ReadRecord {
    export type AsObject = {
        messageId: string,
        num: number,

        metadataMap: Array<[string, string]>,
        receivedAtUnixTsUtc: number,
        payload: Uint8Array | string,
        kafka?: records_ps_records_kafka_pb.Kafka.AsObject,
        Raw: Uint8Array | string,
        PlumberId: string,
    }

    export enum RecordCase {
        RECORD_NOT_SET = 0,
        KAFKA = 100,
    }

}

export class WriteRecord extends jspb.Message { 
    getInput(): string;
    setInput(value: string): WriteRecord;

    getInputMetadataMap(): jspb.Map<string, string>;
    clearInputMetadataMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): WriteRecord.AsObject;
    static toObject(includeInstance: boolean, msg: WriteRecord): WriteRecord.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: WriteRecord, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): WriteRecord;
    static deserializeBinaryFromReader(message: WriteRecord, reader: jspb.BinaryReader): WriteRecord;
}

export namespace WriteRecord {
    export type AsObject = {
        input: string,

        inputMetadataMap: Array<[string, string]>,
    }
}

export class ErrorRecord extends jspb.Message { 
    getOccurredAtUnixTsUtc(): number;
    setOccurredAtUnixTsUtc(value: number): ErrorRecord;
    getError(): string;
    setError(value: string): ErrorRecord;

    getMetadataMap(): jspb.Map<string, Uint8Array | string>;
    clearMetadataMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ErrorRecord.AsObject;
    static toObject(includeInstance: boolean, msg: ErrorRecord): ErrorRecord.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ErrorRecord, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ErrorRecord;
    static deserializeBinaryFromReader(message: ErrorRecord, reader: jspb.BinaryReader): ErrorRecord;
}

export namespace ErrorRecord {
    export type AsObject = {
        occurredAtUnixTsUtc: number,
        error: string,

        metadataMap: Array<[string, Uint8Array | string]>,
    }
}