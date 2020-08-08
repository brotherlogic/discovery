///
//  Generated code. Do not modify.
//  source: discovery.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'discovery.pbenum.dart';

export 'discovery.pbenum.dart';

class RegistryEntry extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegistryEntry', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOS(1, 'ip')
    ..a<$core.int>(2, 'port', $pb.PbFieldType.O3)
    ..aOS(3, 'name')
    ..aOB(4, 'externalPort')
    ..aOS(5, 'identifier')
    ..aOB(6, 'master')
    ..aInt64(7, 'registerTime')
    ..aInt64(8, 'timeToClean')
    ..aInt64(9, 'lastSeenTime')
    ..aOB(10, 'ignoresMaster')
    ..aInt64(11, 'masterTime')
    ..aOB(12, 'weakMaster')
    ..e<RegistryEntry_Version>(13, 'version', $pb.PbFieldType.OE, defaultOrMaker: RegistryEntry_Version.V1, valueOf: RegistryEntry_Version.valueOf, enumValues: RegistryEntry_Version.values)
    ..hasRequiredFields = false
  ;

  RegistryEntry._() : super();
  factory RegistryEntry() => create();
  factory RegistryEntry.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegistryEntry.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegistryEntry clone() => RegistryEntry()..mergeFromMessage(this);
  RegistryEntry copyWith(void Function(RegistryEntry) updates) => super.copyWith((message) => updates(message as RegistryEntry));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegistryEntry create() => RegistryEntry._();
  RegistryEntry createEmptyInstance() => create();
  static $pb.PbList<RegistryEntry> createRepeated() => $pb.PbList<RegistryEntry>();
  @$core.pragma('dart2js:noInline')
  static RegistryEntry getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegistryEntry>(create);
  static RegistryEntry _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get ip => $_getSZ(0);
  @$pb.TagNumber(1)
  set ip($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIp() => $_has(0);
  @$pb.TagNumber(1)
  void clearIp() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get port => $_getIZ(1);
  @$pb.TagNumber(2)
  set port($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPort() => $_has(1);
  @$pb.TagNumber(2)
  void clearPort() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(2);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get externalPort => $_getBF(3);
  @$pb.TagNumber(4)
  set externalPort($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasExternalPort() => $_has(3);
  @$pb.TagNumber(4)
  void clearExternalPort() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get identifier => $_getSZ(4);
  @$pb.TagNumber(5)
  set identifier($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasIdentifier() => $_has(4);
  @$pb.TagNumber(5)
  void clearIdentifier() => clearField(5);

  @$pb.TagNumber(6)
  $core.bool get master => $_getBF(5);
  @$pb.TagNumber(6)
  set master($core.bool v) { $_setBool(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasMaster() => $_has(5);
  @$pb.TagNumber(6)
  void clearMaster() => clearField(6);

  @$pb.TagNumber(7)
  $fixnum.Int64 get registerTime => $_getI64(6);
  @$pb.TagNumber(7)
  set registerTime($fixnum.Int64 v) { $_setInt64(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasRegisterTime() => $_has(6);
  @$pb.TagNumber(7)
  void clearRegisterTime() => clearField(7);

  @$pb.TagNumber(8)
  $fixnum.Int64 get timeToClean => $_getI64(7);
  @$pb.TagNumber(8)
  set timeToClean($fixnum.Int64 v) { $_setInt64(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasTimeToClean() => $_has(7);
  @$pb.TagNumber(8)
  void clearTimeToClean() => clearField(8);

  @$pb.TagNumber(9)
  $fixnum.Int64 get lastSeenTime => $_getI64(8);
  @$pb.TagNumber(9)
  set lastSeenTime($fixnum.Int64 v) { $_setInt64(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasLastSeenTime() => $_has(8);
  @$pb.TagNumber(9)
  void clearLastSeenTime() => clearField(9);

  @$pb.TagNumber(10)
  $core.bool get ignoresMaster => $_getBF(9);
  @$pb.TagNumber(10)
  set ignoresMaster($core.bool v) { $_setBool(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasIgnoresMaster() => $_has(9);
  @$pb.TagNumber(10)
  void clearIgnoresMaster() => clearField(10);

  @$pb.TagNumber(11)
  $fixnum.Int64 get masterTime => $_getI64(10);
  @$pb.TagNumber(11)
  set masterTime($fixnum.Int64 v) { $_setInt64(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasMasterTime() => $_has(10);
  @$pb.TagNumber(11)
  void clearMasterTime() => clearField(11);

  @$pb.TagNumber(12)
  $core.bool get weakMaster => $_getBF(11);
  @$pb.TagNumber(12)
  set weakMaster($core.bool v) { $_setBool(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasWeakMaster() => $_has(11);
  @$pb.TagNumber(12)
  void clearWeakMaster() => clearField(12);

  @$pb.TagNumber(13)
  RegistryEntry_Version get version => $_getN(12);
  @$pb.TagNumber(13)
  set version(RegistryEntry_Version v) { setField(13, v); }
  @$pb.TagNumber(13)
  $core.bool hasVersion() => $_has(12);
  @$pb.TagNumber(13)
  void clearVersion() => clearField(13);
}

class ServiceList extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServiceList', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..pc<RegistryEntry>(1, 'services', $pb.PbFieldType.PM, subBuilder: RegistryEntry.create)
    ..hasRequiredFields = false
  ;

  ServiceList._() : super();
  factory ServiceList() => create();
  factory ServiceList.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServiceList.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServiceList clone() => ServiceList()..mergeFromMessage(this);
  ServiceList copyWith(void Function(ServiceList) updates) => super.copyWith((message) => updates(message as ServiceList));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServiceList create() => ServiceList._();
  ServiceList createEmptyInstance() => create();
  static $pb.PbList<ServiceList> createRepeated() => $pb.PbList<ServiceList>();
  @$core.pragma('dart2js:noInline')
  static ServiceList getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServiceList>(create);
  static ServiceList _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<RegistryEntry> get services => $_getList(0);
}

class Empty extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Empty', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  Empty._() : super();
  factory Empty() => create();
  factory Empty.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Empty.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Empty clone() => Empty()..mergeFromMessage(this);
  Empty copyWith(void Function(Empty) updates) => super.copyWith((message) => updates(message as Empty));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Empty create() => Empty._();
  Empty createEmptyInstance() => create();
  static $pb.PbList<Empty> createRepeated() => $pb.PbList<Empty>();
  @$core.pragma('dart2js:noInline')
  static Empty getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Empty>(create);
  static Empty _defaultInstance;
}

class StateResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('StateResponse', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aInt64(1, 'longestCall')
    ..aOS(2, 'mostFrequent')
    ..a<$core.int>(3, 'frequency', $pb.PbFieldType.O3)
    ..aOS(4, 'count')
    ..hasRequiredFields = false
  ;

  StateResponse._() : super();
  factory StateResponse() => create();
  factory StateResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StateResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  StateResponse clone() => StateResponse()..mergeFromMessage(this);
  StateResponse copyWith(void Function(StateResponse) updates) => super.copyWith((message) => updates(message as StateResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StateResponse create() => StateResponse._();
  StateResponse createEmptyInstance() => create();
  static $pb.PbList<StateResponse> createRepeated() => $pb.PbList<StateResponse>();
  @$core.pragma('dart2js:noInline')
  static StateResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StateResponse>(create);
  static StateResponse _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get longestCall => $_getI64(0);
  @$pb.TagNumber(1)
  set longestCall($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLongestCall() => $_has(0);
  @$pb.TagNumber(1)
  void clearLongestCall() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get mostFrequent => $_getSZ(1);
  @$pb.TagNumber(2)
  set mostFrequent($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMostFrequent() => $_has(1);
  @$pb.TagNumber(2)
  void clearMostFrequent() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get frequency => $_getIZ(2);
  @$pb.TagNumber(3)
  set frequency($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasFrequency() => $_has(2);
  @$pb.TagNumber(3)
  void clearFrequency() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get count => $_getSZ(3);
  @$pb.TagNumber(4)
  set count($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCount() => $_has(3);
  @$pb.TagNumber(4)
  void clearCount() => clearField(4);
}

class StateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('StateRequest', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  StateRequest._() : super();
  factory StateRequest() => create();
  factory StateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  StateRequest clone() => StateRequest()..mergeFromMessage(this);
  StateRequest copyWith(void Function(StateRequest) updates) => super.copyWith((message) => updates(message as StateRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StateRequest create() => StateRequest._();
  StateRequest createEmptyInstance() => create();
  static $pb.PbList<StateRequest> createRepeated() => $pb.PbList<StateRequest>();
  @$core.pragma('dart2js:noInline')
  static StateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StateRequest>(create);
  static StateRequest _defaultInstance;
}

class RegisterRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegisterRequest', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOM<RegistryEntry>(1, 'service', subBuilder: RegistryEntry.create)
    ..aOS(2, 'caller')
    ..aOB(4, 'fanout')
    ..hasRequiredFields = false
  ;

  RegisterRequest._() : super();
  factory RegisterRequest() => create();
  factory RegisterRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegisterRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegisterRequest clone() => RegisterRequest()..mergeFromMessage(this);
  RegisterRequest copyWith(void Function(RegisterRequest) updates) => super.copyWith((message) => updates(message as RegisterRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegisterRequest create() => RegisterRequest._();
  RegisterRequest createEmptyInstance() => create();
  static $pb.PbList<RegisterRequest> createRepeated() => $pb.PbList<RegisterRequest>();
  @$core.pragma('dart2js:noInline')
  static RegisterRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegisterRequest>(create);
  static RegisterRequest _defaultInstance;

  @$pb.TagNumber(1)
  RegistryEntry get service => $_getN(0);
  @$pb.TagNumber(1)
  set service(RegistryEntry v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasService() => $_has(0);
  @$pb.TagNumber(1)
  void clearService() => clearField(1);
  @$pb.TagNumber(1)
  RegistryEntry ensureService() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get caller => $_getSZ(1);
  @$pb.TagNumber(2)
  set caller($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCaller() => $_has(1);
  @$pb.TagNumber(2)
  void clearCaller() => clearField(2);

  @$pb.TagNumber(4)
  $core.bool get fanout => $_getBF(2);
  @$pb.TagNumber(4)
  set fanout($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(4)
  $core.bool hasFanout() => $_has(2);
  @$pb.TagNumber(4)
  void clearFanout() => clearField(4);
}

class RegisterResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegisterResponse', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOM<RegistryEntry>(1, 'service', subBuilder: RegistryEntry.create)
    ..hasRequiredFields = false
  ;

  RegisterResponse._() : super();
  factory RegisterResponse() => create();
  factory RegisterResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegisterResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegisterResponse clone() => RegisterResponse()..mergeFromMessage(this);
  RegisterResponse copyWith(void Function(RegisterResponse) updates) => super.copyWith((message) => updates(message as RegisterResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegisterResponse create() => RegisterResponse._();
  RegisterResponse createEmptyInstance() => create();
  static $pb.PbList<RegisterResponse> createRepeated() => $pb.PbList<RegisterResponse>();
  @$core.pragma('dart2js:noInline')
  static RegisterResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegisterResponse>(create);
  static RegisterResponse _defaultInstance;

  @$pb.TagNumber(1)
  RegistryEntry get service => $_getN(0);
  @$pb.TagNumber(1)
  set service(RegistryEntry v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasService() => $_has(0);
  @$pb.TagNumber(1)
  void clearService() => clearField(1);
  @$pb.TagNumber(1)
  RegistryEntry ensureService() => $_ensure(0);
}

class DiscoverRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DiscoverRequest', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOM<RegistryEntry>(1, 'request', subBuilder: RegistryEntry.create)
    ..aOS(2, 'caller')
    ..hasRequiredFields = false
  ;

  DiscoverRequest._() : super();
  factory DiscoverRequest() => create();
  factory DiscoverRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DiscoverRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DiscoverRequest clone() => DiscoverRequest()..mergeFromMessage(this);
  DiscoverRequest copyWith(void Function(DiscoverRequest) updates) => super.copyWith((message) => updates(message as DiscoverRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DiscoverRequest create() => DiscoverRequest._();
  DiscoverRequest createEmptyInstance() => create();
  static $pb.PbList<DiscoverRequest> createRepeated() => $pb.PbList<DiscoverRequest>();
  @$core.pragma('dart2js:noInline')
  static DiscoverRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DiscoverRequest>(create);
  static DiscoverRequest _defaultInstance;

  @$pb.TagNumber(1)
  RegistryEntry get request => $_getN(0);
  @$pb.TagNumber(1)
  set request(RegistryEntry v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasRequest() => $_has(0);
  @$pb.TagNumber(1)
  void clearRequest() => clearField(1);
  @$pb.TagNumber(1)
  RegistryEntry ensureRequest() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get caller => $_getSZ(1);
  @$pb.TagNumber(2)
  set caller($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCaller() => $_has(1);
  @$pb.TagNumber(2)
  void clearCaller() => clearField(2);
}

class DiscoverResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DiscoverResponse', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOM<RegistryEntry>(1, 'service', subBuilder: RegistryEntry.create)
    ..hasRequiredFields = false
  ;

  DiscoverResponse._() : super();
  factory DiscoverResponse() => create();
  factory DiscoverResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DiscoverResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DiscoverResponse clone() => DiscoverResponse()..mergeFromMessage(this);
  DiscoverResponse copyWith(void Function(DiscoverResponse) updates) => super.copyWith((message) => updates(message as DiscoverResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DiscoverResponse create() => DiscoverResponse._();
  DiscoverResponse createEmptyInstance() => create();
  static $pb.PbList<DiscoverResponse> createRepeated() => $pb.PbList<DiscoverResponse>();
  @$core.pragma('dart2js:noInline')
  static DiscoverResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DiscoverResponse>(create);
  static DiscoverResponse _defaultInstance;

  @$pb.TagNumber(1)
  RegistryEntry get service => $_getN(0);
  @$pb.TagNumber(1)
  set service(RegistryEntry v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasService() => $_has(0);
  @$pb.TagNumber(1)
  void clearService() => clearField(1);
  @$pb.TagNumber(1)
  RegistryEntry ensureService() => $_ensure(0);
}

class ListRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ListRequest', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOS(1, 'caller')
    ..hasRequiredFields = false
  ;

  ListRequest._() : super();
  factory ListRequest() => create();
  factory ListRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ListRequest clone() => ListRequest()..mergeFromMessage(this);
  ListRequest copyWith(void Function(ListRequest) updates) => super.copyWith((message) => updates(message as ListRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListRequest create() => ListRequest._();
  ListRequest createEmptyInstance() => create();
  static $pb.PbList<ListRequest> createRepeated() => $pb.PbList<ListRequest>();
  @$core.pragma('dart2js:noInline')
  static ListRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListRequest>(create);
  static ListRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get caller => $_getSZ(0);
  @$pb.TagNumber(1)
  set caller($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCaller() => $_has(0);
  @$pb.TagNumber(1)
  void clearCaller() => clearField(1);
}

class ListResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ListResponse', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOM<ServiceList>(1, 'services', subBuilder: ServiceList.create)
    ..hasRequiredFields = false
  ;

  ListResponse._() : super();
  factory ListResponse() => create();
  factory ListResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ListResponse clone() => ListResponse()..mergeFromMessage(this);
  ListResponse copyWith(void Function(ListResponse) updates) => super.copyWith((message) => updates(message as ListResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListResponse create() => ListResponse._();
  ListResponse createEmptyInstance() => create();
  static $pb.PbList<ListResponse> createRepeated() => $pb.PbList<ListResponse>();
  @$core.pragma('dart2js:noInline')
  static ListResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListResponse>(create);
  static ListResponse _defaultInstance;

  @$pb.TagNumber(1)
  ServiceList get services => $_getN(0);
  @$pb.TagNumber(1)
  set services(ServiceList v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServices() => $_has(0);
  @$pb.TagNumber(1)
  void clearServices() => clearField(1);
  @$pb.TagNumber(1)
  ServiceList ensureServices() => $_ensure(0);
}

class GetRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('GetRequest', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOS(1, 'job')
    ..aOS(2, 'server')
    ..aOS(3, 'friend')
    ..hasRequiredFields = false
  ;

  GetRequest._() : super();
  factory GetRequest() => create();
  factory GetRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  GetRequest clone() => GetRequest()..mergeFromMessage(this);
  GetRequest copyWith(void Function(GetRequest) updates) => super.copyWith((message) => updates(message as GetRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetRequest create() => GetRequest._();
  GetRequest createEmptyInstance() => create();
  static $pb.PbList<GetRequest> createRepeated() => $pb.PbList<GetRequest>();
  @$core.pragma('dart2js:noInline')
  static GetRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetRequest>(create);
  static GetRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get job => $_getSZ(0);
  @$pb.TagNumber(1)
  set job($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasJob() => $_has(0);
  @$pb.TagNumber(1)
  void clearJob() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get server => $_getSZ(1);
  @$pb.TagNumber(2)
  set server($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasServer() => $_has(1);
  @$pb.TagNumber(2)
  void clearServer() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get friend => $_getSZ(2);
  @$pb.TagNumber(3)
  set friend($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasFriend() => $_has(2);
  @$pb.TagNumber(3)
  void clearFriend() => clearField(3);
}

class GetResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('GetResponse', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..pc<RegistryEntry>(1, 'services', $pb.PbFieldType.PM, subBuilder: RegistryEntry.create)
    ..hasRequiredFields = false
  ;

  GetResponse._() : super();
  factory GetResponse() => create();
  factory GetResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  GetResponse clone() => GetResponse()..mergeFromMessage(this);
  GetResponse copyWith(void Function(GetResponse) updates) => super.copyWith((message) => updates(message as GetResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetResponse create() => GetResponse._();
  GetResponse createEmptyInstance() => create();
  static $pb.PbList<GetResponse> createRepeated() => $pb.PbList<GetResponse>();
  @$core.pragma('dart2js:noInline')
  static GetResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetResponse>(create);
  static GetResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<RegistryEntry> get services => $_getList(0);
}

class UnregisterRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UnregisterRequest', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOM<RegistryEntry>(1, 'service', subBuilder: RegistryEntry.create)
    ..aOB(2, 'fanout')
    ..aOS(3, 'address')
    ..aOS(4, 'caller')
    ..hasRequiredFields = false
  ;

  UnregisterRequest._() : super();
  factory UnregisterRequest() => create();
  factory UnregisterRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UnregisterRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UnregisterRequest clone() => UnregisterRequest()..mergeFromMessage(this);
  UnregisterRequest copyWith(void Function(UnregisterRequest) updates) => super.copyWith((message) => updates(message as UnregisterRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UnregisterRequest create() => UnregisterRequest._();
  UnregisterRequest createEmptyInstance() => create();
  static $pb.PbList<UnregisterRequest> createRepeated() => $pb.PbList<UnregisterRequest>();
  @$core.pragma('dart2js:noInline')
  static UnregisterRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UnregisterRequest>(create);
  static UnregisterRequest _defaultInstance;

  @$pb.TagNumber(1)
  RegistryEntry get service => $_getN(0);
  @$pb.TagNumber(1)
  set service(RegistryEntry v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasService() => $_has(0);
  @$pb.TagNumber(1)
  void clearService() => clearField(1);
  @$pb.TagNumber(1)
  RegistryEntry ensureService() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.bool get fanout => $_getBF(1);
  @$pb.TagNumber(2)
  set fanout($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFanout() => $_has(1);
  @$pb.TagNumber(2)
  void clearFanout() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get address => $_getSZ(2);
  @$pb.TagNumber(3)
  set address($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAddress() => $_has(2);
  @$pb.TagNumber(3)
  void clearAddress() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get caller => $_getSZ(3);
  @$pb.TagNumber(4)
  set caller($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCaller() => $_has(3);
  @$pb.TagNumber(4)
  void clearCaller() => clearField(4);
}

class UnregisterResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UnregisterResponse', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  UnregisterResponse._() : super();
  factory UnregisterResponse() => create();
  factory UnregisterResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UnregisterResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UnregisterResponse clone() => UnregisterResponse()..mergeFromMessage(this);
  UnregisterResponse copyWith(void Function(UnregisterResponse) updates) => super.copyWith((message) => updates(message as UnregisterResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UnregisterResponse create() => UnregisterResponse._();
  UnregisterResponse createEmptyInstance() => create();
  static $pb.PbList<UnregisterResponse> createRepeated() => $pb.PbList<UnregisterResponse>();
  @$core.pragma('dart2js:noInline')
  static UnregisterResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UnregisterResponse>(create);
  static UnregisterResponse _defaultInstance;
}

class LockRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LockRequest', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOS(1, 'job')
    ..aInt64(2, 'lockKey')
    ..aOS(3, 'requestor')
    ..hasRequiredFields = false
  ;

  LockRequest._() : super();
  factory LockRequest() => create();
  factory LockRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LockRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LockRequest clone() => LockRequest()..mergeFromMessage(this);
  LockRequest copyWith(void Function(LockRequest) updates) => super.copyWith((message) => updates(message as LockRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LockRequest create() => LockRequest._();
  LockRequest createEmptyInstance() => create();
  static $pb.PbList<LockRequest> createRepeated() => $pb.PbList<LockRequest>();
  @$core.pragma('dart2js:noInline')
  static LockRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LockRequest>(create);
  static LockRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get job => $_getSZ(0);
  @$pb.TagNumber(1)
  set job($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasJob() => $_has(0);
  @$pb.TagNumber(1)
  void clearJob() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get lockKey => $_getI64(1);
  @$pb.TagNumber(2)
  set lockKey($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLockKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearLockKey() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get requestor => $_getSZ(2);
  @$pb.TagNumber(3)
  set requestor($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRequestor() => $_has(2);
  @$pb.TagNumber(3)
  void clearRequestor() => clearField(3);
}

class LockResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LockResponse', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  LockResponse._() : super();
  factory LockResponse() => create();
  factory LockResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LockResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LockResponse clone() => LockResponse()..mergeFromMessage(this);
  LockResponse copyWith(void Function(LockResponse) updates) => super.copyWith((message) => updates(message as LockResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LockResponse create() => LockResponse._();
  LockResponse createEmptyInstance() => create();
  static $pb.PbList<LockResponse> createRepeated() => $pb.PbList<LockResponse>();
  @$core.pragma('dart2js:noInline')
  static LockResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LockResponse>(create);
  static LockResponse _defaultInstance;
}

class MasterRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('MasterRequest', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOM<RegistryEntry>(1, 'service', subBuilder: RegistryEntry.create)
    ..aInt64(2, 'lockKey')
    ..aOB(3, 'masterElect')
    ..aOB(4, 'fanout')
    ..hasRequiredFields = false
  ;

  MasterRequest._() : super();
  factory MasterRequest() => create();
  factory MasterRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MasterRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  MasterRequest clone() => MasterRequest()..mergeFromMessage(this);
  MasterRequest copyWith(void Function(MasterRequest) updates) => super.copyWith((message) => updates(message as MasterRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MasterRequest create() => MasterRequest._();
  MasterRequest createEmptyInstance() => create();
  static $pb.PbList<MasterRequest> createRepeated() => $pb.PbList<MasterRequest>();
  @$core.pragma('dart2js:noInline')
  static MasterRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MasterRequest>(create);
  static MasterRequest _defaultInstance;

  @$pb.TagNumber(1)
  RegistryEntry get service => $_getN(0);
  @$pb.TagNumber(1)
  set service(RegistryEntry v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasService() => $_has(0);
  @$pb.TagNumber(1)
  void clearService() => clearField(1);
  @$pb.TagNumber(1)
  RegistryEntry ensureService() => $_ensure(0);

  @$pb.TagNumber(2)
  $fixnum.Int64 get lockKey => $_getI64(1);
  @$pb.TagNumber(2)
  set lockKey($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLockKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearLockKey() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get masterElect => $_getBF(2);
  @$pb.TagNumber(3)
  set masterElect($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasMasterElect() => $_has(2);
  @$pb.TagNumber(3)
  void clearMasterElect() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get fanout => $_getBF(3);
  @$pb.TagNumber(4)
  set fanout($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFanout() => $_has(3);
  @$pb.TagNumber(4)
  void clearFanout() => clearField(4);
}

class MasterResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('MasterResponse', package: const $pb.PackageName('discovery'), createEmptyInstance: create)
    ..aOM<RegistryEntry>(1, 'service', subBuilder: RegistryEntry.create)
    ..hasRequiredFields = false
  ;

  MasterResponse._() : super();
  factory MasterResponse() => create();
  factory MasterResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MasterResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  MasterResponse clone() => MasterResponse()..mergeFromMessage(this);
  MasterResponse copyWith(void Function(MasterResponse) updates) => super.copyWith((message) => updates(message as MasterResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MasterResponse create() => MasterResponse._();
  MasterResponse createEmptyInstance() => create();
  static $pb.PbList<MasterResponse> createRepeated() => $pb.PbList<MasterResponse>();
  @$core.pragma('dart2js:noInline')
  static MasterResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MasterResponse>(create);
  static MasterResponse _defaultInstance;

  @$pb.TagNumber(1)
  RegistryEntry get service => $_getN(0);
  @$pb.TagNumber(1)
  set service(RegistryEntry v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasService() => $_has(0);
  @$pb.TagNumber(1)
  void clearService() => clearField(1);
  @$pb.TagNumber(1)
  RegistryEntry ensureService() => $_ensure(0);
}

