# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/execution_envs.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"flyteidl/core/execution_envs.proto\x12\rflyteidl.core\x1a\x1cgoogle/protobuf/struct.proto\"\x92\x01\n\x16\x45xecutionEnvAssignment\x12\x19\n\x08node_ids\x18\x01 \x03(\tR\x07nodeIds\x12\x1b\n\ttask_type\x18\x02 \x01(\tR\x08taskType\x12@\n\rexecution_env\x18\x03 \x01(\x0b\x32\x1b.flyteidl.core.ExecutionEnvR\x0c\x65xecutionEnv\"\xc1\x01\n\x0c\x45xecutionEnv\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n\x04type\x18\x02 \x01(\tR\x04type\x12\x31\n\x06\x65xtant\x18\x03 \x01(\x0b\x32\x17.google.protobuf.StructH\x00R\x06\x65xtant\x12-\n\x04spec\x18\x04 \x01(\x0b\x32\x17.google.protobuf.StructH\x00R\x04spec\x12\x18\n\x07version\x18\x05 \x01(\tR\x07versionB\r\n\x0b\x65nvironmentB\xb8\x01\n\x11\x63om.flyteidl.coreB\x12\x45xecutionEnvsProtoP\x01Z:github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core\xa2\x02\x03\x46\x43X\xaa\x02\rFlyteidl.Core\xca\x02\rFlyteidl\\Core\xe2\x02\x19\x46lyteidl\\Core\\GPBMetadata\xea\x02\x0e\x46lyteidl::Coreb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.core.execution_envs_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\021com.flyteidl.coreB\022ExecutionEnvsProtoP\001Z:github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core\242\002\003FCX\252\002\rFlyteidl.Core\312\002\rFlyteidl\\Core\342\002\031Flyteidl\\Core\\GPBMetadata\352\002\016Flyteidl::Core'
  _globals['_EXECUTIONENVASSIGNMENT']._serialized_start=84
  _globals['_EXECUTIONENVASSIGNMENT']._serialized_end=230
  _globals['_EXECUTIONENV']._serialized_start=233
  _globals['_EXECUTIONENV']._serialized_end=426
# @@protoc_insertion_point(module_scope)
