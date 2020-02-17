///
//  Generated code. Do not modify.
//  source: discovery.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'discovery.pb.dart' as $0;
export 'discovery.pb.dart';

class DiscoveryServiceClient extends $grpc.Client {
  static final _$registerService =
      $grpc.ClientMethod<$0.RegisterRequest, $0.RegisterResponse>(
          '/discovery.DiscoveryService/RegisterService',
          ($0.RegisterRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.RegisterResponse.fromBuffer(value));
  static final _$discover =
      $grpc.ClientMethod<$0.DiscoverRequest, $0.DiscoverResponse>(
          '/discovery.DiscoveryService/Discover',
          ($0.DiscoverRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.DiscoverResponse.fromBuffer(value));
  static final _$listAllServices =
      $grpc.ClientMethod<$0.ListRequest, $0.ListResponse>(
          '/discovery.DiscoveryService/ListAllServices',
          ($0.ListRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.ListResponse.fromBuffer(value));
  static final _$state = $grpc.ClientMethod<$0.StateRequest, $0.StateResponse>(
      '/discovery.DiscoveryService/State',
      ($0.StateRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.StateResponse.fromBuffer(value));

  DiscoveryServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.RegisterResponse> registerService(
      $0.RegisterRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$registerService, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.DiscoverResponse> discover($0.DiscoverRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$discover, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.ListResponse> listAllServices($0.ListRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$listAllServices, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.StateResponse> state($0.StateRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$state, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class DiscoveryServiceBase extends $grpc.Service {
  $core.String get $name => 'discovery.DiscoveryService';

  DiscoveryServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.RegisterRequest, $0.RegisterResponse>(
        'RegisterService',
        registerService_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RegisterRequest.fromBuffer(value),
        ($0.RegisterResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DiscoverRequest, $0.DiscoverResponse>(
        'Discover',
        discover_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DiscoverRequest.fromBuffer(value),
        ($0.DiscoverResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ListRequest, $0.ListResponse>(
        'ListAllServices',
        listAllServices_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ListRequest.fromBuffer(value),
        ($0.ListResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.StateRequest, $0.StateResponse>(
        'State',
        state_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.StateRequest.fromBuffer(value),
        ($0.StateResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.RegisterResponse> registerService_Pre(
      $grpc.ServiceCall call, $async.Future<$0.RegisterRequest> request) async {
    return registerService(call, await request);
  }

  $async.Future<$0.DiscoverResponse> discover_Pre(
      $grpc.ServiceCall call, $async.Future<$0.DiscoverRequest> request) async {
    return discover(call, await request);
  }

  $async.Future<$0.ListResponse> listAllServices_Pre(
      $grpc.ServiceCall call, $async.Future<$0.ListRequest> request) async {
    return listAllServices(call, await request);
  }

  $async.Future<$0.StateResponse> state_Pre(
      $grpc.ServiceCall call, $async.Future<$0.StateRequest> request) async {
    return state(call, await request);
  }

  $async.Future<$0.RegisterResponse> registerService(
      $grpc.ServiceCall call, $0.RegisterRequest request);
  $async.Future<$0.DiscoverResponse> discover(
      $grpc.ServiceCall call, $0.DiscoverRequest request);
  $async.Future<$0.ListResponse> listAllServices(
      $grpc.ServiceCall call, $0.ListRequest request);
  $async.Future<$0.StateResponse> state(
      $grpc.ServiceCall call, $0.StateRequest request);
}

class DiscoveryServiceV2Client extends $grpc.Client {
  static final _$lock = $grpc.ClientMethod<$0.LockRequest, $0.LockResponse>(
      '/discovery.DiscoveryServiceV2/Lock',
      ($0.LockRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.LockResponse.fromBuffer(value));
  static final _$registerV2 =
      $grpc.ClientMethod<$0.RegisterRequest, $0.RegisterResponse>(
          '/discovery.DiscoveryServiceV2/RegisterV2',
          ($0.RegisterRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.RegisterResponse.fromBuffer(value));
  static final _$get = $grpc.ClientMethod<$0.GetRequest, $0.GetResponse>(
      '/discovery.DiscoveryServiceV2/Get',
      ($0.GetRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.GetResponse.fromBuffer(value));
  static final _$unregister =
      $grpc.ClientMethod<$0.UnregisterRequest, $0.UnregisterResponse>(
          '/discovery.DiscoveryServiceV2/Unregister',
          ($0.UnregisterRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.UnregisterResponse.fromBuffer(value));
  static final _$masterElect =
      $grpc.ClientMethod<$0.MasterRequest, $0.MasterResponse>(
          '/discovery.DiscoveryServiceV2/MasterElect',
          ($0.MasterRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.MasterResponse.fromBuffer(value));

  DiscoveryServiceV2Client($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.LockResponse> lock($0.LockRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$lock, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.RegisterResponse> registerV2(
      $0.RegisterRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$registerV2, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.GetResponse> get($0.GetRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.UnregisterResponse> unregister(
      $0.UnregisterRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$unregister, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.MasterResponse> masterElect($0.MasterRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$masterElect, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class DiscoveryServiceV2ServiceBase extends $grpc.Service {
  $core.String get $name => 'discovery.DiscoveryServiceV2';

  DiscoveryServiceV2ServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.LockRequest, $0.LockResponse>(
        'Lock',
        lock_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.LockRequest.fromBuffer(value),
        ($0.LockResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RegisterRequest, $0.RegisterResponse>(
        'RegisterV2',
        registerV2_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RegisterRequest.fromBuffer(value),
        ($0.RegisterResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetRequest, $0.GetResponse>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetRequest.fromBuffer(value),
        ($0.GetResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UnregisterRequest, $0.UnregisterResponse>(
        'Unregister',
        unregister_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UnregisterRequest.fromBuffer(value),
        ($0.UnregisterResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.MasterRequest, $0.MasterResponse>(
        'MasterElect',
        masterElect_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.MasterRequest.fromBuffer(value),
        ($0.MasterResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.LockResponse> lock_Pre(
      $grpc.ServiceCall call, $async.Future<$0.LockRequest> request) async {
    return lock(call, await request);
  }

  $async.Future<$0.RegisterResponse> registerV2_Pre(
      $grpc.ServiceCall call, $async.Future<$0.RegisterRequest> request) async {
    return registerV2(call, await request);
  }

  $async.Future<$0.GetResponse> get_Pre(
      $grpc.ServiceCall call, $async.Future<$0.GetRequest> request) async {
    return get(call, await request);
  }

  $async.Future<$0.UnregisterResponse> unregister_Pre($grpc.ServiceCall call,
      $async.Future<$0.UnregisterRequest> request) async {
    return unregister(call, await request);
  }

  $async.Future<$0.MasterResponse> masterElect_Pre(
      $grpc.ServiceCall call, $async.Future<$0.MasterRequest> request) async {
    return masterElect(call, await request);
  }

  $async.Future<$0.LockResponse> lock(
      $grpc.ServiceCall call, $0.LockRequest request);
  $async.Future<$0.RegisterResponse> registerV2(
      $grpc.ServiceCall call, $0.RegisterRequest request);
  $async.Future<$0.GetResponse> get(
      $grpc.ServiceCall call, $0.GetRequest request);
  $async.Future<$0.UnregisterResponse> unregister(
      $grpc.ServiceCall call, $0.UnregisterRequest request);
  $async.Future<$0.MasterResponse> masterElect(
      $grpc.ServiceCall call, $0.MasterRequest request);
}
