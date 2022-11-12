# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: visitManager.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12visitManager.proto\",\n\x06Invite\x12\x10\n\x08ID_Visit\x18\x01 \x02(\t\x12\x10\n\x08Username\x18\x02 \x02(\t\"F\n\x0eInviteResponse\x12\x10\n\x08ID_Visit\x18\x01 \x02(\t\x12\x10\n\x08Username\x18\x02 \x02(\t\x12\x10\n\x08Response\x18\x03 \x02(\t\"\x1f\n\x06Visits\x12\x15\n\x05Visit\x18\x01 \x03(\x0b\x32\x06.Visit\"=\n\x05Visit\x12\x10\n\x08ID_Visit\x18\x01 \x02(\r\x12\x0f\n\x07ID_Path\x18\x02 \x02(\t\x12\x11\n\tTimestamp\x18\x03 \x01(\x01\"0\n\nInputVisit\x12\x10\n\x08Username\x18\x01 \x02(\t\x12\x10\n\x08Pathname\x18\x02 \x02(\t\"\x15\n\x06Return\x12\x0b\n\x03Ret\x18\x01 \x02(\x05\"\x12\n\x04User\x12\n\n\x02ID\x18\x01 \x02(\t*+\n\nDifficulty\x12\x05\n\x01T\x10\x00\x12\x05\n\x01\x45\x10\x01\x12\x06\n\x02\x45\x45\x10\x02\x12\x07\n\x03\x45\x45\x41\x10\x03\x32\xb3\x01\n\x0bManageVisit\x12%\n\x0b\x41\x64\x64NewVisit\x12\x0b.InputVisit\x1a\x07.Return\"\x00\x12 \n\x0cGetAllVisits\x12\x05.User\x1a\x07.Visits\"\x00\x12\'\n\x11InviteUserToVisit\x12\x07.Invite\x1a\x07.Return\"\x00\x12\x32\n\x14\x41\x63\x63\x65ptOrRefuseInvite\x12\x0f.InviteResponse\x1a\x07.Return\"\x00\x42\x12Z\x10\x61pi_gateway/grpc')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'visitManager_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\020api_gateway/grpc'
  _DIFFICULTY._serialized_start=329
  _DIFFICULTY._serialized_end=372
  _INVITE._serialized_start=22
  _INVITE._serialized_end=66
  _INVITERESPONSE._serialized_start=68
  _INVITERESPONSE._serialized_end=138
  _VISITS._serialized_start=140
  _VISITS._serialized_end=171
  _VISIT._serialized_start=173
  _VISIT._serialized_end=234
  _INPUTVISIT._serialized_start=236
  _INPUTVISIT._serialized_end=284
  _RETURN._serialized_start=286
  _RETURN._serialized_end=307
  _USER._serialized_start=309
  _USER._serialized_end=327
  _MANAGEVISIT._serialized_start=375
  _MANAGEVISIT._serialized_end=554
# @@protoc_insertion_point(module_scope)