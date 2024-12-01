//
//Hyperledger Cacti Plugin - Common Operator Primitives
//
//Contains/describes the Hyperledger Cacti Common Operator Primitives plugin.  These primitives require specific chaincode and weaver relays to be deployed on the network as described at https://hyperledger-cacti.github.io/cacti/weaver/introduction/.
//
//The version of the OpenAPI document: 2.1.0
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.8.0 with parameter "target=ts"
// @generated from file models/view_v1_pb.proto (package org.hyperledger.cacti.plugin.cacti.plugin.copm.core, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ViewAddressV1PB } from "./view_address_v1_pb_pb.js";

/**
 * @generated from message org.hyperledger.cacti.plugin.cacti.plugin.copm.core.ViewV1PB
 */
export class ViewV1PB extends Message<ViewV1PB> {
  /**
   * @generated from field: string network = 1;
   */
  network = "";

  /**
   * @generated from field: org.hyperledger.cacti.plugin.cacti.plugin.copm.core.ViewAddressV1PB view_address = 312477787;
   */
  viewAddress?: ViewAddressV1PB;

  constructor(data?: PartialMessage<ViewV1PB>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "org.hyperledger.cacti.plugin.cacti.plugin.copm.core.ViewV1PB";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "network", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 312477787, name: "view_address", kind: "message", T: ViewAddressV1PB },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ViewV1PB {
    return new ViewV1PB().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ViewV1PB {
    return new ViewV1PB().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ViewV1PB {
    return new ViewV1PB().fromJsonString(jsonString, options);
  }

  static equals(a: ViewV1PB | PlainMessage<ViewV1PB> | undefined, b: ViewV1PB | PlainMessage<ViewV1PB> | undefined): boolean {
    return proto3.util.equals(ViewV1PB, a, b);
  }
}

